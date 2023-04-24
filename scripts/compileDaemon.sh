#!/bin/bash

# check if CompileDaemon is installed
if ! command -v CompileDaemon &> /dev/null
then
    echo "CompileDaemon could not be found"
    echo "installing CompileDaemon"
    go get github.com/githubnemo/CompileDaemon
    go install github.com/githubnemo/CompileDaemon
fi 


CompileDaemon -exclude=".env,.gitignore,*_test.go" -include="*.html" -build="env GOOS=linux CGO_ENABLED=0 go build -o ./bin/$1 ./cmd/admin" -command="./bin/$1"