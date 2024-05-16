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

couponId=$(curl -s -X POST \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $sessionToken" \
    -d "$payload" \
    "$URL/coupon/" | jq -r '.id')

statusCode=$(curl -s -o /dev/null -w "%{http_code}" -X POST \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $sessionToken" \
    -d "$payload" \
    "$URL/coupon/")

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