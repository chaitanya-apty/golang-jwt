# JWT Generation using Golang

Basic HTTP Server which generates token and validates session token

### Prerequisites
 - Ensure that Go is installed
 - Basics of GO

### Installing
- Pull repo
- Open terminal from project root directory
- Run the Server (Listening @ 8081)

        > go run main.go
    * Use API tool like POST MAN to test
    * Project has two end points(GET/POST doesnt matter, its a demo)
        - /login (Hit this API to fetch token)
        - /validate
            * Request Body should contain {token: "PLACE_TOKEN_HERE"}
            * On Success Youll get Verified as Status
## Built With

* [Go Lang](https://golang.org/) - Programming Language
* [JWT library](https://godoc.org/github.com/dgrijalva/jwt-go#example-Parse--Hmac) - Creating Tokens and Validate Session

## Authors

* **Chaitanya Kumar** - [Github](https://github.com/chaitanya-apty)
