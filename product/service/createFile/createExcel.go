package createFile

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"log"
	"strconv"
)

func CreateExcel(filename string, report [][]string){
	alphabetArray := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	f := excelize.NewFile()
	style, _ := f.NewStyle(`{"alignment":{"horizontal":"center"}, 
        "font":{"bold":true,"italic":true}}`)
	_ = f.SetColWidth("Sheet1", alphabetArray[0], alphabetArray[len(report[0])], 20)
	for a,b := range report{
		for x,y := range b{
			_ = f.SetCellStyle("Sheet1", alphabetArray[a] + strconv.Itoa(x), alphabetArray[a] + strconv.Itoa(x), style)
			_ = f.SetCellValue("Sheet1", alphabetArray[a] + strconv.Itoa(x), y)
		}
	}
	if err := f.SaveAs(filename); err != nil {
		log.Fatal(err)
	}
}

