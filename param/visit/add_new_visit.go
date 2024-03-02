package visitparam

type AddNewCardVisitRequest struct {
	CardName  string `json:"card_name"`
	UserAgent map[string]string
}
