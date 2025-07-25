package tuple

type Tuple interface {
	Size() int

	values() []any
}

func Values(t Tuple) []any {
	return t.values()
}
