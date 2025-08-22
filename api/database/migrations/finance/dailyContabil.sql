-- Lançamento (cabeçalho)
CREATE TABLE journal_entries (
  id SERIAL PRIMARY KEY,
  entry_date DATE NOT NULL,
  description TEXT,
  source TEXT,                 -- 'AP','AR','PAYROLL','MANUAL'
  posted BOOLEAN NOT NULL DEFAULT FALSE,
  fk_sourceId INT           -- id do AP/AR/payroll correspondente (opcional)
);

-- Partidas (linhas débito/crédito)
CREATE TABLE journal_lines (
  id SERIAL PRIMARY KEY,
  debit NUMERIC(14,2) NOT NULL DEFAULT 0,
  credit NUMERIC(14,2) NOT NULL DEFAULT 0,
  fk_accountId INT NOT NULL REFERENCES accounts(id),
  fk_costCenterId INT REFERENCES cost_centers(id),
  fk_journalId INT NOT NULL REFERENCES journal_entries(id) ON DELETE CASCADE
  -- CONSTRAINT chk_debit_credit CHECK (NOT (debit = 0 AND credit = 0))
);
