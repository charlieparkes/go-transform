package transform

// Slice applies fn to each element of in and returns a new slice of results.
func Slice[I, O any](in []I, fn func(I) O) []O {
	out := make([]O, len(in))
	for i, v := range in {
		out[i] = fn(v)
	}
	return out
}

// SliceErr applies fn to each element of in and returns a new slice of results.
// The first error from fn stops the loop and is returned with a nil slice.
func SliceErr[I, O any](in []I, fn func(I) (O, error)) ([]O, error) {
	out := make([]O, len(in))
	for i, v := range in {
		o, err := fn(v)
		if err != nil {
			return nil, err
		}
		out[i] = o
	}
	return out, nil
}

// SliceIf applies fn to each element of in and appends the result when the bool
// is true.
func SliceIf[I, O any](in []I, fn func(I) (O, bool)) []O {
	out := make([]O, 0, len(in))
	for _, v := range in {
		o, ok := fn(v)
		if ok {
			out = append(out, o)
		}
	}
	return out
}
