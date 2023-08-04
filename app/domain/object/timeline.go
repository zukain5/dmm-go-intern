package object

import (
	"math"
	"net/url"
	"strconv"
)

type Timeline []*Status

type FindPublicTimelinesParams struct {
	MaxId   string
	SinceId string
	Limit   string
}

func NewFindPublicTimelinesParams(v url.Values) FindPublicTimelinesParams {
	maxId := strconv.Itoa(math.MaxInt64)
	sinceId := "0"
	limit := "40"

	if v.Has("max_id") {
		maxId = v.Get("max_id")
	}

	if v.Has("since_id") {
		sinceId = v.Get("since_id")
	}

	if v.Has("limit") {
		limit = determineLimit(v.Get("limit"))
	}

	return FindPublicTimelinesParams{
		MaxId:   maxId,
		SinceId: sinceId,
		Limit:   limit,
	}
}

func determineLimit(limit string) string {
	l, err := strconv.Atoi(limit)
	if err != nil {
		return "40"
	}

	if l > 80 {
		return "80"
	} else {
		return limit
	}
}
