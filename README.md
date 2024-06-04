# go-cart

This repository contains a setup for a multi-store application. It includes a REST API and a minimal yet fully functional frontend for store management. 

### Content

- [Description & requirements](#description)
- [TO-DO](#to-do)
- [Running locally](#running-locally)
- [Scripts](#scripts)
- [Licence](#licence)

### Description

To enable this API, a gRPC service is required. The goal is for the application to host multiple stores using shared resources. To achieve this, the API leverages another application that provides utilities connected to customers. A simple implementation used during development can be found [here](https://github.com/pstano1/customer-api).

![basic-api-schema](https://github.com/pstano1/diagrams/blob/trunk/pkg/diagrams/go-cart/assets/customerAPISchema.png?raw=true)

### TO-DO:

- frontend:
  - [x] refresh token when expiry is near;
  - [ ] style;
  - [ ] add manuals;
  - [ ] add employees UI;
  - [x] add coupons UI;
  - [ ] add orders UI;
  - [x] translations;
  - [ ] handle errors;
  - [ ] add loaders;
  - [ ] validation;
  - [ ] toasts;
- backend:
  - [x] orders;
  - [x] coupons;
  - [x] permissions;
  - [ ] tests;
  - [x] errors;
  - [x] add goDoc comments;
  - [ ] images storage;
  - [x] add price history;
  - [x] payment processing (stripe);
  - [x] localized product names;
  - [x] BUG: server crashes when user doesn't have any permission;
  - [x] scrape exchange rates in 24h interval (API call in the end);
  - [x] proccessing of coupons to orders;
  - [x] add TAXID to orders;
  - [x] add unit (percentage/currency) to coupon;
  - [ ] invoice/receipt generation (probably can be done by stripe);
  - [ ] mail sending;
  - [ ] webhooks for changing status to paid when stripe gets payment;
- other:
  - [x] swagger;
  - [ ] makefile;
  - [x] .sh script;
  - [ ] README.md;
  - [ ] dockerfile;
  - [ ] investigate dependabot's alerts;

### Running locally

```console
git clone git@github.com:pstano1/go-cart.git
cd ./go-cart
```

### Scripts

#### endpoints testing script

The script runs `cURL` requests to all enpoints end prints result based on received status code. Please note that although all created entries are deleted, API performs soft delete. To delete all artifacts from database run [this script](#database-managing-script)

```console
./scripts/enpointsTests.sh

# result
...
Success: POST /order
Failed: GET /order
...
```

#### database managing script

| Flag              | Action                                    |
| ----------------- | ----------------------------------------- |
| migrate           | creates tables                            |
| create-permission | inserts predefined permissions into table |
| flush             | drops all the tables & creatres new ones  |

```console
go run ./scripts/manage.go <flag>
```

### Licence

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT)
