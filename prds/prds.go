package prds

import "slices"

type Pred[V any] = func(val V) bool

func Eq[V comparable](left V) func(right V) bool {
	return func(right V) bool {
		return left == right
	}
}

func Nq[V comparable](left V) func(right V) bool {
	return func(right V) bool {
		return left != right
	}
}

func Nx[S ~[]E, E comparable](elements S) func(next E) bool {
	pos := 0
	return func(elem E) bool {
		if pos == len(elements) {
			return false
		}
		equals := elements[pos] == elem
		pos++
		return equals
	}
}

func In[S ~[]E, E comparable](elements S) func(elem E) bool {
	return func(elem E) bool {
		return slices.Contains(elements, elem)
	}
}
