package main

import (
	"context"

	"github.com/go-rel/doc/examples/db/migrations"
	"github.com/go-rel/mysql"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/migrator"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var (
		ctx  = context.TODO()
		repo = rel.New(mysql.MustOpen("root@(source:3306)/rel_test?charset=utf8&parseTime=True&loc=Local"))
		m    = migrator.New(repo)
	)

	// Register migrations
	m.Register(20202806225100, migrations.MigrateCreateTodos, migrations.RollbackCreateTodos)

	// Run migrations
	m.Migrate(ctx)
	// OR:
	// m.Rollback(ctx)
}
