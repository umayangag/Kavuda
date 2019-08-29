package daily_mirror

import (
	"kavuda/models"
)

var newsSiteUrl = "http://www.dailymirror.lk/top-storys/155"

type DailyMirrorDecoder struct {
	models.IDecoder
}

func (d DailyMirrorDecoder) GetSourceTitle() string {
	return "Daily Mirror"
}