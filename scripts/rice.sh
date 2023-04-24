#!/bin/bash

# check if rice is installed
if ! command -v rice &> /dev/null
then
    echo "rice could not be found"
    echo "installing rice"
    go install github.com/GeertJohan/go.rice/rice@latest
fi 


if [ "$1" == "admin" ]; then
    cd internal/adminApp/interfaces/views
    rice embed-go
    # cd pkg
    # rice embed-go
    cd ../../../infrastructure/web
    rice embed-go

fi
if [ "$1" == "customer" ]; then
    cd cmd/customer 
    rice embed-go
fi