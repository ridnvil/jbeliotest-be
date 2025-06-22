package excel

import (
	"github.com/xuri/excelize/v2"
)

type ExcelReader interface {
	ReadExcel(filepath string) ([][]string, error)
}

type ExcelReaderImpl struct{}

func (ExcelReaderImpl) ReadExcel(filepath string) ([][]string, error) {
	file, err := excelize.OpenFile(filepath)
	if err != nil {
		return nil, err
	}

	rows, err := file.GetRows("Tableau Superstore")

	var newRows [][]string
	for i, row := range rows {
		if i == 0 {
			continue
		} else {
			newRows = append(newRows, row)
		}
	}

	if err != nil {
		return nil, err
	}

	return newRows, nil
}
