package main

// keep track of last record that i saw
// get base 64 encoded and sent out to the client as a part of the response
// then what needs to happen is that client needs to take that next page
// key and send it back to the application or to the database
// the database will decode it and say okay my next page key or my token or my cursor
// is an ID of 10 and now what we do with that token or that cursor is we take that cursor and we put it back into the query

// Cursor pagination seems super useful in scenarios, where you have:
// a) infinite scroll pagination
// b) a table that updates in real time
// c) both
// draw back:
// you never know what page two or three or four actually is you only know what is the next page after the page that i have already seen
// so i'm going to give you a cursor which is a pointer to a record and I want you to show me the records that come after that record.
// so with cursor-based pagination there is no way to directly address a page, you can not skip to page 10 you can only forward and backward.
//
// more performant, so instead of generating this entire result set
// and throwing away the pages you don't need you allow the database to jump directly to the page you do need
// it's also more resilient to drift even if all of the records in the previous pages were deleted any your users sends you a token they're still going to get the next page of Records
// more complicated, stateful
// cannot directly address pages
// if you have an infinite scroll kind of thing you only ever need to know what the next page you might be fine.
type Cursor struct {
	MaxPageSize   int    `json:"max_page_size"`
	PageToken     string `json:"page_token"`
	NextPageToken string `json:"next_page_token"`
}

// advance
// + you can directly address any single page, if you want to to jump to page five you could you do not have to go through pages 1,2,3 and your
// you can jump directly to page five.
// limit pagesize offset pagesize*(page-1)
// + not very resilient to shifting data

// summary
// easy, stateless
// directly addressable pages
// more prone to drift
// slow for deep pages
type LimitOffset struct {
}
