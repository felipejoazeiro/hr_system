package models

type Customer struct {
	ID			int 			`json:"id" db:"id"`
	Name		string			`json:"name" db:"name"`
	CNPJ_CPF	string			`json:"cnpj_cpf" db:"cnpj_cpf"`
	Email		string			`json:"email" db:"email"`
	Phone		string			`json:"phone" db:"phone"`
	IsActive	bool			`json:"is_active" db:"is_active"`
}

type CreateCustomer struct {
	Name		string			`json:"name" db:"name"`
	CNPJ_CPF	string			`json:"cnpj_cpf" db:"cnpj_cpf"`
	Email		*string			`json:"email" db:"email"`
	Phone		*string			`json:"phone" db:"phone"`
	IsActive	*bool			`json:"is_active" db:"is_active"`
}

type UpdateCustomer struct {
	Name		*string			`json:"name" db:"name"`
	CNPJ_CPF	*string			`json:"cnpj_cpf" db:"cnpj_cpf"`
	Email		*string			`json:"email" db:"email"`
	Phone		*string			`json:"phone" db:"phone"`
	IsActive	*bool			`json:"is_active" db:"is_active"`
}

type ArInvoices struct {
	ID				int			`json:"id" db:"id"`
	Number			string		`json:"number" db:"number"`
	IssueDate		time.Time	`json:"issue_date" db:"issue_date"`
	DueTime			time.Time	`json:"due_time" db:"due_time"`
	Total			int			`json:"total" db:"total"`
	Status			string		`json:"status" db:"status"` // 'PENDING', 'PAID', 'OVERDUE'
	Notes			string		`json:"notes" db:"notes"`
	FkCustomerID	int			`json:"fk_customerId" db:"fk_customerId"`	
	FKCostCenterID	int			`json:"fk_costCenterId" db:"fk_costCenterId"`
}

type CreateArInvoices struct {
	Number			string		`json:"number" db:"number"`
	IssueDate		time.Time	`json:"issue_date" db:"issue_date"`
	DueTime			time.Time	`json:"due_time" db:"due_time"`
	Total			int			`json:"total" db:"total"`
	Status			string		`json:"status" db:"status"` // 'PENDING', 'PAID', 'OVERDUE'
	Notes			string		`json:"notes" db:"notes"`
	FkCustomerID	int			`json:"fk_customerId" db:"fk_customerId"`	
	FKCostCenterID	int			`json:"fk_costCenterId" db:"fk_costCenterId"`
}

type UpdateArInvoices struct {
	Number			*string		`json:"number" db:"number"`
	IssueDate		*time.Time	`json:"issue_date" db:"issue_date"`
	DueTime			*time.Time	`json:"due_time" db:"due_time"`
	Total			*int			`json:"total" db:"total"`
	Status			*string		`json:"status" db:"status"` // 'PENDING', 'PAID', 'OVERDUE'
	Notes			*string		`json:"notes" db:"notes"`
	FkCustomerID	*int			`json:"fk_customerId" db:"fk_customerId"`	
	FKCostCenterID	*int			`json:"fk_costCenterId" db:"fk_costCenterId"`
}

type ArInvoiceItems struct {
	ID				int			`json:"id" db:"id"`
	Description		string		`json:"description" db:"description"`
	Quantity		int			`json:"quantity" db:"quantity"`
	UnitPrice		int			`json:"unit_price" db:"unit_price"`
	FkArInvoiceID	int			`json:"fk_arInvoiceId" db:"fk_arInvoiceId"`
	FkAccountID		int			`json:"fk_accountId" db:"fk_accountId"`
	FKCostCenterID	int			`json:"fk_costCenterId" db:"fk_costCenterId"`
}

type CreateArInvoiceItems struct {
	Description		string		`json:"description" db:"description"`
	Quantity		int			`json:"quantity" db:"quantity"`
	UnitPrice		int			`json:"unit_price" db:"unit_price"`
	FkArInvoiceID	int			`json:"fk_arInvoiceId" db:"fk_arInvoiceId"`
	FkAccountID		int			`json:"fk_accountId" db:"fk_accountId"`
	FKCostCenterID	int			`json:"fk_costCenterId" db:"fk_costCenterId"`
}

type UpdateArInvoiceItems struct {
	Description		*string		`json:"description" db:"description"`
	Quantity		*int			`json:"quantity" db:"quantity"`
	UnitPrice		*int			`json:"unit_price" db:"unit_price"`
	FkArInvoiceID	*int			`json:"fk_arInvoiceId" db:"fk_arInvoiceId"`
	FkAccountID		*int			`json:"fk_accountId" db:"fk_accountId"`
	FKCostCenterID	*int			`json:"fk_costCenterId" db:"fk_costCenterId"`
}

type ArReceipts struct {
	ID 					int			`json:"id" db:"id"`
	ReceiptDate			time.Time	`json:"receipt_date" db:"receipt_date"`
	Amount 				int			`json:"amount" db:"amount"`
	Method 				string		`json:"method" db:"method"`
	FkArInvoiceID		int			`json:"fk_arInvoiceId" db:"fk_arInvoiceId"`
	FkBankAccountId		int			`json:"fk_bankAccountId" db:"fk_bankAccountId"`
	FkBankTransactionId	int		`json:"fk_bankTransactionId" db:"fk_bankTransactionId"`
	Notes				string		`json:"notes" db:"notes"`
}

type CreateArReceipts struct {
	ReceiptDate			time.Time	`json:"receipt_date" db:"receipt_date"`
	Amount 				int			`json:"amount" db:"amount"`
	Method 				string		`json:"method" db:"method"`
	FkArInvoiceID		int			`json:"fk_arInvoiceId" db:"fk_arInvoiceId"`
	FkBankAccountId		int			`json:"fk_bankAccountId" db:"fk_bankAccountId"`
	FkBankTransactionId	int			`json:"fk_bankTransactionId" db:"fk_bankTransactionId"`
	Notes				string		`json:"notes" db:"notes"`
}

type UpdateArReceipts struct {
	ReceiptDate			*time.Time	`json:"receipt_date" db:"receipt_date"`
	Amount 				*int			`json:"amount" db:"amount"`
	Method 				*string		`json:"method" db:"method"`
	FkArInvoiceID		*int			`json:"fk_arInvoiceId" db:"fk_arInvoiceId"`
	FkBankAccountId		*int			`json:"fk_bankAccountId" db:"fk_bankAccountId"`
	FkBankTransactionId	*int			`json:"fk_bankTransactionId" db:"fk_bankTransactionId"`
	Notes				*string		`json:"notes" db:"notes"`
}
