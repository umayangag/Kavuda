package ada_derana

import (
	"kavuda/models"
)

var newsSiteUrl = "http://www.adaderana.lk/hot-news/"

type AdaDeranaDecoder struct {
	models.IDecoder
}

func (d AdaDeranaDecoder) GetSourceTitle() string {
	return "Ada Derana"
}
