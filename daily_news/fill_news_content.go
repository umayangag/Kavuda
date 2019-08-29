package daily_news

import (
	"kavuda/models"
)

func (d DailyNewsDecoder) FillNewsContent(newsItem models.NewsItem) (models.NewsItem, string, error) {
	return models.FillNewsContent(newsItem, ".node-main-content")
}
