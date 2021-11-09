# maxSum

Api GraphQL com uma única mutation: maxSum.
Onde, dada uma lista de números, retorna as posições de uma lista de
números que possuem a maior soma obtida a partir de uma sub-lista contínua não vazia.

## Installation

### Sem Docker

Para rodar o projeto diretamento no seu SO, basta rodar o comando:

```bash
# modo de desenvolvimento:
yarn && yarn dev
# modo de produção
yarn \
&& yarn build \
&& rm -rf node_modules \
&& yarn --production \
&& yarn start
```

### Usando Docker

Para rodar o projeto com o `Docker` basta ter o `docker` e o `docker-compose` instalados e rodar o comando:

```bash
# modo de desenvolvimento:
docker-compose -f docker-compose.dev.yaml up --build
# modo de produção:
docker-compose up --build
```

## Testes

Para rodar os testes, basta rodar o comando:

```bash
yarn test
```

## Usando a API

Uma vez rodando o projeto, ele irá rodar na porta `4000`. Basta acessar `http://localhost:4000` para abrir o playground do GraphQL.

### Request

A api possui somente uma mutation `maxSuum`.

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
