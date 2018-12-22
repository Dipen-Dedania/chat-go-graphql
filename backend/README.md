# README file for backend

## Folder structure

    backend
        |--- api
        |      |-- main.go
        |--- chatConversation
        |      |-- chatInput.go     // Mutaion & Subscription for chat conversation
        |      |-- chatModel.go     // Chat model
        |      |-- chatQueries.go   // Queries realted to chat conversation
        |      |-- chatResolver.go  // Get chat convesation data as per sender_name and receiver_name
        |--- dal
        |      |-- config
        |      |    |-- dbConfig.go // configuration method and model for database
        |      |-- config.json      // config file contain database connection configuration
        |      |-- connection.go    // Connection of database
        |      |-- dbMiddleware.go  
        |--- secrity
        |      |-- token_generated.go
        |      |-- authentication.go
        |--- server
        |      |-- server.go        // main file which executed
        |--- user
        |      |-- userInput.go     // Mutaion & Subscription for user 
        |      |-- userModel.go     // User model
        |      |-- userQueries.go   // Queries realted to user 
        |      |-- userResolver.go  // Get all data of user
        |---generated.go
        |--- gqlgen.yaml
        |--- model_gen.go
        |--- resolver.go
        |--- schema.graphql         // schema for application