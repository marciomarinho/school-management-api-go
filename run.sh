#!/bin/bash

# Start up Docker Compose
echo "Starting Docker Compose..."
docker-compose up -d

# Wait for DynamoDB Local to start
echo "Waiting for DynamoDB Local to start..."
sleep 5

# Delete existing tables
echo "Deleting existing tables..."
aws dynamodb delete-table --table-name students --endpoint-url http://localhost:8000
aws dynamodb delete-table --table-name courses --endpoint-url http://localhost:8000

# Wait for tables to be deleted
echo "Waiting for tables to be deleted..."
sleep 5

# Create students table
echo "Creating students table..."
aws dynamodb create-table \
    --table-name students \
    --attribute-definitions AttributeName=ID,AttributeType=S \
    --key-schema AttributeName=ID,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --endpoint-url http://localhost:8000

# Create courses table
echo "Creating courses table..."
aws dynamodb create-table \
    --table-name courses \
    --attribute-definitions AttributeName=ID,AttributeType=S \
    --key-schema AttributeName=ID,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --endpoint-url http://localhost:8000

echo "Tables created successfully!"

go run .