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

	newsNodes := doc.Find(".col")
	var newsItems []models.NewsItem
	for _, node := range newsNodes.Nodes {
		nodeDoc := goquery.NewDocumentFromNode(node)
		newsDate, _ := nodeDoc.Find(".article_date").First().Html()

		if newsDate != "" {
			extractedUrl, _ := nodeDoc.Find("a").First().Attr("href")
			if extractedUrl!="/" {
				title := nodeDoc.Find("a").First().Nodes[0].FirstChild.Data
				url := commons.FixUrl(extractedUrl, newsSiteUrl)

				newsItems = append(newsItems, models.NewsItem{
					Title: title,
					Link:  url,
					Date:  utils.ExtractPublishedDate("January 02, 2006, 3:04 pm", newsDate),
				})
			}
		}
	}

	return newsItems, nil
}
