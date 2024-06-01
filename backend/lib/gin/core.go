package main

type Paging struct {
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	Total      int    `json:"total"`
	FakeCursor string `json:"fake_cursor"`
	NextCursor string `json:"next_cursor"`
}

func (p *Paging) Process() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit < 1 {
		p.Limit = 10
	}
	if p.Limit >= 200 {
		p.Limit = 200
	}
}

type successResponse struct {
	Data   any `json:"data"`
	Paging any `json:"paging"`
	Extra  any `json:"extra"`
}

func SuccessResponse(data, paging, extra any) *successResponse {
	return &successResponse{
		Data:   data,
		Paging: paging,
		Extra:  extra,
	}
}

func ResponseData(data any) *successResponse {
	return SuccessResponse(data, nil, nil)
}
