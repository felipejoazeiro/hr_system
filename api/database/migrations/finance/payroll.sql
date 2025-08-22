-- Componentes salariais (provento/desconto)
CREATE TABLE payroll_components (
  id SERIAL PRIMARY KEY,
  code TEXT UNIQUE NOT NULL,              -- ex: SAL_BASE, INSS, VT, VR, FGTS
  name TEXT NOT NULL,
  nature TEXT NOT NULL,                   -- 'EARNING' (provento) | 'DEDUCTION' (desconto)
  fk_accountId INT REFERENCES accounts(id), -- conta contábil para lançamento
  is_active BOOLEAN NOT NULL DEFAULT TRUE
);

-- Histórico salarial por funcionário (mudanças de base)
CREATE TABLE employee_salary_history (
  id SERIAL PRIMARY KEY,
  employee_id INT NOT NULL REFERENCES employees(id),
  base_salary NUMERIC(14,2) NOT NULL,
  valid_from DATE NOT NULL,
  valid_to DATE
);

-- Folhas (competência)
CREATE TABLE payroll_runs (
  id SERIAL PRIMARY KEY,
  period_start DATE NOT NULL,
  period_end DATE NOT NULL,
  status TEXT NOT NULL DEFAULT 'OPEN',    -- OPEN, CLOSED, POSTED
  notes TEXT
);

-- Itens da folha por funcionário
CREATE TABLE payroll_items (
  id SERIAL PRIMARY KEY,
  amount NUMERIC(14,2) NOT NULL,          -- positivo p/ EARNING, negativo p/ DEDUCTION (ou guarde sinal separado)
  fk_payrollRunId INT NOT NULL REFERENCES payroll_runs(id) ON DELETE CASCADE,
  fk_employeeId INT NOT NULL REFERENCES employees(id),
  fk_componentId INT NOT NULL REFERENCES payroll_components(id),
  fk_costCenterId INT REFERENCES cost_centers(id)
);
