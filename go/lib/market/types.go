package market

type LotusMarket struct {
	ApiUrl       string
	AccessToken  string
	ClientApiUrl string
}

type LotusJsonRpcParams struct {
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int           `json:"id"`
}

type DealCid struct {
	DealCid string `json:"/"`
}

type Deal struct {
	State       int     `json:"State"`
	Message     string  `json:"Message"`
	ProposalCid DealCid `json:"ProposalCid"`
}

type JsonRpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type MarketListIncompleteDeals struct {
	Id      int           `json:"id"`
	JsonRpc string        `json:"jsonrpc"`
	Result  []Deal        `json:"result"`
	Error   *JsonRpcError `json:"error"`
}
