package the_island

import (
	"kavuda/models"
)

var newsSiteUrl = "http://www.island.lk/index.php?page_cat=news-section&page=news-section&code_title=39"
var newsSiteUrl1 = "http://www.island.lk/index.php?page_cat=news-section&page=news-section&code_title=48"

type TheIslandDecoder struct {
	models.IDecoder
}

func (d TheIslandDecoder) GetSourceTitle() string {
	return "The Island"
}
