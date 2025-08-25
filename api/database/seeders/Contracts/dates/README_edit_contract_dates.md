# Testes de Edição de Contract Dates

Este diretório contém arquivos JSON para testar o endpoint de edição de `contract_dates`.

## Arquivos Disponíveis:

### 1. `test_edit_contract_dates.json`
**Edição completa** - Atualiza todas as datas do contrato
```json
{
  "date_initial": "2024-01-15T08:00:00Z",
  "date_limit": "2025-01-15T17:00:00Z",
  "date_guarantee": "2027-01-15T17:00:00Z",
  "date_proposal": "2023-11-20T14:30:00Z",
  "date_budget": "2023-12-01T16:00:00Z",
  "date_tables": "2023-12-15T10:00:00Z"
}
```

### 2. `test_edit_contract_dates_partial.json`
**Edição parcial** - Atualiza apenas datas de início e fim
```json
{
  "date_initial": "2024-02-01T08:00:00Z",
  "date_limit": "2025-06-30T17:00:00Z"
}
```

### 3. `test_edit_contract_dates_guarantee.json`
**Edição de garantia** - Atualiza apenas a data de garantia
```json
{
  "date_guarantee": "2026-12-31T23:59:59Z"
}
```

### 4. `test_edit_contract_dates_planning.json`
**Edição de planejamento** - Atualiza datas de proposta, orçamento e tabelas
```json
{
  "date_proposal": "2023-10-15T09:00:00Z",
  "date_budget": "2023-11-01T14:00:00Z",
  "date_tables": "2023-11-15T16:30:00Z"
}
```

### 5. `test_edit_contract_dates_execution.json`
**Edição de execução** - Atualiza datas relacionadas à execução do contrato
```json
{
  "date_initial": "2024-03-01T08:00:00Z",
  "date_limit": "2026-02-28T17:00:00Z",
  "date_guarantee": "2029-02-28T17:00:00Z"
}
```

## Como Usar:

### Usando curl:
```bash
# Edição completa do contract dates com ID 1
curl -X PATCH http://localhost:8080/api/contracts/dates/1 \
  -H "Content-Type: application/json" \
  -d @test_edit_contract_dates.json

# Edição parcial
curl -X PATCH http://localhost:8080/api/contracts/dates/1 \
  -H "Content-Type: application/json" \
  -d @test_edit_contract_dates_partial.json

# Edição de garantia
curl -X PATCH http://localhost:8080/api/contracts/dates/1 \
  -H "Content-Type: application/json" \
  -d @test_edit_contract_dates_guarantee.json

# Edição de planejamento
curl -X PATCH http://localhost:8080/api/contracts/dates/1 \
  -H "Content-Type: application/json" \
  -d @test_edit_contract_dates_planning.json

# Edição de execução
curl -X PATCH http://localhost:8080/api/contracts/dates/1 \
  -H "Content-Type: application/json" \
  -d @test_edit_contract_dates_execution.json
```

### Usando Postman/Insomnia:
1. Configure o método como PATCH
2. URL: `http://localhost:8080/api/contracts/dates/{id}` (substitua {id} pelo ID do contract dates)
3. Headers: `Content-Type: application/json`
4. Body: Copie o conteúdo de um dos arquivos JSON

### Usando HTTPie:
```bash
# Edição completa
http PATCH localhost:8080/api/contracts/dates/1 < test_edit_contract_dates.json

# Edição parcial
http PATCH localhost:8080/api/contracts/dates/1 < test_edit_contract_dates_partial.json
```

## Estrutura dos Dados:

Cada JSON segue a estrutura do modelo `EditContractDates`:
- **date_initial**: String (ISO 8601) - Data de início do contrato
- **date_limit**: String (ISO 8601) - Data limite/fim do contrato
- **date_guarantee**: String (ISO 8601) - Data de garantia
- **date_proposal**: String (ISO 8601) - Data da proposta
- **date_budget**: String (ISO 8601) - Data do orçamento
- **date_tables**: String (ISO 8601) - Data das tabelas

## Cronologia Típica das Datas:

1. **date_proposal** (mais antiga) - Quando a proposta foi feita
2. **date_budget** - Quando o orçamento foi aprovado
3. **date_tables** - Quando as tabelas foram definidas
4. **date_initial** - Início da execução do contrato
5. **date_limit** - Fim previsto do contrato
6. **date_guarantee** (mais recente) - Fim do período de garantia

## Notas Importantes:

- **Campos opcionais**: Todos os campos são opcionais na edição. Apenas os campos enviados serão atualizados
- **Formato de data**: Use formato ISO 8601 com timezone UTC (YYYY-MM-DDTHH:mm:ssZ)
- **Validações esperadas**: 
  - date_initial deve ser anterior a date_limit
  - date_guarantee normalmente é posterior a date_limit
  - date_proposal normalmente é anterior a date_initial
  - Todas as datas devem estar em formato válido
- **Resposta**: O endpoint deve retornar o ID do registro modificado
- **Erro**: Se nenhum campo for enviado, deve retornar erro "nenhum campo para atualizar"

## Cenários de Teste:

1. **Teste 1**: Edição completa com todas as datas
2. **Teste 2**: Edição parcial com apenas início e fim
3. **Teste 3**: Extensão da garantia
4. **Teste 4**: Atualização das datas de planejamento
5. **Teste 5**: Prorrogação do contrato (date_limit e date_guarantee)
6. **Teste 6**: Tentar editar com JSON vazio (deve dar erro)
7. **Teste 7**: Tentar editar ID inexistente (deve dar erro 404)
8. **Teste 8**: Validação de data inválida (date_initial > date_limit)
9. **Teste 9**: Validação de formato de data inválido
10. **Teste 10**: Datas em fusos horários diferentes

## Exemplos de Cenários Reais:

### Prorrogação de Contrato:
```json
{
  "date_limit": "2025-12-31T17:00:00Z",
  "date_guarantee": "2028-12-31T17:00:00Z"
}
```

### Antecipação de Início:
```json
{
  "date_initial": "2024-01-01T08:00:00Z"
}
```

### Extensão de Garantia:
```json
{
  "date_guarantee": "2030-06-30T23:59:59Z"
}
```

### Atualização de Planejamento:
```json
{
  "date_proposal": "2023-09-01T10:00:00Z",
  "date_budget": "2023-10-15T14:00:00Z"
}
```
