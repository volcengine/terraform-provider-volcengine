package common

type RequestConvertType int

const (
	ConvertDefault RequestConvertType = iota
	ConvertWithN
	ConvertListUnique
	ConvertListN
	ConvertSingleN
	ConvertJsonObject
	ConvertJsonArray
	ConvertJsonObjectArray
	//ConvertWithFilter
	//ConvertListFilter
)

type RequestConvertMode int

const (
	RequestConvertAll RequestConvertMode = iota
	RequestConvertInConvert
	RequestConvertIgnore
)

type RequestContentType int

const (
	ContentTypeDefault RequestContentType = iota
	ContentTypeJson
)

type ServiceCategory int

const (
	DefaultServiceCategory ServiceCategory = iota
	ServiceBypass
)

type SpecialParamType int

const (
	DomainParam SpecialParamType = iota
	HeaderParam
	PathParam
	UrlParam
	FilePathParam
)
