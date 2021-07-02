package createFile

import (
	"encoding/csv"
	"os"
)

func CreateCSV(filename string, report [][]string){
	file, _ := os.Create(filename)
	csvWriter := csv.NewWriter(file)
	_ = csvWriter.WriteAll(report)
	csvWriter.Flush()
	_ = file.Close()
}


