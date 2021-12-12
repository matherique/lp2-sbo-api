package dto

type ReqNewBook struct {
	Name        string `json:"name"`
	AuthorId    int    `json:"author_id"`
	ShortDesc   string `json:"short_desc"`
	Longdesc    string `json:"long_desc"`
	CategoryId  string `json:"category_id"`
	PaperBack   int    `json:"paperback"`
	PublisherId int    `json:"publisher_id"`
}

type ResNewBook struct{}

type NewBook struct{}
