# Testes de Edição de Contract Info

Este diretório contém arquivos JSON para testar o endpoint de edição de `contract_info`.

## Arquivos Disponíveis:

### 1. `test_edit_contract_info.json`
**Edição completa** - Atualiza todos os campos do contract info
```json
{
  "construction": "Edifício Comercial Atualizado",
  "cap": "CAP-EDIT-2024-001",
  "process": "PROC-EDIT-001",
  "info_pmco": "Informações atualizadas do projeto...",
  "max_employee": 85,
  "address": "Av. Editada, 1500",
  "nro": 1500,
  "complement": "Torre Editada",
  "phone": 11999888777,
  "state": "RJ",
  "city": "Rio de Janeiro",
  "cep": "20000-000",
  "email": "editado@empresa.com",
  "contact": "Gerente Editado",
  "fk_employee": 3
}
```

### 2. `test_edit_contract_info_partial.json`
**Edição parcial** - Atualiza apenas alguns campos básicos
```json
{
  "construction": "Nome da Construção Parcialmente Editado",
  "max_employee": 120,
  "phone": 21987654321
}
```

### 3. `test_edit_contract_info_address.json`
**Edição de endereço** - Atualiza apenas campos relacionados ao endereço
```json
{
  "address": "Nova Rua dos Endereços, 2024",
  "nro": 2024,
  "complement": "Novo Complemento",
  "state": "MG",
  "city": "Belo Horizonte",
  "cep": "30000-000"
}
```

### 4. `test_edit_contract_info_contact.json`
**Edição de contato** - Atualiza apenas informações de contato
```json
{
  "email": "novo.email@empresa.com",
  "contact": "Novo Contato Responsável",
  "phone": 85999888777,
  "fk_employee": 5
}
```

## Como Usar:

### Usando curl:
```bash
# Edição completa do contract info com ID 1
curl -X PATCH http://localhost:8080/api/contracts/info/1 \
  -H "Content-Type: application/json" \
  -d @test_edit_contract_info.json

# Edição parcial
curl -X PATCH http://localhost:8080/api/contracts/info/1 \
  -H "Content-Type: application/json" \
  -d @test_edit_contract_info_partial.json

# Edição de endereço
curl -X PATCH http://localhost:8080/api/contracts/info/1 \
  -H "Content-Type: application/json" \
  -d @test_edit_contract_info_address.json

# Edição de contato
curl -X PATCH http://localhost:8080/api/contracts/info/1 \
  -H "Content-Type: application/json" \
  -d @test_edit_contract_info_contact.json
```

### Usando Postman/Insomnia:
1. Configure o método como PATCH
2. URL: `http://localhost:8080/api/contracts/info/{id}` (substitua {id} pelo ID do contract info)
3. Headers: `Content-Type: application/json`
4. Body: Copie o conteúdo de um dos arquivos JSON

### Usando HTTPie:
```bash
# Edição completa
http PATCH localhost:8080/api/contracts/info/1 < test_edit_contract_info.json

# Edição parcial
http PATCH localhost:8080/api/contracts/info/1 < test_edit_contract_info_partial.json
```

## Estrutura dos Dados:

Cada JSON segue a estrutura do modelo `EditContractInfo`:
- **construction**: String - Nome da construção/obra
- **cap**: String - Código CAP do projeto
- **process**: String - Número do processo
- **info_pmco**: String - Informações do PMCO
- **max_employee**: Integer - Número máximo de funcionários
- **address**: String - Endereço da obra
- **nro**: Integer - Número do endereço
- **complement**: String - Complemento do endereço
- **phone**: Integer - Telefone de contato
- **state**: String - Estado (2 caracteres)
- **city**: String - Cidade
- **cep**: String - CEP
- **email**: String - Email de contato
- **contact**: String - Nome do contato responsável
- **fk_employee**: Integer - ID do funcionário responsável

## Notas Importantes:

- **Campos opcionais**: Todos os campos são opcionais na edição. Apenas os campos enviados serão atualizados
- **Validações esperadas**: 
  - Email deve ter formato válido
  - CEP deve ter formato válido (XXXXX-XXX)
  - State deve ter exatamente 2 caracteres
  - Phone deve ser um número válido
  - fk_employee deve referenciar um funcionário existente
- **Resposta**: O endpoint deve retornar o ID do registro modificado
- **Erro**: Se nenhum campo for enviado, deve retornar erro "nenhum campo para atualizar"

## Cenários de Teste:

1. **Teste 1**: Edição completa com todos os campos
2. **Teste 2**: Edição parcial com apenas alguns campos básicos
3. **Teste 3**: Edição focada no endereço
4. **Teste 4**: Edição focada nas informações de contato
5. **Teste 5**: Tentar editar com JSON vazio (deve dar erro)
6. **Teste 6**: Tentar editar ID inexistente (deve dar erro 404)
7. **Teste 7**: Validação de email inválido
8. **Teste 8**: Validação de fk_employee inexistente
