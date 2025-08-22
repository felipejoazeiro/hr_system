package repository

import "app/internal/finance/models"

type FinanceApRepository interface {
	GetVendors(id int) (models.Vendors, error)
	GetAllVendors() ([]models.Vendors, error)
	CreateVendors(vendor models.CreateVendors) (int, error)
	EditVendors(id int, vendor models.EditVendors) error
	DeactivateVendor(id int) error
	ReactivateVendor(id int) error

	GetApInvoices(id int) (models.ApInvoices, error)
	GetAllApInvoices() ([]models.ApInvoices, error)
	CreateApInvoices(invoice models.CreateApInvoices) (int, error)
	EditApInvoices(id int, invoice models.EditApInvoices) error
	DeactivateApInvoice(id int) error
	ReactivateApInvoice(id int) error

	GetApInvoiceItems(invoiceID int) ([]models.ApInvoiceItems, error)
	GetAllApInvoiceItems() ([]models.ApInvoiceItems, error)
	CreateApInvoiceItems(items []models.CreateApInvoiceItems) ([]int, error)
	EditApInvoiceItems(id int, item models.EditApInvoiceItems) error
	DeactivateApInvoiceItems(id int) error
	ReactivateApInvoiceItems(id int) error

	GetApPayments(id int) (models.ApPayments, error)
	GetAllApPayments() ([]models.ApPayments, error)
	CreateApPayments(payment models.CreateApPayments) (int, error)
	EditApPayments(id int, payment models.EditApPayments) error
	DeactivateApPayments(id int) error
	ReactivateApPayments(id int) error
}