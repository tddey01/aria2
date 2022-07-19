package market

import "github.com/tddey01/aria2/lib/web"
import "encoding/json"

func (lotusMarket *LotusMarket) LotusGetDeals() ([]Deal, error) {
	var params []interface{}
	jsonRpcParams := LotusJsonRpcParams{
		JsonRpc: LOTUS_JSON_RPC_VERSION,
		Method:  LOTUS_MARKET_LIST_INCOMPLETE_DEALS,
		Params:  params,
		Id:      LOTUS_JSON_RPC_ID,
	}

	response, err := web.HttpGet(lotusMarket.ApiUrl, lotusMarket.AccessToken, jsonRpcParams)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	deals := &MarketListIncompleteDeals{}
	err = json.Unmarshal(response, deals)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return deals.Result, nil
}
