package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getCategories(c *gin.Context) {

	embed := c.QueryArray("embed")
	categories := make([]*DBCategory, 0)
	_, feeds := Find(embed, "feeds")
	if feeds {
		if err := db.NewSelect().Model(&categories).Relation("Feeds").Order("cat_order asc").Scan(ctx); err != nil {
			panic(err)
		}
	} else {
		if err := db.NewSelect().Model(&categories).Order("cat_order asc").Scan(ctx); err != nil {
			panic(err)
		}
	}

	c.IndentedJSON(http.StatusOK, mapCategories(categories))
}

func getFeeds(c *gin.Context) {

	feeds := make([]*DBFeed, 0)
	if err := db.NewSelect().Model(&feeds).
		Where("id IN (25,768,769)").
		Scan(ctx); err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, mapFeeds(feeds))
}

func getFeed(c *gin.Context) {

	embed := c.QueryArray("embed")
	_, entries := Find(embed, "entries")
	feed := new(DBFeed)
	if entries {
		if err := db.NewSelect().Model(feed).Where("id = ?", c.Param("id")).Relation("Entries").Scan(ctx); err != nil {
			panic(err)
		}
	} else {
		if err := db.NewSelect().Model(feed).Where("id = ?", c.Param("id")).Scan(ctx); err != nil {
			panic(err)
		}
	}

	c.IndentedJSON(http.StatusOK, mapFeed(feed))
}

func getFeedEntries(c *gin.Context) {

	entries := make([]DBEntry, 0)
	if err := db.NewSelect().Model(&entries).Where("feed_id = ?", c.Param("id")).Scan(ctx); err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, entries)
}

func getEntry(c *gin.Context) {

	entry := new(DBEntry)
	if err := db.NewSelect().Model(entry).Where("id = ?", c.Param("id")).Scan(ctx); err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, mapEntry(entry))
}

func patchEntry(c *gin.Context) {
	var json PatchEntry
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	entry := DBEntry{
		ToRead: json.Read,
	}
	_, err := db.NewUpdate().
		Model(&entry).
		Set("to_read = ? ", entry.ToRead).
		Where("id = ?", c.Param("id")).
		Exec(ctx)
	if err != nil {
		panic(err)
	}
}
