package utils

import (
	"strconv"
	"strings"
)

type QueryParam struct {
	Filters  map[string]string
	Page     int64
	Limit    int64
	Ordering []string
	Search   string
}

func ParseQueryParam(queryParams map[string][]string) (*QueryParam, []string) {
	params := QueryParam{
		Filters:  make(map[string]string),
		Page:     1,
		Limit:    10,
		Ordering: []string{},
		Search:   "",
	}
	var errStr []string
	var err error

	for key, value := range queryParams {
		if key == "page" {
			params.Page, err = strconv.ParseInt(value[0], 10, 64)
			if err != nil {
				errStr = append(errStr, "invalid `page` param")
			}
			continue
		}

		if key == "limit" {
			params.Limit, err = strconv.ParseInt(value[0], 10, 64)
			if err != nil {
				errStr = append(errStr, "invalid `limit` param")
			}
			continue
		}

		if key == "search" {
			params.Search = value[0]
			continue
		}
		if key == "ordering" {
			params.Ordering = strings.Split(value[0], ",")
			continue
		}
		params.Filters[key] = value[0]
	}
	return &params, errStr
}

type UserIdQueryParam struct {
	Filters  map[string]string
    UserId   string
	Ordering []string
	Search   string
}

func ParseUserIdQueryParam(queryParams map[string][]string) ([]*UserIdQueryParam, []string) {
	var paramsList []*UserIdQueryParam
	var errStr []string

	for _, value := range queryParams["user_id"] {
		params := UserIdQueryParam{
			Filters:  make(map[string]string),
			UserId:   value,
			Ordering: []string{},
			Search:   "",
		}

		for key, value := range queryParams {
			if key == "search" {
				params.Search = value[0]
				continue
			}
			if key == "ordering" {
				params.Ordering = strings.Split(value[0], ",")
				continue
			}
			params.Filters[key] = value[0]
		}

		paramsList = append(paramsList, &params)
	}

	if len(paramsList) == 0 {
		errStr = append(errStr, "user_id parameter is required")
	}

	return paramsList, errStr
}

type PostIdQueryParam struct {
	Filters  map[string]string
    PostId   string
	Ordering []string
	Search   string
}

func ParsePostIdQueryParam(queryParams map[string][]string) ([]*PostIdQueryParam, []string) {
	var paramsList []*PostIdQueryParam
	var errStr []string

	for _, value := range queryParams["post_id"] {
		params := PostIdQueryParam{
			Filters:  make(map[string]string),
			PostId:   value,
			Ordering: []string{},
			Search:   "",
		}

		for key, value := range queryParams {
			if key == "search" {
				params.Search = value[0]
				continue
			}
			if key == "ordering" {
				params.Ordering = strings.Split(value[0], ",")
				continue
			}
			params.Filters[key] = value[0]
		}

		paramsList = append(paramsList, &params)
	}

	if len(paramsList) == 0 {
		errStr = append(errStr, "post_id parameter is required")
	}

	return paramsList, errStr
}