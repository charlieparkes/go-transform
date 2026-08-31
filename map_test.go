package transform

import (
	"errors"
	"maps"
	"strconv"
	"testing"

	"github.com/charlieparkes/go-testsize"
)

func TestMap(t *testing.T) {
	t.Parallel()
	testsize.Small(t)

	got := Map(map[string]int{"a": 1, "b": 2, "c": 3}, func(k string, n int) (string, string) {
		return k, strconv.Itoa(n)
	})
	want := map[string]string{"a": "1", "b": "2", "c": "3"}
	if !maps.Equal(got, want) {
		t.Fatalf("Map() = %v, want %v", got, want)
	}
}

func TestMapNewKey(t *testing.T) {
	t.Parallel()
	testsize.Small(t)

	got := Map(map[string]int{"a": 1, "b": 2, "c": 3}, func(k string, n int) (int, string) {
		return n, k
	})
	want := map[int]string{1: "a", 2: "b", 3: "c"}
	if !maps.Equal(got, want) {
		t.Fatalf("Map() = %v, want %v", got, want)
	}
}

func TestMapEmpty(t *testing.T) {
	t.Parallel()
	testsize.Small(t)

	got := Map(map[string]int{}, func(k string, n int) (string, int) {
		return k, n * 2
	})
	if got == nil || len(got) != 0 {
		t.Fatalf("Map() = %v, want empty map", got)
	}
}

func TestMapErr(t *testing.T) {
	t.Parallel()
	testsize.Small(t)

	got, err := MapErr(map[string]int{"a": 1, "b": 2, "c": 3}, func(k string, n int) (string, string, error) {
		return k, strconv.Itoa(n), nil
	})
	if err != nil {
		t.Fatalf("MapErr() error = %v", err)
	}
	want := map[string]string{"a": "1", "b": "2", "c": "3"}
	if !maps.Equal(got, want) {
		t.Fatalf("MapErr() = %v, want %v", got, want)
	}
}

func TestMapErrNewKey(t *testing.T) {
	t.Parallel()
	testsize.Small(t)

	got, err := MapErr(map[string]int{"a": 1, "b": 2, "c": 3}, func(k string, n int) (int, string, error) {
		return n, k, nil
	})
	if err != nil {
		t.Fatalf("MapErr() error = %v", err)
	}
	want := map[int]string{1: "a", 2: "b", 3: "c"}
	if !maps.Equal(got, want) {
		t.Fatalf("MapErr() = %v, want %v", got, want)
	}
}

func TestMapErrEmpty(t *testing.T) {
	t.Parallel()
	testsize.Small(t)

	got, err := MapErr(map[string]int{}, func(k string, n int) (string, int, error) {
		return k, n * 2, nil
	})
	if err != nil {
		t.Fatalf("MapErr() error = %v", err)
	}
	if got == nil || len(got) != 0 {
		t.Fatalf("MapErr() = %v, want empty map", got)
	}
}

func TestMapErrStopsOnError(t *testing.T) {
	t.Parallel()
	testsize.Small(t)

	wantErr := errors.New("boom")
	got, err := MapErr(map[string]int{"a": 1, "b": 2, "c": 3}, func(k string, n int) (string, int, error) {
		if n == 2 {
			return "", 0, wantErr
		}
		return k, n, nil
	})
	if !errors.Is(err, wantErr) {
		t.Fatalf("MapErr() error = %v, want %v", err, wantErr)
	}
	if got != nil {
		t.Fatalf("MapErr() = %v, want nil", got)
	}
}

func TestMapIf(t *testing.T) {
	t.Parallel()
	testsize.Small(t)

	got := MapIf(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}, func(k string, n int) (string, string, bool) {
		return k, strconv.Itoa(n), n%2 == 0
	})
	want := map[string]string{"b": "2", "d": "4"}
	if !maps.Equal(got, want) {
		t.Fatalf("MapIf() = %v, want %v", got, want)
	}
}

func TestMapIfNewKey(t *testing.T) {
	t.Parallel()
	testsize.Small(t)

	got := MapIf(map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}, func(k string, n int) (int, string, bool) {
		return n, k, n%2 == 0
	})
	want := map[int]string{2: "b", 4: "d"}
	if !maps.Equal(got, want) {
		t.Fatalf("MapIf() = %v, want %v", got, want)
	}
}

func TestMapIfEmpty(t *testing.T) {
	t.Parallel()
	testsize.Small(t)

	got := MapIf(map[string]int{}, func(k string, n int) (string, int, bool) {
		return k, n, true
	})
	if got == nil || len(got) != 0 {
		t.Fatalf("MapIf() = %v, want empty map", got)
	}
}

func TestMapIfNoneMatch(t *testing.T) {
	t.Parallel()
	testsize.Small(t)

	got := MapIf(map[string]int{"a": 1, "c": 3, "e": 5}, func(k string, n int) (string, int, bool) {
		return k, n, n%2 == 0
	})
	if got == nil || len(got) != 0 {
		t.Fatalf("MapIf() = %v, want empty map", got)
	}
}
