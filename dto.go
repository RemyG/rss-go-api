package main

import (
	"time"

	"github.com/uptrace/bun"
)

type Category struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	CatOrder int    `json:"cat_order"`
	Feeds    []Feed `json:"feeds,omitempty"`
}

type Feed struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Link        string  `json:"link,omitempty"`
	BaseLink    string  `json:"base_link,omitempty"`
	Description string  `json:"description,omitempty"`
	CategoryID  int64   `json:"category_id,omitempty"`
	Valid       bool    `json:"valid,omitempty"`
	Viewframe   bool    `json:"viewframe,omitempty"`
	CatOrder    int     `json:"category_order,omitempty"`
	Entries     []Entry `json:"entries,omitempty"`
}

type Entry struct {
	ID          int64     `json:"id"`
	Published   time.Time `json:"published_date"`
	Updated     time.Time `json:"updated_date"`
	Link        string    `json:"link"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	Read        bool      `json:"read"`
	Content     string    `json:"content"`
	FeedID      int64     `json:"feed_id"`
	Favourite   bool      `json:"favourite"`
	ToRead      bool      `json:"to_read"`
}

type DBCategory struct {
	bun.BaseModel `bun:"rss_category,alias:c"`
	ID            int64
	Name          string
	CatOrder      int
	Feeds         []*DBFeed `bun:"rel:has-many,join:id=category_id"`
}

type DBFeed struct {
	bun.BaseModel `bun:"rss_feed,alias:f"`
	ID            int64
	Link          string
	BaseLink      string
	Title         string
	Description   string
	Updated       time.Time
	ToUpdate      bool
	MarkNewToRead bool
	CategoryID    int64
	Valid         bool
	Viewframe     bool
	CatOrder      int
	Entries       []*DBEntry `bun:"rel:has-many,join:id=feed_id"`
}

type DBEntry struct {
	bun.BaseModel `bun:"rss_entry,alias:e"`
	ID            int64
	Published     time.Time
	Updated       time.Time
	Link          string
	Title         string
	Description   string
	Author        string
	Read          bool
	Content       string
	FeedID        int64
	Favourite     bool
	ToRead        bool
}

type PatchEntry struct {
	Read      bool `json:"read"`
	Favourite bool `json:"favourite"`
}
