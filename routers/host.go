package routers

type HostInfo struct {
	SwanMinerVersion string `json:"swan_miner_version"`
	OperatingSystem  string `json:"operating_system"`
	Architecture     string `json:"architecture"`
	CPUnNumber       int    `json:"cpu_number"`
}

const (
	MajorVersion = 2
	MinorVersion = 5
	FixVersion   = 0
	CommitHash   = ""
	URL_HOST_GET_HOST_INFO = "/miner/host/info"
)