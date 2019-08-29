package daily_mirror

import (
	"kavuda/models"
)

func (d DailyMirrorDecoder) FillNewsContent(newsItem models.NewsItem) (models.NewsItem, string, error) {
	return models.FillNewsContent(newsItem,".inner-content")
}
