package the_island

import (
	"GIG/scripts/crawlers/utils/clean_html"
	"kavuda/models"
)

func (d TheIslandDecoder) FillNewsContent(newsItem models.NewsItem) (models.NewsItem, string, error) {
	return models.FillNewsContent(newsItem, "#left_video_area", clean_html.HtmlCleaner{
		Config: clean_html.Config{
			IgnoreElements: []string{"h1"},
			IgnoreStrings:  []string{"Add new comment", "Print Edition", "Send to Friend"},
			IgnoreClasses:  []string{"article_info_col"},
		}})
}
