package sequence

import "math"

func Chunks[T any](xs []T, n int) [][]T {
	if len(xs) == 0 {
		return nil
	}
	if n <= 0 {
		return [][]T{xs}
	}
	nchunks := math.Ceil(float64(len(xs)) / float64(n))
	ret := make([][]T, 0, int(nchunks))
	for i := 0; i < len(xs); i += n {
		var subxs []T
		if i+n <= len(xs) {
			subxs = xs[i : i+n]
		} else {
			subxs = xs[i:]
		}
		ret = append(ret, subxs)
	}
	return ret
}
