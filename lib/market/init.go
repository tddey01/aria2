package market

import "github.com/tddey01/aria2/logger"

var log *logger.Logger

func init() {
	log = logger.InitLog()
}

const (
	LOTUS_JSON_RPC_ID      = 7878
	LOTUS_JSON_RPC_VERSION = "2.0"
)

const (
	LOTUS_MARKET_GET_ASK               = "Filecoin.MarketGetAsk"
	LOTUS_MARKET_IMPORT_DATA           = "Filecoin.MarketImportDealData"
	LOTUS_MARKET_LIST_INCOMPLETE_DEALS = "Filecoin.MarketListIncompleteDeals"
)
