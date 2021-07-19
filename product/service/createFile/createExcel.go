package createFile

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"log"
	"strconv"
)

func CreateExcel(filename string, report [][]string){
	f := excelize.NewFile()
	style, _ := f.NewStyle(`{"alignment":{"horizontal":"center"}, 
        "font":{"bold":true}}`)
	for a,b := range report{
		for x,y := range b{
			row,col := ChangIndexToAxis(a, x)
			_ = f.SetCellStyle("Sheet1", row+col, row+col, style)
			_ = f.SetCellValue("Sheet1", row+col, y)
		}
	}
	if err := f.SaveAs(filename); err != nil {
		log.Fatal(err)
	}
}


func ChangIndexToAxis (intIndexX int,intIndexY int ) (string, string){
	var arr = [...]string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	intIndexY = intIndexY+1
	resultY := ""
	for true {
		if intIndexY <= 26 {
			resultY = resultY + arr[intIndexY-1]
			break
		}
		mo := intIndexY % 26
		resultY = arr[mo-1] + resultY
		shang := intIndexY/26
		if shang <= 26 {
			resultY = arr[shang-1] + resultY
			break
		}
		intIndexY = shang
	}
	return resultY,strconv.Itoa(intIndexX+1)
}


