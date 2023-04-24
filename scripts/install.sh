#!/bin/bash

echo "Installing..."
# check if there .env file exists
if [ ! -f .env ]; then
    echo "No .env file found. Creating one..."
    cp .env.example .env
    
    # then ask database details and save them to .env file
    echo "Please enter your database details:"
    read -p "Database name: " dbname
    read -p "Database user: " dbuser
    read -p "Database password: " dbpass
    read -p "Database host: " dbhost
    read -p "Database port: " dbport
    # sed append to .env file if it exists
    sed -i "/DB_DATABASE=/c\DB_DATABASE=$dbname" .env
    sed -i "/DB_USERNAME=/c\DB_USERNAME=$dbuser" .env
    sed -i "/DB_PASSWORD=/c\DB_PASSWORD=$dbpass" .env
    sed -i "/DB_HOST=/c\DB_HOST=$dbhost" .env
    sed -i "/DB_PORT=/c\DB_PORT=$dbport" .env

    # then ask for admin details and save them to .env file
    echo "Please enter your admin details:"
    read -p "Admin username: " adminuser
    read -p "Admin password: " adminpass
    
    # sed append to .env file if it exists
    sed -i "/ADMIN_USERNAME=/c\ADMIN_USERNAME=$adminuser" .env
    sed -i "/ADMIN_PASSWORD=/c\ADMIN_PASSWORD=$adminpass" .env


fi