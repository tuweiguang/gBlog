package session

const (
	SessionExist   = iota // 存在但未过期
	SessionNoexist        // 不存在
	SessionExpire         // 存在但过期
)

const SessionMapSize = 8
