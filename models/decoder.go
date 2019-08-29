package models

import (
	"GIG/app/models"
	"GIG/commons/request_handlers"
	"GIG/scripts/crawlers/utils"
	"GIG/scripts/crawlers/utils/clean_html"
	"GIG/scripts/entity_handlers"
	"golang.org/x/net/html"
	"strings"
)

type IDecoder interface {
	ExtractNewsItems() ([]NewsItem, error)
	FillNewsContent(newsItem NewsItem) (NewsItem, string, error)
	GetSourceTitle() string
}

func FillNewsContent(newsItem NewsItem, contentClass string) (NewsItem, string, error) {
	resp, err := request_handlers.GetRequest(newsItem.Link)
	if err != nil {
		return newsItem, "", err
	}

	newsDoc, err := utils.HTMLStringToDoc(resp)
	if err != nil {
		return newsItem, "", err
	}

	newsSelection := newsDoc.Find(contentClass).First()
	newsHtml, err := newsSelection.Html()
	if err != nil {
		return newsItem, "", err
	}

	news, err := html.Parse(strings.NewReader(newsHtml))
	if err != nil {
		return newsItem, "", err
	}

	//clean html code by removing unwanted information
	var imageList []models.Upload
	newsItem.Content, _, imageList = clean_html.CleanHTML(newsItem.Link, news)

	for _, image := range imageList {
		go func(payload models.Upload) {
			entity_handlers.UploadImage(payload)
		}(image)
	}

	return newsItem, newsSelection.Text(), nil
}