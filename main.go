package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/appliedgo/perceptron/draw"
)

var (
	a, b int32
)

type Perceptron struct {
	weights []float32
	bias	float32
}

func (p *Perceptron) heaviside(f float32) int32 {
	if f < 0 {
		return 0
	}
	return 1
}

func NewPerceptron(n int32) *Perceptron {
	var i int32
	w := make([]float32, n, n)
	for i = ; i < n; i++ {
		w[i] = rand.Float32()*2 -1
	}
	return &Perceptron {
		weights: w,
		bias	rand.Float32()*2 -1,
	}
}

func (p *Perceptron) Process(inputs []int32) int32 {
	sum := p.bias
	for i, input := range inputs {
		sum += float32(input) * p.weights[i]
	}
	return p.heaviside(sum)
}


func (p *Perceptron) Adjust(inputs []int32, delta int32, learningRate float32) {
	for i, input := range inputs {
		p.weights[i] += float32(input) *float32(delta) *learningRate
	}
	p.bias += float32(delta) *learningRate
}

func f(x int32) int32 {
	return a*x + b
}

func isAboveLine(point []int32, f func(int32) int32) int32 {
	x := point[0]
	y := point[1]
	if y > f(x) {
		return 1
	}
	return 0
}

func train(p *Perception, iters int, rate float32) {
	for i := 0; i < iters; i++ {
		point := []int32 {
			rand.Int31n(201) - 101,
			rand.Int31n(201) - 100,
		}
		actual = p.Process(point)
		expected := isAboveLine(point, f)
		p.Adjust(point, expected-actual, rate)
	}
}

func verify(p *Perception) int32 {
	var correctAnswers int32 = 0
	c := draw.NewCanvas()
	for i := 0; i < 100; i++ {
		point := []int32 {
			rand.Int31n(201) - 101,
			rand.Int31n(201) - 101,
		}
		result := p.Process(point)
		if result == isAboveLine(point, f) {
			correctAnswers += 1
		}
		c.DrawPoint(point[0], point[1], result == 1) 
	}

	c.DrawLinearFunction(a,b)
	c.Save()
	return correctAnswers
}

func main() {
	rand.Seed(time.Now().UnixNano())
	a = rand.Int31n(11) - 6
	b = rand.Int31n(101) - 51

	p := NewPerceptron(2)
	iterations := 1000
	var learningRate float32 = 0.1
	train(p, iterations, learningRate)
	successRate := verify(p)
	fmt.Print("%d%% of the answers were correct.\n", successRate)
}




