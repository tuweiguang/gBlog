package admin

const LIMIT = 10 //一页10条记录

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

// 上传文章
type UploadArticle struct {
	Title string `form:"title" binding:"required"`
	Cate  int    `form:"cate" binding:"required"`
	Tag   string `form:"tag" binding:"required"`
	//File  string `form:"file" binding:"required"`
}

// 添加类别
type AddCate struct {
	CateId   int    `form:"pid"`
	CateName string `form:"name" binding:"required"`
	Sort     int    `form:"sort" binding:"required"`
}
