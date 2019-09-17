package utils

import (
	"GIG/app/models"
	"GIG/scripts/crawlers/utils"
	"GIG/scripts/entity_handlers"
	"fmt"
)

func ProcessAndSaveEntity(entity models.Entity, textContent string) {
	//NER extraction
	fmt.Println("		Running NER on the text content...")
	entityTitles, err := utils.ExtractEntityNames(textContent)
	if err != nil {
		fmt.Println(err,entity.Title)
		fmt.Println(entityTitles)
	}
	fmt.Println("		NER completed successfully.")

	var entities []models.Entity

	for _, entityObject := range entityTitles {
		// uncomment the following line to enable normalizing entity names
		// normalizedName, err := utils.NormalizeName(entityObject.EntityName)
		normalizedName := entityObject.EntityName
		if err == nil {
			entities = append(entities, models.Entity{Title: normalizedName}.AddCategory(entityObject.Category))
		}
	}

	entity, err = entity_handlers.AddEntitiesAsLinks(entity, entities)
	//save to db
	entity, saveErr := entity_handlers.CreateEntity(entity)
	if saveErr != nil {
		fmt.Println(saveErr.Error(), entity.Title)
	}
	fmt.Println("		News Article Saved.", entity.Title)
}
