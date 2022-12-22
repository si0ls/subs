package stl

// GSIField represents a GSI block field.
// It extends Field.
type GSIField Field

const (
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
)
