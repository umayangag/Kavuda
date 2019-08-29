package daily_news

import (
	"kavuda/models"
)

var newsSiteUrl1 = "http://www.dailynews.lk/category/local"
var newsSiteUrl2 = "http://www.dailynews.lk/category/political"
var newsSiteUrl3 = "http://www.dailynews.lk/category/business"

type DailyNewsDecoder struct {
	models.IDecoder
}

func (d DailyNewsDecoder) GetSourceTitle() string {
	return "Daily News"
}
