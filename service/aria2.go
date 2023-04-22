package service

import (
	"fmt"
	"github.com/tddey01/aria2/config"
	"github.com/tddey01/aria2/model"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/tddey01/aria2/utils"
)

var aria2Client *Aria2Client

var aria2Service *Aria2Service

func AdminOfflineDeal() {

	aria2Client = SetAndCheckAria2Config()

	//logs.GetLogger().Info("swan token:", swanClient.SwanToken)
	go aria2CheckDownloadStatus()
	go aria2StartDownload()
}

func aria2StartDownload() {
	for {
		log.Info(">>>>>>>>>>>> Start...")
		aria2Service.StartDownload(aria2Client)
		log.Info(">>>>>>>>>>>>  Sleeping...")
		time.Sleep(time.Minute)
	}
}

func aria2CheckDownloadStatus() {
	for {
		log.Info(">>>>>>>>>>>>  Start...")
		aria2Service.CheckDownloadStatus(aria2Client)
		log.Info(">>>>>>>>>>>>  Sleeping...")
		time.Sleep(time.Second * 30)
	}
}

func SetAndCheckAria2Config() *Aria2Client {

	aria2Host := config.GetConfig().Aria2.Aria2Host
	aria2Port := config.GetConfig().Aria2.Aria2Port
	aria2Secret := config.GetConfig().Aria2.Aria2Secret

	if len(aria2Host) == 0 {
		log.Fatal("please set config:aria2->aria2_host")
	}

	aria2Client = GetAria2Client(aria2Host, aria2Secret, aria2Port)

	return aria2Client
}

type Aria2Service struct {
	MinerFid string
}

func (aria2Service *Aria2Service) CheckDownloadStatus4Deal(aria2Client *Aria2Client, deal *model.FilSwan, gid string) {

	Table := config.GetConfig().Mysql.Table
	aria2Status := aria2Client.GetDownloadStatus(gid)
	if aria2Status == nil {
		log.Info(deal, DEAL_STATUS_DOWNLOAD_FAILED, "get download status failed for gid:"+gid, "no response from aria2")
		return
	}

	if aria2Status.Error != nil {
		log.Info(deal, DEAL_STATUS_DOWNLOAD_FAILED, "get download status failed for gid:"+gid, aria2Status.Error.Message)
		return
	}

	if len(aria2Status.Result.Files) != 1 {
		log.Info(deal, DEAL_STATUS_DOWNLOAD_FAILED, "get download status failed for gid:"+gid, "wrong file amount")
		return
	}

	result := aria2Status.Result
	file := result.Files[0]
	filePath := file.Path
	fileSize := utils.GetInt64FromStr(file.Length)
	Drive := deal.Drive
	msg := fmt.Sprintf("current status:,%s,%s", result.Status, result.ErrorMessage)
	log.Info(deal.DownloadUrl, "  ", deal.GId, "  ", msg)
	switch result.Status {
	case ARIA2_TASK_STATUS_ERROR:
		log.Info(deal, DEAL_STATUS_DOWNLOAD_FAILED, &filePath, result.Status, "download gid:"+gid, result.ErrorCode, result.ErrorMessage)
	case ARIA2_TASK_STATUS_ACTIVE, ARIA2_TASK_STATUS_WAITING:
		fileSizeDownloaded := utils.GetFileSize(filePath)
		completedLen := utils.GetInt64FromStr(file.CompletedLength)
		var completePercent float64 = 0
		if fileSize > 0 {
			completePercent = float64(completedLen) / float64(fileSize) * 100
		}
		downloadSpeed := utils.GetInt64FromStr(result.DownloadSpeed) / 1024
		fileSizeDownloaded = fileSizeDownloaded / 1024
		note := fmt.Sprintf("downloading, complete: %.2f%%, speed: %dKiB, downloaded:%dKiB, %s, download gid:%s", completePercent, downloadSpeed, fileSizeDownloaded, result.Status, gid)
		log.Info(deal.DownloadUrl, " ", note)
		if result.Status == ARIA2_TASK_STATUS_WAITING {
			msg = fmt.Sprintf("waiting to download,%s,%s", result.Status, result.ErrorMessage)
			log.Info(deal.DownloadUrl, "  ", deal.GId, "  ", msg)
		}
	case ARIA2_TASK_STATUS_COMPLETE:
		fileSizeDownloaded := utils.GetFileSize(filePath)
		log.Info(deal, "  downloaded")
		log.Info(deal, "  下载完成  ", fileSizeDownloaded)
		log.Info(deal.FileSize, "==", fileSizeDownloaded)
		if fileSizeDownloaded >= 0 {
			if err := model.UpdateSetDownload2(deal, gid, filePath, Drive, Table, fileSizeDownloaded); err != nil {
				return
			}
			log.Info(deal, DEAL_STATUS_DOWNLOADED, &filePath, "download gid:"+gid)
		} else {
			log.Info(deal, DEAL_STATUS_DOWNLOAD_FAILED, &filePath, "file not found on its download path", "download gid:"+gid)
		}
	default:
		log.Info(deal, DEAL_STATUS_DOWNLOAD_FAILED, &filePath, result.Status, "download gid:"+gid, result.ErrorCode, result.ErrorMessage)
	}
}

func (aria2Service *Aria2Service) StartDownload4Deal(deal *model.FilSwan, aria2Client *Aria2Client) {

	Table := config.GetConfig().Mysql.Table
	log.Info(deal, "start downloading")
	urlInfo, err := url.Parse(deal.DownloadUrl)
	if err != nil {
		log.Info(deal, DEAL_STATUS_DOWNLOAD_FAILED, "parse source file url error,", err.Error())
		return
	}

	log.Info("下载文件大小 ", deal.FileSize, "   ", deal.DownloadUrl)

	outFilename := urlInfo.Path
	if strings.HasPrefix(urlInfo.RawQuery, "filename=") {
		outFilename = strings.TrimPrefix(urlInfo.RawQuery, "filename=")
		outFilename = filepath.Join(urlInfo.Path, outFilename)
	}
	_, outFilename1 := filepath.Split(outFilename)
	outFilename = strings.TrimLeft(outFilename, "/")
	outDir := strings.TrimSuffix(deal.DiskPath, "/")
	filePath := outDir + "/" + outFilename
	Drive := deal.Drive
	s := -1
	if IsExist(filePath) {
		log.Info(deal, s, DEAL_STATUS_DOWNLOADED, filePath, outFilename+", the car file already exists, skip downloading it")

		file, _ := os.Stat(filePath)

		err = model.UpdateSetDownload3(deal, outFilename1, filePath, Drive, Table, file.Size())
		if err != nil {
			return
		}
		return
	}
	aria2Download := aria2Client.DownloadFile(deal.DownloadUrl, outDir, outFilename)
	if err = model.UpdateSetDownload1(deal, aria2Download.Gid, Drive, Table, outFilename1); err != nil { //  1 4
		log.Error("改状态失败")
		return
	}

	if aria2Download == nil {
		log.Info(deal, DEAL_STATUS_DOWNLOAD_FAILED, "no response when asking aria2 to download")
		return
	}

	if aria2Download.Error != nil {
		log.Info(deal, DEAL_STATUS_DOWNLOAD_FAILED, aria2Download.Error.Message)
		return
	}

	if aria2Download.Gid == "" {
		log.Info(deal, DEAL_STATUS_DOWNLOAD_FAILED, "no gid returned when asking aria2 to download")
		return
	}

	aria2Service.CheckDownloadStatus4Deal(aria2Client, deal, aria2Download.Gid)
}

func (aria2Service *Aria2Service) StartDownload(aria2Client *Aria2Client) {

	Table := config.GetConfig().Mysql.Table
	downloadingDeals, err := model.GetAll(Table)
	if err != nil {
		log.Error("获取错误")
		return
	}
	log.Debug("download task limit :", config.GetConfig().Aria2.Aria2Task)

	countDownloadingDeals := len(downloadingDeals)
	if countDownloadingDeals >= config.GetConfig().Aria2.Aria2Task {
		return
	}

	Locked, err := model.GeTLocked(Table)
	if err != nil {
		return
	}
	log.Info("downloading >>>>>>>>> ", len(Locked))
	if len(Locked) >= config.GetConfig().Aria2.Aria2Task {
		log.Infof("当前任务大于：%d 停止接新任务", config.GetConfig().Aria2.Aria2Task)
		return
	}
	for i := 1; i <= config.GetConfig().Aria2.Aria2Task-countDownloadingDeals; i++ {
		Locked1, err1 := model.GeTLocked(Table)
		if err1 != nil {
			return
		}

		if len(Locked1) >= config.GetConfig().Aria2.Aria2Task {
			log.Info("超出任务最下载限制===================")
			return
		}
		log.Info("开始下载")
		deal2Download, err := model.GetFindOne(Table) // 1  3
		if err != nil {
			log.Error(err)
			break
		}
		if deal2Download == nil {
			log.Info("No offline deal to download")
			break
		}
		aria2Service.StartDownload4Deal(deal2Download, aria2Client)
		time.Sleep(time.Second)
	}

}

func (aria2Service *Aria2Service) CheckDownloadStatus(aria2Client *Aria2Client) {

	Table := config.GetConfig().Mysql.Table
	downloadingDeals, err := model.GeTGId(Table)
	if err != nil {
		return
	}
	for _, deal := range downloadingDeals {
		gid := deal.GId
		if gid == "" {
			log.Error(deal, DEAL_STATUS_DOWNLOAD_FAILED, "download gid not found in offline_deals.note")
			continue
		}
		aria2Service.CheckDownloadStatus4Deal(aria2Client, deal, gid)
	}
}

func IsExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil || os.IsExist(err)
}
