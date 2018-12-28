# Server side (Go-Cockroachdb-Graphql-Gqlgen) 

## Prerequisite

1. [Install golang](https://golang.org/dl/) [I use go 1.11 vesrion]
2. Setup GOPATH [Link1](https://golang.org/doc/code.html#GOPATH) and [Link2](https://github.com/golang/go/wiki/GOPATH)
3. [Install cockroachdb](https://www.cockroachlabs.com/docs/stable/install-cockroachdb-windows.html)
4. Check Binary file of cockroachdb : `.\cockroach.exe` run in cmd
5. Start node in cockroachdb : `cockroach start --insecure` run in cmd where cockroachdb is intalled
6. Test cluster : `cockroach sql --insecure` run in cmd where cockroachdb is intalled
7. Create database : create database dbname;(here I took chatApp)
8. Show database : show datbases;
9. Clone repo in gopath src folder (Can't run outside gopath because gqlgen not support outside gopath)
10. Install gqlgen (go get -u github.com/99designs/gqlgen) [I use gqlgen 1.0 beta version]
11. If there is a new vesrion of gqlgen than do (gqlgen -v) 

## Folder structure of server side

    backendserver
        |--- app
        |     |--- dal
        |     |     |-- config.json         // database configuration file
        |     |     |-- connection.go       // Database connection file
        |     |     |-- dbConfig.go         // Load config file
        |     |--- graph
        |     |     |-- generated.go        // Auto generated file which handle all the methods 
        |     |--- model
        |     |     |-- models_gen.go       // Model of chat application
        |     |--- resolver
        |     |     |-- chatResolver.go     // Chat resolver for operation relate to chat
        |     |     |-- userResolver.go     // User resolver for operation relate to user
        |     |     |-- resolver.go
        |--- server
        |     |--server.go                  // main execution file
        |--- gqlgen.yml                     // Show all mapping
        |--- schema.graphql                 // Schema for chat app