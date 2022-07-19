package lotus

import (
	"fmt"
	"strings"

	"github.com/tddey01/aria2/lib/client"
	"github.com/tddey01/aria2/lib/logs"
)

func IsWalletVerified(wallet string) (bool, error) {
	wallet = strings.Trim(wallet, " ")
	if wallet == "" {
		err := fmt.Errorf("invalid wallet")
		logs.GetLogger().Error(err)
		return false, err
	}

	cmd := "lotus-shed verifreg check-client " + wallet

	result, err := client.ExecOsCmd(cmd, true)
	if err != nil {
		logs.GetLogger().Error(err)

		if strings.Contains(err.Error(), "is not a verified client") {
			return false, nil
		}

		return false, err
	}

	if strings.Contains(result, "is not a verified client") {
		return false, nil
	}

	return true, nil
}
