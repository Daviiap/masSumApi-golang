# maxSum

Api GraphQL com uma única mutation: maxSum.
Onde, dada uma lista de números, retorna as posições de uma lista de
números que possuem a maior soma obtida a partir de uma sub-lista contínua não vazia.

## Installation

### Sem Docker

Para rodar o projeto diretamento no seu SO, basta rodar o comando:

```bash
# modo de desenvolvimento:
go run server.go
# modo de produção
go build server.go && ./server
```

### Usando Docker

Para rodar o projeto com o `Docker` basta ter o `docker` e o `docker-compose` instalados e rodar o comando:

```bash
# modo de produção:
docker-compose up --build
```

## Testes

Para rodar os testes, basta, de dentro da pasta `__tests__` rodar o comando:

```bash
go test 
```

## Usando a API

Uma vez rodando o projeto, ele irá rodar na porta `8080`. Basta acessar `http://localhost:8080` para abrir o playground do GraphQL.

### Request

A api possui somente uma mutation `maxSum`.

#### `Exemplo`

`Request:`

```graphql
mutation {
	maxSum(list: [1, 2, -1, 3, 4, 5, -10]) {
		sum
		positions
	}
}
```

`Response:`

```json
{
	"data": {
		"maxSum": {
			"sum": 14,
			"positions": [1, 2, 3, 4, 5, 6]
		}
	}
}
```
