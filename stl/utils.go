package stl

// cutPad returns a copy of b with length n.
// If b is longer than n, it is truncated.
// If b is shorter than n, it is right padded with p.
func cutPad(b []byte, n int, p rune) []byte {
	if len(b) > n {
		return b[:n]
	} else if len(b) < n {
		bc := make([]byte, n)
		copy(bc, b)
		for i := len(b); i < n; i++ {
			bc[i] = byte(p)
		}
		return bc
	}
	return b
}

// appendNonNilErrs appends err to errs if err is not nil.
func appendNonNilErrs(errs []error, err ...error) []error {
	for _, e := range err {
		if e != nil {
			errs = append(errs, e)
		}
	}
	return errs
}

// trimRight returns a copy of b with trailing v removed.
func trimRight(b []byte, v byte) []byte {
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] != v {
			return b[:i+1]
		}
	}

	return nil
}
