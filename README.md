# Agile Engine - Accounting notebook App
### Agile Engine Task test for interview


## Dependencies used
* Gomod for dependency management
* Gingonic for routing and web context
* Google UUID for UUID generation

## How To Run
```cd ./build```

```./accountnotebook```

Server will listen at localhost:8080

## Endpoints (cURL)

### Create Transaction

```
curl --location --request POST 'http://localhost:8080/accounting/account/transaction' \
--header 'Content-Type: application/json' \
--data-raw '{
    "type":"credit",
    "Amount": 500
}'
```

### List Transactions

```
curl --location --request GET 'http://localhost:8080/accounting/account/transaction' \
--header 'Content-Type: application/json' \
--data-raw ''
```

### Get Transaction
```
curl --location --request GET 'http://localhost:8080/accounting/account/transaction/99c54bde-40b2-11eb-941e-8c859056ebf9'
```

### Get Balance
```
curl --location --request GET 'http://localhost:8080/accounting/account/balance'
```



## Testing
Disclaimer: as in memory DB, tests need to be run in order (must not happen in real life, I know mocked services should be used)

Service has been tested. Tests can be found under src/api/service/impl/accountsservice_test.go

Coverage is at 44%, which would not be acceptable under real life conditions

## Disclaimer 2
I've been working with Java for the last 2 years, so sorry if any naming convention is not compliant

