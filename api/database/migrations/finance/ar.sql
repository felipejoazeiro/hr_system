-- Clientes
CREATE TABLE customers (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  cnpj_cpf TEXT,
  email TEXT,
  phone TEXT,
  is_active BOOLEAN NOT NULL DEFAULT TRUE
);

-- Faturas a receber
CREATE TABLE ar_invoices (
  id SERIAL PRIMARY KEY,
  number TEXT,
  issue_date DATE NOT NULL,
  due_date DATE NOT NULL,
  total NUMERIC(14,2) NOT NULL,
  status TEXT NOT NULL DEFAULT 'OPEN', -- OPEN, PARTIALLY_RECEIVED, RECEIVED, CANCELED
  fk_customerId INT NOT NULL REFERENCES customers(id),
  fk_costCenterId INT REFERENCES cost_centers(id),
  notes TEXT
);

-- Itens da fatura (receita por conta)
CREATE TABLE ar_invoice_items (
  id SERIAL PRIMARY KEY,
  description TEXT NOT NULL,
  quantity NUMERIC(14,2) NOT NULL DEFAULT 1,
  unit_price NUMERIC(14,2) NOT NULL DEFAULT 0,
  fk_arInvoiceId INT NOT NULL REFERENCES ar_invoices(id) ON DELETE CASCADE,
  fk_accountId INT NOT NULL REFERENCES accounts(id), -- receita
  fk_costCenterId INT REFERENCES cost_centers(id)
);

-- Recebimentos (podem ser parciais)
CREATE TABLE ar_receipts (
  id SERIAL PRIMARY KEY,
  receipt_date DATE NOT NULL,
  amount NUMERIC(14,2) NOT NULL,
  method TEXT NOT NULL,                 -- PIX, BOLETO, CARTAO, etc.
  fk_arInvoiceId INT NOT NULL REFERENCES ar_invoices(id),
  fk_bankAccountId INT REFERENCES bank_accounts(id),
  fk_bankTransactionId INT REFERENCES bank_transactions(id),
  notes TEXT
);
