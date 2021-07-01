package bill

import (
	"github.com/jung-kurt/gofpdf"
)

func GeneratePdf(filename string, order []string) (*gofpdf.Fpdf,error) {
	pdf := gofpdf.New("P", "mm", "A6", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 11)
	//pdf.ImageOptions(
	//	"Untitled_Artwork.jpg",
	//	35, 10,
	//	30, 30,
	//	false,
	//	gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
	//	0,
	//	"",
	//)
	for _,i := range order{
		pdf.CellFormat(0, 70, i, "0", 0, "CM", false, 0, "")
	}
	pdf.Close()
	return pdf,pdf.OutputFileAndClose(filename)
}
