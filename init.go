package main

import "github.com/tddey01/aria2/logger"

var log *logger.Logger

func init() {
	log = logger.InitLog()
}
