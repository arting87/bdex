package sync

import logger "bdex.in/bdex/bdex-ex-backend/log"

var log *logger.SugaredLogger

func setupSyncLog() {
	log = logger.Named("sync")
}
