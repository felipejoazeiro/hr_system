-- Criação das tabelas para o sistema de contratos
-- Baseado nos modelos em api/internal/contracts/models/contractModels.go

-- Tabela: contract_dates
CREATE TABLE contract_dates (
    id SERIAL PRIMARY KEY,
    date_initial TIMESTAMP NOT NULL,
    date_limit TIMESTAMP NOT NULL,
    date_guarantee TIMESTAMP NOT NULL,
    date_proposal TIMESTAMP NOT NULL,
    date_budget TIMESTAMP NOT NULL,
    date_tables TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela: contract_rh_info
CREATE TABLE contract_rh_info (
    id SERIAL PRIMARY KEY,
    hour_limit TIME NOT NULL,
    minutes_limit TIME NOT NULL,
    days_first_exp INTEGER NOT NULL,
    days_second_exp INTEGER NOT NULL,
    data_init TIMESTAMP NOT NULL,
    pay_extra_hour BOOLEAN NOT NULL DEFAULT FALSE,
    manual_stitch BOOLEAN NOT NULL DEFAULT FALSE,
    pays_breakfast BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela: contract_discount
CREATE TABLE contract_discount (
    id SERIAL PRIMARY KEY,
    dics_identifier INTEGER NOT NULL DEFAULT 0,
    disc_service INTEGER NOT NULL DEFAULT 0,
    disc_transport INTEGER NOT NULL DEFAULT 0,
    disc_tranp_employee INTEGER NOT NULL DEFAULT 0,
    disc_labor INTEGER NOT NULL DEFAULT 0,
    disc_material INTEGER NOT NULL DEFAULT 0,
    disc_field INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela: contract_info
CREATE TABLE contract_info (
    id SERIAL PRIMARY KEY,
    construction VARCHAR(255) NOT NULL,
    cap VARCHAR(50) NOT NULL,
    process VARCHAR(255) NOT NULL,
    info_pmco TEXT,
    max_employee INTEGER NOT NULL,
    address VARCHAR(255) NOT NULL,
    nro INTEGER NOT NULL,
    complement VARCHAR(100),
    phone BIGINT NOT NULL,
    state VARCHAR(2) NOT NULL,
    city VARCHAR(100) NOT NULL,
    cep VARCHAR(10) NOT NULL,
    email VARCHAR(255) NOT NULL,
    contact VARCHAR(255) NOT NULL,
    fk_employee INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_contract_info_fk_employee ON contract_info(fk_employee);

-- Tabela: contract_values
CREATE TABLE contract_values (
    id SERIAL PRIMARY KEY,
    acronym_values VARCHAR(50) NOT NULL,
    bdi_service INTEGER NOT NULL DEFAULT 0,
    bdi_material INTEGER NOT NULL DEFAULT 0,
    bid_labor INTEGER NOT NULL DEFAULT 0,
    entry_table BOOLEAN NOT NULL DEFAULT FALSE,
    send_email BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Tabela principal: contracts
CREATE TABLE contracts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(50) NOT NULL UNIQUE,
    research BOOLEAN NOT NULL DEFAULT FALSE,
    uses_cpf BOOLEAN NOT NULL DEFAULT FALSE,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    fk_contract_contact INTEGER,
    fk_contract_info INTEGER NOT NULL,
    fk_contract_dates INTEGER NOT NULL,
    fk_contract_values INTEGER NOT NULL,
    fk_contract_discount INTEGER NOT NULL,
    fk_contract_rh INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    -- Foreign Keys
    CONSTRAINT fk_contracts_info FOREIGN KEY (fk_contract_info) REFERENCES contract_info(id) ON DELETE CASCADE,
    CONSTRAINT fk_contracts_dates FOREIGN KEY (fk_contract_dates) REFERENCES contract_dates(id) ON DELETE CASCADE,
    CONSTRAINT fk_contracts_values FOREIGN KEY (fk_contract_values) REFERENCES contract_values(id) ON DELETE CASCADE,
    CONSTRAINT fk_contracts_discount FOREIGN KEY (fk_contract_discount) REFERENCES contract_discount(id) ON DELETE CASCADE,
    CONSTRAINT fk_contracts_rh FOREIGN KEY (fk_contract_rh) REFERENCES contract_rh_info(id) ON DELETE CASCADE
);

-- Indexes
CREATE INDEX idx_contracts_name ON contracts(name);
CREATE INDEX idx_contracts_code ON contracts(code);
CREATE INDEX idx_contracts_is_active ON contracts(is_active);
CREATE INDEX idx_contracts_fk_contract_info ON contracts(fk_contract_info);
CREATE INDEX idx_contracts_fk_contract_dates ON contracts(fk_contract_dates);
CREATE INDEX idx_contracts_fk_contract_values ON contracts(fk_contract_values);
CREATE INDEX idx_contracts_fk_contract_discount ON contracts(fk_contract_discount);
CREATE INDEX idx_contracts_fk_contract_rh ON contracts(fk_contract_rh);

-- Comentários das tabelas
COMMENT ON TABLE contracts IS 'Tabela principal para armazenar informações dos contratos';
COMMENT ON TABLE contract_dates IS 'Tabela para armazenar as datas relacionadas aos contratos';
COMMENT ON TABLE contract_rh_info IS 'Tabela para armazenar informações de RH dos contratos';
COMMENT ON TABLE contract_discount IS 'Tabela para armazenar informações de desconto dos contratos';
COMMENT ON TABLE contract_info IS 'Tabela para armazenar informações detalhadas dos contratos';
COMMENT ON TABLE contract_values IS 'Tabela para armazenar valores e configurações dos contratos';
