package models

// 学生
type Student struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Id         string `json:"id"`
	CreateTime string `json:"createTime"`
	Email      string `json:"email"`
}
