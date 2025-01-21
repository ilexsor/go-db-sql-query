package internal

import (
	log "github.com/sirupsen/logrus"
)

func HandleError(err error) {
	if err != nil {
		log.WithFields(log.Fields{
			"err_message": err,
		}).Error("error appears")
		return
	}
}

func CheckOpenedDBError(err error) {
	if err != nil {
		log.WithFields(log.Fields{
			"err_message": err,
		}).Fatal("error appears")
		return
	}
}
