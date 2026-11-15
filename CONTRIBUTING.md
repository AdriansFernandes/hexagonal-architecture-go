# Como contribuir

Contribuições para corrigir erros, melhorar as explicações ou ampliar o exemplo são bem-vindas. Antes de começar uma mudança grande, abra uma issue para alinhar o escopo.

## Preparar o ambiente

Use a versão de Go indicada em `go.mod`. Clone o repositório e confirme que os pacotes compilam:

```sh
go test ./...
go vet ./...
```

Atualmente não há testes automatizados. Para mudanças de comportamento, inclua testes que cubram o caso alterado e os erros relevantes.

## Enviar uma contribuição

1. Crie uma branch para a mudança.
2. Mantenha o escopo concentrado em uma correção ou melhoria.
3. Formate os arquivos Go com `gofmt`.
4. Execute `go test ./...` e `go vet ./...`.
5. Atualize o README ou a documentação quando o comportamento ou a arquitetura mudar.
6. Abra um pull request com o motivo da mudança, um resumo do que foi alterado e os comandos executados.

Se um comando não puder ser executado, informe isso no pull request. Não inclua credenciais, dados pessoais ou exemplos com informações reais.

## Padrões do projeto

- Preserve a direção das dependências: domínio e aplicação não importam HTTP nem o adaptador de memória.
- Mantenha as regras de negócio no domínio e a coordenação dos fluxos na aplicação.
- Coloque interfaces junto ao código que precisa delas; os adaptadores implementam essas interfaces implicitamente.
- Faça cada módulo registrar suas rotas por `RegisterRoutes`.
- Evite dependências externas quando a biblioteca padrão atender ao exemplo.
- Mantenha explícitas as limitações didáticas do projeto; ele não processa transações financeiras reais.

## Relatos de problemas

Abra uma issue com a versão de Go, os passos para reproduzir o problema, o resultado esperado e o resultado observado. Para vulnerabilidades, siga [SECURITY.md](SECURITY.md) e não publique detalhes exploráveis em uma issue aberta.
