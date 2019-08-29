package daily_news

import (
	"kavuda/models"
)

var newsSiteUrl = "http://www.dailynews.lk/"

type DailyNewsDecoder struct {
	models.IDecoder
}

func (d DailyNewsDecoder) GetSourceTitle() string {
	return "Daily News"
}
