package bill

import (
	"github.com/jung-kurt/gofpdf"
)

func GeneratePdf(filename string, order []string) ( *gofpdf.Fpdf,error) {
	pdf := gofpdf.New("P", "mm", "A7", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.ImageOptions(
		"avatar.jpg",
		80, 20,
		0, 0,
		false,
		gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
		0,
		"",
	)
	for _,i := range order{
		pdf.CellFormat(190, 7, i, "0", 0, "CM", false, 0, "")
	}
	pdf.Close()
	return pdf,pdf.OutputFileAndClose(filename)
}
