package stl

import (
	"fmt"
	"io"
)

// File is the representation of a STL file.
// The file comprises one General Subtitle Information (GSI) block and a
// number of Text and Timing Information (TTI) blocks.
type File struct {
	GSI *GSIBlock
	TTI []*TTIBlock
}

// CreateFile returns a new stl.File.
func NewFile() *File {
	return &File{}
}

// Decode reads and decodes the STL file from r.
func (f *File) Decode(r io.Reader) (warns []error, err error) {
	f.GSI = NewGSIBlock()

	gsiWarns, gsiErr := f.GSI.Decode(r)
	if gsiErr != nil {
		return nil, gsiErr
	}
	warns = appendErrs(warns, gsiWarns...)

	var i int = 0
	for {
		tti := NewTTIBlock()
		ttiErr := tti.Decode(r)
		if ttiErr == io.EOF {
			break
		} else if ttiErr != nil {
			if ttiErr, ok := ttiErr.(*TTIEncodingError); ok {
				return warns, setTTIEncodingErrBlock(ttiErr, i)
			}
			panic(fmt.Errorf("unexpected error type: %T", ttiErr))
		}
		f.TTI = append(f.TTI, tti)
		i++
	}

	return
}

// Encode encodes and writes the STL file to w.
func (f *File) Encode(w io.Writer) error {
	if err := f.GSI.Encode(w); err != nil {
		return err
	}
	for i, tti := range f.TTI {
		if err := tti.Encode(w); err != nil {
			if ttiErr, ok := err.(*TTIEncodingError); ok {
				return setTTIEncodingErrBlock(ttiErr, i)
			}
			panic(fmt.Errorf("unexpected error type: %T", err))
		}
	}
	return nil
}

func (f *File) Validate() []error {
	var errs []error
	//todo: validation
	return errs
}
