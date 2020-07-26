## GraphQL started with postgres as the repository

### generate code after any schema change

`gqlgen generate`

then delete the resolver.schema.go file coz its compaints redeclared

### learnings

- implement middleware the right way
- use the environemnt variables
- graphql with generation using gql go
- how to validate jwt validation and get the json web keys for certifcate
- logging middleware


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