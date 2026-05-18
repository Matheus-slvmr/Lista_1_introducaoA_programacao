# Servidor HTTP em GoLang - Sistema de Gerenciamento de Usuários

Um servidor HTTP desenvolvido em GoLang com integração ao banco de dados PostgreSQL para gerenciamento de contas de usuários.

## 📋 Requisitos

Antes de começar, certifique-se de ter instalado:

- **Go** (versão 1.24.1 ou superior)
  - Download: https://golang.org/dl
- **PostgreSQL** (versão 12 ou superior)
  - Download: https://www.postgresql.org/download
- **Git** (opcional, para clonar o repositório)

## 🚀 Como Rodar o Sistema

### 1. Clonar/Acessar o Projeto

```bash
cd AA1
```

### 2. Configurar o Banco de Dados

#### 2.1 Criar o banco de dados PostgreSQL

```sql
CREATE DATABASE saude_ex;
```

#### 2.2 Conectar ao banco e criar as tabelas

```sql
\c saude_ex

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    born_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 3. Configurar as Variáveis de Ambiente

Crie um arquivo `.env` na pasta `app/` com as seguintes informações:

```bash
DB_USER=postgres
DB_PASSWORD=seu_password_aqui(postgre)
DB_NAME=saude_ex
DB_HOST=localhost
DB_PORT=5432
```

**Exemplo completo do arquivo `app/.env`:**

```
DB_USER=postgres
DB_PASSWORD=postgre
DB_NAME=saude_ex
DB_HOST=localhost
DB_PORT=5432
```

> ⚠️ **IMPORTANTE:** O arquivo `.env` deve estar na pasta `app/`, não na raiz do projeto.

### 4. Instalar Dependências

Na pasta raiz do projeto (AA1), execute:

```bash
go mod download
```

### 5. Rodar o Servidor

Entre na pasta `app` e execute:

```bash
cd app
go run main.go
```

Você deverá ver a mensagem:

```
Conexão com o banco de dados estabelecida com sucesso!
Servidor rodando em: http://169.254.3.248:3000/
```

### 6. Acessar a Aplicação

Abra seu navegador e acesse:

```
http://localhost:3000/
```

## 📱 Funcionalidades

O sistema oferece as seguintes rotas e funcionalidades:

| Rota | Método | Descrição |
|------|--------|-----------|
| `/` | GET | Página inicial com sumário de opções |
| `/forms/createAcc.html` | GET | Formulário para criar uma nova conta |
| `/form` | POST | Endpoint para processar criação de conta |
| `/forms/login.html` | GET | Formulário de login |
| `/login` | POST | Endpoint para processar login |
| `/forms/updateAcc.html` | GET | Formulário para atualizar dados da conta |
| `/updateAccount` | POST | Endpoint para processar atualização de conta |
| `/forms/deleteAcc.html` | GET | Formulário para deletar conta |
| `/deleteAccount` | POST | Endpoint para processar exclusão de conta |
| `/hello` | GET | Endpoint de teste |

## 🗂️ Estrutura do Projeto

```
AA1/
├── .env                          # Variáveis de ambiente (⚠️ NÃO fazer commit)
├── go.mod                        # Dependências Go
├── go.sum                        # Hash das dependências
├── README.md                     # Este arquivo
│
├── app/
│   ├── .env                      # ⭐ Arquivo de configuração (DEVE estar aqui)
│   ├── main.go                   # Arquivo principal do servidor
│   │
│   ├── handlers/                 # Tratadores de rotas
│   │   ├── helloHandler.go
│   │   ├── formHandler.go
│   │   ├── loginHandler.go
│   │   ├── updateAccountHandler.go
│   │   └── deleteAccountHandler.go
│   │
│   ├── utils/                    # Utilitários
│   │   ├── connDB.go             # Conexão com banco de dados
│   │   ├── DB.go                 # Inicialização do banco
│   │   ├── validateUser.go       # Validação de usuários
│   │   ├── getUserByEmail.go     # Buscar usuário
│   │   ├── dellUser.go           # Deletar usuário
│   │   ├── updateUser.go         # Atualizar usuário
│   │   ├── encrypt.go            # Criptografia de senhas
│   │
│   └── static/                   # Arquivos estáticos (Frontend)
│       ├── index.html            # Página inicial
│       ├── profile.html          # Página de perfil
│       │
│       ├── forms/                # Formulários HTML
│       │   ├── createAcc.html
│       │   ├── login.html
│       │   ├── updateAcc.html
│       │   └── deleteAcc.html
│       │
│       ├── style/                # Folhas de estilo CSS
│       │   ├── index.style.css
│       │   ├── login.style.css
│       │   ├── CreateAcc.style.css
│       │   ├── updateAcc.style.css
│       │   ├── deleteAcc.style.css
│       │   └── profile.style.css
│       │
│       └── pictures/             # Imagens
│           └── logo.svg
│
└── pictures/                     # Imagens adicionais (não utilizada)
```

## 🔧 Troubleshooting

### Erro: "Erro ao carregar o arquivo .env"

**Solução:** Certifique-se de que o arquivo `.env` está na pasta `app/` e não na raiz.

```
❌ Errado: AA1/.env
✅ Correto: AA1/app/.env
```

### Erro: "Porta 3000 já está em uso"

**Solução:** Feche outros aplicativos que estão usando a porta 3000 ou modifique a porta no arquivo `app/main.go`:

```go
port := "8080"  // Altere para outra porta
```

### Erro: "Connection refused" ao banco de dados

**Solução:** Verifique se:
1. PostgreSQL está rodando
2. Credenciais no `.env` estão corretas
3. O banco de dados `saude_ex` foi criado

### Erro: "404 page not found" ao acessar o site

**Solução:** Certifique-se de que:
1. A pasta `static` está dentro de `app/`
2. O servidor está rodando corretamente

## 📦 Dependências do Projeto

```
github.com/joho/godotenv v1.5.1    # Carregamento de .env
github.com/lib/pq v1.10.9          # Driver PostgreSQL
```

## 🔐 Segurança

- As senhas são criptografadas antes de serem armazenadas
- Valide sempre os inputs do usuário
- Nunca faça commit do arquivo `.env` (adicione ao `.gitignore`)

## 📝 Exemplo de Uso

1. Acesse `http://localhost:3000/`
2. Clique em "Crie sua conta"
3. Preencha o formulário com seus dados
4. Clique em "Concluir"
5. Retorne à página inicial e faça login com suas credenciais

## 👨‍💻 Desenvolvido por

Matheus Silva Moreira-Universidade Federal de Goiás (UFG) - Engenharia de Software
Introdução à Programação - AA1

## 📄 Licença

Este projeto é fornecido como material educacional.

---

**Dúvidas?** Verifique se seguiu todos os passos acima e consulte a seção Troubleshooting.
