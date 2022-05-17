package common

type ResourceResponseHandler func() (map[string]interface{}, map[string]ResponseConvert, error)
