# PAPER ID CHALLENGE

## study case

Topic: The user has a balance in the application wallet and the balance wants to be disbursed.

- Write code in Golang
- Write API (only 1 endpoint) for disbursement case only
- User data and balances can be stored as hard coded or database

## requirement

- go 1.22.1

## dependencies

- fiber
- golang-jwt
- gorm
- gorm sqlite
- xendit
- decimal

## dev dependencies

- air

## to start

1. register to xendit & get API key with payout write access
2. create `.env` file & put values accordingly, including xendit's api key (example in `.env.example`)
3. on command line execute `air` command

# Docs

## architecture/decision

- Refer to rough C4 diagram [here](https://github.com/samm021/sitemate-challenge-node-cli)

## TODOs

- Fix xendit payout creation
- Improve DB trx on creating transaction
- Complete API for receiving xendit webhook
- Improve error handling
- Proper user's token validation
