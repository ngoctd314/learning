package main

import (
	"log"

	"github.com/xuri/excelize/v2"
)

type style struct {
	excelizeStyle *excelize.Style
}

func newStyle(opts ...styleOpt) *style {
	result := &style{
		excelizeStyle: &excelize.Style{},
	}

	for _, opt := range opts {
		opt(result.excelizeStyle)
	}

	return result
}

func (s *style) id(f *excelize.File) int {
	id, err := f.NewStyle(s.excelizeStyle)
	if err != nil {
		log.Println(err)
		return -1
	}
	return id
}

type styleOpt func(*excelize.Style)

func withCenter() styleOpt {
	return func(s *excelize.Style) {
		if s.Alignment == nil {
			s.Alignment = &excelize.Alignment{}
		}
		s.Alignment.Horizontal = "center"
	}
}

func withWrapText() styleOpt {
	return func(s *excelize.Style) {
		if s.Alignment == nil {
			s.Alignment = &excelize.Alignment{}
		}
		s.Alignment.WrapText = true
	}
}
