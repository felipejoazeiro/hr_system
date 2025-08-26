package repository

type FinanceDailyContabilRepository interface {
	GetAllJournalEntries() ([]models.JournalEntry, error)
	CreateJournalEntry(entry models.CreateJournalEntry) (int, error)
	UpdateJournalEntry(id int, entry models.EditJournalEntry) (int, error)
	DeactivateJournalEntry(id int) error
	ReactivateJournalEntry(id int) error

	GetAllJournalLines() ([]models.JournalLine, error)
	CreateJournalLine(line models.CreateJournalLine) (int, error)
	UpdateJournalLine(id int, line models.EditJournalLine) (int, error)
	DeactivateJournalLine(id int) error
	ReactivateJournalLine(id int) error
}