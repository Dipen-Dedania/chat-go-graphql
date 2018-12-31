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
10. Install cockroachdb driver in go (go get github.com/lib/pq) for cockroachdb driver
11. Install gorilla mux (go get github.com/gorilla/mux)
12. Install gqlgen (go get -u github.com/99designs/gqlgen) [I use gqlgen 1.0 beta version]
13. If there is a new vesrion of gqlgen than do (gqlgen -v) it will generate generated.go file,resolver.go file and models_gen.go file as per new version

## Package used

1. github.com/gorilla/mux  = As router carry handler request
2. github.com/lib/pq       = For cockroachdb
3. github.com/gorilla/websocket = For handling socket 
4. github.com/99designs/gqlgen = Graphql library 
5. github.com/99designs/gqlgen/handler

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

## Database 

- This example auto generate tables.
- Create tables as per configuration. 
    [You can find my database configuration file here](https://github.com/AkhilxNair/chat-go-graphql/blob/master/backendserver/app/dal/config.json)
- Change configuration as per your db 

### Features

1. Add user
2. Retrive user
3. Realtime add user
4. Realtime chatconversation

### Getting started

1. go run server/server.go
       (It 'll start server)
2. Open  http://localhost:8080/ and run below queries
3. Can test services backendside using graphql playground

    3(a): Real time add user listen
    ```
            subscription realtimeUseradd{
            userjoined{
                name
            }
        }
    ```
    3(b): AddUser
    ```
            mutation adduser($name: String!){
                joinUser(name: $name){
                    name
                }
            }
    ```
    3(c): Userlist
    ```      
            query userlist{
            users{
                id
                name
                createdAt
            }
        }
    ```
    3(d): RealTime chat listen
    ```
            subscription realtimechat($id: String!){
            messagePosted(id: $id){
                sender_name
                receiver_name
                message
                createdAt
            }
        }
    ```
    3(e): Chat conversation
    ```
            mutation addMessage($sender_name: String!,$receiver_name: String!,$message: String!){
            postMessage(sender_name: $sender_name,receiver_name: $receiver_name,message: $message){
                sender_name
                receiver_name
                message
            }
        }
    ```
    3(f): Retrive chatconversation between two users
    ```
            query allchats($sender_name: String!,$receiver_name: String!){
            chats(sender_name: $sender_name,receiver_name: $receiver_name){
                sender_name
                receiver_name
                message
                createdAt
            }
        }
    ```



@here, [You can find frontend code] (https://github.com/AkhilxNair/chat-go-graphql/tree/master/frontend)