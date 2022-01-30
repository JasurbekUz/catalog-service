#!/bin/bash

export POSTGRES_USER=jasurbek
export POSTGRES_PASSWORD=1001
export POSTGRES_DATABASE=catalogdb

migrate -database "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@localhost:5432/$POSTGRES_DATABASE?sslmode=disable" -path="../migrations" up