package the_island

import (
	"GIG/commons"
	"GIG/commons/request_handlers"
	utils2 "GIG/scripts/crawlers/utils"
	"github.com/PuerkitoBio/goquery"
	"kavuda/models"
	"kavuda/utils"
)

func (d TheIslandDecoder) ExtractNewsItems() ([]models.NewsItem, error) {
	//get the page
	resp, err := request_handlers.GetRequest(newsSiteUrl)
	if err != nil {
		return nil, err
	}
	//convert html string to doc for element selection
	doc, err := utils2.HTMLStringToDoc(resp)
	if err != nil {
		return nil, err
	}

	var newsLinks []string

	newsNodes := doc.Find(".col")
	var newsItems []models.NewsItem
	for _, node := range newsNodes.Nodes {
		nodeDoc := goquery.NewDocumentFromNode(node)
		dateString, _ := nodeDoc.Find(".article_date").First().Html()

		if dateString != "" {
			extractedUrl, _ := nodeDoc.Find("a").First().Attr("href")
			if extractedUrl != "/" {
				title := nodeDoc.Find("a").First().Nodes[0].FirstChild.Data
				url := commons.FixUrl(extractedUrl, newsSiteUrl)

				if !commons.StringInSlice(newsLinks, url) { // if the link is not already enlisted before
					newsLinks = append(newsLinks, url)

					newsItems = append(newsItems, models.NewsItem{
						Title: title,
						Link:  url,
						Date:  utils.ExtractPublishedDate("January 02, 2006, 3:04 pm", dateString),
					})
				}
			}
		}
	}

	return newsItems, nil
}
