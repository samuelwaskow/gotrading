package nn

import "gonum.org/v1/gonum/mat"

type NN interface {
	Train(x, y mat.Matrix)
	Evaluate(x, y mat.Matrix)
	Predict(x mat.Matrix)
}
