package stl

import "fmt"

// FieldError is an error that occurred on a field.
type FieldError struct {
	error
	field Field
}

// Error returns the error message.
func (e *FieldError) Error() string {
	return fmt.Sprintf("%s: %s", e.field, e.error.Error())
}

// Unwrap returns the underlying error.
func (e *FieldError) Unwrap() error {
	return e.error
}

// Field returns the concerned field.
func (e *FieldError) Field() Field {
	return e.field
}

// Filed represents a block field.
type Field string

// GSIField represents a GSI block field.
// It extends Field.
type GSIField Field

// TTIField represents a TTI block field.
// It extends Field.
type TTIField Field

const (
	FieldUnknown Field = "<unknown>" // Unknown field

	GSIFieldCPN GSIField = "CPN" // Code Page Number
	GSIFieldDFC GSIField = "DSC" // Disk Format Code
	GSIFieldDSC GSIField = "DSC" // Display Standard Code
	GSIFieldCCT GSIField = "CCT" // Character Code Table number
	GSIFieldLC  GSIField = "LC"  // Language Code
	GSIFieldOPT GSIField = "OPT" // Original Program Title
	GSIFieldOET GSIField = "OET" // Original Episode Title
	GSIFieldTPT GSIField = "TPT" // Translated Program Title
	GSIFieldTET GSIField = "TET" // Translated Episode Title
	GSIFieldTN  GSIField = "TN"  // Translator's Name
	GSIFieldTCD GSIField = "TCD" // Translator's Contact Details
	GSIFieldSLR GSIField = "SLR" // Subtitle List Reference Code
	GSIFieldCD  GSIField = "CD"  // Creation Date
	GSIFieldRD  GSIField = "RD"  // Revision Date
	GSIFieldRN  GSIField = "RN"  // Revision Number
	GSIFieldTNB GSIField = "TNB" // Total Number of Text and Timing Information (TTI) blocks
	GSIFieldTNS GSIField = "TNS" // Total Number of Subtitles
	GSIFieldTNG GSIField = "TNG" // Total Number of Subtitle Groups
	GSIFieldMNC GSIField = "MNC" // Maximum Number of Displayable Characters in any text row
	GSIFieldMNR GSIField = "MNR" // Maximum Number of Displayable Rows
	GSIFieldTCS GSIField = "TCS" // Time Code: Status
	GSIFieldTCP GSIField = "TCP" // Time Code: Start-of-Program
	GSIFieldTCF GSIField = "TCF" // Time Code: First In-Cue
	GSIFieldTND GSIField = "TND" // Total Number of Disks
	GSIFieldDSN GSIField = "DSN" // Disk Sequence Number
	GSIFieldCO  GSIField = "CO"  // Country of Origin
	GSIFieldPUB GSIField = "PUB" // Publisher
	GSIFieldEN  GSIField = "EN"  // Editor's Name
	GSIFieldECD GSIField = "ECD" // Editor's Contact
	GSIFieldUDA GSIField = "UDA" // User-Defined Area

	TTIFieldSGN TTIField = "SGN" // Subtitle Group Number
	TTIFieldSN  TTIField = "SN"  // Subtitle Number
	TTIFieldEBN TTIField = "EBN" // Extension Block Number
	TTIFieldCS  TTIField = "CS"  // Cumulative Status
	TTIFieldTCI TTIField = "TCI" // Time Code In
	TTIFieldTCO TTIField = "TCO" // Time Code Out
	TTIFieldVP  TTIField = "VP"  // Vertical Position
	TTIFieldJC  TTIField = "JC"  // Justification Code
	TTIFieldCF  TTIField = "CF"  // Comment Flag
	TTIFieldTF  TTIField = "TF"  // Text Field
)
