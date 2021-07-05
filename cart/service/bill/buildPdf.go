package bill

import (
	"github.com/pondnarawich/gofpdf"
)

func GeneratePdf(filename string, order []string){
	pdf := gofpdf.New("p", "pt", "proud", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 2)
	pdf.Image(
		"./service/bill/Untitled_Artwork.jpg",
		15, 10,
		30, 30,
		false,
		"",
		0,
		"",
	)
	pdf.CellFormat(0, 15, "", "", 1, "CM", false, 0, "")
	for _,i := range order{
		pdf.CellFormat(0, 3, i, "", 1, "CM", false, 0, "")
	}
	pdf.Close()
	_ = pdf.OutputFileAndClose(filename)
}
