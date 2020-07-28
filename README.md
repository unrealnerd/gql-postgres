## GraphQL starter with postgres as the repository

### generate code after any schema change

`gqlgen generate`

OR for using `go generate`, add this line in resolver.go `//go:generate go run github.com/99designs/gqlgen` and then generate recursively

`go generate ./...`

then implement the method that talk to repository and call that from schema.resolvers.go

### learnings

- implement middleware the right way
- use the environemnt variables
- graphql with generation using gql go
- how to validate jwt validation and get the json web keys for certifcate
- logging middleware
- handling optional graphql arguments
- error handling


### vscode settings

```jsonc
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}",
            "env": {
                "inventoryDBConnectionString":"postgres://kammfkje:<password>@satao.db.elephantsql.com:5432/kammfkje"
            },
            "args": []
        }
    ]
}
```