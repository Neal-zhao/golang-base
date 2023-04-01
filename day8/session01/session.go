package main

import (
	uuid "github.com/satori/go.uuid"
)

const (
	sessionCookieName = "session_id"
)
var MgrObj Mgr
type Mgr interface {
	Init(addr string,options ...string)
	GetSessionData(id string) (sd SessionData,ok bool)
	SetSessionData(uid string ) (err bool)
	DeleteSessionData(uid string ) (err bool)
	CreateSessionData() (sd SessionData)
}
//字段不兼容 就新开个接口
type SessionData interface {
	GetID() string
	Set(key string,val interface{})
	Get(key string) (val interface{},err error)
	Del(key string)
	Save()
}
func InitMgr(name string,addr string,options ...string) {
	//var MgrObj Mgr
	switch name {
	case "memory":
		MgrObj = NewMemoryMgr()
	case "redis":
		MgrObj = NewRedisMgr()
	}
	MgrObj.Init("127.0.0.1","","0")
}
func CreateUUid() string  {
	uid := uuid.NewV4()
	return uid.String()
}
func SessionNew(name string) (sd SessionData )  {
	uid := CreateUUid()
	//sd = NewSessionDataRedis(uid)
	////var MgrObj Mgr
	switch name {
	case "memory":
		sd = NewSessionDataMem(uid)
	case "redis":
		sd = NewSessionDataRedis(uid)
	}
	return sd
	//return
}

