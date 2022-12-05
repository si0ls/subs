package stl

import "io"

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
	warns = appendErrs(warns, wrapGSIEncodingErrs(gsiWarns...)...)

	var i int = 0
	for {
		tti := NewTTIBlock()
		ttiErr := tti.Decode(r)
		if ttiErr == io.EOF {
			break
		} else if ttiErr != nil {
			return warns, &TTIEncodingErr{error: ttiErr, block: i}
		}
		f.TTI = append(f.TTI, tti)
		i++
	}

	return
}

// Encode encodes and writes the STL file to w.
func (f *File) Encode(w io.Writer) error {
	if err := f.GSI.Encode(w); err != nil {
		return &GSIEncodingErr{error: err}
	}
	for i, tti := range f.TTI {
		if err := tti.Encode(w); err != nil {
			return &TTIEncodingErr{error: err, block: i}
		}
	}
	return nil
}

func (f *File) Validate() []error {
	var errs []error
	//todo: validation
	return errs
}
