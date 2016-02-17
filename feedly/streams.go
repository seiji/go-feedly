package feedly

import (
	"net/url"
)

type APIStreams struct {
	client *Client
}

type StreamIds struct {
	Ids          []string `json:"ids"`
	Continuation string   `json:"continuation,omitempty"`
}

type StreamContents struct {
	Continuation string `json:"continuation"`
	Id           string `json:"id"`
	Items        []Entry `json:"items"`
	Updated      int64  `json:"updated"`
}

type StreamOptions struct {
	Count        uint   `url:"count,omitempty"`
	Ranked       string `url:"ranked,omitempty"`
	UnreadOnly   bool   `url:"unreadOnly,omitempty"`
	NewerThan    int64  `url:"newerThan,omitempty"`
	Continuation string `url:"continuation,omitempty"`
}

func (a *APIStreams) Ids(streamId string, opt *StreamOptions) (*StreamIds, *Response, error) {
	rel := "streams/" + url.QueryEscape(streamId) + "/ids"
	rel, err := addOptions(rel, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := a.client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}
	ids := new(StreamIds)

	res, err := a.client.Do(req, ids)
	if err != nil {
		return nil, res, err
	}

	return ids, res, nil
}

func (a *APIStreams) Contents(streamId string, opt *StreamOptions) (*StreamContents, *Response, error) {
	rel := "streams/" + url.QueryEscape(streamId) + "/contents"
	rel, err := addOptions(rel, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := a.client.NewRequest("GET", rel, nil)
	if err != nil {
		return nil, nil, err
	}
	contents := new(StreamContents)

	res, err := a.client.Do(req, contents)
	if err != nil {
		return nil, res, err
	}

	return contents, res, nil
}
