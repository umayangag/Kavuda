package the_nation

import (
	"GIG/commons"
	"GIG/commons/request_handlers"
	utils2 "GIG/scripts/crawlers/utils"
	"github.com/PuerkitoBio/goquery"
	"kavuda/models"
	"kavuda/utils"
)

func (d TheNationDecoder) ExtractNewsItems() ([]models.NewsItem, error) {
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

	newsNodes := doc.Find(".rss_item")
	var newsItems []models.NewsItem
	for _, node := range newsNodes.Nodes {
		nodeDoc := goquery.NewDocumentFromNode(node)

		url, _ := nodeDoc.Find("a").First().Attr("href")

		if !commons.StringInSlice(newsLinks, url) { // if the link is not already enlisted before
			newsLinks = append(newsLinks, url)

			dateString := nodeDoc.Find("small").First().Nodes[0].FirstChild.Data
			snippet := nodeDoc.Find("p").First().Nodes[0].FirstChild.Data
			title, _ := nodeDoc.Find("a").First().Attr("title")

			newsItems = append(newsItems, models.NewsItem{
				Title:   title,
				Link:    url,
				Date:    utils.ExtractPublishedDate("on January 02, 2006 at 3:04 pm", dateString),
				Snippet: snippet,
			})
		}
	}

	return newsItems, nil
}
