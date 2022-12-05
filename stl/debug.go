package stl

import "fmt"

func PrintGSI(gsi *GSIBlock) {
	fmt.Printf("--- GSI ---\n")
	fmt.Println("CPN (Code Page Number):", gsi.CPN)
	fmt.Println("DFC (Disk Format Code):", gsi.DFC)
	fmt.Println("DSC (Display Standard Code):", gsi.DSC)
	fmt.Println("CCT (Character Code Table number):", gsi.CCT)
	fmt.Println("LC (Language Code):", gsi.LC)
	fmt.Println("OPT (Original Program Title):", gsi.OPT)
	fmt.Println("OET (Original Episode Title):", gsi.OET)
	fmt.Println("TPT (Translated Program Title):", gsi.TPT)
	fmt.Println("TET (Translated Episode Title):", gsi.TET)
	fmt.Println("TN (Translator Name):", gsi.TN)
	fmt.Println("TCD (Translator Contact Details):", gsi.TCD)
	fmt.Println("SLR (Subtitle List Reference Code):", gsi.SLR)
	fmt.Println("CD (Creation Date):", gsi.CD)
	fmt.Println("RD (Revision Date):", gsi.RD)
	fmt.Println("RN (Revision Number):", gsi.RN)
	fmt.Println("TNB (Total Number of TTI blocks):", gsi.TNB)
	fmt.Println("TNS (Total Number of Subtitles):", gsi.TNS)
	fmt.Println("TNG (Total Number of Subtitle Groups):", gsi.TNG)
	fmt.Println("PNC (Maximum Number of Displayable Characters):", gsi.MNC)
	fmt.Println("MNR (Maximum Number of Displayable Rows):", gsi.MNR)
	fmt.Println("TCS (Time Code: Status):", gsi.TCS)
	fmt.Println("TCP (Time Code: Start-of-Program):", gsi.TCP)
	fmt.Println("TCF (Time Code: First In-Cue):", gsi.TCF)
	fmt.Println("TND (Total Number of Disks):", gsi.TND)
	fmt.Println("DSN (Disk Sequence Number):", gsi.DSN)
	fmt.Println("CO (Country of Origin):", gsi.CO)
	fmt.Println("PUB (Publisher):", gsi.PUB)
	fmt.Println("EN (Editor's Name):", gsi.EN)
	fmt.Println("ECD (Editor's Contact Details):", gsi.ECD)
	fmt.Println("UDA (User-Defined Area):", gsi.UDA)
	fmt.Println("Framerate (additional):", gsi.Framerate())
}

func PrintTTI(tti *TTIBlock, cct CharacterCodeTable) {
	fmt.Printf("--- TTI ---\n")
	fmt.Printf("SGN (Subtitle Group Number): %d\n", tti.SGN)
	fmt.Printf("SN (Subtitle Number): %d\n", tti.SN)
	fmt.Printf("EBN (Extension Block Number): %d\n", tti.EBN)
	fmt.Printf("CS (Cumulative Status): %s\n", tti.CS)
	fmt.Printf("TCI (Time Code In): %v\n", tti.TCI)
	fmt.Printf("TCO (Time Code Out): %v\n", tti.TCO)
	fmt.Printf("VP (Vertical Positioning): %d\n", tti.VP)
	fmt.Printf("JC (Justification Code): %s\n", tti.JC)
	fmt.Printf("CF (Comment Flag): %s\n", tti.CF)
	fmt.Printf("Terminated by space: %t\n", tti.terminatedBySpace)
	fmt.Printf("Text (replace Text Field):\n")
	t, err := tti.Text(cct)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t)
	}
}
