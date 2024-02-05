package dtos

type TemplateDtoReq struct {
	Data string `json:"data"`
}

type TemplateDtoRes struct {
	Data interface{} `json:"data"`
}
