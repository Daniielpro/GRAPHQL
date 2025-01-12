# Simple GraphQL API in Go

This project is a simple implementation of a GraphQL API using Go (`golang`). It allows for basic queries such as `hello`.
---

## Features

- **GraphQL Support**: Perform queries to retrieve data efficiently.
- **Basic resolvers**: Implement resolvers for `hello`
---

## Project Structure

```plaintext
GRAPHQL/
├── main.go # Main server file
├── go.mod # Go module dependencies
├── go.sum # Dependency checksum
└── README.md # Documentation file
```
## Prerequisites
Before running the project, make sure you have installed:

 Go (version 1.19 or higher)

## Instructions for Using the Project
1. Clone the Repository
Clone the project to your local machine with:

 git clone https://github.com/Daniielpro/GRAPHQL.git

 cd simple-graphql-go

2. Install Dependencies
Initialize and download the dependencies Required dependencies:
 
 go mod init

3. Run the Server
Start the GraphQL server:
 
 go run main.go

## How to Make Queries
Interact with the API
You can use tools like curl, Postman, or any GraphQL client to make queries.

Query Example:

Invoke-WebRequest -Uri "http://localhost:8080/graphql" `
                  -Method POST `
                  -ContentType "application/json" `
                  -Body '{"query": "{ hello }"}'

## Troubleshooting
Common Problems
Module not found: Make sure you have initialized the go.mod file by running:

 go mod init <module-name>

## Author
EDWIN PROAÑO
GitHub: Daniielpro10


