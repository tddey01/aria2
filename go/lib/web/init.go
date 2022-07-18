package web

import "github.com/tddey01/aria2/logger"

var log *logger.Logger

func init() {
	log = logger.InitLog()
}

const HTTP_CONTENT_TYPE_FORM = "application/x-www-form-urlencoded"
const HTTP_CONTENT_TYPE_JSON = "application/json; charset=UTF-8"
