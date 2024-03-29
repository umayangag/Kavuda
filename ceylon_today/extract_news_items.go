package ceylon_today

import (
	"GIG/commons"
	"GIG/commons/request_handlers"
	"encoding/json"
	"errors"
	models2 "kavuda/ceylon_today/models"
	"kavuda/models"
)

func (d CeylonTodayDecoder) ExtractNewsItems() ([]models.NewsItem, error) {
	//get the page
	resp, err := request_handlers.GetRequest(newsSiteUrl)
	if err != nil {
		return nil, err
	}
	var (
		newsItemsResponse models2.NewsItemsResponse
		newsItems         []models.NewsItem
	)
	if err := json.Unmarshal([]byte(resp), &newsItemsResponse); err != nil {
		return nil, err
	}

	if newsItemsResponse.SuccessMessage != "OK" {
		return nil, errors.New("request success message not received")
	}

	var newsLinks []string

	//create news item list from news item responses
	for _, newsItemResponse := range newsItemsResponse.Data {
		url := singleNewsUrl + newsItemResponse.NewsId
		if !commons.StringInSlice(newsLinks, url) { // if the link is not already enlisted before
			newsLinks = append(newsLinks, url)
			newsItem := models.NewsItem{
				Link: url,
			}
			newsItems = append(newsItems, newsItem)
		}
	}

	return newsItems, nil
}
