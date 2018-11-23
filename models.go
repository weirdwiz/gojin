package main

// Article : used for the Article
type Article struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

var articles = []Article{
	Article{
		Title:       "Title 1",
		Content:     "wowee",
		Author:      "M'Boi MG",
		Description: "M'boi go to the market",
	},
	Article{
		Title:       "myboad",
		Content:     "lulz",
		Author:      "GBoi",
		Description: "wowee beyoch",
	},
}
