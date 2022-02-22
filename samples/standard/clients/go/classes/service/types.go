// THIS FILE IS A GENERATED CODE.
// DO NOT EDIT THIS CODE BY YOUR OWN HANDS
// generated version: (devel)

package types

import (
	"time"
)

type GetArticleRequest struct {
	ID int
}

type GetArticleResponse struct {
	Body        string
	Group       []string
	ID          int
	RequestTime time.Time
}
