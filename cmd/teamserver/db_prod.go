// +build !dev

package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kcarretto/paragon/ent"
	"github.com/kcarretto/paragon/ent/migrate"
)

func getClient(ctx context.Context) *ent.Client {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	ipAddress := os.Getenv("DB_IP_ADDRESS")
	drv, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/ent", user, password, ipAddress))
	if err != nil {
		panic(err)
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	client := ent.NewClient(ent.Driver(drv))

	if err := client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
		panic(err)
	}

	return client
}
