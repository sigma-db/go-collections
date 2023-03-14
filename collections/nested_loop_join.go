package collections

type JoinEqualityPredicate[R, S any, T comparable] struct {
	Left  Offset[R, T]
	Right Offset[S, T]
}

func (p *JoinEqualityPredicate[R, S, T]) Evaluate(r *R, s *S) bool {
	return *p.Left.Get(r) == *p.Right.Get(s)
}

func newJoinEqualityPredicate[R, S any, T comparable](s Selector[R, T], t Selector[S, T]) *JoinEqualityPredicate[R, S, T] {
	return &JoinEqualityPredicate[R, S, T]{NewOffset(s), NewOffset(t)}
}

func NestedLoopJoin[R, S any, T comparable](r []R, s []S, p Selector[R, T], q Selector[S, T]) []Pair[R, S] {
	result := make([]Pair[R, S], 0, len(r)*len(s))
	predicate := newJoinEqualityPredicate(p, q)
	for _, r := range r {
		for _, s := range s {
			if predicate.Evaluate(&r, &s) {
				result = append(result, Pair[R, S]{r, s})
			}
		}
	}
	return result
}
