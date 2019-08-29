package models

type NewsItemsResponse struct {
	SuccessMessage string             `json:"SuccessMessage"`
	ErrorMessage   string             `json:"ErrorMessage"`
	IsSuccess      bool               `json:"isSuccess"`
	Data           []NewsItemResponse `json:"data"`
}

type NewsItemResponse struct {
	NewsId string `json:"News_ID"`
}
