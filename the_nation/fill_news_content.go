package the_nation

import (
	"GIG/scripts/crawlers/utils/clean_html"
	"kavuda/models"
)

func (d TheNationDecoder) FillNewsContent(newsItem models.NewsItem) (models.NewsItem, string, error) {
	return models.FillNewsContent(newsItem, ".news-content", clean_html.HtmlCleaner{
	})
}
