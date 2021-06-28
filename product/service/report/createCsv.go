package report

import (
	"encoding/csv"
	"log"
	"os"
)

func CreateCsv(data [][]string) *os.File{
	result, err := os.Create("report.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvWriter := csv.NewWriter(result)
	_ = csvWriter.WriteAll(data)
	csvWriter.Flush()
	_ = result.Close()
	return result
}
