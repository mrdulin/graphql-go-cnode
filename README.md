# graphql-go-cnode

A GraphQL APIs for https://cnodejs.org community built with [graphql-go](https://github.com/graphql-go/graphql) package.

### Environment variables

Put following environment variables into `.env` file

```
API_BASE_URL=https://cnodejs.org/api/v1
GRAPHQL_PATH=/graphql
PORT=3001
```

### Usage

Start GraphQL API server:
```
make run
```

build to binary:
```
make build
```

Run unit tests:
```
make unit_test
```

