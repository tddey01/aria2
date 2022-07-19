package routers

import (
	"github.com/tddey01/aria2/service"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

func HostManager(router *gin.RouterGroup) {
	router.GET(URL_HOST_GET_HOST_INFO)
	router.GET(URL_HOST_GET_HOST_INFO, GetSwanMinerVersion)
}

func GetSwanMinerVersion(c *gin.Context) {
	info := getSwanMinerHostInfo()
	c.JSON(http.StatusOK, gin.H{
		"Code":200,
		"Msg":info,
	})
}

func getSwanMinerHostInfo() *HostInfo {
	info := new(HostInfo)
	info.SwanMinerVersion = service.GetVersion()
	info.OperatingSystem = runtime.GOOS
	info.Architecture = runtime.GOARCH
	info.CPUnNumber = runtime.NumCPU()
	return info
}
