package stl

// cutPad returns a copy of s with length n.
// If s is longer than n, it is truncated.
// If s is shorter than n, it is right padded with c.
func cutPad(b []byte, n int, c rune) []byte {
	if len(b) > n {
		return b[:n]
	} else if len(b) < n {
		bc := make([]byte, n)
		copy(bc, b)
		for i := len(b); i < n; i++ {
			bc[i] = byte(c)
		}
		return bc
	}
	return b
}

// appendErrs appends err to errs if err is not nil.
func appendErrs(errs []error, err ...error) []error {
	for _, e := range err {
		if e != nil {
			errs = append(errs, e)
		}
	}
	return errs
}

func wrapGSIEncodingErrs(err ...error) []error {
	var errs []error
	for _, e := range err {
		if e != nil {
			errs = append(errs, &GSIEncodingErr{error: e})
		}
	}
	return errs
}

func wrapGTIEncodingErrs(block int, err ...error) []error {
	var errs []error
	for _, e := range err {
		if e != nil {
			errs = append(errs, &TTIEncodingErr{error: e, block: block})
		}
	}
	return errs
}
