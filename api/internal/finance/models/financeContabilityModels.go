package models

type CostCenters struct {
	ID   		string 		`json:"id" db:"id"`
	Name 		string 		`json:"name" db:"name"`	
	Code 		string 		`json:"code" db:"code"`
	isActive 	bool 		`json:"is_active" db:"is_active"`
}

type CreateCostCenters struct {
	Name 		string 		`json:"name" db:"name"`	
	Code 		string 		`json:"code" db:"code"`
	isActive 	bool 		`json:"is_active" db:"is_active"`
}

type EditCostCenters struct {
	Name		*string 	`json:"name" db:"name"`
	Code		*string 	`json:"code" db:"code"`
}

type Account struct {
	ID   		string 		`json:"id" db:"id"`
	Name 		string 		`json:"name" db:"name"`
	Code 		string 		`json:"code" db:"code"`
	Type		string		`json:"type" db:"type"`
	isActive 	bool 		`json:"is_active" db:"is_active"`
}

type CreateAccount struct {
	Name		string		`json:"name" db:"name"`
	Code		string		`json:"code" db:"code"`
	Type		string		`json:"type" db:"type"`
	isActive	bool		`json:"is_active" db:"is_active"`
}

type EditAccount struct {
	Name		*string 	`json:"name" db:"name"`
	Code		*string 	`json:"code" db:"code"`
	Type		*string 	`json:"type" db:"type"`
	isActive	*bool		`json:"is_active" db:"is_active"`
}

type BankAccount struct {
	ID   		string 		`json:"id" db:"id"`
	BankName	string		`json:"bank_name" db:"bank_name"`
	AccountNumber	string	`json:"account_number" db:"account_number"`
	Agency 		string		`json:"agency" db:"agency"`
	Currency		string		`json:"currency" db:"currency"`
	isActive 	bool 		`json:"is_active" db:"is_active"`
}

type CreateBankAccount struct {
	BankName		string		`json:"bank_name" db:"bank_name"`
	AccountNumber	string		`json:"account_number" db:"account_number"`
	Agency 		string		`json:"agency" db:"agency"`
	Currency		string		`json:"currency" db:"currency"`
	isActive 	bool 		`json:"is_active" db:"is_active"`
}

type EditBankAccount struct {
	BankName		*string		`json:"bank_name" db:"bank_name"`
	AccountNumber	*string		`json:"account_number" db:"account_number"`
	Agency 			*string		`json:"agency" db:"agency"`
	Currency		*string		`json:"currency" db:"currency"`
	isActive		*bool		`json:"is_active" db:"is_active"`
}

type BankTransaction struct {
	ID              string  `json:"id" db:"id"`
	BankAccountID   string  `json:"bank_account_id" db:"bank_account_id"`
	Amount          float64 `json:"amount" db:"amount"`
	Currency       string  `json:"currency" db:"currency"`
	TransactionDate time.Time `json:"transaction_date" db:"transaction_date"`
	Description     string  `json:"description" db:"description"`
}

type CreateBankTransaction struct {
	BankAccountID   string  `json:"bank_account_id" db:"bank_account_id"`
	Amount          float64 `json:"amount" db:"amount"`
	Currency       string  `json:"currency" db:"currency"`
	TransactionDate time.Time `json:"transaction_date" db:"transaction_date"`
	Description     string  `json:"description" db:"description"`
}

type EditBankTransaction struct {
	BankAccountID   *string  `json:"bank_account_id" db:"bank_account_id"`
	Amount          *float64 `json:"amount" db:"amount"`
	Currency       *string  `json:"currency" db:"currency"`
	TransactionDate *time.Time `json:"transaction_date" db:"transaction_date"`
	Description     *string  `json:"description" db:"description"`
}