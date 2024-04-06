package cardparam

type IsExistRequest struct {
	Name string `json:"card_name"`
}

type IsExistResponse struct {
	IsExist bool `json:"is_exist"`
}
