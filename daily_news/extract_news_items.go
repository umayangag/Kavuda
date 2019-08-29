package daily_news

import (
	"GIG/commons"
	"GIG/commons/request_handlers"
	utils2 "GIG/scripts/crawlers/utils"
	"github.com/PuerkitoBio/goquery"
	"kavuda/models"
	"kavuda/utils"
	"strings"
)

func (d DailyNewsDecoder) ExtractNewsItems() ([]models.NewsItem, error) {
	//get the page
	resp, err := request_handlers.GetRequest(newsSiteUrl)
	if err != nil {
		panic(err)
	}
	//convert html string to doc for element selection
	doc, err := utils2.HTMLStringToDoc(resp)
	if err != nil {
		panic(err)
	}

	newsNodes := doc.Find(".field-content")
	var newsItems []models.NewsItem
	for _, node := range newsNodes.Nodes {
		nodeDoc := goquery.NewDocumentFromNode(node)
		extractedUrl, exist := nodeDoc.Find("a").First().Attr("href")

		if exist { // if url found
			title := nodeDoc.Find("a").First().Nodes[0].FirstChild.Data
			if title != "img" { //is valid news link
				url := commons.FixUrl(extractedUrl, newsSiteUrl)
				extractDate := strings.Split(url, "/")
				dateString := extractDate[0] + " " + extractDate[1] + " " + extractDate[2]

				newsItems = append(newsItems, models.NewsItem{
					Title: title,
					Link:  url,
					Date:    utils.ExtractPublishedDate("2006 01 02", dateString),
				})
			}
		}
	}

	return newsItems, nil
}
