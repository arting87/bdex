package controller

import logger "bdex.in/bdex/bdex-ex-backend/log"

var log *logger.SugaredLogger

func SetupLog() {
	log = logger.Named("controller")
}
