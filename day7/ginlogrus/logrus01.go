package main

import "github.com/sirupsen/logrus"

func logrus_Test()  {
	logrus.WithFields(logrus.Fields{
		"method":"get",
		"action":"index",
	}).Warn("这是一个请求的一些信息")

}