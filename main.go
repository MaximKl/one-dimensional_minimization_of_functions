package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
)

type result struct {
	name   string
	points [][2]float64
}

func main() {
	goldenReceiver := make(chan result)
	dichotomousReceiver := make(chan result)

	go goldenSectionSearch(1, 2, 0.0001, goldenReceiver)
	go dichotomousSearch(1, 2, 0.0001, dichotomousReceiver)

	for runtime.NumGoroutine() != 1 {
		select {
		case result := <-goldenReceiver:
			getBestPoint("Golden Section method", result)
			writePoints(fmt.Sprintf("Golden_Section%s", result.name), result.points)
		case result := <-dichotomousReceiver:
			getBestPoint("Dichotomous method", result)
			writePoints(fmt.Sprintf("Dichotomous%s", result.name), result.points)
		}
	}
}

func calcFunction(x float64) float64 {
	return math.Pow(x, 3) - 0.3*math.Pow(x, 2) - 2.97*x
}

func dichotomousSearch(a, b int, e float64, receiver chan result) {
	s := &search{
		a: float64(a),
		b: float64(b),
		e: e}
	s.calcDichotomous()
	for !(math.Abs(s.b-s.a) <= e) {
		if s.f1 >= s.f2 {
			s.savePoint(s.x1, s.f1)
			s.a = s.ms
			s.calcDichotomous()
		} else {
			s.savePoint(s.x2, s.f2)
			s.b = s.ms
			s.calcDichotomous()
		}
	}
	s.calcMiddleOfSegment()
	s.savePoint(s.ms, calcFunction(s.ms))
	receiver <- result{
		name:   fmt.Sprintf("(%v,%v)", a, b),
		points: s.points}
}

func goldenSectionSearch(a, b int, e float64, receiver chan result) {
	s := &search{
		a: float64(a),
		b: float64(b)}
	s.calcX1Gold()
	s.calcX2Gold()
	s.f1 = calcFunction(s.x1)
	s.f2 = calcFunction(s.x2)
	for !(math.Abs(s.b-s.a) <= e) {
		if s.f1 >= s.f2 {
			s.savePoint(s.x1, s.f1)
			s.a = s.x1
			s.x1 = s.x2
			s.f1 = s.f2
			s.calcX2Gold()
			s.f2 = calcFunction(s.x2)
		} else {
			s.savePoint(s.x2, s.f2)
			s.b = s.x2
			s.x2 = s.x1
			s.f2 = s.f1
			s.calcX1Gold()
			s.f1 = calcFunction(s.x1)
		}
	}
	s.calcMiddleOfSegment()
	s.savePoint(s.ms, calcFunction(s.ms))
	receiver <- result{
		name:   fmt.Sprintf("(%v,%v)", a, b),
		points: s.points}
}

func getBestPoint(methodName string, result result) {
	fmt.Printf("----Best points of %s-----\n%s of %s X: %v\n", methodName, methodName, result.name, result.points[len(result.points)-1 : len(result.points)][0][0])
	fmt.Printf("%s of %s F: %g\n", methodName, result.name, result.points[len(result.points)-1 : len(result.points)][0][1])
	fmt.Printf("%s of %s K: %v\n", methodName, result.name, len(result.points)-1)
}

func writePoints(fileName string, points [][2]float64) {
	stringPoints := ""
	fullFileName := fmt.Sprintf("%s.txt", fileName)
	for _, p := range points {
		stringPoints = fmt.Sprintf("%s(%v, %v)\n", stringPoints, p[0], p[1])
	}
	err := os.WriteFile(fullFileName, []byte(stringPoints), 0777)
	if err != nil {
		fmt.Println("\nError while was writing into a file")
		return
	}
	fmt.Printf("All intermediate results have been successfully written to %s\n\n", fullFileName)
}
