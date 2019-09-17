package utils

import (
	"GIG/app/models"
	models2 "kavuda/models"
)

/*
Decode NewsItem to entity
 */
func EntityFromNews(newsItem models2.NewsItem, category string) models.Entity {
	return models.Entity{
		Title:     newsItem.Title,
		SourceURL: newsItem.Link,
		UpdatedAt: newsItem.Date,
	}.SetAttribute("", models.Value{
		Type:     "html",
		RawValue: newsItem.Content,
	}).SetAttribute("date", models.Value{
		Type:     "date",
		RawValue: newsItem.Date.String(),
	}).SetAttribute("author", models.Value{
		Type:     "string",
		RawValue: newsItem.Author,
	}).SetAttribute("snippet", models.Value{
		Type:     "string",
		RawValue: newsItem.Snippet,
	}).AddCategory("News").AddCategory(category)

}
