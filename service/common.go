package service

import (
	"fmt"
	"github.com/tddey01/aria2/config"
	"time"

	"github.com/tddey01/aria2/lib/client"
	"github.com/tddey01/aria2/lib/utils"
)

const ARIA2_TASK_STATUS_ERROR = "error"
const ARIA2_TASK_STATUS_WAITING = "waiting"
const ARIA2_TASK_STATUS_ACTIVE = "active"
const ARIA2_TASK_STATUS_COMPLETE = "complete"

const DEAL_STATUS_CREATED = "Created"
const DEAL_STATUS_WAITING = "Waiting"
const DEAL_STATUS_SUSPENDING = "Suspending"

const DEAL_STATUS_DOWNLOADING = "Downloading"
const DEAL_STATUS_DOWNLOADED = "Downloaded"
const DEAL_STATUS_DOWNLOAD_FAILED = "DownloadFailed"

const DEAL_STATUS_IMPORT_READY = "ReadyForImport"
const DEAL_STATUS_IMPORTING = "FileImporting"
const DEAL_STATUS_IMPORTED = "FileImported"
const DEAL_STATUS_IMPORT_FAILED = "ImportFailed"
const DEAL_STATUS_ACTIVE = "DealActive"

const ONCHAIN_DEAL_STATUS_ERROR = "StorageDealError"
const ONCHAIN_DEAL_STATUS_ACTIVE = "StorageDealActive"
const ONCHAIN_DEAL_STATUS_NOTFOUND = "StorageDealNotFound"
const ONCHAIN_DEAL_STATUS_WAITTING = "StorageDealWaitingForData"
const ONCHAIN_DEAL_STATUS_ACCEPT = "StorageDealAcceptWait"
const ONCHAIN_DEAL_STATUS_AWAITING = "StorageDealAwaitingPreCommit"

const LOTUS_IMPORT_NUMNBER = 20 //Max number of deals to be imported at a time
const LOTUS_SCAN_NUMBER = 100   //Max number of deals to be scanned at a time

var aria2Client *client.Aria2Client

var aria2Service *Aria2Service


func AdminOfflineDeal() {
	aria2Service = GetAria2Service()
	aria2Client = SetAndCheckAria2Config()

	//logs.GetLogger().Info("swan token:", swanClient.SwanToken)
	go aria2StartDownload()
}

func SetAndCheckAria2Config() *client.Aria2Client {
	aria2DownloadDir := config.GetConfig().Aria2.Aria2DownloadDir
	aria2Host := config.GetConfig().Aria2.Aria2Host
	aria2Port := config.GetConfig().Aria2.Aria2Port
	aria2Secret := config.GetConfig().Aria2.Aria2Secret

	if !utils.IsDirExists(aria2DownloadDir) {
		err := fmt.Errorf("aria2 down load dir:%s not exits, please set config:aria2->aria2_download_dir", aria2DownloadDir)
		log.Fatal(err)
	}

	if len(aria2Host) == 0 {
		log.Fatal("please set config:aria2->aria2_host")
	}

	aria2Client = client.GetAria2Client(aria2Host, aria2Secret, aria2Port)

	return aria2Client
}






func aria2StartDownload() {
	for {
		log.Info("Start...")
		aria2Service.StartDownload(aria2Client)
		log.Info("Sleeping...")
		time.Sleep(time.Minute)
	}
}


