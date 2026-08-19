package preds

import (
	"slices"
)

type Pred[V any] func(val V) bool

func (p Pred[V]) Or(right Pred[V]) Pred[V] {
	return func(val V) bool {
		return p(val) || right(val)
	}
}

func (p Pred[V]) And(right Pred[V]) Pred[V] {
	return func(val V) bool {
		return p(val) && right(val)
	}
}

func (p Pred[V]) Then(right Pred[V]) Pred[V] {
	return func(val V) bool {
		return !p(val) || right(val)
	}
}

func (p Pred[V]) Equiv(right Pred[V]) Pred[V] {
	return func(val V) bool {
		return (!p(val) || right(val)) &&
			(!right(val) || p(val))
	}
}

func (p Pred[V]) Not(predicate Pred[V]) Pred[V] {
	return func(val V) bool {
		return !predicate(val)
	}
}

func Equal[C comparable](left C) Pred[C] {
	return func(right C) bool {
		return left == right
	}
}

func NotEqual[C comparable](left C) Pred[C] {
	return func(right C) bool {
		return left != right
	}
}

func Find[C comparable](elements ...C) Pred[C] {
	return func(elem C) bool {
		return slices.Contains(elements, elem)
	}
}

func NotFind[C comparable](elements ...C) Pred[C] {
	return func(elem C) bool {
		return !slices.Contains(elements, elem)
	}
}

func Iter[C comparable](elements ...C) Pred[C] {
	pos := 0
	return func(elem C) bool {
		if pos == len(elements) {
			return false
		}
		equals := elements[pos] == elem
		pos++
		return equals
	}
}
