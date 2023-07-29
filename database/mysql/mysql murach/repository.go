package main

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r repository) listingInvoice(ctx context.Context) []InvoiceModel {
	var rs []InvoiceModel

	// r.db.Find(&rs)

	var tmp InvoiceModel
	r.db.First(&tmp)
	r.db.Limit(1).Find(&tmp)

	return rs
}

func (r repository) getInvoiceCredit(ctx context.Context) *InvoiceCredit {
	var rs InvoiceCredit

	r.db.Table(InvoiceModel{}.TableName()).Select("invoice_id, invoice_total, credit_total + payment_total AS total_credits").Where("invoice_id = ?", 89).Scan(&rs)
	// r.db.Take(&rs, 17)

	return &rs
}

func (r repository) getInvoiceBetweenDates(ctx context.Context, start, end time.Time) []InvoiceModel {
	var rs []InvoiceModel

	r.db.
		Table(InvoiceTableName).
		Select("invoice_number, invoice_date, invoice_total").
		Where("invoice_date BETWEEN ? AND ?",
			start.Format("2006-01-02"),
			end.Format("2006-01-02")).
		Scan(&rs)

	return rs
}
