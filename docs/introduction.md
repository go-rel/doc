# Introduction

REL is golang orm-ish database layer for layered architecture. It's testable and comes with it's own test library. REL also features extendable query builder that allows you to write query using builder or plain sql.

## Features

- Testable repository with builtin reltest package.
- Seamless nested transactions.
- Elegant, yet extendable query builder with mix of syntactic sugar.
- Supports Eager loading.
- Composite Primary Key.
- Multi adapter.
- Soft Deletion.
- Pagination.
- Schema Migration.

## Install

```
go get github.com/go-rel/rel
```

## Why rel

### Easy to test

Most (if not all) orm for golang is written as a chainable API, meaning all of the query need to be called before performing actual action as a chain of method invocations. example:

```go
db.Where("id = ?", 1).First(&user)
```

Chainable api is very hard to be unit tested without writing a wrapper. One way to make it testable is to make an interface that also acts as a wrapper, which is usually ends up as its own repository package resides somewhere in your project:

```go
// mockable interface.
type UserRepository interface {
	Find(user *User, id int) error
}

// actual implementation
type userRepository struct {
	db *DB
}

func (ur userRepository) Find(user *User, id int) error {
	return db.Where("id = ?", 1).First(&user)
}
```

Compared to other orm, REL api is built with testability in mind. REL uses interface to define contract of every database query or execution, all while making a chainable query possible. The ultimate **goal of REL is to be your database package** without the needs of making your own wrapper.

### Seamless transactions

When starting a transaction using builtin `database/sql` package or other orm, special transaction instance that's similar to non-transaction instance is returned,
and all execution needs to be called using that instance to be in transaction:

```go
// start transaction
tx := db.Begin()

// do some database operations in the transaction (use 'tx' from this point, not 'db')
tx.Create(...)

// Or commit the transaction
tx.Commit()
```

Now what happens if you want to introduce a shared function, that might execute it's operation on a transaction or on a regular connection depending on the caller?
you may end up with a function that may not only accepts conventional `context` argument, but also a `tx` or connection argument.

```go
tx := db.Begin()
tx.Create(...)

// UpdateOtherEntity(txOrDb TxOrDbInterface, ...)
UpdateOtherEntity(tx, ...)

tx.Commit()
```

REL attempts to solve that problem by managing transaction state using context, the REL repository can decide whether to use transaction or not based on provided context, performs commit or rollback in case of exception, and even supports nested transaction (see: [transactions](/transactions)).