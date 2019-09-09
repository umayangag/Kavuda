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
	localNewsItems, err := extractNewItems(newsSiteUrl1)
	if err != nil {
		return nil, err
	}
	politicalNewsItems, err := extractNewItems(newsSiteUrl2)
	if err != nil {
		return nil, err
	}
	businessNewsItems, err := extractNewItems(newsSiteUrl3)
	if err != nil {
		return nil, err
	}

	allNewsItems := append(localNewsItems, politicalNewsItems...)
	allNewsItems = append(allNewsItems, businessNewsItems...)

	return allNewsItems, err
}

func extractNewItems(siteUrl string) ([]models.NewsItem, error) {
	//get the page
	resp, err := request_handlers.GetRequest(siteUrl)
	if err != nil {
		return nil, err
	}

	//convert html string to doc for element selection
	doc, err := utils2.HTMLStringToDoc(resp)
	if err != nil {
		return nil, err
	}

	var newsLinks []string

	newsNodes := doc.Find(".field-content")
	var newsItems []models.NewsItem
	for _, node := range newsNodes.Nodes {
		nodeDoc := goquery.NewDocumentFromNode(node)
		extractedUrl, exist := nodeDoc.Find("a").First().Attr("href")

		if exist { // if url found
			title := nodeDoc.Find("a").First().Nodes[0].FirstChild.Data
			if title != "img" { //is valid news link
				url := commons.FixUrl(extractedUrl, siteUrl)

				if !commons.StringInSlice(newsLinks, url) && !strings.Contains(url,"#comment"){ // if the link is not already enlisted before
					newsLinks = append(newsLinks, url)

					extractDate := strings.Split(extractedUrl, "/")
					dateString := extractDate[1] + " " + extractDate[2] + " " + extractDate[3]
					newsItems = append(newsItems, models.NewsItem{
						Title: title,
						Link:  url,
						Date:  utils.ExtractPublishedDate("2006 01 02", dateString),
					})
				}
			}
		}
	}

	return newsItems, nil
}
