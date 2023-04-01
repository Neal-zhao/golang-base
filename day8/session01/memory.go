package main

import (
	"fmt"
	"sync"
)

type SessionDataMem struct {
	ID string
	RWLock sync.RWMutex
	Data map[string]interface{}
}
func NewSessionDataMem(uid string) *SessionDataMem {
	sd := &SessionDataMem{
		ID: uid,
		Data: make(map[string]interface{},10),
	}
	sd.Data["session_id"] = uid
	return sd
}
func (s *SessionDataMem)GetID() string {
	return s.ID
}
func (s *SessionDataMem)Set(key string,val interface{})  {
	s.RWLock.RLock()
	defer s.RWLock.RUnlock()
	s.Data[key] = val
}
func (s *SessionDataMem)Get(key string) (val interface{},err error)  {
	val,ok := s.Data[key];
	if ok{
		err = fmt.Errorf("invalid key %s",key)
		return
	}
	return
}
func (s *SessionDataMem)Del(key string) {
	s.RWLock.Lock()
	defer s.RWLock.Unlock()
	delete(s.Data,key)
}
func (s *SessionDataMem)Save()  {}

type MgrMemory struct {
	SessionData map[string]SessionData
	Lock sync.RWMutex
}
func NewMemoryMgr() *MgrMemory {
	return &MgrMemory{
		SessionData: make(map[string]SessionData,1024),
	}
}
func (m *MgrMemory)CreateSessionData() (sd SessionData )  {
	uid := CreateUUid()
	sd = NewSessionDataMem(uid)
	return
}
func (m *MgrMemory)Init(addr string,options ...string)  {}
func (m *MgrMemory)GetSessionData(id string) (sd SessionData,ok bool) {
	if v,ok := m.SessionData[id];ok{
		return v,ok
	}
	return nil,ok
}
func (m *MgrMemory)SetSessionData(uid string ) (err bool) {
	if _,ok := m.SessionData[uid];ok{
		return ok
	}
	//uid := m.CreateUUid()
	m.SessionData[uid] = NewSessionDataMem(uid)
	return true
}
func (m *MgrMemory)DeleteSessionData(uid string ) (err bool) {
	if _,ok := m.SessionData[uid];!ok{
		fmt.Printf("%s不存在",uid)
		return ok
	}
	delete(m.SessionData,uid)
	return true
}
