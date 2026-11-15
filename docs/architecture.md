# Arquitetura do exemplo

O projeto demonstra uma API HTTP de transferência. O núcleo aplica as regras de conta e coordena o caso de uso; os adaptadores traduzem os protocolos e o armazenamento.

## Pacotes

| Pacote | Responsabilidade | Depende de |
| --- | --- | --- |
| `internal/modules/transfers/domain` | Conta, regras de saldo e erros de negócio | Biblioteca padrão |
| `internal/modules/transfers/domain/ports` | Contrato `AccountRepository` usado pela aplicação | `domain` |
| `internal/modules/transfers/application` | Busca as contas, aplica as regras e solicita a gravação | `domain` e `domain/ports` |
| `internal/modules/transfers/adapters/persistence/memory` | Implementa o repositório com um map em memória | `domain` e `domain/ports` |
| `internal/modules/transfers/entrypoints/httpapi` | Lê a requisição, chama a aplicação e converte erros em status HTTP | `application`, `domain` e `platform/httpserver` |
| `internal/platform/httpserver` | Funções comuns de JSON, health check e configuração do servidor | Biblioteca padrão |
| `internal/bootstrap` | Instancia os adaptadores e registra os handlers no router | Implementações concretas dos módulos e `platform/httpserver` |
| `cmd/api` | Inicia o servidor HTTP | `bootstrap` e `platform/httpserver` |

O caminho da chamada de transferência é:

```text
POST /transfers
  → entrypoints/httpapi.Handler
  → application.TransferService
  → domain.Account
  → domain/ports.AccountRepository
  → adapters/persistence/memory.Accounts
```

## Portas e composição

`AccountRepository` fica em `domain/ports` porque é o contrato que o caso de uso precisa para buscar e salvar contas. O adaptador em memória satisfaz esse contrato por implementação implícita: ele tem os métodos `Get` e `Save` com as assinaturas exigidas.

O handler HTTP chama `TransferService` diretamente. O exemplo não cria uma interface de entrada para o serviço porque só há um adaptador de entrada e essa abstração não é necessária para demonstrar a separação.

`bootstrap.NewRouter` recebe valores que implementam `RouteRegister`. Cada handler registra os caminhos do seu próprio módulo; o router adiciona o endpoint compartilhado `/health`. `bootstrap.Bootstrap` instancia o repositório, o serviço e o handler e passa o handler ao router.

## Armazenamento em memória

O adaptador protege o map com `sync.RWMutex` para permitir acesso concorrente sem corrida de memória. `Save` valida todas as contas antes de atualizar o map. O exemplo não oferece uma transação que cubra toda a sequência de leitura, alteração e gravação do caso de uso; portanto, não deve ser usado como implementação de saldo em produção.

Os dados são reiniciados quando o processo termina. A API não tem autenticação, banco de dados, observabilidade nem mecanismos de recuperação de middleware.
