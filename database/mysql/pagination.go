package main

type Pagination struct {
	MaxPageSize   int    `json:"max_page_size"`
	PageToken     string `json:"page_token"`
	NextPageToken string `json:"next_page_token"`
}
