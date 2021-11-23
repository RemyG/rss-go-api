package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

var db *bun.DB
var ctx context.Context

func main() {

	ctx = context.Background()

	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(psqlInfo)))
	db = bun.NewDB(pgdb, pgdialect.New())

	defer db.Close()

	router := gin.Default()

	categories := router.Group("/categories")
	{
		categories.GET("", getCategories)
	}

	feeds := router.Group("/feeds")
	{
		feeds.GET("", getFeeds)
		feeds.GET("/:id", getFeed)
		feeds.GET("/:id/entries", getFeedEntries)
	}

	entries := router.Group("/entries")
	{
		entries.GET("/:id", getEntry)
		entries.PATCH("/:id", patchEntry)
	}

	router.Run("localhost:8080")
}
