package go_web_utils

import (
	"container/list"
)

type requests struct {
	SeenUrls list.List
}

var UserRequests requests
