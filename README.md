## Learn GraphQL with Golang

### Installation

We need to install the graphql package `github.com/graph-gophers/graphql-go`. Install the package using `go get -u github.com/graph-gophers/graphql-go`

### Setting up the schema
We will build out the schema for our trivial todo app GraphQL API. Therefore we will be defining our schema and resolvers in `schema.go` and `resolver.go` respectively. The `graphql-go` will parse and validate the schema and resolvers for us. It will panic in case of any errors.

### Setting up the server

We will setup the graphql endpoints and register their respective handlers.

### Test
Run the main binary found inside `cmd` directory, it wil start a GraphQL API server on port `:8080`. You can use the sample curl request to test that it is responding to graphql queries as expected.

```Bash
 curl -X POST -H "Content-Type: application/json" --data '{ "query": "{ tasks {id name description startDate endDate done} }" }' http://localhost:8080/tasks
```
