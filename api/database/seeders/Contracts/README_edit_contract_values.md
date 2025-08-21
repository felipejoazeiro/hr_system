# Testes de Edição de Contract Values

Este diretório contém arquivos JSON para testar o endpoint de edição de `contract_values`.

## Arquivos Disponíveis:

### 1. `test_edit_contract_values.json`
**Edição completa** - Atualiza todos os campos do contract values
```json
{
  "acronym_values": "EDIT-TEST-001",
  "bdi_service": 30,
  "bdi_material": 25,
  "bdi_labor": 35,
  "entry_table": false,
  "send_email": true
}
```

### 2. `test_edit_contract_values_partial.json`
**Edição parcial** - Atualiza apenas campos de BDI
```json
{
  "bdi_service": 28,
  "bdi_material": 22
}
```

### 3. `test_edit_contract_values_flags.json`
**Edição de flags** - Atualiza acronym e configurações booleanas
```json
{
  "acronym_values": "UPDATED-ACRONYM",
  "entry_table": true,
  "send_email": false
}
```

## Como Usar:

### Usando curl:
```bash
# Edição completa do contract values com ID 1
curl -X PATCH http://localhost:8080/api/contracts/values/1 \
  -H "Content-Type: application/json" \
  -d @test_edit_contract_values.json

# Edição parcial
curl -X PATCH http://localhost:8080/api/contracts/values/1 \
  -H "Content-Type: application/json" \
  -d @test_edit_contract_values_partial.json
```

### Usando Postman/Insomnia:
1. Configure o método como PATCH
2. URL: `http://localhost:8080/api/contracts/values/{id}` (substitua {id} pelo ID do contract values)
3. Headers: `Content-Type: application/json`
4. Body: Copie o conteúdo de um dos arquivos JSON

### Usando HTTPie:
```bash
# Edição completa
http PATCH localhost:8080/api/contracts/values/1 < test_edit_contract_values.json

# Edição parcial
http PATCH localhost:8080/api/contracts/values/1 < test_edit_contract_values_partial.json
```

## Estrutura dos Dados:

Cada JSON segue a estrutura do modelo `EditContractValues`:
- **acronym_values**: String - Sigla/acrônimo dos valores
- **bdi_service**: Integer - BDI para serviços
- **bdi_material**: Integer - BDI para materiais  
- **bdi_labor**: Integer - BDI para mão de obra
- **entry_table**: Boolean - Se utiliza tabela de entrada
- **send_email**: Boolean - Se envia email automático

## Notas Importantes:

- **Campos opcionais**: Todos os campos são opcionais na edição. Apenas os campos enviados serão atualizados
- **Validações**: O endpoint deve validar se o ID existe antes de tentar atualizar
- **Resposta**: O endpoint deve retornar o objeto atualizado ou o ID do registro modificado
- **Erro**: Se nenhum campo for enviado, deve retornar erro "Missing fields for update"

## Cenários de Teste:

1. **Teste 1**: Edição completa com todos os campos
2. **Teste 2**: Edição parcial com apenas alguns campos
3. **Teste 3**: Edição apenas de flags booleanas
4. **Teste 4**: Tentar editar com JSON vazio (deve dar erro)
5. **Teste 5**: Tentar editar ID inexistente (deve dar erro 404)
