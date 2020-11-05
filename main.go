package main

import (
	"meh/router"
	"github.com/gin-gonic/gin"
	
	_"meh/logger"
	"github.com/sirupsen/logrus"
	
	
)

func main() {

	logrus.WithFields(logrus.Fields{"User": "walrus",}).Info("This is an info message")
	logrus.WithFields(logrus.Fields{"User": "walrus",}).Error("This is an error message")
	logrus.WithFields(logrus.Fields{"User": "walrus",}).Warn("This is a warning message")
	logrus.WithFields(logrus.Fields{"User": "walrus",}).Debug("This is a debug message")
	
	ginRouter := gin.Default()
	router.RegisterRoutes(ginRouter)
	
}
