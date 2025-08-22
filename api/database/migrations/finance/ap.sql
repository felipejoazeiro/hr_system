-- Fornecedores
CREATE TABLE vendors (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  cnpj_cpf TEXT,
  email TEXT,
  phone TEXT,
  is_active BOOLEAN NOT NULL DEFAULT TRUE
);

-- Faturas/boletos a pagar (documentos)
CREATE TABLE ap_invoices (
  id SERIAL PRIMARY KEY,
  number TEXT,                        -- número do doc
  issue_date DATE NOT NULL,
  due_date DATE NOT NULL,
  total NUMERIC(14,2) NOT NULL,
  status TEXT NOT NULL DEFAULT 'OPEN', -- OPEN, PARTIALLY_PAID, PAID, CANCELED
  notes TEXT,
  fk_costCenterId INT REFERENCES cost_centers(id),
  fk_vendorId INT NOT NULL REFERENCES vendors(id)
);

-- Itens da fatura (rateio por conta contábil)
CREATE TABLE ap_invoice_items (
  id SERIAL PRIMARY KEY,
  description TEXT NOT NULL,
  quantity NUMERIC(14,2) NOT NULL DEFAULT 1,
  unit_price NUMERIC(14,2) NOT NULL DEFAULT 0,
  fk_accountId INT NOT NULL REFERENCES accounts(id), -- despesa/custo
  fk_invoiceId INT NOT NULL REFERENCES ap_invoices(id) ON DELETE CASCADE,
  fk_costCenterId INT REFERENCES cost_centers(id)
);

-- Pagamentos realizados contra AP (podem quitar parcial)
CREATE TABLE ap_payments (
  id SERIAL PRIMARY KEY,
  payment_date DATE NOT NULL,
  amount NUMERIC(14,2) NOT NULL,
  method TEXT NOT NULL,                -- PIX, TED, BOLETO, CARTAO, etc.
  fk_invoiceId INT NOT NULL REFERENCES ap_invoices(id),
  fk_bankAccountId INT REFERENCES bank_accounts(id),
  fk_bankTransactionId INT REFERENCES bank_transactions(id), -- opcional para conciliação
  notes TEXT
);
