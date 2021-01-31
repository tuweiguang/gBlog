package session

import (
	"fmt"
	"gBlog/commom/config"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
)

type MemoryMgr struct {
	lock          sync.RWMutex
	sessionExpire int              // session过期时间
	sessionMap    map[string]int64 // sessionId -> expire time
}

var mOnce sync.Once
var memoryMgr MemoryMgr

func NewMemoryMgr() *MemoryMgr {
	mOnce.Do(func() {
		memoryMgr.sessionMap = make(map[string]int64, SessionMapSize)
		memoryMgr.sessionExpire = config.GetSessionConfig().Expire
	})

	return &memoryMgr
}

func (s *MemoryMgr) CheckSession(sessionId string) int {
	s.lock.RLock()
	defer s.lock.RUnlock()

	session, ok := s.sessionMap[sessionId]
	if !ok {
		return SessionNoexist
	}

	now := time.Now().Unix()
	if now-session < 0 {
		return SessionExist
	}
	return SessionExpire
}

func (s *MemoryMgr) DelSession(sessionId string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.sessionMap, sessionId)
}

func (s *MemoryMgr) CreateSessoin() string {
	s.lock.Lock()
	defer s.lock.Unlock()

	// 构造一个sessionId
	sessionId := uuid.NewV4().String()

	ts := time.Now().Add(time.Duration(s.sessionExpire) * time.Second).Unix()
	s.sessionMap[sessionId] = ts
	return sessionId
}

func (s *MemoryMgr) PrintSession() {
	s.lock.RLock()
	defer s.lock.RUnlock()

	now := time.Now().Unix()
	fmt.Printf("sessionMap:")
	for id, ts := range s.sessionMap {
		if ts-now <= 0 {
			delete(s.sessionMap, id)
			return
		}
		fmt.Printf("%v:%v ", id, ts-now)
	}
	fmt.Println()
}
