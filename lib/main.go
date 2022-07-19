package main

import (
	"github.com/tddey01/aria2/lib/client/web"
	"github.com/tddey01/aria2/lib/logs"
)

func main() {
	response, err := web.HttpGetNoToken("https://calibration-api.filscout.com/api/v1/storagedeal/6666", nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(response)
}
