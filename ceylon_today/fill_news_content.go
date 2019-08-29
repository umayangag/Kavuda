package ceylon_today

import (
	"GIG/commons/request_handlers"
	"encoding/json"
	models2 "kavuda/ceylon_today/models"
	"kavuda/models"
	"kavuda/utils"
)

func (d CeylonTodayDecoder) FillNewsContent(newsItem models.NewsItem) (models.NewsItem, string, error) {
	singleNewsResult, err := request_handlers.GetRequest(newsItem.Link)
	if err != nil {
		panic(err)
	}

	var singleNewsResponse models2.SingleNewsResponse
	if err = json.Unmarshal([]byte(singleNewsResult), &singleNewsResponse); err != nil {
		panic(err)
	}

	newsContent := singleNewsResponse.Data[0]

	newsItem = models.NewsItem{
		Title:   newsContent.Title,
		Snippet: newsContent.Snippet,
		Link:    newsItem.Link,
		Content: newsContent.HtmlContent,
		Date:    utils.ExtractPublishedDate("2006-01-02 15:04:05", newsContent.PublishDate),
		Author:  newsContent.AuthorName,
	}

	return newsItem, newsItem.Content, nil
}
