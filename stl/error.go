package stl

import "fmt"

type GSIEncodingErr struct {
	error
}

func (e *GSIEncodingErr) Error() string {
	return e.error.Error()
}

type TTIEncodingErr struct {
	error
	block int
}

func (e *TTIEncodingErr) Error() string {
	return fmt.Sprintf("TTI block %d: %s", e.block, e.error.Error())
}

func (e *TTIEncodingErr) Block() int {
	return e.block
}
