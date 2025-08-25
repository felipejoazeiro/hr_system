# Testes de Edição de Contract RH Info

Este diretório contém arquivos JSON para testar o endpoint de edição de informações de RH de contratos. O sistema gerencia políticas de trabalho, experiência e benefícios dos funcionários.

## 📋 Arquivos de Teste Disponíveis

### 1. `test_edit_contract_rh_info.json`
**Edição completa** - Atualiza todas as configurações de RH
```json
{
  "hour_limit": "2024-01-01T08:00:00Z",
  "minutes_limit": "2024-01-01T17:30:00Z",
  "days_first_exp": 90,
  "days_second_exp": 180,
  "data_init": "2024-01-15T08:00:00Z",
  "pay_extra_hour": true,
  "manual_stitch": false,
  "pays_breakfast": true
}
```

### 2. `test_edit_contract_rh_info_partial.json`
**Edição parcial** - Atualiza apenas políticas de experiência e horas extras
```json
{
  "days_first_exp": 120,
  "days_second_exp": 240,
  "pay_extra_hour": true
}
```

### 3. `test_edit_contract_rh_info_hours.json`
**Configuração de horários** - Atualiza limites de horário de trabalho
```json
{
  "hour_limit": "2024-01-01T07:30:00Z",
  "minutes_limit": "2024-01-01T18:00:00Z"
}
```

### 4. `test_edit_contract_rh_info_benefits.json`
**Configuração de benefícios** - Atualiza políticas de benefícios
```json
{
  "pay_extra_hour": false,
  "manual_stitch": true,
  "pays_breakfast": false
}
```

### 5. `test_edit_contract_rh_info_experience.json`
**Períodos de experiência** - Atualiza apenas prazos de experiência
```json
{
  "days_first_exp": 60,
  "days_second_exp": 120
}
```

### 6. `test_edit_contract_rh_info_init_date.json`
**Data de início** - Atualiza data de início das políticas
```json
{
  "data_init": "2024-02-01T08:00:00Z"
}
```

## 🚀 Como Usar

### Usando curl:
```bash
# Edição completa das configurações de RH
curl -X PATCH http://localhost:8080/api/contracts/rh-info/1 \
  -H "Content-Type: application/json" \
  -d @test_edit_contract_rh_info.json

# Edição parcial
curl -X PATCH http://localhost:8080/api/contracts/rh-info/1 \
  -H "Content-Type: application/json" \
  -d @test_edit_contract_rh_info_partial.json

# Configuração de horários
curl -X PATCH http://localhost:8080/api/contracts/rh-info/1 \
  -H "Content-Type: application/json" \
  -d @test_edit_contract_rh_info_hours.json
```

### Usando Postman/Insomnia:
1. **Método**: PATCH
2. **URL**: `http://localhost:8080/api/contracts/rh-info/{id}`
3. **Headers**: `Content-Type: application/json`
4. **Body**: Conteúdo de um dos arquivos JSON

### Usando HTTPie:
```bash
http PATCH localhost:8080/api/contracts/rh-info/1 < test_edit_contract_rh_info.json
```

## 📊 Estrutura dos Dados

### Campos de Configuração (todos opcionais):
| Campo | Tipo | Descrição |
|-------|------|-----------|
| `hour_limit` | DateTime | Limite máximo de horas de trabalho |
| `minutes_limit` | DateTime | Limite de minutos para controle |
| `days_first_exp` | Integer | Dias do primeiro período de experiência |
| `days_second_exp` | Integer | Dias do segundo período de experiência |
| `data_init` | DateTime | Data de início das políticas de RH |
| `pay_extra_hour` | Boolean | Se paga horas extras |
| `manual_stitch` | Boolean | Se permite registro manual |
| `pays_breakfast` | Boolean | Se paga café da manhã |

## 🏢 Categorias de Configuração

### ⏰ **Controle de Horário**
- **hour_limit**: Horário limite para trabalho
- **minutes_limit**: Controle de minutos para tolerância
- **data_init**: Início da vigência das regras

### 📅 **Períodos de Experiência**
- **days_first_exp**: Primeiro período (geralmente 45-90 dias)
- **days_second_exp**: Segundo período (geralmente 90-180 dias)

### 💰 **Benefícios e Pagamentos**
- **pay_extra_hour**: Política de horas extras
- **pays_breakfast**: Benefício de café da manhã

### 📝 **Controles Operacionais**
- **manual_stitch**: Permite edição manual de registros

## ⚠️ Validações e Regras

### Formatos Esperados:
- **DateTime**: ISO 8601 com timezone UTC (YYYY-MM-DDTHH:mm:ssZ)
- **Integer**: Números inteiros positivos para dias
- **Boolean**: true/false para políticas

### Regras de Negócio:
- ✅ days_first_exp deve ser menor que days_second_exp
- ✅ hour_limit deve ser anterior a minutes_limit no mesmo dia
- ✅ Todos os campos são opcionais na edição
- ✅ data_init não pode ser no futuro
- ✅ Períodos de experiência: máximo 180 dias cada

## 🧪 Cenários de Teste

| Teste | Descrição | Arquivo |
|-------|-----------|---------|
| **T1** | Configuração completa de RH | `test_edit_contract_rh_info.json` |
| **T2** | Ajuste de experiência e horas extras | `test_edit_contract_rh_info_partial.json` |
| **T3** | Mudança de horários de trabalho | `test_edit_contract_rh_info_hours.json` |
| **T4** | Alteração de benefícios | `test_edit_contract_rh_info_benefits.json` |
| **T5** | Ajuste de períodos de experiência | `test_edit_contract_rh_info_experience.json` |
| **T6** | Mudança de data de início | `test_edit_contract_rh_info_init_date.json` |
| **T7** | JSON vazio (deve dar erro) | `{}` |
| **T8** | ID inexistente (erro 404) | Qualquer arquivo com ID inválido |
| **T9** | Período de experiência inválido | `{"days_first_exp": 200}` |
| **T10** | Data futura inválida | `{"data_init": "2030-01-01T08:00:00Z"}` |

## 📈 Exemplos Práticos

### Configuração Padrão para Empresa:
```json
{
  "hour_limit": "2024-01-01T08:00:00Z",
  "minutes_limit": "2024-01-01T17:48:00Z",
  "days_first_exp": 45,
  "days_second_exp": 90,
  "pay_extra_hour": true,
  "pays_breakfast": true
}
```

### Configuração para Trabalho Flexível:
```json
{
  "hour_limit": "2024-01-01T07:00:00Z",
  "minutes_limit": "2024-01-01T19:00:00Z",
  "manual_stitch": true,
  "pay_extra_hour": false
}
```

### Período de Experiência Estendido:
```json
{
  "days_first_exp": 90,
  "days_second_exp": 180
}
```

### Configuração Sem Benefícios:
```json
{
  "pay_extra_hour": false,
  "pays_breakfast": false,
  "manual_stitch": false
}
```

## 📝 Notas Importantes

- **Horários**: Use horários base (data fictícia) para representar limites de horário
- **Experiência**: Respeite a legislação trabalhista para períodos máximos
- **Benefícios**: Mudanças podem impactar folha de pagamento
- **Auditoria**: Alterações em políticas de RH devem ser registradas
- **Integração**: Configurações afetam sistema de ponto e folha de pagamento

## ⏰ Exemplos de Horários

### Horário Comercial:
- **hour_limit**: `08:00:00Z` (início)
- **minutes_limit**: `17:48:00Z` (fim com tolerância)

### Horário Estendido:
- **hour_limit**: `06:00:00Z` (início cedo)
- **minutes_limit**: `20:00:00Z` (fim tardio)

### Turno Noturno:
- **hour_limit**: `22:00:00Z` (início noite)
- **minutes_limit**: `06:00:00Z` (fim manhã seguinte)
