# wallet system backend

A simple centralized wallet app backend system.

## Requirements 
### Functional
#### User
1. User can create an account
2. User can view their account info (wallet balance, transaction history)

#### Wallet
1. User can deposit money into their wallet
2. User can withdraw money from their wallet
3. User can transfer money to other user's wallet

### Non-functional 
1. Performance: 
2. Data Consistency: 
   - Ensure that all transaction operations (deposits, withdrawals, transfers) are atomic.
3. Maintainability & Extensibility:
   - The codebase should be well-documented, modular, and easy to understand. 

## Design

### Language
Golang are chosen as the programming language for this coding test due to the following reasons:
1. Concurrency: 
2. Performance: 
3. Strong Typing: 

### Storage
For this coding test, in-memory storage would used to store user data and transactions.
However, an interface `Storage` were provided that can be easily replaced with a more robust database system like PostgreSQL.

## Setup

## Work plan

## Review and Improve
