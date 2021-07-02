package bill

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

func GeneratePdf(filename string, order []string){
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
	fmt.Println("bill information",order)
	for z,i := range order{
		pdf.CellFormat(0, float64(z*20), i, "0", 0, "CM", false, 0, "")
	}
	pdf.Close()
	_ = pdf.OutputFileAndClose(filename)
}
