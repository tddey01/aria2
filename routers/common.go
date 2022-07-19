package routers

import (
	"github.com/tddey01/aria2/comm"

	"github.com/gin-gonic/gin"
)

func HostManager(router *gin.RouterGroup) {
	router.GET(comm.URL_HOST_GET_HOST_INFO)
	//router.GET(comm.URL_HOST_GET_HOST_INFO, GetSwanMinerVersion)
}

//func GetSwanMinerVersion(c *gin.Context) {
//	info := getSwanMinerHostInfo()
//	c.JSON(http.StatusOK, common.CreateSuccessResponse(info))
//}
//
//func getSwanMinerHostInfo() *models.HostInfo {
//	//info := new(models.HostInfo)
//	//info.SwanMinerVersion = common.GetVersion()
//	//info.OperatingSystem = runtime.GOOS
//	//info.Architecture = runtime.GOARCH
//	//info.CPUnNumber = runtime.NumCPU()
//	return info
//}
