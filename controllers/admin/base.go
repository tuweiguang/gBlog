package admin

// 后台登录信息
type LoginInfo struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// 后台添加新用户信息
type RegisterInfo struct {
	Username   string `form:"username" binding:"required"`
	Password   string `form:"password"`
	RePassword string `form:"repassword"`
	Email      string `form:"email" binding:"required"`
}
