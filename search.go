package main

import (
	"math"
)

type search struct {
	a      float64
	b      float64
	e      float64
	x1     float64
	x2     float64
	ms     float64
	f1     float64
	f2     float64
	points [][2]float64
}

func (s *search) calcMiddleOfSegment() {
	s.ms = (s.a + s.b) / 2
}

func (s *search) calcTheta() float64 {
	return (1 + math.Sqrt(5)) / 2
}

func (s *search) calcX1Gold() {
	s.x1 = s.a + (1/math.Pow(s.calcTheta(), 2))*(s.b-s.a)
}

func (s *search) calcX2Gold() {
	s.x2 = s.a + (1/s.calcTheta())*(s.b-s.a)
}

func (s *search) calcX1Dichotomous() {
	s.x1 = s.ms - (s.e / 2)
}

func (s *search) calcX2Dichotomous() {
	s.x2 = s.ms + (s.e / 2)
}

func (s *search) calcDichotomous() {
	s.calcMiddleOfSegment()
	s.calcX1Dichotomous()
	s.calcX2Dichotomous()
	s.f1 = calcFunction(s.x1)
	s.f2 = calcFunction(s.x2)
}

func (s *search) savePoint(x, f float64) {
	s.points = append(s.points, [2]float64{x, f})
}
