package randutil

func RandSliceInt64(vals []int64) []int64 {
	for i := len(vals) - 1; i > 0; i-- {
		randNum := RandInt(i + 1)

		vals[i], vals[randNum] = vals[randNum], vals[i]
	}

	return vals
}
