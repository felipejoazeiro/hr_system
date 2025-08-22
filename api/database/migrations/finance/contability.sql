-- Centros de custo (facilita rateio e relatórios por área)
CREATE TABLE cost_centers (
  id SERIAL PRIMARY KEY,
  code TEXT UNIQUE NOT NULL,
  name TEXT NOT NULL,
  is_active BOOLEAN NOT NULL DEFAULT TRUE
);

-- Contas contábeis (Plano de Contas)
CREATE TABLE accounts (
  id SERIAL PRIMARY KEY,
  code TEXT UNIQUE NOT NULL,         -- ex: 1.1.1.01
  name TEXT NOT NULL,
  type TEXT NOT NULL,                -- 'ASSET','LIABILITY','EQUITY','REVENUE','EXPENSE'
  is_active BOOLEAN NOT NULL DEFAULT TRUE
);

-- Contas bancárias
CREATE TABLE bank_accounts (
  id SERIAL PRIMARY KEY,
  bank_name TEXT NOT NULL,
  account_number TEXT NOT NULL,
  agency TEXT,
  currency TEXT NOT NULL DEFAULT 'BRL',
  is_active BOOLEAN NOT NULL DEFAULT TRUE
);

-- Lançamentos bancários (extrato/conciliação)
CREATE TABLE bank_transactions (
  id SERIAL PRIMARY KEY,
  tx_date DATE NOT NULL,
  description TEXT,
  amount NUMERIC(14,2) NOT NULL,     -- + crédito, - débito (ou vice-versa, escolha um padrão)
  external_ref TEXT,                 -- ID do banco, NSU, etc.
  reconciled BOOLEAN NOT NULL DEFAULT FALSE,
  fk_bankAccountId INT NOT NULL REFERENCES bank_accounts(id)
);
