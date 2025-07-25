package tuple

type Tuple1[T1 any] struct {
	Val1 T1
}

func (Tuple1[T1]) Size() int { return 1 }

func (t Tuple1[T1]) Values() T1 {
	return t.Val1
}

func (t Tuple1[T1]) values() []any {
	values := make([]any, 1)
	values[0] = t.Val1

	return values
}

func New1[T1 any](v1 T1) *Tuple1[T1] {
	return &Tuple1[T1]{
		Val1: v1,
	}
}

var _ Tuple = *new(Tuple1[any])

type Tuple2[T1, T2 any] struct {
	Val1 T1
	Val2 T2
}

func (Tuple2[T1, T2]) Size() int { return 2 }

func (t Tuple2[T1, T2]) Values() (T1, T2) {
	return t.Val1, t.Val2
}

func (t Tuple2[T1, T2]) values() []any {
	values := make([]any, 2)
	values[0] = t.Val1
	values[1] = t.Val2

	return values
}

func New2[T1, T2 any](v1 T1, v2 T2) *Tuple2[T1, T2] {
	return &Tuple2[T1, T2]{
		Val1: v1,
		Val2: v2,
	}
}

var _ Tuple = *new(Tuple2[any, any])

type Tuple3[T1, T2, T3 any] struct {
	Val1 T1
	Val2 T2
	Val3 T3
}

func (Tuple3[T1, T2, T3]) Size() int { return 3 }

func (t Tuple3[T1, T2, T3]) Values() (T1, T2, T3) {
	return t.Val1, t.Val2, t.Val3
}

func (t Tuple3[T1, T2, T3]) values() []any {
	values := make([]any, 3)
	values[0] = t.Val1
	values[1] = t.Val2
	values[2] = t.Val3

	return values
}

func New3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3) *Tuple3[T1, T2, T3] {
	return &Tuple3[T1, T2, T3]{
		Val1: v1,
		Val2: v2,
		Val3: v3,
	}
}

var _ Tuple = *new(Tuple3[any, any, any])

type Tuple4[T1, T2, T3, T4 any] struct {
	Val1 T1
	Val2 T2
	Val3 T3
	Val4 T4
}

func (Tuple4[T1, T2, T3, T4]) Size() int { return 4 }

func (t Tuple4[T1, T2, T3, T4]) Values() (T1, T2, T3, T4) {
	return t.Val1, t.Val2, t.Val3, t.Val4
}

func (t Tuple4[T1, T2, T3, T4]) values() []any {
	values := make([]any, 4)
	values[0] = t.Val1
	values[1] = t.Val2
	values[2] = t.Val3
	values[3] = t.Val4

	return values
}

func New4[T1, T2, T3, T4 any](v1 T1, v2 T2, v3 T3, v4 T4) *Tuple4[T1, T2, T3, T4] {
	return &Tuple4[T1, T2, T3, T4]{
		Val1: v1,
		Val2: v2,
		Val3: v3,
		Val4: v4,
	}
}

var _ Tuple = *new(Tuple4[any, any, any, any])

type Tuple5[T1, T2, T3, T4, T5 any] struct {
	Val1 T1
	Val2 T2
	Val3 T3
	Val4 T4
	Val5 T5
}

func (Tuple5[T1, T2, T3, T4, T5]) Size() int { return 5 }

func (t Tuple5[T1, T2, T3, T4, T5]) Values() (T1, T2, T3, T4, T5) {
	return t.Val1, t.Val2, t.Val3, t.Val4, t.Val5
}

func (t Tuple5[T1, T2, T3, T4, T5]) values() []any {
	values := make([]any, 5)
	values[0] = t.Val1
	values[1] = t.Val2
	values[2] = t.Val3
	values[3] = t.Val4
	values[4] = t.Val5

	return values
}

func New5[T1, T2, T3, T4, T5 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5) *Tuple5[T1, T2, T3, T4, T5] {
	return &Tuple5[T1, T2, T3, T4, T5]{
		Val1: v1,
		Val2: v2,
		Val3: v3,
		Val4: v4,
		Val5: v5,
	}
}

var _ Tuple = *new(Tuple5[any, any, any, any, any])

type Tuple6[T1, T2, T3, T4, T5, T6 any] struct {
	Val1 T1
	Val2 T2
	Val3 T3
	Val4 T4
	Val5 T5
	Val6 T6
}

func (Tuple6[T1, T2, T3, T4, T5, T6]) Size() int { return 6 }

func (t Tuple6[T1, T2, T3, T4, T5, T6]) Values() (T1, T2, T3, T4, T5, T6) {
	return t.Val1, t.Val2, t.Val3, t.Val4, t.Val5, t.Val6
}

func (t Tuple6[T1, T2, T3, T4, T5, T6]) values() []any {
	values := make([]any, 6)
	values[0] = t.Val1
	values[1] = t.Val2
	values[2] = t.Val3
	values[3] = t.Val4
	values[4] = t.Val5
	values[5] = t.Val6

	return values
}

func New6[T1, T2, T3, T4, T5, T6 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6) *Tuple6[T1, T2, T3, T4, T5, T6] {
	return &Tuple6[T1, T2, T3, T4, T5, T6]{
		Val1: v1,
		Val2: v2,
		Val3: v3,
		Val4: v4,
		Val5: v5,
		Val6: v6,
	}
}

var _ Tuple = *new(Tuple6[any, any, any, any, any, any])

type Tuple7[T1, T2, T3, T4, T5, T6, T7 any] struct {
	Val1 T1
	Val2 T2
	Val3 T3
	Val4 T4
	Val5 T5
	Val6 T6
	Val7 T7
}

func (Tuple7[T1, T2, T3, T4, T5, T6, T7]) Size() int { return 7 }

func (t Tuple7[T1, T2, T3, T4, T5, T6, T7]) Values() (T1, T2, T3, T4, T5, T6, T7) {
	return t.Val1, t.Val2, t.Val3, t.Val4, t.Val5, t.Val6, t.Val7
}

func (t Tuple7[T1, T2, T3, T4, T5, T6, T7]) values() []any {
	values := make([]any, 7)
	values[0] = t.Val1
	values[1] = t.Val2
	values[2] = t.Val3
	values[3] = t.Val4
	values[4] = t.Val5
	values[5] = t.Val6
	values[6] = t.Val7

	return values
}

func New7[T1, T2, T3, T4, T5, T6, T7 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7) *Tuple7[T1, T2, T3, T4, T5, T6, T7] {
	return &Tuple7[T1, T2, T3, T4, T5, T6, T7]{
		Val1: v1,
		Val2: v2,
		Val3: v3,
		Val4: v4,
		Val5: v5,
		Val6: v6,
		Val7: v7,
	}
}

var _ Tuple = *new(Tuple7[any, any, any, any, any, any, any])

type Tuple8[T1, T2, T3, T4, T5, T6, T7, T8 any] struct {
	Val1 T1
	Val2 T2
	Val3 T3
	Val4 T4
	Val5 T5
	Val6 T6
	Val7 T7
	Val8 T8
}

func (Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Size() int { return 8 }

func (t Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) Values() (T1, T2, T3, T4, T5, T6, T7, T8) {
	return t.Val1, t.Val2, t.Val3, t.Val4, t.Val5, t.Val6, t.Val7, t.Val8
}

func (t Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]) values() []any {
	values := make([]any, 8)
	values[0] = t.Val1
	values[1] = t.Val2
	values[2] = t.Val3
	values[3] = t.Val4
	values[4] = t.Val5
	values[5] = t.Val6
	values[6] = t.Val7
	values[7] = t.Val8

	return values
}

func New8[T1, T2, T3, T4, T5, T6, T7, T8 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8) *Tuple8[T1, T2, T3, T4, T5, T6, T7, T8] {
	return &Tuple8[T1, T2, T3, T4, T5, T6, T7, T8]{
		Val1: v1,
		Val2: v2,
		Val3: v3,
		Val4: v4,
		Val5: v5,
		Val6: v6,
		Val7: v7,
		Val8: v8,
	}
}

var _ Tuple = *new(Tuple8[any, any, any, any, any, any, any, any])

type Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any] struct {
	Val1 T1
	Val2 T2
	Val3 T3
	Val4 T4
	Val5 T5
	Val6 T6
	Val7 T7
	Val8 T8
	Val9 T9
}

func (Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Size() int { return 9 }

func (t Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) Values() (T1, T2, T3, T4, T5, T6, T7, T8, T9) {
	return t.Val1, t.Val2, t.Val3, t.Val4, t.Val5, t.Val6, t.Val7, t.Val8, t.Val9
}

func (t Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]) values() []any {
	values := make([]any, 9)
	values[0] = t.Val1
	values[1] = t.Val2
	values[2] = t.Val3
	values[3] = t.Val4
	values[4] = t.Val5
	values[5] = t.Val6
	values[6] = t.Val7
	values[7] = t.Val8
	values[8] = t.Val9

	return values
}

func New9[T1, T2, T3, T4, T5, T6, T7, T8, T9 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9) *Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9] {
	return &Tuple9[T1, T2, T3, T4, T5, T6, T7, T8, T9]{
		Val1: v1,
		Val2: v2,
		Val3: v3,
		Val4: v4,
		Val5: v5,
		Val6: v6,
		Val7: v7,
		Val8: v8,
		Val9: v9,
	}
}

var _ Tuple = *new(Tuple9[any, any, any, any, any, any, any, any, any])

type Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any] struct {
	Val1  T1
	Val2  T2
	Val3  T3
	Val4  T4
	Val5  T5
	Val6  T6
	Val7  T7
	Val8  T8
	Val9  T9
	Val10 T10
}

func (Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Size() int { return 10 }

func (t Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) Values() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) {
	return t.Val1, t.Val2, t.Val3, t.Val4, t.Val5, t.Val6, t.Val7, t.Val8, t.Val9, t.Val10
}

func (t Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]) values() []any {
	values := make([]any, 10)
	values[0] = t.Val1
	values[1] = t.Val2
	values[2] = t.Val3
	values[3] = t.Val4
	values[4] = t.Val5
	values[5] = t.Val6
	values[6] = t.Val7
	values[7] = t.Val8
	values[8] = t.Val9
	values[9] = t.Val10

	return values
}

func New10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10) *Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10] {
	return &Tuple10[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10]{
		Val1:  v1,
		Val2:  v2,
		Val3:  v3,
		Val4:  v4,
		Val5:  v5,
		Val6:  v6,
		Val7:  v7,
		Val8:  v8,
		Val9:  v9,
		Val10: v10,
	}
}

var _ Tuple = *new(Tuple10[any, any, any, any, any, any, any, any, any, any])

type Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11 any] struct {
	Val1  T1
	Val2  T2
	Val3  T3
	Val4  T4
	Val5  T5
	Val6  T6
	Val7  T7
	Val8  T8
	Val9  T9
	Val10 T10
	Val11 T11
}

func (Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Size() int { return 11 }

func (t Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) Values() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11) {
	return t.Val1, t.Val2, t.Val3, t.Val4, t.Val5, t.Val6, t.Val7, t.Val8, t.Val9, t.Val10, t.Val11
}

func (t Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]) values() []any {
	values := make([]any, 11)
	values[0] = t.Val1
	values[1] = t.Val2
	values[2] = t.Val3
	values[3] = t.Val4
	values[4] = t.Val5
	values[5] = t.Val6
	values[6] = t.Val7
	values[7] = t.Val8
	values[8] = t.Val9
	values[9] = t.Val10
	values[10] = t.Val11

	return values
}

func New11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11) *Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11] {
	return &Tuple11[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11]{
		Val1:  v1,
		Val2:  v2,
		Val3:  v3,
		Val4:  v4,
		Val5:  v5,
		Val6:  v6,
		Val7:  v7,
		Val8:  v8,
		Val9:  v9,
		Val10: v10,
		Val11: v11,
	}
}

var _ Tuple = *new(Tuple11[any, any, any, any, any, any, any, any, any, any, any])

type Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12 any] struct {
	Val1  T1
	Val2  T2
	Val3  T3
	Val4  T4
	Val5  T5
	Val6  T6
	Val7  T7
	Val8  T8
	Val9  T9
	Val10 T10
	Val11 T11
	Val12 T12
}

func (Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Size() int { return 12 }

func (t Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) Values() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12) {
	return t.Val1, t.Val2, t.Val3, t.Val4, t.Val5, t.Val6, t.Val7, t.Val8, t.Val9, t.Val10, t.Val11, t.Val12
}

func (t Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]) values() []any {
	values := make([]any, 12)
	values[0] = t.Val1
	values[1] = t.Val2
	values[2] = t.Val3
	values[3] = t.Val4
	values[4] = t.Val5
	values[5] = t.Val6
	values[6] = t.Val7
	values[7] = t.Val8
	values[8] = t.Val9
	values[9] = t.Val10
	values[10] = t.Val11
	values[11] = t.Val12

	return values
}

func New12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12) *Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12] {
	return &Tuple12[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12]{
		Val1:  v1,
		Val2:  v2,
		Val3:  v3,
		Val4:  v4,
		Val5:  v5,
		Val6:  v6,
		Val7:  v7,
		Val8:  v8,
		Val9:  v9,
		Val10: v10,
		Val11: v11,
		Val12: v12,
	}
}

var _ Tuple = *new(Tuple12[any, any, any, any, any, any, any, any, any, any, any, any])

type Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13 any] struct {
	Val1  T1
	Val2  T2
	Val3  T3
	Val4  T4
	Val5  T5
	Val6  T6
	Val7  T7
	Val8  T8
	Val9  T9
	Val10 T10
	Val11 T11
	Val12 T12
	Val13 T13
}

func (Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Size() int { return 13 }

func (t Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) Values() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13) {
	return t.Val1, t.Val2, t.Val3, t.Val4, t.Val5, t.Val6, t.Val7, t.Val8, t.Val9, t.Val10, t.Val11, t.Val12, t.Val13
}

func (t Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]) values() []any {
	values := make([]any, 13)
	values[0] = t.Val1
	values[1] = t.Val2
	values[2] = t.Val3
	values[3] = t.Val4
	values[4] = t.Val5
	values[5] = t.Val6
	values[6] = t.Val7
	values[7] = t.Val8
	values[8] = t.Val9
	values[9] = t.Val10
	values[10] = t.Val11
	values[11] = t.Val12
	values[12] = t.Val13

	return values
}

func New13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13) *Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13] {
	return &Tuple13[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13]{
		Val1:  v1,
		Val2:  v2,
		Val3:  v3,
		Val4:  v4,
		Val5:  v5,
		Val6:  v6,
		Val7:  v7,
		Val8:  v8,
		Val9:  v9,
		Val10: v10,
		Val11: v11,
		Val12: v12,
		Val13: v13,
	}
}

var _ Tuple = *new(Tuple13[any, any, any, any, any, any, any, any, any, any, any, any, any])

type Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14 any] struct {
	Val1  T1
	Val2  T2
	Val3  T3
	Val4  T4
	Val5  T5
	Val6  T6
	Val7  T7
	Val8  T8
	Val9  T9
	Val10 T10
	Val11 T11
	Val12 T12
	Val13 T13
	Val14 T14
}

func (Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Size() int { return 14 }

func (t Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) Values() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14) {
	return t.Val1, t.Val2, t.Val3, t.Val4, t.Val5, t.Val6, t.Val7, t.Val8, t.Val9, t.Val10, t.Val11, t.Val12, t.Val13, t.Val14
}

func (t Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]) values() []any {
	values := make([]any, 14)
	values[0] = t.Val1
	values[1] = t.Val2
	values[2] = t.Val3
	values[3] = t.Val4
	values[4] = t.Val5
	values[5] = t.Val6
	values[6] = t.Val7
	values[7] = t.Val8
	values[8] = t.Val9
	values[9] = t.Val10
	values[10] = t.Val11
	values[11] = t.Val12
	values[12] = t.Val13
	values[13] = t.Val14

	return values
}

func New14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14) *Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14] {
	return &Tuple14[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14]{
		Val1:  v1,
		Val2:  v2,
		Val3:  v3,
		Val4:  v4,
		Val5:  v5,
		Val6:  v6,
		Val7:  v7,
		Val8:  v8,
		Val9:  v9,
		Val10: v10,
		Val11: v11,
		Val12: v12,
		Val13: v13,
		Val14: v14,
	}
}

var _ Tuple = *new(Tuple14[any, any, any, any, any, any, any, any, any, any, any, any, any, any])

type Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15 any] struct {
	Val1  T1
	Val2  T2
	Val3  T3
	Val4  T4
	Val5  T5
	Val6  T6
	Val7  T7
	Val8  T8
	Val9  T9
	Val10 T10
	Val11 T11
	Val12 T12
	Val13 T13
	Val14 T14
	Val15 T15
}

func (Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Size() int {
	return 15
}

func (t Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) Values() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15) {
	return t.Val1, t.Val2, t.Val3, t.Val4, t.Val5, t.Val6, t.Val7, t.Val8, t.Val9, t.Val10, t.Val11, t.Val12, t.Val13, t.Val14, t.Val15
}

func (t Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]) values() []any {
	values := make([]any, 15)
	values[0] = t.Val1
	values[1] = t.Val2
	values[2] = t.Val3
	values[3] = t.Val4
	values[4] = t.Val5
	values[5] = t.Val6
	values[6] = t.Val7
	values[7] = t.Val8
	values[8] = t.Val9
	values[9] = t.Val10
	values[10] = t.Val11
	values[11] = t.Val12
	values[12] = t.Val13
	values[13] = t.Val14
	values[14] = t.Val15

	return values
}

func New15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15) *Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15] {
	return &Tuple15[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15]{
		Val1:  v1,
		Val2:  v2,
		Val3:  v3,
		Val4:  v4,
		Val5:  v5,
		Val6:  v6,
		Val7:  v7,
		Val8:  v8,
		Val9:  v9,
		Val10: v10,
		Val11: v11,
		Val12: v12,
		Val13: v13,
		Val14: v14,
		Val15: v15,
	}
}

var _ Tuple = *new(Tuple15[any, any, any, any, any, any, any, any, any, any, any, any, any, any, any])

type Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16 any] struct {
	Val1  T1
	Val2  T2
	Val3  T3
	Val4  T4
	Val5  T5
	Val6  T6
	Val7  T7
	Val8  T8
	Val9  T9
	Val10 T10
	Val11 T11
	Val12 T12
	Val13 T13
	Val14 T14
	Val15 T15
	Val16 T16
}

func (Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Size() int {
	return 16
}

func (t Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) Values() (T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16) {
	return t.Val1, t.Val2, t.Val3, t.Val4, t.Val5, t.Val6, t.Val7, t.Val8, t.Val9, t.Val10, t.Val11, t.Val12, t.Val13, t.Val14, t.Val15, t.Val16
}

func (t Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]) values() []any {
	values := make([]any, 16)
	values[0] = t.Val1
	values[1] = t.Val2
	values[2] = t.Val3
	values[3] = t.Val4
	values[4] = t.Val5
	values[5] = t.Val6
	values[6] = t.Val7
	values[7] = t.Val8
	values[8] = t.Val9
	values[9] = t.Val10
	values[10] = t.Val11
	values[11] = t.Val12
	values[12] = t.Val13
	values[13] = t.Val14
	values[14] = t.Val15
	values[15] = t.Val16

	return values
}

func New16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16 any](v1 T1, v2 T2, v3 T3, v4 T4, v5 T5, v6 T6, v7 T7, v8 T8, v9 T9, v10 T10, v11 T11, v12 T12, v13 T13, v14 T14, v15 T15, v16 T16) *Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16] {
	return &Tuple16[T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, T13, T14, T15, T16]{
		Val1:  v1,
		Val2:  v2,
		Val3:  v3,
		Val4:  v4,
		Val5:  v5,
		Val6:  v6,
		Val7:  v7,
		Val8:  v8,
		Val9:  v9,
		Val10: v10,
		Val11: v11,
		Val12: v12,
		Val13: v13,
		Val14: v14,
		Val15: v15,
		Val16: v16,
	}
}

var _ Tuple = *new(Tuple16[any, any, any, any, any, any, any, any, any, any, any, any, any, any, any, any])
