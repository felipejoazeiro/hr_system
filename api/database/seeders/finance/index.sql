CREATE INDEX idx_ap_invoices_vendor ON ap_invoices(fk_vendorId, status, due_date);
CREATE INDEX idx_ar_invoices_customer ON ar_invoices(fk_customerId, status, due_date);
CREATE INDEX idx_payroll_items_run_emp ON payroll_items(fk_payrollRunId, fk_employeeId);
CREATE INDEX idx_bank_tx_account_date ON bank_transactions(fk_bankAccountId, tx_date);
CREATE INDEX idx_bank_transactions_reconciled ON bank_transactions(reconciled);