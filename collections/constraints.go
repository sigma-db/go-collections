package collections

import "golang.org/x/exp/constraints"

type Integer = constraints.Integer
type Float = constraints.Float
type Number interface{ Integer | Float }
