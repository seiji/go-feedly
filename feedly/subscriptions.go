package feedly

type APISubscriptions struct {
	client *Client
}

type Subscription struct {
	Categories      []Category  `json:"categories"`
	ContentType     string  `json:"contentType"`
	IconUrl         string  `json:"iconUrl"`
	Id              string  `json:"id"`
	Partial         bool    `json:"partial"`
	Subscribers     float64 `json:"subscribers"`
	Title           string  `json:"title"`
	Topics          []string  `json:"topics"`
	Updated         float64 `json:"updated"`
	Velocity        float64 `json:"velocity"`
	VisualUrl       string  `json:"visualUrl"`
	Website         string  `json:"website"`
}

func (a *APISubscriptions) Get() ([]Subscription, *Response, error) {
	rel := "subscriptions"

	req, err := a.client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}

	subscriptions := new([]Subscription)

	res, err := a.client.Do(req, subscriptions)
	if err != nil {
		return nil, res, err
	}

	return *subscriptions, res, nil
}

