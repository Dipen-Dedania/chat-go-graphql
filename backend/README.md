# README file for backend

## Prerequisite
1. [Install golang](https://golang.org/dl/)
2. Setup GOPATH [Link1](https://golang.org/doc/code.html#GOPATH) and [Link2](https://github.com/golang/go/wiki/GOPATH)
3. [Install cockroachdb](https://www.cockroachlabs.com/docs/stable/install-cockroachdb-windows.html)
4. Check Binary file of cockroachdb : `.\cockroach.exe` run in cmd
5. Start node in cockroachdb : `cockroach start --insecure` run in cmd where cockroachdb is intalled
6. Test cluster : `cockroach sql --insecure` run in cmd where cockroachdb is intalled
7. create database : create database dbname;
8. test database : show databases;
8. [Install gqlgen] (go get )

## Folder structure

    backend
        |--- dal
        |      |-- config
        |      |    |-- dbConfig.go // configuration method and model for database
        |      |-- config.json      // config file contain database connection configuration
        |      |-- connection.go    // Connection of database 
        |--- server
        |      |-- server.go        // main file which executed
        |--- generated.go
        |--- gqlgen.yaml
        |--- model_gen.go
        |--- resolver.go
        |--- schema.graphql         // schema for application
        |--- userresolver.go        // resolver relate to user
        |--- chatresolver .go       // resolver relate to chatconversation

