# Deploy-Hub

Deploy-Hub e uma CLI em Go para organizar e executar deploys simples em servidores via SSH. A ideia e manter um arquivo `deploy.yaml` com seus servidores, servicos e comandos de deploy, e depois disparar tudo pelo terminal de forma interativa ou direta.

O projeto usa [Cobra](https://github.com/spf13/cobra) para os comandos da CLI, [Survey](https://github.com/AlecAivazis/survey) para prompts interativos e `golang.org/x/crypto/ssh` para executar comandos remotos.

## O que ele faz

- Inicializa um projeto criando um `deploy.yaml`.
- Cadastra, lista, atualiza e remove servidores.
- Cadastra, lista, atualiza e remove servicos de deploy.
- Executa comandos remotos via SSH no servidor associado ao servico.
- Permite escolher servicos interativamente quando o nome nao e informado.

## Como funciona

O Deploy-Hub gira em torno de um arquivo local chamado `deploy.yaml`.

Nele voce define:

- `project`: nome do projeto.
- `servers`: servidores remotos disponiveis.
- `services`: servicos que apontam para um servidor, um diretorio remoto e uma lista de comandos.

Quando voce roda um deploy, a CLI:

1. Carrega o `deploy.yaml`.
2. Busca o servico solicitado.
3. Busca o servidor associado ao servico.
4. Abre uma conexao SSH.
5. Executa cada comando no formato `cd <path> && <command>`.

## Requisitos

- Go 1.25 ou superior.
- Acesso SSH ao servidor de destino.
- Um arquivo `deploy.yaml` configurado na raiz do projeto onde o comando sera executado.

## Instalacao

Clone o repositorio:

```bash
git clone https://github.com/HTTPauloGoncalves/Deploy-Hub.git
cd Deploy-Hub
```

Baixe as dependencias:

```bash
go mod download
```

Rode em modo desenvolvimento:

```bash
go run .
```

Ou compile o binario:

```bash
go build -o deploy-hub
```

No Windows:

```powershell
go build -o deploy-hub.exe
```

## Inicio rapido

Crie o arquivo de configuracao:

```bash
go run . init MeuProjeto
```

Cadastre um servidor:

```bash
go run . server add
```

Cadastre um servico de deploy:

```bash
go run . service add
```

Execute o deploy:

```bash
go run . deploy nome-do-servico
```

Se voce rodar `deploy` sem informar o servico, a CLI mostra uma lista interativa:

```bash
go run . deploy
```

## Exemplo de `deploy.yaml`

```yaml
project: MeuProjeto
servers:
  producao:
    host: 192.168.0.10
    user: root
    password: sua-senha-aqui
    auth: password
    port: 22
services:
  api:
    server: producao
    path: /var/www/minha-api
    command:
      - git pull origin main
      - go build -o app
      - systemctl restart minha-api
```

Neste exemplo, o comando:

```bash
go run . deploy api
```

executaria no servidor:

```bash
cd /var/www/minha-api && git pull origin main
cd /var/www/minha-api && go build -o app
cd /var/www/minha-api && systemctl restart minha-api
```

## Comandos disponiveis

### `init`

Inicializa um projeto criando o arquivo `deploy.yaml`.

```bash
go run . init MeuProjeto
```

Se o arquivo ja existir, ele nao sobrescreve a configuracao atual.

### `server add`

Adiciona um servidor de forma interativa.

```bash
go run . server add
```

Campos solicitados:

| Campo | Descricao |
| --- | --- |
| Nome do servidor | Apelido usado dentro do `deploy.yaml` |
| Host | IP ou dominio do servidor |
| Usuario | Usuario SSH |
| Autenticacao | `password` ou `key` |
| Senha | Usada quando a autenticacao for `password` |
| Porta | Porta SSH, normalmente `22` |

Observacao: atualmente o deploy usa conexao por senha. A opcao `key` ja aparece no cadastro, mas a execucao por chave ainda precisa ser implementada.

### `server list`

Lista os servidores cadastrados.

```bash
go run . server list
```

### `server update`

Atualiza um servidor existente usando prompts interativos.

```bash
go run . server update
```

### `server delete`

Remove um servidor cadastrado.

```bash
go run . server delete
```

### `service add`

Adiciona um servico de deploy.

```bash
go run . service add
```

Campos solicitados:

| Campo | Descricao |
| --- | --- |
| Nome do servico | Nome usado para executar o deploy |
| Servidor | Servidor previamente cadastrado |
| Path na VPS | Diretorio remoto onde os comandos serao executados |
| Comandos | Lista de comandos de deploy |

Para finalizar a lista de comandos, digite:

```text
done
```

### `service list`

Lista os servicos cadastrados e seus comandos.

```bash
go run . service list
```

### `service update`

Atualiza servidor, path e comandos de um servico existente.

```bash
go run . service update
```

### `service delete`

Remove um servico de deploy.

```bash
go run . service delete
```

### `deploy`

Executa o deploy de um servico.

```bash
go run . deploy api
```

Tambem e possivel chamar sem argumento e escolher o servico em uma lista:

```bash
go run . deploy
```

### `logs`

O comando `logs` existe na CLI, mas ainda esta como placeholder.

```bash
go run . logs
```

Atualmente ele apenas imprime uma mensagem simples no terminal.

## Estrutura do projeto

```text
.
|-- cmd/
|   |-- deploy.go
|   |-- init.go
|   |-- logs.go
|   |-- root.go
|   |-- server.go
|   `-- service.go
|-- internal/
|   |-- config/
|   |   |-- models.go
|   |   `-- yaml.go
|   |-- deploy.go
|   |-- init.go
|   |-- logs.go
|   |-- server.go
|   |-- service.go
|   `-- sshClient.go
|-- main.go
|-- go.mod
`-- README.md
```

## Seguranca

O arquivo `deploy.yaml` pode conter senhas de servidores. Por isso:

- Nao commite `deploy.yaml` em repositorios publicos.
- Use `.gitignore` para manter esse arquivo fora do Git.
- Prefira senhas fortes e usuarios com permissoes limitadas.
- Evite rodar comandos destrutivos sem revisar a lista do servico.

O `.gitignore` do projeto ja inclui `deploy.yaml`.

## Desenvolvimento

Rode os testes e a validacao basica dos pacotes:

```bash
go test ./...
```

Formate o codigo:

```bash
gofmt -w .
```

## Possiveis proximos passos

- Implementar autenticacao por chave SSH.
- Implementar o comando `logs`.
- Adicionar suporte a arquivos `.env`.
- Melhorar mensagens de erro e validacoes dos prompts.
- Adicionar testes para manipulacao do `deploy.yaml`.
- Permitir configurar um caminho customizado para o arquivo de deploy.

## Licenca

Este projeto possui um arquivo `LICENSE`, mas ele ainda esta vazio. Adicione os termos de licenca antes de distribuir ou publicar oficialmente.
