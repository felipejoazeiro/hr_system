package models

type Vendors struct {
	ID        int    		`json:"id" db:"id"`
	Name      string 		`json:"name" db:"name"`
	CNPJ_CPF  string 		`json:"cnpj_cpf" db:"cnpj_cpf"`
	Email     string 		`json:"email" db:"email"`
	Phone     string 		`json:"phone" db:"phone"`
	IsActive  bool   		`json:"is_active" db:"is_active"`
}

type CreateVendors struct {
	Name     string 		`json:"name" db:"name"`
	CNPJ_CPF string 		`json:"cnpj_cpf" db:"cnpj_cpf"`
	Email    string 		`json:"email" db:"email"`
	Phone    string 		`json:"phone" db:"phone"`
}

type EditVendors struct {
	Name		*string 	`json:"name" db:"name"`
	CNPJ_CPF	*string 	`json:"cnpj_cpf" db:"cnpj_cpf"`
	Email		*string 	`json:"email" db:"email"`
	Phone		*string 	`json:"phone" db:"phone"`
	IsActive	*bool  	 	`json:"is_active" db:"is_active"`
}

type ApInvoices struct {
	ID 				int			`json:"id" db:"id"`
	Number 			string		`json:"number" db:"number"`	
	IssueDate 		time.Time	`json:"issue_date" db:"issue_date"`
	DueDate 		time.Time	`json:"due_date" db:"due_date"`
	Total 			int			`json:"total" db:"total"`
	Status			string		`json:"status" db:"status"`
	Notes			string		`json:"notes" db:"notes"`
	FKCostCenterID	int			`json:"fk_CostCenterId" db:"fk_CostCenterId"`
	FkVendorId		int			`json:"fk_VendorId" db:"fk_VendorId"`
}

type CreateApInvoices struct {
	Number 			string		`json:"number" db:"number"`	
	IssueDate 		time.Time	`json:"issue_date" db:"issue_date"`
	DueDate 		time.Time	`json:"due_date" db:"due_date"`
	Total 			int			`json:"total" db:"total"`
	Status			string		`json:"status" db:"status"`
	Notes			string		`json:"notes" db:"notes"`
	FKCostCenterID	int			`json:"fk_CostCenterId" db:"fk_CostCenterId"`
	FkVendorId		int			`json:"fk_VendorId" db:"fk_VendorId"`
}

type EditApInvoices struct {
	Number 			*string		`json:"number" db:"number"`	
	IssueDate 		*time.Time	`json:"issue_date" db:"issue_date"`
	DueDate 		*time.Time	`json:"due_date" db:"due_date"`
	Total 			*int			`json:"total" db:"total"`
	Status			*string		`json:"status" db:"status"`
	Notes			*string		`json:"notes" db:"notes"`
	FKCostCenterID	*int			`json:"fk_CostCenterId" db:"fk_CostCenterId"`
	FkVendorId		*int			`json:"fk_VendorId" db:"fk_VendorId"`
}

type ApInvoiceItems struct {
	ID 				int			`json:"id" db:"id"`
	Description		string		`json:"description" db:"description"`
	Quantity		int			`json:"quantity" db:"quantity"`
	UnitPrice		float64		`json:"unit_price" db:"unit_price"`
	FkAccountID		int			`json:"fk_AccountId" db:"fk_AccountId"`
	FkArInvoiceID 	int			`json:"fk_InvoiceId" db:"fk_InvoiceId"`
	FKCostCenterID 	int			`json:"fk_CostCenterId" db:"fk_CostCenterId"`
}

type CreateApInvoiceItems struct {
	Description		string		`json:"description" db:"description"`
	Quantity		int			`json:"quantity" db:"quantity"`
	UnitPrice		float64		`json:"unit_price" db:"unit_price"`
	FkAccountID		int			`json:"fk_AccountId" db:"fk_AccountId"`
	FkArInvoiceID 	int			`json:"fk_InvoiceId" db:"fk_InvoiceId"`
	FKCostCenterID 	int			`json:"fk_CostCenterId" db:"fk_CostCenterId"`
}

type EditApInvoiceItems struct {
	Description		*string		`json:"description" db:"description"`
	Quantity		*int			`json:"quantity" db:"quantity"`
	UnitPrice		*float64		`json:"unit_price" db:"unit_price"`
	FkAccountID		*int			`json:"fk_AccountId" db:"fk_AccountId"`
	FkArInvoiceID	*int			`json:"fk_InvoiceId" db:"fk_InvoiceId"`
	FKCostCenterID	*int			`json:"fk_CostCenterId" db:"fk_CostCenterId"`
}

type ApPayments struct {
	ID 					int				`json:"id" db:"id"`	
	PaymentDate			time.Time		`json:"payment_date" db:"payment_date"`
	Amount				float64			`json:"amount" db:"amount"`
	Method				string			`json:"method" db:"method"` // 'CASH', 'BANK_TRANSFER', 'CHECK'
	Notes				string			`json:"notes" db:"notes"`	
	FkInvoiceId			int				`json:"fk_invoiceId" db:"fk_invoiceId"`
	FkBankAccountId		int				`json:"fk_bankAccountId" db:"fk_bankAccountId"`
	FkBankTransactionId	int				`json:"fk_bankTransactionId" db:"fk_bankTransactionId"`
}

type CreateApPayments struct {
	PaymentDate			time.Time		`json:"payment_date" db:"payment_date"`
	Amount				float64			`json:"amount" db:"amount"`
	Method				string			`json:"method" db:"method"` // 'CASH', 'BANK_TRANSFER', 'CHECK'
	Notes				string			`json:"notes" db:"notes"`	
	FkInvoiceId			int				`json:"fk_invoiceId" db:"fk_invoiceId"`
	FkBankAccountId		int				`json:"fk_bankAccountId" db:"fk_bankAccountId"`
	FkBankTransactionId	int				`json:"fk_bankTransactionId" db:"fk_bankTransactionId"`
}

type EditApPayments struct {
	PaymentDate			*time.Time		`json:"payment_date" db:"payment_date"`
	Amount				*float64			`json:"amount" db:"amount"`
	Method				*string			`json:"method" db:"method"` // 'CASH', 'BANK_TRANSFER', 'CHECK'
	Notes				*string			`json:"notes" db:"notes"`	
	FkInvoiceId			*int				`json:"fk_invoiceId" db:"fk_invoiceId"`
	FkBankAccountId		*int				`json:"fk_bankAccountId" db:"fk_bankAccountId"`
	FkBankTransactionId	*int				`json:"fk_bankTransactionId" db:"fk_bankTransactionId"`
}
