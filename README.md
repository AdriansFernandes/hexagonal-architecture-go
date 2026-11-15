# Arquitetura hexagonal em Go

Uma API pequena de transferências, construída para mostrar como domínio, aplicação, portas e adaptadores se conectam em Go.

O código acompanha o [tutorial guiado do Bendev Oficial](https://bendevoficial.com/learning/aprendendo-hexagonal-em-go). A trilha explica as decisões arquiteturais e os recursos da linguagem usados em cada etapa.

## O que o exemplo faz

- Expõe `GET /health` e `POST /transfers`.
- Transfere saldo entre duas contas de demonstração, `a` e `b`.
- Mantém as contas em um repositório em memória protegido para acesso concorrente.
- Organiza a composição no pacote `internal/bootstrap` e permite que cada handler registre suas rotas.

O saldo usa inteiros para manter o exemplo simples. Não há banco de dados, autenticação ou regras de moeda. O mutex protege o map contra acesso concorrente, mas a transferência inteira não é uma transação. Os valores voltam ao estado inicial quando o processo reinicia; não use o projeto para processar transações reais.

## Requisitos e execução

Use a versão de Go indicada em [`go.mod`](go.mod) ou uma versão mais recente. Na raiz do repositório, execute:

```sh
go run ./cmd/api
```

O servidor escuta na porta `8080`. Consulte o health check:

```sh
curl http://localhost:8080/health
```

Faça uma transferência de 500 unidades da conta `a` para a conta `b`:

```sh
curl -X POST http://localhost:8080/transfers \
  -H 'Content-Type: application/json' \
  -d '{"fromId":"a","toId":"b","amount":500}'
```

No Prompt de Comando do Windows (`cmd.exe`), use aspas duplas e escape as aspas do JSON:

```bat
curl.exe -X POST "http://localhost:8080/transfers" -H "Content-Type: application/json" -d "{\"fromId\":\"a\",\"toId\":\"b\",\"amount\":500}"
```

No PowerShell, use `curl.exe` e aspas simples ao redor do JSON:

```powershell
curl.exe -X POST "http://localhost:8080/transfers" -H "Content-Type: application/json" -d '{"fromId":"a","toId":"b","amount":500}'
```

Uma transferência válida retorna HTTP `201`:

```json
{"message":"transfer completed"}
```

## Organização do código

```text
cmd/
└── api/
    └── main.go

internal/
├── bootstrap/
│   ├── app.go
│   └── router.go
├── modules/
│   └── transfers/
│       ├── adapters/
│       │   └── persistence/
│       │       └── memory/repository.go
│       ├── application/service.go
│       ├── domain/
│       │   ├── account.go
│       │   ├── errors.go
│       │   └── ports/repository.go
│       └── entrypoints/
│           └── httpapi/
│               ├── dto.go
│               ├── errors.go
│               ├── handler.go
│               └── routes.go
└── platform/
    └── httpserver/
        ├── health.go
        ├── response.go
        └── server.go
```

`domain` contém as regras de conta. `application` coordena a transferência por meio da porta `AccountRepository`. O adaptador `memory` implementa essa porta. `entrypoints/httpapi` converte HTTP em chamadas da aplicação e traduz os erros para respostas HTTP. `bootstrap` conecta as implementações concretas, enquanto `platform/httpserver` reúne funções HTTP compartilhadas.

O fluxo de uma transferência é:

```text
HTTP → handler → serviço de aplicação → domínio
                         ↓
                 porta AccountRepository
                         ↓
                 adaptador em memória
```

O código do núcleo depende do contrato do repositório. O adaptador implementa esse contrato, e o bootstrap escolhe qual implementação será usada. O router recebe handlers que implementam `RouteRegister`, então cada módulo registra seus próprios caminhos.

Para uma descrição das responsabilidades e dependências de cada pacote, consulte [docs/architecture.md](docs/architecture.md).

## Tutorial e referências

- [Tutorial: arquitetura hexagonal em Go](https://bendevoficial.com/learning/aprendendo-hexagonal-em-go)
- [Referência de Go no Bendev Oficial](https://bendevoficial.com/reference)
- [Interfaces em Go](https://bendevoficial.com/reference/go/interfaces)
- [Ponteiros e endereços em Go](https://bendevoficial.com/reference/go/pointers)
- [Métodos e tipos em Go](https://bendevoficial.com/reference/go/methods-and-types)
- [Literais compostos em Go](https://bendevoficial.com/reference/go/composite-literals)
- [Slices em Go](https://bendevoficial.com/reference/go/slices)
- [Maps e o padrão comma-ok](https://bendevoficial.com/reference/go/maps)
- [Identificador em branco (`_`) em Go](https://bendevoficial.com/reference/go/blank-identifier)
- [Go Language Specification](https://go.dev/ref/spec)
- [Effective Go](https://go.dev/doc/effective_go)
- [A Tour of Go](https://go.dev/tour/)

## Verificações

```sh
go test ./...
go vet ./...
```

`go test ./...` compila todos os pacotes e executa os testes disponíveis. O projeto ainda não contém testes automatizados.

## Contribuições

Leia [CONTRIBUTING.md](CONTRIBUTING.md) antes de abrir uma contribuição. Para relatar uma vulnerabilidade, siga as orientações em [SECURITY.md](SECURITY.md).

## Licença

Este projeto está sob a licença [MIT](LICENSE).
