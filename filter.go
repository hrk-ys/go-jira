package jira

import "fmt"

type FilterService struct {
	client *Client
}

type Filter struct {
	ID          string `json:"id,omitempty" structs:"id,omitempty"`
	Self        string `json:"self,omitempty" structs:"self,omitempty"`
	Name        string `json:"name,omitempty" structs:"name,omitempty"`
	Description string `json:"description,omitempty" structs:"description,omitempty"`
	Jql         string `json:"jql,omitempty" structs:"jql,omitempty"`
	SearchURL   string `json:"searchUrl,omitempty" structs:"searchUrl,omitempty"`
}

func (s *FilterService) GetFilter(filterID int) (*Filter, *Response, error) {
	apiEndpoint := fmt.Sprintf("rest/api/2/filter/%v", boardID)

	req, err := s.client.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	filter := new(Filter)
	resp, err := s.client.Do(req, filter)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}

	return filter, resp, nil
}
