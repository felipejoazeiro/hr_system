package repository

import "app/internal/finance/models"

type FinanceArRepository interface {
	CreateCustomers(customer models.CreateCustomer) (int, error)
	UpdateCustomers(customer models.UpdateCustomer) (int, error)
	DeleteCustomers(customerId int) (int, error)

	CreateArInvoices(invoice models.CreateArInvoices) (int, error)
	UpdateArInvoices(invoice models.UpdateArInvoices) (int, error)
	DeleteArInvoices(invoiceId int) (int, error)

	CreateArInvoiceItems(item models.CreateArInvoiceItems) (int, error)
	UpdateArInvoiceItems(item models.UpdateArInvoiceItems) (int, error)
	DeleteArInvoiceItems(itemId int) (int, error)

	CreateArReceipts(receipt models.CreateArReceipts) (int, error)
	UpdateArReceipts(receipt models.UpdateArReceipts) (int, error)
	DeleteArReceipts(receiptId int) (int, error)
}