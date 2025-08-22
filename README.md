# API minimalista em Go

> **Status:** Fase 1 — Projeto de aprendizado com foco em manutenibilidade e performance, com suporte à configuração de múltiplos servidores.

## Highlights

- Não utiliza nenhum framework
- Escalável com o httprouter
- Fácil configuração utilizando arquivo json
- Múltiplos servidores com regras e rotas independentes específicas e/ou globais

---

## Por que minimalista?

> Menos overhead para a manutenção do código e ganhos de performance.

Além do motivo principal, o objetivo do projeto é me familiarizar com a linguagem Go e seus primitivos HTTP.  
Essa escolha por uma abordagem minimalista, sem frameworks, atende a esses objetivos principais e também traz outras vantagens:

- **Código curto:** Mais fácil de ler, entender e manter
- **Menor barreira de entrada:** Novos contribuidores não precisam aprender um framework primeiro; só Go

---

## Funcionamento

Ao ser executado o programa irá ler as configurações de servidor a partir de um arquivo JSON.  
O arquivo de configuração deve estar formatado em JSON conforme o exemplo neste README.

Dados sobre os servidores (endereço ip, porta, timeouts, métodos e rotas) são dinamicamente carregados e validados.

Por fim, a execução dos servidores é feita de forma simultânea, utilizando _goroutines_ e _channels_

---

## Quick start

**Pré-requisitos:** Go 1.22+

```bash
# clonar
git clone https://github.com/re-miranda/go_http_api.git
cd go_http_api

# certificar que módulos estão disponíveis
go mod download

# rodar com as configurações padrão (api_config.json)
make run

# rodar especificando seu próprio arquivo de configuração
./bin/api -config api_config.json

# teste
make test

# limpar
make clean
```
---

## Configuração

A API facilita a configuração por meio de arquivos JSON.

Neste repositório é forneçido um arquivo de exemplo: [api_config.json](api_config.json)  
Nele está definido a criação de dois servidores que rodam simultaneamente em diferentes portas.

Servidores herdam configurações globais caso não especificadas.
> Por esse motivo é obrigatório ter todas as configurações globais no arquivo JSON.

### Exemplo

```json
{
  "Global": {
    "Name": "Go API",
    "ReadTimeout": 5,
    "WriteTimeout": 10,
    "IdleTimeout": 60
  },
  "Servers": [
    {
      "Name": "Reverse API",
      "Host": "0.0.0.0",
      "Port": "8080",
      "Routes": [...]
    },
    {
      "Name": "Ping API",
      "Host": "0.0.0.0",
      "Port": "8081",
      "ReadTimeout": 10,
      "WriteTimeout": 20,
      "IdleTimeout": 120,
      "Routes": [...]
    }
  ]
}
```

---

##  API Endpoints (make run)

### Servidor 1 (Porta 8080) - Reverse API
- **GET /healthz** → `{"status": "ok"}`
- **POST /v1/reverse** → `{"input": "your_input", "output": "your_input_reversed"}`

### Servidor 2 (Porta 8081) - Ping API
- **GET /healthz** → `{"status": "ok"}`
- **GET /v1/ping** → `pong`

### Formato de Erro

Respostas com status diferente de 2xx (Ex.: 400, 401, etc) utilizam estrutura JSON:
```json
{
  "error": "<HTTP status description>",
  "details": ["optional", "error", "details"]
}
```

Exemplos:

- 400 Bad Request (JSON inválido):
```json
{
  "error": "Bad Request",
  "details": ["json: unknown field \"wrong_field\""]
}
```

- 405 Method Not Allowed:
```json
{
  "error": "Method not allowed"
}
```

---

## Licença

[LICENSE](LICENSE) BSD-2-Clause © 2025 Renato Miranda Goncalves.
