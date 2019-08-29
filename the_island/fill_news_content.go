package the_island

import (
	"kavuda/models"
)

func (d TheIslandDecoder) FillNewsContent(newsItem models.NewsItem) (models.NewsItem, string, error) {
	return models.FillNewsContent(newsItem, ".node-main-content")
}
