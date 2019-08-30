package main

import (
	"fmt"
	"kavuda/daily_mirror"
	"kavuda/models"
	"kavuda/utils"
)

func main() {
	crawl(daily_mirror.DailyMirrorDecoder{})
	//crawl(ceylon_today.CeylonTodayDecoder{})
	//crawl(daily_news.DailyNewsDecoder{})
	//crawl(the_island.TheIslandDecoder{})
}

func crawl(decoder models.IDecoder) {
	//extract news items from site
	newsItems, err := decoder.ExtractNewsItems()
	if err != nil {
		panic(err)
	}
	fmt.Println("News links extracted...")
	fmt.Println(len(newsItems), "news items found\n ")

	//for each news article
	fmt.Println("Reading News...")
	for _, newsItem := range newsItems {

		fmt.Println("	News: ", newsItem.Link)
		newsItem, contentString, err := decoder.FillNewsContent(newsItem)
		if err != nil {
			panic(err)
		}

		fmt.Println("		Reading News Article Completed.")

		//decode to entity
		entity := utils.EntityFromNews(newsItem, decoder.GetSourceTitle())

		//save entity with NER processing
		utils.ProcessAndSaveEntity(entity, contentString)
	}
}
