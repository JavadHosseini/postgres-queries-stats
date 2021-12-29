package cnst

const (
	Type                = "type"
	Service             = "service"
	Database            = "database"
	Handler             = "handler"
	Star                = "*"
	PathRouteGroupV1    = "/api/v1"
	COLON               = ":"
	PathListQueries     = "/list/"
	InitDBError         = "db error: #2, could not initialize database: %v"
	ServerPort          = "3000"
	CloseDBError        = "db error: #1, could not close database: %v"
	Empty               = ""
	ErrNoDirection      = "no direction found"
	ErrPageNotValid     = "page header is not valid"
	ErrPageSizeNotValid = "pagesize header is not valid"
)
