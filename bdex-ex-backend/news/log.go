package news

import logger "bdex.in/bdex/bdex-ex-backend/log"

var log *logger.SugaredLogger

func SetUpLog() {
	log = logger.Named("News")
}
