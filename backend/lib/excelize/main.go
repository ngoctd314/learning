package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

type allStyles struct {
	center   int
	wrapText int
}

func main() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	sheet := "Chi tiếtttt"
	_, err := f.NewSheet(sheet)
	if err != nil {
		log.Println(err)
		return
	}
	f.DeleteSheet("Sheet1")
	f.SaveAs("test.xlsx")

}

func groupSummary() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	index, err := f.NewSheet("Tổng hợp")
	if err != nil {
		log.Println(err)
		return
	}
	sheet := "Tổng hợp"

	// Set value of a cell.
	f.MergeCell(sheet, "A2", "H2")
	f.SetCellValue(sheet, "A2", "Moshop")
	f.MergeCell(sheet, "A3", "H3")
	f.SetCellValue(sheet, "A3", "From: xxx, To yyy")

	columnNum := rune(8)
	columnName := []string{"Tổng số Tickets", "Urgent", "High", "Medium", "Low", "Số Tickets cảnh báo đúng", "Số Tickets cảnh báo sai", "Vi phạm SLA"}
	columnValue := []any{12, 0, 0, 0, 4, 10, 2, 8}

	var i rune
	for i = 0; i < columnNum; i++ {
		fmt.Println(string('A'+i) + "4")
		f.SetCellValue(sheet, string('A'+i)+"4", columnName[i])
		f.SetCellValue(sheet, string('A'+i)+"5", columnValue[i])
	}

	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func groupDetail() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	index, err := f.NewSheet("Chi tiết")
	if err != nil {
		log.Println(err)
		return
	}
	sheet := "Chi tiết"

	f.MergeCell(sheet, "A1", "L1")
	f.SetCellValue(sheet, "A1", "Chi tiết tình trạng xử lý TM của MOSHOP")
	style, _ := f.NewStyle(&excelize.Style{
		// Font: &excelize.Font{
		// 	Color: "00B946",
		// },
		// Fill: excelize.Fill{
		// 	Type:    "pattern",
		// 	Color:   []string{"00B946"},
		// 	Pattern: 1,
		// },
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
	})
	f.SetCellStyle(sheet, "A1", "A1", style)

	f.MergeCell(sheet, "A2", "F2")
	f.SetCellValue(sheet, "A2", "Thông tin issue")
	f.SetCellStyle(sheet, "A2", "A2", style)

	f.MergeCell(sheet, "G2", "L2")
	f.SetCellValue(sheet, "G2", "Thông tin SLA")
	f.SetCellStyle(sheet, "G2", "G2", style)

	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

type cellData struct {
	data    any
	styleID int
}

type table struct {
	f           *excelize.File
	sheet       string
	topLeft     [2]rune
	bottomRight [2]rune
	data        [][]cellData
}

func newTable(f *excelize.File, sheet string, topLeft, bottomRight [2]rune, data [][]cellData) (*table, error) {
	if bottomRight[1] < topLeft[1] {
		return nil, errors.New("invalid topLeft, bottomRight")
	}

	if bottomRight[1]-topLeft[1] < rune(len(data))-1 {
		return nil, errors.New("table out of bound")
	}

	return &table{
		f:           f,
		sheet:       sheet,
		topLeft:     topLeft,
		bottomRight: bottomRight,
		data:        data,
	}, nil
}

func (t *table) draw() error {
	for _, records := range t.data {
		if !t.hasNext() {
			break
		}
		cells := t.nextCell()
		if len(cells) != len(records) {
			return errors.New("table size is not fit")
		}
		for i := range cells {
			t.f.SetCellValue(t.sheet, cells[i], records[i].data)
			if records[i].styleID != 0 {
				t.f.SetCellStyle(t.sheet, cells[i], cells[i], records[i].styleID)
			}
		}
	}

	return nil
}

func (t *table) nextCell() []string {
	var cells []string
	var i rune
	for i = 0; i <= t.bottomRight[0]-t.topLeft[0]; i++ {
		cell := fmt.Sprintf("%s%d", string(t.topLeft[0]+i), t.topLeft[1])
		cells = append(cells, cell)
	}
	t.topLeft[1]++

	return cells
}

func (t *table) hasNext() bool {
	return t.bottomRight[1] >= t.topLeft[1]
}
