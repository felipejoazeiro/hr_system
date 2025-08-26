package models

import "time"

type JournalEntries struct {
	ID          int       `json:"id" db:"id"`
	EntryDate	time.Time  `json:"entry_date" db:"entry_date"`
	Description string    `json:"description" db:"description"`
	Source      string    `json:"source" db:"source"`
	Posted 		bool	  `json:"posted" db:"posted"`
	FkSourceId  int       `json:"fk_sourceId" db:"fk_sourceId"`
}

type CreateJournalEntry struct {
	EntryDate   time.Time `json:"entry_date" db:"entry_date"`
	Description string    `json:"description" db:"description"`
	Source      string    `json:"source" db:"source"`
	Posted      bool      `json:"posted" db:"posted"`
	FkSourceId  int       `json:"fk_sourceId" db:"fk_sourceId"`
}

type EditJournalEntry struct {
	EntryDate   *time.Time `json:"entry_date" db:"entry_date"`
	Description *string    `json:"description" db:"description"`
	Source      *string    `json:"source" db:"source"`
	Posted      *bool      `json:"posted" db:"posted"`
	FkSourceId  *int       `json:"fk_sourceId" db:"fk_sourceId"`
}

type JournalLines struct {
	ID          		int       `json:"id" db:"id"`
	Debit      			float64   `json:"debit" db:"debit"`
	Credit     			float64   `json:"credit" db:"credit"`
	FkAccountId 		int       `json:"fk_accountId" db:"fk_accountId"`
	FkCostCenterId		int       `json:"fk_costCenterId" db:"fk_costCenterId"`
	FkJournalId 		int       `json:"fk_journalId" db:"fk_journalId"`
}

type CreateJournalLine struct {
	Debit       float64   `json:"debit" db:"debit"`
	Credit      float64   `json:"credit" db:"credit"`
	FkAccountId int       `json:"fk_accountId" db:"fk_accountId"`
	FkCostCenterId int    `json:"fk_costCenterId" db:"fk_costCenterId"`
	FkJournalId int       `json:"fk_journalId" db:"fk_journalId"`
}

type EditJournalLine struct {
	Debit       	*float64   	`json:"debit" db:"debit"`
	Credit      	*float64   	`json:"credit" db:"credit"`
	FkAccountId 	*int       	`json:"fk_accountId" db:"fk_accountId"`
	FkCostCenterId 	*int    	`json:"fk_costCenterId" db:"fk_costCenterId"`
	FkJournalId 	*int       	`json:"fk_journalId" db:"fk_journalId"`
}