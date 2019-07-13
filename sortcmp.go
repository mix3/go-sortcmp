package sortcmp

type CompareFunc func() int

type compare struct {
	funcs []CompareFunc
}

func Compare(funcs ...CompareFunc) bool {
	for _, f := range funcs {
		if f() == 0 {
			continue
		}
		return 0 < f()
	}
	return false
}
