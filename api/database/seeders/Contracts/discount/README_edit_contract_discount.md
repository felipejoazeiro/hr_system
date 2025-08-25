# Testes de Edição de Contract Discount

Este diretório contém arquivos JSON para testar o endpoint de edição de descontos de contratos. O sistema permite aplicar diferentes tipos de descontos em categorias específicas do contrato.

## 📋 Arquivos de Teste Disponíveis

### 1. `test_edit_contract_discount.json`
**Edição completa** - Atualiza todos os campos de desconto
```json
{
  "disc_identifier": 15.5,
  "disc_service": 10.0,
  "disc_transport": 8.0,
  "disc_tranp_employee": 12.0,
  "disc_labor": 7.5,
  "disc_material": 5.0,
  "disc_field": 20.0
}
```

### 2. `test_edit_contract_discount_partial.json`
**Edição parcial** - Atualiza apenas alguns descontos específicos
```json
{
  "disc_service": 12.0,
  "disc_material": 8.0
}
```

### 3. `test_edit_contract_discount_dates.json`
**Descontos de categoria única** - Foca em descontos de transporte
```json
{
  "disc_transport": 15.0,
  "disc_tranp_employee": 18.0
}
```

### 4. `test_edit_contract_discount_value_type.json`
**Descontos operacionais** - Atualiza serviços, mão de obra e campo
```json
{
  "disc_service": 14.0,
  "disc_labor": 11.0,
  "disc_field": 25.0
}
```

### 5. `test_edit_contract_discount_deactivate.json`
**Zerar descontos** - Remove todos os descontos (valores zero)
```json
{
  "disc_identifier": 0.0,
  "disc_service": 0.0,
  "disc_transport": 0.0,
  "disc_tranp_employee": 0.0,
  "disc_labor": 0.0,
  "disc_material": 0.0,
  "disc_field": 0.0
}
```

## 🚀 Como Usar

### Usando curl:
```bash
# Edição completa de descontos
curl -X PATCH http://localhost:8080/api/contracts/discount/1 \
  -H "Content-Type: application/json" \
  -d @test_edit_contract_discount.json

# Edição parcial de descontos específicos
curl -X PATCH http://localhost:8080/api/contracts/discount/1 \
  -H "Content-Type: application/json" \
  -d @test_edit_contract_discount_partial.json
```

### Usando Postman/Insomnia:
1. **Método**: PATCH
2. **URL**: `http://localhost:8080/api/contracts/discount/{id}`
3. **Headers**: `Content-Type: application/json`
4. **Body**: Conteúdo de um dos arquivos JSON

### Usando HTTPie:
```bash
http PATCH localhost:8080/api/contracts/discount/1 < test_edit_contract_discount.json
```

## 📊 Estrutura dos Dados

### Campos de Desconto (todos opcionais):
| Campo | Tipo | Descrição |
|-------|------|-----------|
| `disc_identifier` | Float | Desconto para identificadores/códigos |
| `disc_service` | Float | Desconto para serviços |
| `disc_transport` | Float | Desconto para transporte geral |
| `disc_tranp_employee` | Float | Desconto para transporte de funcionários |
| `disc_labor` | Float | Desconto para mão de obra |
| `disc_material` | Float | Desconto para materiais |
| `disc_field` | Float | Desconto para trabalho de campo |

## 💡 Categorias de Desconto

### 🔧 **Operacionais**
- **disc_service**: Descontos em serviços prestados
- **disc_labor**: Descontos em mão de obra
- **disc_field**: Descontos em trabalhos de campo

### 🚛 **Logística**
- **disc_transport**: Descontos em transporte geral
- **disc_tranp_employee**: Descontos em transporte de pessoal

### 📦 **Materiais**
- **disc_material**: Descontos em materiais e insumos

### 🏷️ **Administrativos**
- **disc_identifier**: Descontos em identificadores/códigos

## ✅ Validações e Regras

### Valores Aceitos:
- **Tipo**: Números decimais (float)
- **Faixa**: 0.0 a 100.0 (porcentagem)
- **Formato**: Até 2 casas decimais

### Regras de Negócio:
- ✅ Todos os campos são opcionais
- ✅ Valores devem ser não-negativos
- ✅ Valores acima de 100% podem gerar alertas
- ✅ Apenas campos enviados são atualizados

## 🧪 Cenários de Teste

| Teste | Descrição | Arquivo |
|-------|-----------|---------|
| **T1** | Edição completa de todos os descontos | `test_edit_contract_discount.json` |
| **T2** | Edição parcial (serviços e materiais) | `test_edit_contract_discount_partial.json` |
| **T3** | Foco em transporte | `test_edit_contract_discount_dates.json` |
| **T4** | Descontos operacionais | `test_edit_contract_discount_value_type.json` |
| **T5** | Zerar todos os descontos | `test_edit_contract_discount_deactivate.json` |
| **T6** | JSON vazio (deve dar erro) | `{}` |
| **T7** | ID inexistente (erro 404) | Qualquer arquivo com ID inválido |
| **T8** | Valores negativos (deve dar erro) | `{"disc_service": -5.0}` |
| **T9** | Valores muito altos | `{"disc_service": 150.0}` |
| **T10** | Campos inválidos | `{"invalid_field": 10.0}` |

## 📈 Exemplos Práticos

### Desconto para Contrato de Grande Porte:
```json
{
  "disc_service": 12.0,
  "disc_labor": 8.0,
  "disc_material": 5.0,
  "disc_transport": 10.0
}
```

### Desconto Sazonal (apenas serviços):
```json
{
  "disc_service": 15.0,
  "disc_field": 20.0
}
```

### Promoção de Materiais:
```json
{
  "disc_material": 18.0
}
```

### Remoção de Descontos:
```json
{
  "disc_service": 0.0,
  "disc_labor": 0.0
}
```

## ⚠️ Notas Importantes

- **Precisão**: Use até 2 casas decimais para evitar problemas de arredondamento
- **Validação**: O sistema deve validar se os valores estão dentro da faixa aceitável
- **Auditoria**: Mudanças em descontos devem ser logadas para auditoria
- **Cálculo**: Descontos são aplicados sobre valores base do contrato
- **Hierarquia**: Verifique a ordem de aplicação dos descontos no sistema
