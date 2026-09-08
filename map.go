package transform

// Map applies fn to each key and value of in and returns a new map of results.
func Map[KI, KO comparable, VI, VO any](in map[KI]VI, fn func(KI, VI) (KO, VO)) map[KO]VO {
	if in == nil {
		return nil
	}
	out := make(map[KO]VO, len(in))
	for k, v := range in {
		ko, vo := fn(k, v)
		out[ko] = vo
	}
	return out
}

// MapErr applies fn to each key and value of in and returns a new map of results.
// The first error from fn stops the loop and is returned with a nil map.
func MapErr[KI, KO comparable, VI, VO any](in map[KI]VI, fn func(KI, VI) (KO, VO, error)) (map[KO]VO, error) {
	if in == nil {
		return nil, nil
	}
	out := make(map[KO]VO, len(in))
	for k, v := range in {
		ko, vo, err := fn(k, v)
		if err != nil {
			return nil, err
		}
		out[ko] = vo
	}
	return out, nil
}

// MapIf applies fn to each key and value of in and includes the result when the
// bool is true.
func MapIf[KI, KO comparable, VI, VO any](in map[KI]VI, fn func(KI, VI) (KO, VO, bool)) map[KO]VO {
	if in == nil {
		return nil
	}
	out := make(map[KO]VO, len(in))
	for k, v := range in {
		ko, vo, ok := fn(k, v)
		if ok {
			out[ko] = vo
		}
	}
	return out
}
