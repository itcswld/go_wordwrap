package xls

import (
	"log"
	"strings"
	"wordwrap/kit"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func GetXlsData(path string, sheet string) (data []string) {

	f, err := excelize.OpenFile(path)
	if err != nil {
		log.Fatalf("openfile err : %s , path: %s", err, path)
	}
	rows := f.GetRows(sheet)
	var rowData []string
	for _, row := range rows {
		rowData = append(rowData, row...)
	}

	for _, str := range rowData {
		if str != "" {
			t := strings.TrimSpace(str)
			rows := kit.WordWrap([]rune(t), 16)
			data = append(data, rows...)
		}
	}

	return data
}
