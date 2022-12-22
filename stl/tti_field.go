package stl

// TTIField represents a TTI block field.
// It extends Field.
type TTIField Field

const (
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
