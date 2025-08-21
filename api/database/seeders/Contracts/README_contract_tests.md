# Exemplos de JSON para Teste da API de Contratos

Este diretório contém arquivos JSON para testar o endpoint `createContract` da API.

## Arquivos Disponíveis:

### 1. `test_create_contract.json`
**Contrato básico de teste** - Exemplo simples para validação inicial da API
- Tipo: Geral/Teste
- Funcionários: 75
- Duração: 12 meses
- Características: Research ativo, usa CPF, todos os benefícios de RH

### 2. `test_create_contract_residential.json` 
**Contrato Residencial de Luxo** - Projeto residencial de alto padrão
- Tipo: Residencial Premium
- Funcionários: 120
- Duração: 24 meses
- Características: Sem research, usa CPF, benefícios completos

### 3. `test_create_contract_infrastructure.json`
**Contrato de Infraestrutura Metropolitana** - Linha de metrô
- Tipo: Infraestrutura Pública
- Funcionários: 350
- Duração: 36 meses
- Características: Research ativo, não usa CPF, benefícios básicos

## Como Usar:

### Usando curl:
```bash
curl -X POST http://localhost:8080/api/contracts \
  -H "Content-Type: application/json" \
  -d @test_create_contract.json
```

### Usando Postman/Insomnia:
1. Configure o método como POST
2. URL: `http://localhost:8080/api/contracts`
3. Headers: `Content-Type: application/json`
4. Body: Copie o conteúdo de um dos arquivos JSON

### Usando HTTPie:
```bash
http POST localhost:8080/api/contracts < test_create_contract.json
```

## Estrutura dos Dados:

Cada JSON segue a estrutura do modelo `CreateFullContractRequest`:
- **Dados principais**: name, code, research, uses_cpf, is_active
- **Info**: Informações detalhadas do contrato
- **Dates**: Datas importantes do projeto
- **Values**: Valores e configurações de BDI
- **Discount**: Descontos por categoria
- **RH Info**: Configurações de recursos humanos

## Validações Esperadas:

- Código único do contrato
- Datas coerentes (inicial < limite < garantia)
- Valores de desconto e BDI válidos
- Referência válida de funcionário (fk_employee)
- Formato correto de email e telefone
