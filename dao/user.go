package dao

import (
	"bookManager/model"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserDAOInterface interface {
	Add(*model.User) error
	Has(*model.User) error
	VerifyAccountPassword(*model.User) error
	UpdateToken(*model.User) (error, string)
	GetToken(*model.User) (err error)
}

// 封装一个db结构体，方便实现各种入库方法
type UserDAO struct {
	db *gorm.DB
}

// 将全局的mysql连接对象封装成自定义对象
func NewUserDAO(dbm *gorm.DB) UserDAOInterface {
	return &UserDAO{dbm}
}

// user注册实现
func (UserDAO *UserDAO) Add(req *model.User) error {
	if tx := UserDAO.db.Create(req); tx.Error != nil {
		return tx.Error
	}
	return nil
}

// 用户是否存在
func (UserDAO *UserDAO) Has(req *model.User) error {
	//如果用户已经存在则报错
	tx := UserDAO.db.Where("username = ?", req.Username).First(req)
	if tx.Error != nil && tx.Error == gorm.ErrRecordNotFound {
		return nil
	}
	//其他情况则表示用户已存在
	return errors.New("用户已存在")
}

// 验证账号密码
func (UserDAO *UserDAO) VerifyAccountPassword(req *model.User) error {
	tx := UserDAO.db.Where(&req).First(&req)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// 更新token
func (UserDAO *UserDAO) UpdateToken(req *model.User) (err error, Token string) {
	Token = uuid.New().String()
	if tx := UserDAO.db.Model(req).Update("token", Token); tx.Error != nil {
		return tx.Error, ""
	}
	return nil, Token
}

// 获取Token
func (UserDAO *UserDAO) GetToken(req *model.User) (err error) {
	//将请求头中的token查询是否存在
	tx := UserDAO.db.Where("token = ?", req.Token).First(req)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return errors.New("Token不存在")
	}
	return nil
}
