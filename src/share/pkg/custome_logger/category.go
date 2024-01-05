package custome_logger

type Category string
type SubCategory string
type ExtraKey string

const (
	General    Category = "General"
	IO         Category = "IO"
	Internal   Category = "Internal"
	Mongo      Category = "Mongo"
	RabbitMq   Category = "RabbitMq"
	Postgres   Category = "Postgres"
	Redis      Category = "Redis"
	Validation Category = "Validation"
	API        Category = "API"
	Prometheus Category = "Prometheus"
)

const (
	// General
	Startup         SubCategory = "Startup"
	ExternalService SubCategory = "ExternalService"
	// API
	RequestLog SubCategory = "RequestLog"
	// Mongo
	Migration SubCategory = "Migration"
	Seed      SubCategory = "Seed"
	Select    SubCategory = "Select"
	Rollback  SubCategory = "Rollback"
	Update    SubCategory = "Update"
	Delete    SubCategory = "Delete"
	Insert    SubCategory = "Insert"
	// Database
	Connect SubCategory = "Connect"
	Close   SubCategory = "Close"
	// RabitMq
	CreateChannel SubCategory = "CreateChannel"
	// Internal
	HashPassword        SubCategory = "HashPassword"
	DefaultRoleNotFound SubCategory = "DefaultRoleNotFound"

	// Validation
	MobileValidation   SubCategory = "MobileValidation"
	PasswordValidation SubCategory = "PasswordValidation"

	// IO
	RemoveFile SubCategory = "RemoveFile"
)

const (
	AppName      ExtraKey = "AppName"
	LoggerName   ExtraKey = "Logger"
	ClientIp     ExtraKey = "ClientIp"
	HostIp       ExtraKey = "HostIp"
	Method       ExtraKey = "Method"
	StatusCode   ExtraKey = "StatusCode"
	BodySize     ExtraKey = "BodySize"
	Path         ExtraKey = "Path"
	Latency      ExtraKey = "Latency"
	RequestBody  ExtraKey = "RequestBody"
	ResponseBody ExtraKey = "ResponseBody"
	ErrorMessage ExtraKey = "ErrorMessage"
)
