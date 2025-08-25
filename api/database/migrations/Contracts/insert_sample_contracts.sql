-- Inserção de 10 contratos de exemplo
-- Script para popular as tabelas de contratos com dados de teste

-- Inserir dados na tabela contract_dates
INSERT INTO contract_dates (date_initial, date_limit, date_guarantee, date_proposal, date_budget, date_tables) VALUES
('2024-01-15 08:00:00', '2024-12-31 18:00:00', '2025-06-30 23:59:59', '2023-11-20 14:30:00', '2023-12-05 16:45:00', '2023-12-10 10:00:00'),
('2024-02-01 09:00:00', '2025-01-31 17:00:00', '2025-07-31 23:59:59', '2023-12-10 15:20:00', '2023-12-20 11:30:00', '2024-01-05 14:15:00'),
('2024-03-10 07:30:00', '2024-11-30 18:30:00', '2025-05-30 23:59:59', '2024-01-05 13:45:00', '2024-01-15 09:20:00', '2024-02-01 16:00:00'),
('2024-04-01 08:15:00', '2025-03-31 17:45:00', '2025-09-30 23:59:59', '2024-02-15 12:30:00', '2024-02-28 14:50:00', '2024-03-15 11:30:00'),
('2024-05-15 09:30:00', '2024-12-15 16:30:00', '2025-06-15 23:59:59', '2024-03-20 16:15:00', '2024-04-10 13:40:00', '2024-04-25 15:20:00'),
('2024-06-01 08:45:00', '2025-05-31 18:15:00', '2025-11-30 23:59:59', '2024-04-25 14:20:00', '2024-05-15 10:45:00', '2024-05-30 12:10:00'),
('2024-07-10 07:00:00', '2025-01-10 17:30:00', '2025-07-10 23:59:59', '2024-05-30 11:50:00', '2024-06-20 15:30:00', '2024-07-05 09:40:00'),
('2024-08-01 08:30:00', '2025-07-31 18:00:00', '2026-01-31 23:59:59', '2024-06-15 13:25:00', '2024-07-05 16:10:00', '2024-07-25 14:45:00'),
('2024-09-15 09:15:00', '2025-03-15 17:15:00', '2025-09-15 23:59:59', '2024-07-20 15:40:00', '2024-08-10 11:20:00', '2024-08-30 13:55:00'),
('2024-10-01 08:00:00', '2025-09-30 18:30:00', '2026-03-31 23:59:59', '2024-08-25 12:15:00', '2024-09-15 14:35:00', '2024-09-30 16:25:00');

-- Inserir dados na tabela contract_rh_info
INSERT INTO contract_rh_info (hour_limit, minutes_limit, days_first_exp, days_second_exp, data_init, pay_extra_hour, manual_stitch, pays_breakfast) VALUES
('08:00:00', '00:00:00', 45, 90, '2024-01-15 08:00:00', true, false, true),
('08:30:00', '15:00:00', 30, 60, '2024-02-01 09:00:00', false, true, false),
('08:00:00', '00:00:00', 60, 120, '2024-03-10 07:30:00', true, false, true),
('07:45:00', '30:00:00', 45, 90, '2024-04-01 08:15:00', true, true, false),
('08:15:00', '00:00:00', 30, 90, '2024-05-15 09:30:00', false, false, true),
('08:00:00', '45:00:00', 60, 120, '2024-06-01 08:45:00', true, false, false),
('07:30:00', '00:00:00', 45, 75, '2024-07-10 07:00:00', false, true, true),
('08:30:00', '30:00:00', 30, 60, '2024-08-01 08:30:00', true, false, true),
('08:00:00', '15:00:00', 45, 90, '2024-09-15 09:15:00', true, true, false),
('08:00:00', '00:00:00', 60, 90, '2024-10-01 08:00:00', false, false, true);

-- Inserir dados na tabela contract_discount
INSERT INTO contract_discount (dics_identifier, disc_service, disc_transport, disc_tranp_employee, disc_labor, disc_material, disc_field) VALUES
(5, 10, 8, 12, 15, 7, 9),
(3, 8, 6, 10, 12, 5, 7),
(7, 12, 10, 15, 18, 9, 11),
(4, 9, 7, 11, 14, 6, 8),
(6, 11, 9, 13, 16, 8, 10),
(5, 10, 8, 12, 15, 7, 9),
(8, 15, 12, 18, 20, 10, 13),
(3, 7, 5, 9, 11, 4, 6),
(6, 13, 10, 15, 17, 8, 11),
(4, 8, 6, 10, 13, 5, 7);

-- Inserir dados na tabela contract_info
INSERT INTO contract_info (construction, cap, process, info_pmco, max_employee, address, nro, complement, phone, state, city, cep, email, contact, fk_employee) VALUES
('Construção Residencial Vila Nova', 'CAP-2024-001', 'PROC-RES-001', 'Projeto residencial de alto padrão com 150 unidades', 50, 'Rua das Flores', 123, 'Bloco A', 11987654321, 'SP', 'São Paulo', '01234-567', 'contrato1@empresa.com', 'João Silva', 1),
('Edifício Comercial Center Plaza', 'CAP-2024-002', 'PROC-COM-002', 'Complexo comercial com 20 andares e estacionamento subterrâneo', 80, 'Av. Paulista', 1500, 'Torre Norte', 11876543210, 'SP', 'São Paulo', '01310-100', 'contrato2@empresa.com', 'Maria Santos', 2),
('Infraestrutura Rodoviária BR-101', 'CAP-2024-003', 'PROC-ROD-003', 'Duplicação de 50km da rodovia com pontes e viadutos', 120, 'Rod. BR-101', 0, 'KM 45', 21987654321, 'RJ', 'Rio de Janeiro', '23456-789', 'contrato3@empresa.com', 'Carlos Oliveira', 3),
('Hospital Regional Norte', 'CAP-2024-004', 'PROC-HOS-004', 'Construção de hospital com 200 leitos e centro cirúrgico', 150, 'Rua da Saúde', 500, 'Setor Hospitalar', 85987654321, 'CE', 'Fortaleza', '60000-123', 'contrato4@empresa.com', 'Ana Costa', 4),
('Shopping Boulevard', 'CAP-2024-005', 'PROC-SHO-005', 'Centro comercial com 300 lojas e praça de alimentação', 200, 'Av. dos Estados', 2000, 'Loja Âncora', 47987654321, 'SC', 'Florianópolis', '88000-456', 'contrato5@empresa.com', 'Pedro Almeida', 5),
('Condomínio Residencial Jardins', 'CAP-2024-006', 'PROC-CON-006', 'Condomínio fechado com 80 casas e área de lazer completa', 60, 'Rua dos Jardins', 800, 'Portaria Principal', 31987654321, 'MG', 'Belo Horizonte', '30000-789', 'contrato6@empresa.com', 'Lucia Ferreira', 6),
('Ponte Estaiada Rio Verde', 'CAP-2024-007', 'PROC-PON-007', 'Ponte estaiada com 800m de extensão sobre o rio', 90, 'Marginal Rio Verde', 0, 'Acesso Norte', 62987654321, 'GO', 'Goiânia', '74000-012', 'contrato7@empresa.com', 'Roberto Lima', 7),
('Escola Técnica Profissionalizante', 'CAP-2024-008', 'PROC-ESC-008', 'Campus educacional com laboratórios e biblioteca moderna', 70, 'Rua da Educação', 300, 'Bloco Administrativo', 81987654321, 'PE', 'Recife', '50000-345', 'contrato8@empresa.com', 'Fernanda Souza', 8),
('Terminal Rodoviário Metropolitano', 'CAP-2024-009', 'PROC-TER-009', 'Terminal com 30 plataformas e área comercial integrada', 110, 'Av. Central', 1200, 'Terminal A', 51987654321, 'RS', 'Porto Alegre', '90000-678', 'contrato9@empresa.com', 'Gabriel Torres', 9),
('Aeroporto Regional Cargo', 'CAP-2024-010', 'PROC-AER-010', 'Pista de pouso e terminal de cargas com hangar', 180, 'Zona Aeroportuária', 100, 'Terminal de Cargas', 65987654321, 'MT', 'Cuiabá', '78000-901', 'contrato10@empresa.com', 'Patricia Rocha', 10);

-- Inserir dados na tabela contract_values
INSERT INTO contract_values (acronym_values, bdi_service, bdi_material, bdi_labor, entry_table, send_email) VALUES
('RES-001', 25, 20, 30, true, true),
('COM-002', 28, 22, 32, false, true),
('ROD-003', 30, 25, 35, true, false),
('HOS-004', 27, 23, 33, true, true),
('SHO-005', 26, 21, 31, false, false),
('CON-006', 24, 19, 29, true, true),
('PON-007', 32, 27, 37, true, false),
('ESC-008', 23, 18, 28, false, true),
('TER-009', 29, 24, 34, true, true),
('AER-010', 31, 26, 36, true, false);

-- Inserir dados na tabela principal contracts
INSERT INTO contracts (name, code, research, uses_cpf, is_active, fk_contract_info, fk_contract_dates, fk_contract_values, fk_contract_discount, fk_contract_rh) VALUES
('Contrato Residencial Vila Nova', 'CONT-RES-001', false, true, true, 1, 1, 1, 1, 1),
('Contrato Edifício Comercial Center', 'CONT-COM-002', true, false, true, 2, 2, 2, 2, 2),
('Contrato Infraestrutura Rodoviária', 'CONT-ROD-003', true, true, true, 3, 3, 3, 3, 3),
('Contrato Hospital Regional', 'CONT-HOS-004', false, true, true, 4, 4, 4, 4, 4),
('Contrato Shopping Boulevard', 'CONT-SHO-005', true, false, true, 5, 5, 5, 5, 5),
('Contrato Condomínio Jardins', 'CONT-CON-006', false, true, true, 6, 6, 6, 6, 6),
('Contrato Ponte Estaiada', 'CONT-PON-007', true, true, false, 7, 7, 7, 7, 7),
('Contrato Escola Técnica', 'CONT-ESC-008', false, false, true, 8, 8, 8, 8, 8),
('Contrato Terminal Rodoviário', 'CONT-TER-009', true, true, true, 9, 9, 9, 9, 9),
('Contrato Aeroporto Cargo', 'CONT-AER-010', true, false, false, 10, 10, 10, 10, 10);

-- Verificar os dados inseridos
SELECT 'Contracts inserted:' as info, COUNT(*) as total FROM contracts;
SELECT 'Contract dates inserted:' as info, COUNT(*) as total FROM contract_dates;
SELECT 'Contract RH info inserted:' as info, COUNT(*) as total FROM contract_rh_info;
SELECT 'Contract discount inserted:' as info, COUNT(*) as total FROM contract_discount;
SELECT 'Contract info inserted:' as info, COUNT(*) as total FROM contract_info;
SELECT 'Contract values inserted:' as info, COUNT(*) as total FROM contract_values;
