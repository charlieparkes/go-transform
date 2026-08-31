package transform

import (
	"errors"
	"slices"
	"strconv"
	"testing"

	"github.com/charlieparkes/go-testsize"
)

func TestSlice(t *testing.T) {
	t.Parallel()
	testsize.Small(t)

	got := Slice([]int{1, 2, 3}, strconv.Itoa)
	want := []string{"1", "2", "3"}
	if !slices.Equal(got, want) {
		t.Fatalf("Slice() = %v, want %v", got, want)
	}
}

func TestSliceEmpty(t *testing.T) {
	t.Parallel()
	testsize.Small(t)

	got := Slice([]int{}, func(n int) int { return n * 2 })
	if got == nil || len(got) != 0 {
		t.Fatalf("Slice() = %v, want empty slice", got)
	}
}

func TestSliceErr(t *testing.T) {
	t.Parallel()
	testsize.Small(t)

	got, err := SliceErr([]int{1, 2, 3}, func(n int) (string, error) {
		return strconv.Itoa(n), nil
	})
	if err != nil {
		t.Fatalf("SliceErr() error = %v", err)
	}
	want := []string{"1", "2", "3"}
	if !slices.Equal(got, want) {
		t.Fatalf("SliceErr() = %v, want %v", got, want)
	}
}

func TestSliceErrEmpty(t *testing.T) {
	t.Parallel()
	testsize.Small(t)

	got, err := SliceErr([]int{}, func(n int) (int, error) {
		return n * 2, nil
	})
	if err != nil {
		t.Fatalf("SliceErr() error = %v", err)
	}
	if got == nil || len(got) != 0 {
		t.Fatalf("SliceErr() = %v, want empty slice", got)
	}
}

func TestSliceErrStopsOnError(t *testing.T) {
	t.Parallel()
	testsize.Small(t)

	wantErr := errors.New("boom")
	got, err := SliceErr([]int{1, 2, 3}, func(n int) (int, error) {
		if n == 2 {
			return 0, wantErr
		}
		return n, nil
	})
	if !errors.Is(err, wantErr) {
		t.Fatalf("SliceErr() error = %v, want %v", err, wantErr)
	}
	if got != nil {
		t.Fatalf("SliceErr() = %v, want nil", got)
	}
}

func TestSliceIf(t *testing.T) {
	t.Parallel()
	testsize.Small(t)

	got := SliceIf([]int{1, 2, 3, 4}, func(n int) (string, bool) {
		return strconv.Itoa(n), n%2 == 0
	})
	want := []string{"2", "4"}
	if !slices.Equal(got, want) {
		t.Fatalf("SliceIf() = %v, want %v", got, want)
	}
}

func TestSliceIfEmpty(t *testing.T) {
	t.Parallel()
	testsize.Small(t)

	got := SliceIf([]int{}, func(n int) (int, bool) {
		return n, true
	})
	if got == nil || len(got) != 0 {
		t.Fatalf("SliceIf() = %v, want empty slice", got)
	}
}

func TestSliceIfNoneMatch(t *testing.T) {
	t.Parallel()
	testsize.Small(t)

	got := SliceIf([]int{1, 3, 5}, func(n int) (int, bool) {
		return n, n%2 == 0
	})
	if got == nil || len(got) != 0 {
		t.Fatalf("SliceIf() = %v, want empty slice", got)
	}
}
