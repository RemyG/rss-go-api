package main

func mapCategories(cs []*DBCategory) (categories []Category) {
	categories = make([]Category, len(cs))
	for i, c := range cs {
		categories[i] = mapCategory(c)
	}
	return
}

func mapCategory(c *DBCategory) (category Category) {
	category = Category{
		ID:       c.ID,
		CatOrder: c.CatOrder,
		Name:     c.Name,
		Feeds:    mapFeeds(c.Feeds),
	}
	return
}

func mapFeeds(fs []*DBFeed) (feeds []Feed) {
	feeds = make([]Feed, len(fs))
	for i, f := range fs {
		feeds[i] = mapFeed(f)
	}
	return
}

func mapFeed(f *DBFeed) (feed Feed) {
	feed = Feed{
		ID:          f.ID,
		Title:       f.Title,
		Link:        f.Link,
		BaseLink:    f.BaseLink,
		Description: f.Description,
		CategoryID:  f.CategoryID,
		Valid:       f.Valid,
		Viewframe:   f.Viewframe,
		CatOrder:    f.CatOrder,
		Entries:     mapEntries(f.Entries),
	}
	return
}

func mapLightFeed(f *DBFeed) (feed Feed) {
	feed = Feed{
		ID:    f.ID,
		Title: f.Title,
	}
	return
}

func mapEntries(es []*DBEntry) (entries []Entry) {
	entries = make([]Entry, len(es))
	for i, e := range es {
		entries[i] = mapEntry(e)
	}
	return
}

func mapEntry(e *DBEntry) (entry Entry) {
	entry = Entry{
		ID:          e.ID,
		Title:       e.Title,
		Published:   e.Published,
		Updated:     e.Updated,
		Link:        e.Link,
		Description: e.Description,
		Author:      e.Author,
		Read:        e.Read,
		Content:     e.Content,
		FeedID:      e.FeedID,
		Favourite:   e.Favourite,
		ToRead:      e.ToRead,
	}
	return
}

func mapLightEntry(e *DBEntry) (entry Entry) {
	entry = Entry{
		ID:    e.ID,
		Title: e.Title,
	}
	return
}
