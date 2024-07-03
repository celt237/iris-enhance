package internal

const (
	ParamTypePath   string = "path"
	ParamTypeQuery  string = "query"
	ParamTypeHeader string = "header"
	ParamTypeCookie string = "cookie"
	ParamTypeBody   string = "body"
)

//type DataTypeEnum string

const (
	DataTypeString string = "string"
	DataTypeNumber string = "number"
	DataTypeInt    string = "integer"
	DataTypeBool   string = "boolean"
	DataTypeArray  string = "array"
	DataTypeObject string = "object"
)

const (
	ZServiceTag     = "zService"
	ZReplyTypeTag   = "zResult"
	ZSummaryTag     = "zSummary"
	ZDescriptionTag = "zDescription"
	ZTagsTag        = "zTags"
	ZParamTag       = "zParam"
	ZReplyDataTag   = "zResultData"
	ZAcceptTag      = "zAccept"
	ZProduceTag     = "zProduce"
	ZRouterTag      = "zRouter"
)

const (
	SwaggerSummaryTag     = "@Summary"
	SwaggerDescriptionTag = "@Description"
	SwaggerTagsTag        = "@Tags"
	SwaggerParamTag       = "@Param"
	SwaggerSuccessTag     = "@Success"
	SwaggerFailureTag     = "@Failure"
	SwaggerAcceptTag      = "@Accept"
	SwaggerProduceTag     = "@Produce"
	SwaggerRouterTag      = "@Router"
)

const DefaultErrorCode = "500"
