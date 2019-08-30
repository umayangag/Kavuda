package the_nation

import (
	"kavuda/models"
)

var newsSiteUrl = "https://nation.lk/edition/"

type TheNationDecoder struct {
	models.IDecoder
}

func (d TheNationDecoder) GetSourceTitle() string {
	return "The Nation"
}
