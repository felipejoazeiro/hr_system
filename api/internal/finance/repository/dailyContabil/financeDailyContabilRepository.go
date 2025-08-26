package repository

type FinanceDailyContabilRepository struct {
	db *sql.DB
}

func NewFinanceDailyContabilRepository(db *sql.DB) *FinanceDailyContabilRepository {
	return &FinanceDailyContabilRepository{
		db: db,
	}
}

func (r *FinanceDailyContabilRepository) GetAllJournalEntries() ([]models.JournalEntries, error) {
	var journalEntries []models.JournalEntries
	err := r.db.Select(&journalEntries, "SELECT * FROM journal_entries")
	if err != nil {
		return nil, err
	}
	return journalEntries, nil
}

func (r *FinanceDailyContabilRepository) GetJournalEntryByID(id int) (*models.JournalEntries, error) {
	var journalEntry models.JournalEntries
	err := r.db.Get(&journalEntry, "SELECT * FROM journal_entries WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &journalEntry, nil
}

func (r *FinanceDailyContabilRepository) CreateJournalEntry(input models.CreateJournalEntry) (int, error) {
	query := `INSERT INTO journal_entries (entry_date, description, source, posted, fk_sourceId) VALUES (?, ?, ?, ?, ?)`
	result, err := r.db.Exec(query, input.EntryDate, input.Description, input.Source, input.Posted, input.FkSourceId)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *FinanceDailyContabilRepository) EditJournalEntry(id int, input models.EditJournalEntry) error {
	setParts := []string{}
	args := []interface{}{}
	argIndex := 1

	if input.EntryDate != nil {
		setParts = append(setParts, fmt.Sprintf("entry_date = ?"))
		args = append(args, *input.EntryDate)
		argIndex++
	}
	if input.Description != nil {
		setParts = append(setParts, fmt.Sprintf("description = ?"))
		args = append(args, *input.Description)
		argIndex++
	}
	if input.Source != nil {
		setParts = append(setParts, fmt.Sprintf("source = ?"))
		args = append(args, *input.Source)
		argIndex++
	}
	if input.Posted != nil {
		setParts = append(setParts, fmt.Sprintf("posted = ?"))
		args = append(args, *input.Posted)
		argIndex++
	}
	if input.FkSourceId != nil {
		setParts = append(setParts, fmt.Sprintf("fk_sourceId = ?"))
		args = append(args, *input.FkSourceId)
		argIndex++
	}

	if len(setParts) == 0 {
		return nil
	}

	query := fmt.Sprintf("UPDATE journal_entries SET %s WHERE id = ?", strings.Join(setParts, ", "))
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *FinanceDailyContabilRepository) GetAllJournalLines() ([]models.JournalLines, error) {
	var journalLines []models.JournalLines
	err := r.db.Select(&journalLines, "SELECT * FROM journal_lines")
	if err != nil {
		return nil, err
	}
	return journalLines, nil
}

func (r *FinanceDailyContabilRepository) GetJournalLineByID(id int) (*models.JournalLines, error) {
	var journalLine models.JournalLines
	err := r.db.Get(&journalLine, "SELECT * FROM journal_lines WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &journalLine, nil
}

func (r *FinanceDailyContabilRepository) CreateJournalLine(input models.CreateJournalLine) (int, error) {
	query := `INSERT INTO journal_lines (entry_id, description, amount) VALUES (?, ?, ?)`
	result, err := r.db.Exec(query, input.EntryID, input.Description, input.Amount)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *FinanceDailyContabilRepository) EditJournalLine(id int, input models.EditJournalLine) error {
	setParts := []string{}
	args := []interface{}{}
	argIndex := 1

	if input.Debit != nil {
		setParts = append(setParts, fmt.Sprintf("debit = ?"))
		args = append(args, *input.Debit)
		argIndex++
	}
	if input.Credit != nil {
		setParts = append(setParts, fmt.Sprintf("credit = ?"))
		args = append(args, *input.Credit)
		argIndex++
	}
	if input.FkAccountId != nil {
		setParts = append(setParts, fmt.Sprintf("fk_account_id = ?"))
		args = append(args, *input.FkAccountId)
		argIndex++
	}
	if input.FkCostCenterId != nil {
		setParts = append(setParts, fmt.Sprintf("fk_cost_center_id = ?"))
		args = append(args, *input.FkCostCenterId)
		argIndex++
	}
	if input.FkJournalId != nil {
		setParts = append(setParts, fmt.Sprintf("fk_journal_id = ?"))
		args = append(args, *input.FkJournalId)
		argIndex++
	}

	if len(setParts) == 0 {
		return nil
	}

	query := fmt.Sprintf("UPDATE journal_lines SET %s WHERE id = ?", strings.Join(setParts, ", "))
	args = append(args, id)

	_, err := r.db.Exec(query, args...)
	return err
}
