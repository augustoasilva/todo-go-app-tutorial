# TODO Go App Tutorial

> Um passo a passo para aprender Go construindo uma API REST de forma evolutiva e didática.

## Sobre o Projeto

Este repositório faz parte da série de artigos “[Construindo Sua Primeira API em Go do Zero](#)”, disponível no meu blog.
A ideia é construir juntos uma API de TODOs simples, mas evoluindo aos poucos para abordar conceitos, boas práticas e curiosidades
do ecossistema Go. Começamos com o básico e vamos aumentando a complexidade a cada novo artigo!

**Seja você iniciante ou alguém querendo ver como uma aplicação real evolui em Go, seja bem-vindo(a)!**

## Motivação

- Explorar as melhores práticas do desenvolvimento Go na vida real, de forma incremental.
- Mostrar que o padrão “MVC clássico” não se encaixa 100% em uma API Go, mas pode ser adaptado e depois
evoluído, até mesmo  se formos sair de uma aplicação de API para uma aplicação Fullstack.
- Criar um material didático e prático, fugindo dos tutoriais que venho encontrando.

## Arquitetura e Organização

A arquitetura inicial é uma releitura pessoal do MVC, adaptada ao mundo Go:

- **Model**: Representa a estrutura dos dados (TODO), incluindo regras de persistência/acesso.
- **Service**: Camada que centraliza as regras de negócio e faz o meio-campo entre o modelo e os handlers.
- **Handler**: Responsável por receber requisições HTTP e servir as respostas JSON.  
  (Em Go, handlers são mais idiomáticos que “controllers”).

Estrutura inicial de pastas:

```
todo-go-app-tutorial/
│
├── cmd/
│ └── api/ # main.go (ponto de entrada da aplicação)
├── internal/
│ └── todo/
│ ├── model.go # Modelos e acesso a dados
│ ├── service.go # Lógica de negócio
│ └── handler.go # Handlers (HTTP handlers)
├── pkg/ # Código reutilizável (ex: integração com BD)
├── go.mod
└── go.sum
```

> Detalhes sobre cada pasta e arquivo estão  primeiro artigo da série!

## Como Rodar

Pré-requisitos: [Go instalado](https://go.dev/dl/)

Depois basa clonar e executar o arquivo main

```sh
git clone https://github.com/seu-usuario/todo-go-app-tutorial.git
cd todo-go-app-tutorial/cmd/api
go run cmd/api/main.go
```

A API vai subir na porta 8080 e ai basta acessar usando Postman ou até mesmo o seu navegador.

## Evolução da Série
O projeto começa simples e vai ficando mais robusto e profissional à medida que os artigos evoluem.

Alguns temas que planejo de serem abordados:
- Boas práticas de organização de código em Go
- Evoluindo o projeto com rotas, validações, middlewares
- Introdução a testes em Go
- Persistência com banco de dados
- Deploy e melhores práticas de produção

Fique ligado(a) no blog para novidades e novas branches!

## Contribuindo
Sugestões, correções e dúvidas são bem-vindas!
Abra uma issue ou faça um pull request :rocket:

## Links Importantes:

### Artigos originais da série

- 


Feito com 💙 por Augusto Silva
