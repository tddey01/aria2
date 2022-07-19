package service

import "fmt"

const (
	EMPTY_STRING         = ""
	DEFAULT_SELECT_LIMIT = "100"

	STORAGE_SERVER_TYPE_WEB_SERVER  = "web server"
	STORAGE_SERVER_TYPE_IPFS_SERVER = "ipfs server"

	SWAN_API_STATUS_SUCCESS = "success"
	SWAN_API_STATUS_FAIL    = "fail"

	TASK_TYPE_VERIFIED = "verified"
	TASK_TYPE_REGULAR  = "regular"

	TASK_STATUS_ASSIGNED              = "Assigned"
	TASK_STATUS_DEAL_SENT             = "DealSent"
	TASK_STATUS_PROGRESS_WITH_FAILURE = "ProgressWithFailure"

	TASK_FAST_RETRIEVAL_NO  = 0
	TASK_FAST_RETRIEVAL_YES = 1

	TASK_BID_MODE_MANUAL = 0 // allocate miner manually after creating task
	TASK_BID_MODE_AUTO   = 1 // allocate miner by market matcher after creating task
	TASK_BID_MODE_NONE   = 2 // set miner when creating task

	TASK_IS_PUBLIC = 1

	CAR_FILE_STATUS_CREATED  = "Created"
	CAR_FILE_STATUS_ASSIGNED = "Assigned"

	OFFLINE_DEAL_STATUS_ASSIGNED = "Assigned"
	OFFLINE_DEAL_STATUS_CREATED  = "Created"

	EPOCH_PER_HOUR = 120



	AuthorizationHeaderKey = "Authorization"

	TASK_SOURCE_ID_DEFAULT      = 0
	TASK_SOURCE_ID_SWAN         = 1
	TASK_SOURCE_ID_SWAN_CLIENT  = 2
	TASK_SOURCE_ID_SWAN_FS3     = 3
	TASK_SOURCE_ID_SWAN_PAYMENT = 4
	TASK_SOURCE_ID_OTHER        = 5



	WALLET_NON_VERIFIED_MESSAGE = "Not a Verified Client"

	LOTUS_AUTH_READ  = "read"
	LOTUS_AUTH_WRITE = "write"
	LOTUS_AUTH_SIGN  = "sign"
	LOTUS_AUTH_ADMIN = "admin"

	LOTUS_TRANSFER_TYPE_MANUAL = "manual"

	MAX_AUTO_BID_COPY_NUMBER = 8

	DURATION_DEFAULT = 1512000

	DURATION_MIN = 518400
	DURATION_MAX = 1540000

	HTTP_API_TIMEOUT_SECOND = 30
	URL_HOST_GET_COMMON    = "/common"
	URL_HOST_GET_HOST_INFO = "/miner/host/info"

	ERROR_LAUNCH_FAILED   = "Swan provider launch failed."
	INFO_ON_HOW_TO_CONFIG = "For more information about how to config, please check https://docs.filswan.com/run-swan-provider/config-swan-provider"

	UPDATE_OFFLINE_DEAL_STATUS_FAIL = "failed to update offline deal status"
	NOT_UPDATE_OFFLINE_DEAL_STATUS  = "no need to update deal status in swan"
)

const (
	MajorVersion = 2
	MinorVersion = 5
	FixVersion   = 0
	CommitHash   = ""
)

func GetVersion() string {
	if CommitHash != "" {
		return fmt.Sprintf("swan-miner-v%v.%v.%v-%s", MajorVersion, MinorVersion, FixVersion, CommitHash)
	} else {
		return fmt.Sprintf("swan-miner-v%v.%v.%v", MajorVersion, MinorVersion, FixVersion)
	}
}