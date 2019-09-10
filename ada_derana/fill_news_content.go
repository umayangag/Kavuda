package ada_derana

import (
	"GIG/scripts/crawlers/utils/clean_html"
	"kavuda/models"
)

func (d AdaDeranaDecoder) FillNewsContent(newsItem models.NewsItem) (models.NewsItem, string, error) {
	return models.FillNewsContent(newsItem, ".news-content", clean_html.HtmlCleaner{
	})
}
