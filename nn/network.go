package nn

import (
	"gonum.org/v1/gonum/mat"
)

type Config struct {
	epochs    int
	batchsize int
	eta       float64
}

type MLP struct {
	numlayers int
	sizes     []int
	biases    []*mat.Dense
	weights   []*mat.Dense
	config    Config
}
