# Adapters

Rel uses adapter in order to generate and execute query to a database, below is the list of available adapters currently supported.

| Adapter    | Package                                | Godoc                                                                                                                                     |
| ---------- | -------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------- |
| MySQL      | github.com/go-rel/mysql    | [![GoDoc](https://godoc.org/github.com/go-rel/mysql?status.svg)](https://godoc.org/github.com/go-rel/mysql)       |
| PostgreSQL | github.com/go-rel/postgres | [![GoDoc](https://godoc.org/github.com/go-rel/postgres?status.svg)](https://godoc.org/github.com/go-rel/postgres) |
| MSSQL      | github.com/go-rel/mssql                | [![GoDoc](https://godoc.org/github.com/go-rel/mssql?status.svg)](https://godoc.org/github.com/go-rel/mssql)                               |
| SQLite3    | github.com/go-rel/sqlite3  | [![GoDoc](https://godoc.org/github.com/go-rel/sqlite3?status.svg)](https://godoc.org/github.com/go-rel/sqlite3)   |

## Using Primary Replica Connections

REL Read Write separation for primary replica connections by using intermediary adapters.

```go
package main

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-rel/primaryreplica"
	"github.com/go-rel/mysql"
	"github.com/go-rel/rel"
)
func main() {
	// open mysql connections.
	// note: `clientFoundRows=true` is required for update and delete to works correctly.
	adapter := primaryreplica.New(
		mysql.MustOpen("root@(source:3306)/rel_test?charset=utf8&parseTime=True&loc=Local"),
		mysql.MustOpen("root@(replica1:3306)/rel_test?charset=utf8&parseTime=True&loc=Local"),
		mysql.MustOpen("root@(replica2:3306)/rel_test?charset=utf8&parseTime=True&loc=Local"),
	)
	defer adapter.Close()

	// initialize REL's repo.
	repo := rel.New(adapter)
	repo.Ping(context.TODO())
}
```

### Load Balancing of Replicas

REL only implements a very primitive load balancing for multiple replicas. For large scale application we recommend you to use external load balancing solution.
