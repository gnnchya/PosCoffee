package bill

import (
	"fmt"
	"github.com/pondnarawich/gofpdf"
)

func GeneratePdf(filename string, order []string){
	pdf := gofpdf.New("", "", "proud", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 2)
	//pdf.ImageOptions(
	//	"Untitled_Artwork.jpg",
	//	0, 0,
	//	50, 50,
	//	false,
	//	gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
	//	0,
	//	"",
	//)
	fmt.Println("bill information",order)
	for z,i := range order{
		pdf.CellFormat(0, float64(z)*2, i, "", 0, "L", false, 0, "")
		//pdf.Write(float64(z)*2, i)
	}
	pdf.Close()
	_ = pdf.OutputFileAndClose(filename)
}
