#!/bin/bash

# Test script for import-users endpoint
# Make sure the auth-service is running on port 8080

echo "Testing POST /import-users endpoint..."

# Test with valid CSV file
echo "1. Testing with valid CSV file..."
curl -X POST \
  -F "file=@test_data/sample_users.csv" \
  http://localhost:8080/import-users

echo -e "\n\n2. Testing with invalid file type..."
curl -X POST \
  -F "file=@test_data/sample_users.csv" \
  -H "Content-Type: application/json" \
  http://localhost:8080/import-users

echo -e "\n\n3. Testing without file..."
curl -X POST \
  http://localhost:8080/import-users

echo -e "\n\n4. Testing with GET method (should fail)..."
curl -X GET \
  http://localhost:8080/import-users

echo -e "\n\nTest completed!" 