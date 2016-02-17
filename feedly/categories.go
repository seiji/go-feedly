package feedly

type APICategories struct {
	client *Client
}

type Category struct {
	Id      string  `json:"id"`
	Label   string  `json:"label"`
}

func (a *APICategories) Get() ([]Category, *Response, error) {
	rel := "categories"

	req, err := a.client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	categories := new([]Category)

	res, err := a.client.Do(req, categories)
	if err != nil {
		return nil, res, err
	}

	return *categories, res, nil
}
