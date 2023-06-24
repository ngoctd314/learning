package main

import "time"

type InvoiceModel struct {
	InvoiceID      int       `gorm:"column:invoice_id"`
	VendorID       int       `gorm:"column:vendor_id"`
	InvoiceNumber  string    `gorm:"column:invoice_number"`
	InvoiceDate    time.Time `gorm:"column:invoice_date"`
	InvoiceTotal   float64   `gorm:"column:invoice_total"`
	PaymentTotal   float64   `gorm:"column:payment_total"`
	CreditTotal    float64   `gorm:"column:credit_total"`
	TermIds        int       `gorm:"column:term_ids"`
	InvoiceDueDate time.Time `gorm:"column:invoice_due_date"`
	PaymentDate    time.Time `gorm:"column:payment_date"`
}

const InvoiceTableName = "invoices"

func (InvoiceModel) TableName() string {
	return InvoiceTableName
}

type InvoiceCredit struct {
	InvoiceID    int     `gorm:"column:invoice_id"`
	InvoiceTotal float64 `gorm:"column:invoice_total"`
	TotalCredits float64 `gorm:"column:total_credits"`
}
