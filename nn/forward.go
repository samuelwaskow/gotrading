package nn

import "gonum.org/v1/gonum/mat"

/**
 * forward takes input data and returns the 'activation' and 'z'
 * from each layer - z = w.x +b and a = sigmoid(z)
 **/
func (n *MLP) forward(x mat.Matrix) (as, zs []mat.Matrix) {

	as = append(as, x)

	return

}
