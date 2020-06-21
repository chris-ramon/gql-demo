# gql-demo

gql-demo is working demo for exploring the usage of libraries that enables a GraphQL server with database access that can mantain type safety and leverage code generation.

- Custom ORM: [`sqlboiler`](https://github.com/volatiletech/sqlboiler)
- GraphQL: [`gqlgen`](https://github.com/99designs/gqlgen) 
- GraphQL UI: [`graphql-playground`](https://github.com/prisma-labs/graphql-playground/tree/master/packages/graphql-playground-react)

## Getting Started

### Development Machine Prerequisites

- [mysql](https://www.mysql.com/downloads/)
- [sqlboiler](https://github.com/volatiletech/sqlboiler#download)
- [gqlgen](https://gqlgen.com/getting-started/#setup-project)

### Running Development Environment
```bash
bin/start.sh
```
### Demo

#### Query
![query](https://user-images.githubusercontent.com/1000404/85218395-8736c500-b35f-11ea-9003-6356caf35d71.gif)


#### Subscription
![subscription](https://user-images.githubusercontent.com/1000404/85218515-d7625700-b360-11ea-8591-08313403712d.gif)

