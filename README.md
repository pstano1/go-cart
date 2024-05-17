# go-cart

### Content

- [TO-DO](#to-do)
- [Running locally](#running-locally)
- [Scripts](#scripts)
- [Licence](#licence)

### TO-DO:

- frontend:
  - [] refresh token when expiry is near;
  - [] style;
  - [] add manuals;
  - [] add employees UI;
  - [] translations;
  - [] handle errors;
  - [] add loaders;
  - [] validation;
  - [] toasts;
- backend:
  - [x] orders;
  - [x] coupons;
  - [] permissions;
  - [] tests;
  - [] errors;
  - [] add comments;
  - [] images storage;
  - [] add price history;
- other:
  - [] swagger;
  - [] makefile;
  - [] .sh script;
  - [] README.md;

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

|Flag|Action|
|----|------|
|migrate|creates tables|
|create-permission|inserts predefined permissions into table|
|flush|drops all the tables & creatres new ones|

```console
go run ./scripts/manage.go <flag>
```

### Licence
