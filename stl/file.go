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
	warns = appendNonNilErrs(warns, gsiWarns...)

	var i int = 0
	for {
		tti := NewTTIBlock()
		err := tti.Decode(r)
		if err == io.EOF {
			break
		} else if err != nil {
			if ttiErr, ok := err.(*TTIError); ok {
				ttiErr.setBlockNumber(i)
				return warns, ttiErr
			}
			panic(fmt.Errorf("unexpected error type: %T", err))
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
			if ttiErr, ok := err.(*TTIError); ok {
				ttiErr.setBlockNumber(i)
			}
			panic(fmt.Errorf("unexpected error type: %T", err))
		}
	}
	return nil
}

// Validate validates GSI block.
// It returns a slice of warnings and an error if any.
// Warnings are returned for each field that is invalid, warnings can flagged
// as fatal if they are considered to be fatal to further file processing.
// An error is returned if a field is invalid and prevents validation of
// further fields.
func (f *File) Validate() ([]error, error) {
	var errs []error

	if f.GSI == nil {
		panic(fmt.Errorf("GSI block is nil"))
	}

	gsiWarns, gsiErr := f.GSI.Validate()
	if gsiErr != nil {
		return nil, gsiErr
	}

	errs = appendNonNilErrs(errs, gsiWarns...)

	if len(f.TTI) == 0 {
		return errs, validateErr(ErrNoTTIBlocks, nil, true)
	}

	for i, tti := range f.TTI {
		if tti == nil {
			panic(fmt.Errorf("TTI block %d is nil", i))
		}

		ttiErrs := tti.Validate(f.GSI.Framerate(), f.GSI.DSC, f.GSI.MNR)
		setTTIErrsBlockNumber(ttiErrs, i)
		errs = appendNonNilErrs(errs, ttiErrs...)
	}

	return errs, nil
}
