package common

type PageCall func(map[string]interface{}) ([]interface{}, error)

func WithPageOffsetQuery(condition map[string]interface{}, limitParam string, pageParam string, limit int, start int, call PageCall) (data []interface{}, err error) {
	offset := start
	for {
		var d []interface{}
		condition[limitParam] = limit
		condition[pageParam] = offset
		d, err = call(condition)
		if err != nil {
			return data, err
		}
		data = append(data, d...)
		if len(d) < limit {
			break
		}
		offset = offset + limit
	}
	return data, err
}

func WithPageNumberQuery(condition map[string]interface{}, pageSizeParam string, pageNumParam string, pageSize int, initPageNumber int, call PageCall) (data []interface{}, err error) {
	pageNumber := initPageNumber
	for {
		var d []interface{}
		condition[pageSizeParam] = pageSize
		condition[pageNumParam] = pageNumber
		d, err = call(condition)
		if err != nil {
			return data, err
		}
		data = append(data, d...)
		if len(d) < pageSize {
			break
		}
		pageNumber = pageNumber + 1
	}
	return data, err
}

func WithSimpleQuery(condition map[string]interface{}, call PageCall) (data []interface{}, err error) {
	var d []interface{}
	d, err = call(condition)
	if err != nil {
		return data, err
	}
	data = append(data, d...)
	return
}

type NextTokenCall func(map[string]interface{}) ([]interface{}, string, error)

type DecodeNextToken func(source string) string

func WithNextTokenQuery(condition map[string]interface{}, maxResultsParam string, nextTokenParam string, maxResults int, decode DecodeNextToken, call NextTokenCall) (data []interface{}, err error) {
	var (
		nextToken string
	)
	for {
		var d []interface{}
		condition[maxResultsParam] = maxResults
		if nextToken != "" {
			if decode != nil {
				nextToken = decode(nextToken)
			}
			condition[nextTokenParam] = nextToken
		}
		d, nextToken, err = call(condition)
		if err != nil {
			return data, err
		}
		data = append(data, d...)
		if nextToken == "" {
			break
		}
	}
	return data, err
}
