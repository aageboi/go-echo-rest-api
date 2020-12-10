package models

type (
	// Images list cover image
	Images struct {
		URL     string `bson:"url" json:"url"`
		Caption string `bson:"caption" json:"caption"`
	}

	// ArticleDetail detail article
	ArticleDetail struct {
		ID         int      `bson:"id" json:"id"`
		Title      string   `bson:"title" json:"title"`
		Isi        string   `bson:"content" json:"content"`
		Date       string   `bson:"created_at" json:"created_at"`
		Images     []Images `bson:"image" json:"image"`
		Slug 	   string   `bson:"slug" json:"slug"`
	}

	// ArticleList detail article for list
	ArticleList struct {
		ID         int      `bson:"id" json:"id"`
		Title      string   `bson:"title" json:"title"`
		Date       string   `bson:"created_at" json:"created_at"`
		Images     []Images `bson:"image" json:"image"`
		Slug 	   string   `bson:"slug" json:"slug"`
	}
)
