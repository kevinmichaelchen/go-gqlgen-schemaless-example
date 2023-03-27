# go-gqlgen-schemaless-example

## Motivation

This example shows a GraphQL API (built in Go using the [gqlgen](https://github.com/99designs/gqlgen)
library), which uses a custom scalar to encode explicit fields that are 
deliberately omitted from the GraphQL schema.

See this [issue](https://github.com/99designs/gqlgen/issues/1758) for reference.

### Why would you want to do this?

Why would you ever do this? Isn't the whole point of GraphQL that types are 
strong and explicit?

Most of the time, they are. Sometimes, however, produce requirements aren't 
clear, and you want to avoid having to accumulate either a messy API or having
to break your API to clean it up.

## Getting started

### Start the server

Start the GraphQL server on port 8081:

```shell
make run
```

### Hit the server

Visit [localhost:8081](http://localhost:8081/) and run the following mutation:

```graphql
mutation CreateFoo(
  $attributes: Map!
) {
  createFoo(
    input: {
      attributes: $attributes
    }
  ) {
    name
  }
}
```

with variables:

```json
{
  "attributes": {"settings": {"toggle1": true}}
}
```