package model

type User struct {
	ID       int     `gorm:"primaryKey" json:"id"`
	Username string  `json:"username" binding:"required"`
	Password string  `json:"password" binding:"required"`
	Token    string  `json:"token"`
	Books    []*Book `gorm:"many2many:book_users"`
}

// 自定义表名
func (*User) TableName() string {
	return "user"
}
