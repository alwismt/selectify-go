#!/bin/bash
if [ -d "internal/$1/domains/$2" ]; then
	echo "Domain $2 already exists"
	exit 1
fi
mkdir -p "internal/$1/domains/$2/entities"
echo "package entity" > "internal/$1/domains/$2/entities/$2.go"

mkdir -p "internal/$1/domains/$2/repositories"
echo "package repository" > "internal/$1/domains/$2/repositories/$2_repository.go"

mkdir -p "internal/$1/domains/$2/services"
echo "package service" > "internal/$1/domains/$2/services/$2_service.go"

echo "Domain $2 created"