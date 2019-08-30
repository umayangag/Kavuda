package daily_mirror

import (
	"GIG/scripts/crawlers/utils/clean_html"
	"kavuda/models"
)

func (d DailyMirrorDecoder) FillNewsContent(newsItem models.NewsItem) (models.NewsItem, string, error) {
	return models.FillNewsContent(newsItem,".inner-content", clean_html.HtmlCleaner{})
}
