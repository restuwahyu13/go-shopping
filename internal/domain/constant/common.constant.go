package cons

const (
	DEV  = "development"
	STAG = "staging"
	PROD = "production"
	TEST = "test"

	API = "/api/v1"

	ACTIVE   = true
	INACTIVE = false

	EMPTY           = ""
	Nil             = iota
	InvalidUUID     = "00000000-0000-0000-0000-000000000000"
	DEFAULT_ERR_MSG = "API is busy please try again later!"

	PERCENTAGE = "percentage"
	FIXED      = "fixed"
)

const (
	CREATED  = "created"
	PENDING  = "pending"
	SENT     = "sent"
	FAILED   = "failed"
	EXPIRED  = "expired"
	REFUNDED = "refunded"
	SUCCEED  = "succeed"

	WAITING   = "waiting"
	PROCESS   = "process"
	PICKUP    = "pick up"
	DELIVERED = "delivered"
	RETURNED  = "returned"
	COMPLETED = "completed"
)
