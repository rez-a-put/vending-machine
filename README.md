# Vending Machine
This repo contains a simple application using Go. This project uses a Vending Machine implementation as an example.

# Features
- Get Item List :
    Getting list of items
- Buy Items :
    To buy items based on inputted nominals
- Add Item :
    To add new item that want to be sold in the vending machine
- Modify Item :
    To change name or price of an item
- Remove Item :
    To remove an item from being shown when call get item list

# Installation
1. Clone the repository
    ```bash
        git clone https://github.com/rez-a-put/vending-machine.git
    ```
2. Change into project directory
    ```bash
        cd vending-machine
    ```
3. Set up your .env file based on .env.example
4. Set up your vendor folder
    ```bash
        go mod vendor
    ```
5. Install Go Migrate
    ```bash
        go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
    ```
6. Migrate tables
    ```bash
        migrate -database "dbdriver://dbuser:dbpass@tcp(127.0.0.1:3306)/dbname" -path "db/migrations" up
    ```

# Run the project
1. Open terminal
2. Go to project folder
3. Build application
    ```bash
        go build
    ```
4. Run application from terminal or run using go command
    ```bash
        ./vending-machine
    ```
    ```bash
        go run main.go
    ```

# Testing
1. Open terminal
2. Go to project folder
3. Run Go test
    ```bash
        go test -v ./...
    ```

# Contributing
1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a merge request