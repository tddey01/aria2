package ipfs

import (
	"fmt"

	"github.com/tddey01/aria2/lib/client/web"
	"github.com/tddey01/aria2/lib/constants"
	"github.com/tddey01/aria2/lib/logs"
	"github.com/tddey01/aria2/lib/utils"
)

func IpfsUploadFileByWebApi(apiUrl, filefullpath string) (*string, error) {
	response, err := web.HttpUploadFileByStream(apiUrl, filefullpath)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	fileHash := utils.GetFieldStrFromJson(response, "Hash")
	//logs.GetLogger().Info(carFileHash)

	if fileHash == constants.EMPTY_STRING {
		err := fmt.Errorf("cannot get file hash from response:%s", response)
		//logs.GetLogger().Error(err)
		return nil, err
	}

	return &fileHash, nil
}

func Export2CarFile(apiUrl, fileHash string, carFileFullPath string) error {
	apiUrlFull := utils.UrlJoin(apiUrl, "api/v0/dag/export")
	apiUrlFull = apiUrlFull + "?arg=" + fileHash + "&progress=false"
	carFileContent, err := web.HttpPostNoToken(apiUrlFull, "")
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	bytesWritten, err := utils.CreateFileWithByteContents(carFileFullPath, carFileContent)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	logs.GetLogger().Info(bytesWritten, " bytes have been written to:", carFileFullPath)
	return nil
}
