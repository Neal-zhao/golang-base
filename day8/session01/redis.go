package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"sync"
	"time"
)
type SessionDataRedis struct {
	ID string
	RWLock sync.RWMutex
	Data map[string]interface{}
	client *redis.Client
	expired int	//过期时间
}
func NewSessionDataRedis(uid string) *SessionDataRedis {
	sd := &SessionDataRedis{
		ID: uid,
		Data: make(map[string]interface{},10),
	}
	sd.Data["session_id"] = uid
	return sd
}
func (s *SessionDataRedis)Save()  {
	val,err := json.Marshal(s.Data)
	if err != nil {
		fmt.Printf("marshal data failed err:%v\n",err)
		return
	}
	//保存数据到 redis
	s.client.Set(s.ID,val,time.Second * time.Duration(s.expired))
}
func (s *SessionDataRedis)GetID() string {
	return s.ID
}
func (s *SessionDataRedis)Set(key string,val interface{})  {
	s.RWLock.RLock()
	defer s.RWLock.RUnlock()
	s.Data[key] = val
}
func (s *SessionDataRedis)Get(key string) (val interface{},err error)  {
	val,ok := s.Data[key];
	if ok{
		err = fmt.Errorf("invalid key %s",key)
		return
	}
	return
}
func (s *SessionDataRedis)Del(key string) {
	s.RWLock.Lock()
	defer s.RWLock.Unlock()
	delete(s.Data,key)
}

type RedisMgr struct {
	SessionData map[string]SessionData
	rwLock sync.RWMutex
	client *redis.Client
}
func NewRedisMgr() *MgrMemory {
	return &MgrMemory{
		SessionData: make(map[string]SessionData,1024),
	}
}
func (r *RedisMgr) CreateSessionData() (sd SessionData) {
	uuid := CreateUUid()
	sd = NewSessionDataRedis(uuid)
	return
}
func (r *RedisMgr)Init(addr string,options ...string)  {
	var (
		password string
		db string
	)
	if len(options) == 1 {
		password = options[0]
	}
	if len(options) == 2 {
		db = options[1]
	}
	dbi,err := strconv.Atoi(db)
	if err != nil {
		dbi = 0
	}
	r.client = redis.NewClient(&redis.Options{
		Addr: addr,
		Password: password,
		DB: dbi,
	})
	_,err = r.client.Ping().Result()
	if err != nil {
		panic(err)
	}
}
func (r *RedisMgr)LoadFromRedis(sessionId string) (err error) {
	//连接redis
	value,err := r.client.Get(sessionId).Result()
	if err != nil {
		fmt.Printf("没有对应的值 %s",sessionId)
		return
	}
	err = json.Unmarshal([]byte(value),&r.SessionData)
	if err != nil {
		return
	}
	return
	//根据sessionId拿到数据
	//数据反序列化
}
func (r *RedisMgr)GetSessionData(id string) (sd SessionData,ok bool) {
	//r.Session中必须依据从redis里那出来
	if r.SessionData == nil {
		err := r.LoadFromRedis(id)
		if err != nil {
			return nil,false
		}
	}
	//r.Session[sessionId] 拿到数据
	r.rwLock.RLock()
	defer r.rwLock.RUnlock()
	if v,ok := r.SessionData[id];ok{
		return v,ok
	}
	return nil,ok
}
func (r *RedisMgr)SetSessionData(uid string ) (err bool) {
	if _,ok := r.SessionData[uid];ok{
		return ok
	}
	//uid := m.CreateUUid()
	r.SessionData[uid] = NewSessionDataRedis(uid)
	return true
}
func (r *RedisMgr)DeleteSessionData(uid string ) (err bool) {
	if _,ok := r.SessionData[uid];!ok{
		fmt.Printf("%s不存在",uid)
		return ok
	}
	delete(r.SessionData,uid)
	return true
}
