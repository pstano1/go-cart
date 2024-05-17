#!/bin/bash

# Adjust those variables before runnning the script
username="admin"
password="Admin1234mat@"
URL="http://localhost:8000/api"

customerTag="dev"
tokenType="Bearer"

interpretStatus() {
    statusCode=$1
    message=$2
    
    if [[ $statusCode -ge 200 && $statusCode -lt 300 ]]; then
        echo -e "\033[0;32m Success: $message"
    else
        echo -e "\033[0;31m Failed: $message"
    fi
}

# /customer/id/{tag}
# gets customerId based on a provided tag
customerId=$(curl -s -X GET "$URL/customer/id/${customerTag}" | jq -r '.id')
statusCode=$(curl -s -o /dev/null -w "%{http_code}" -X GET "$URL/customer/id/${customerTag}")
interpretStatus "$statusCode" "/customer/id/${customerTag}"

# /user/signin
# sign user in
payload=$(cat <<EOF
{
  "username": "$username",
  "password": "$password"
}
EOF
)
sessionToken=$(curl -s -X POST \
    -H "Content-Type: application/json" \
    -d "$payload" \
    "$URL/user/signin" | jq -r '.sessionToken')

statusCode=$(curl -s -o /dev/null -w "%{http_code}" -X POST \
    -H "Content-Type: application/json" \
    -d "$payload" \
    "$URL/user/signin")

interpretStatus "$statusCode" "/user/signin"

# /user/refresh
# refresh session token
statusCode=$(curl -s -o /dev/null -w "%{http_code}" -X POST \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $sessionToken" \
    "$URL/user/refresh")

interpretStatus "$statusCode" "/user/refresh"

# POST /product/category
# Create a product category
payload=$(cat <<EOF
{
  "name": "$(openssl rand -hex 8)",
  "customerId": "$customerId"
}
EOF
)

response=$(curl -s -w "\n%{http_code}" -X POST \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $sessionToken" \
    -d "$payload" \
    "$URL/product/category")

body=$(echo "$response" | sed '$d')
categoryId=$(echo "$body" | jq -r '.id')
statusCode=$(echo "$response" | tail -n 1)

interpretStatus "$statusCode" "POST /product/category"

# GET /product/category
# Get product category(ies)
statusCode=$(curl -s -o /dev/null -w "%{http_code}" -X GET \
    -H "Content-Type: application/json" \
    "$URL/product/category")

interpretStatus "$statusCode" "GET /product/category"

# PUT /product/category
# Update a product category
payload=$(cat <<EOF
{
  "id": "$categoryId",
  "name": "$(openssl rand -hex 8)",
  "customerId": "$customerId"
}
EOF
)

response=$(curl -s -w "\n%{http_code}" -X PUT \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $sessionToken" \
    -d "$payload" \
    "$URL/product/category")

statusCode=$(echo "$response" | tail -n 1)

interpretStatus "$statusCode" "PUT /product/category"

# GET /product
# Get product(s)
statusCode=$(curl -s -o /dev/null -w "%{http_code}" -X GET \
    -H "Content-Type: application/json" \
    "$URL/product/")

interpretStatus "$statusCode" "GET /product"

# POST /product
# Create a product
payload=$(cat <<EOF
{
  "name": "product #0",
  "categories": [
    "test-1"
  ],
  "prices": {
    "PLN": 20,
    "EUR": 4.5
  },
  "descriptions": {
    "PL": "Lorem ipsum...",
    "EN": "Lorem ipsum..."
  },
  "customerId": "$customerId"
}
EOF
)

response=$(curl -s -w "\n%{http_code}" -X POST \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $sessionToken" \
    -d "$payload" \
    "$URL/product/")

body=$(echo "$response" | sed '$d')
productId=$(echo "$body" | jq -r '.id')
statusCode=$(echo "$response" | tail -n 1)

interpretStatus "$statusCode" "POST /product"

# PUT /product
# Update a product
payload=$(cat <<EOF
{
  "name": "product #123",
  "categories": [
    "test",
    "test-1"
  ],
  "prices": {
    "PLN": 20,
    "EUR": 4.5
  },
  "descriptions": {
    "PL": "Lorem ipsum...",
    "EN": "Lorem ipsum..."
  },
  "customerId": "$customerId",
  "id": "$productId"
}
EOF
)

response=$(curl -s -w "\n%{http_code}" -X PUT \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $sessionToken" \
    -d "$payload" \
    "$URL/product/")

statusCode=$(echo "$response" | tail -n 1)

interpretStatus "$statusCode" "PUT /product"

# DELETE /product
# DELETE a product
statusCode=$(curl -s -o /dev/null -w "%{http_code}" -X DELETE \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $sessionToken" \
    "$URL/product/$productId?customerId=$customerId")

interpretStatus "$statusCode" "DELETE /product/{id}"

# DELETE /product/category
# DELETE a product category
statusCode=$(curl -s -o /dev/null -w "%{http_code}" -X DELETE \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $sessionToken" \
    "$URL/product/category/$categoryId?customerId=$customerId")

interpretStatus "$statusCode" "DELETE /product/category/{id}"

# POST /coupon
# Create a coupon
payload=$(cat <<EOF
{
  "promoCode": "TESTCOUPON",
  "amount": 20,
  "customerId": "$customerId"
}
EOF
)

response=$(curl -s -w "\n%{http_code}" -X POST \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $sessionToken" \
    -d "$payload" \
    "$URL/coupon/")

body=$(echo "$response" | sed '$d')
couponId=$(echo "$body" | jq -r '.id')
statusCode=$(echo "$response" | tail -n 1)

interpretStatus "$statusCode" "POST /coupon"

# GET /coupon
# Get coupon(s)
statusCode=$(curl -s -o /dev/null -w "%{http_code}" -X GET \
    -H "Content-Type: application/json" \
    "$URL/coupon/")

interpretStatus "$statusCode" "GET /coupon"

# PUT /coupon
# Update a coupon
payload=$(cat <<EOF
{
  "id": "$couponId",
  "promoCode": "TESTCOUPON",
  "amount": 20,
  "customerId": "$customerId"
}
EOF
)

statusCode=$(curl -s -o /dev/null -w "%{http_code}" -X PUT \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $sessionToken" \
    -d "$payload" \
    "$URL/coupon/")

interpretStatus "$statusCode" "PUT /coupon"

# DELETE /coupon/{id}
# Delete a coupon
statusCode=$(curl -s -o /dev/null -w "%{http_code}" -X DELETE \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $sessionToken" \
    "$URL/coupon/$couponId?customerId=$customerId")

interpretStatus "$statusCode" "DELETE /coupon/{id}"

# POST /order
# Create an order
payload=$(cat <<EOF
{
  "customerId": "$customerId",
  "totalCost": 100.00,
  "currency": "PLN",
  "country": "PL",
  "city": "Warszawa",
  "postalCode": "00-902",
  "address": "ul. Wiejska 4",
  "basket": {
    "test-id": {
      "price":    50,
      "currency": "PLN",
      "quantity": 1,
      "name":     "Product #0"
    },
    "test-id-1": {
      "price":    50,
      "currency": "PLN",
      "quantity": 1,
      "name":     "Product #1"
    }
  }
}
EOF
)

response=$(curl -s -w "\n%{http_code}" -X POST \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $sessionToken" \
    -d "$payload" \
    "$URL/order/")

body=$(echo "$response" | sed '$d')
orderId=$(echo "$body" | jq -r '.id')
statusCode=$(echo "$response" | tail -n 1)

interpretStatus "$statusCode" "POST /order"

# GET /order
# Get order(s)
statusCode=$(curl -s -o /dev/null -w "%{http_code}" -X GET \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $sessionToken" \
    "$URL/order/?customerId=$customerId")

interpretStatus "$statusCode" "GET /order"

# PUT /order
# Update an order
payload=$(cat <<EOF
{
  "id": "$orderId",
  "customerId": "$customerId",
  "totalCost": 100.00,
  "currency": "PLN",
  "country": "PL",
  "city": "Warszawa",
  "postalCode": "00-902",
  "address": "ul. Wiejska 4",
  "status": "paid",
  "basket": {
    "test-id": {
      "price":    50,
      "currency": "PLN",
      "quantity": 1,
      "name":     "Product #0"
    },
    "test-id-1": {
      "price":    50,
      "currency": "PLN",
      "quantity": 1,
      "name":     "Product #1"
    }
  }
}
EOF
)

statusCode=$(curl -s -o /dev/null -w "%{http_code}" -X PUT \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $sessionToken" \
    -d "$payload" \
    "$URL/order/")

interpretStatus "$statusCode" "PUT /order"

# DELETE /order/{id}
# Delete an order
statusCode=$(curl -s -o /dev/null -w "%{http_code}" -X DELETE \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $sessionToken" \
    "$URL/order/$orderId?customerId=$customerId")

interpretStatus "$statusCode" "DELETE /order/{id}"