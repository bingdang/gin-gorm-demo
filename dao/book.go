package dao

import (
	"bookManager/model"
	"errors"
	"gorm.io/gorm"
	"log"
)

type BookDAOInterface interface {
	List() ([]*model.Book, error)
	Add(*model.Book) error
	Has(*model.Book) (*model.Book, bool, error)
	Get(*model.Book) (*model.Book, bool, error)
	Update(*model.Book) error
	Delete(*model.Book) error
}

// 封装一个db结构体，方便实现各种入库方法
type BookDAO struct {
	db *gorm.DB
}

// 将全局的mysql连接对象封装成自定义对象
func NewBookDAO(dbm *gorm.DB) BookDAOInterface {
	return &BookDAO{dbm}
}

// 查询书籍列表
func (BookDAO *BookDAO) List() (books []*model.Book, err error) {
	//同时加载与用户相关联的 Users 数据。
	tx := BookDAO.db.Preload("Users").Find(&books)
	if tx.Error != nil {
		log.Printf("查询书籍列表失败%s", tx.Error.Error())
		return nil, errors.New("查询书籍列表失败" + tx.Error.Error())
	}
	return books, nil
}

// 新增书籍
func (BookDAO *BookDAO) Add(book *model.Book) error {
	//一对多或多对多创建时，如果book中有user数据时会自动关联创建user数据
	//在真实场景中 创建书籍不可能同时需要创建用户。所以用Omit跳过关联创建Users
	tx := BookDAO.db.Omit("Users").Create(book)
	if tx.Error != nil {
		log.Printf("创建书籍失败%s", tx.Error.Error())
		return errors.New("创建书籍失败" + tx.Error.Error())
	}

	//书籍关联用户
	if len(book.Users) > 0 {
		err := BookDAO.db.Model(book).Association("Users").Append(book.Users)
		if err != nil {
			log.Printf("书籍关联用户失败%s", err.Error())
			return errors.New("书籍关联用户失败: " + err.Error())
		}
	}

	return nil
}

// 根据书名查询单个
func (BookDAO *BookDAO) Has(req *model.Book) (*model.Book, bool, error) {
	resp := &model.Book{
		Name: req.Name,
	}
	//为了拿来创建书之前看看存不存在。如果存在就不创建所以不用关联查询
	tx := BookDAO.db.Where("name = ?", req.Name).First(resp)
	//如果报错是为空，那么返回false
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, false, tx.Error
	}
	if tx.Error != nil {
		log.Printf("基于书名查询书籍失败%s", tx.Error.Error())
		return nil, false, errors.New("基于书名查询书籍失败" + tx.Error.Error())
	}
	return resp, true, errors.New("书籍已存在")
}

// 根据id查询单个
func (BookDAO *BookDAO) Get(req *model.Book) (*model.Book, bool, error) {
	resp := &model.Book{
		ID: req.ID,
	}
	tx := BookDAO.db.Where("id = ?", req.ID).First(resp)
	//如果报错是为空，那么返回false
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, false, errors.New("书籍不存在")
	}
	if tx.Error != nil {
		log.Printf("基于id查询书籍失败%s", tx.Error.Error())
		return nil, false, errors.New("基于id查询书籍失败" + tx.Error.Error())
	}
	return resp, true, nil
}

// 更新书籍
func (BookDAO *BookDAO) Update(book *model.Book) error {
	tx := BookDAO.db.Model(&model.Book{}).Where("id = ?", book.ID).Updates(book)
	if tx.Error != nil {
		log.Printf("更新书籍失败%s", tx.Error.Error())
		return errors.New("更新书籍失败" + tx.Error.Error())
	}

	if len(book.Users) > 0 {
		//书籍更换用户
		err := BookDAO.db.Model(book).Association("Users").Replace(book.Users)
		if err != nil {
			log.Printf("书籍关联更新用户失败%s", err.Error())
			return errors.New("书籍关联更新用户失败: " + err.Error())
		}
	} else {
		//如果用户不属于任何人，则清空关联
		err := BookDAO.db.Model(book).Association("Users").Clear()
		if err != nil {
			log.Printf("书籍关联更新用户失败%s", err.Error())
			return errors.New("书籍关联更新用户失败: " + err.Error())
		}
	}
	return nil
}

// 删除书籍
func (BookDAO *BookDAO) Delete(book *model.Book) error {
	//先删除中间表的关联关系
	err := BookDAO.db.Model(&book).Association("Users").Clear()
	if err != nil {
		log.Printf("删除书籍关联关系失败%s", err.Error())
		return errors.New("删除书籍关联关系失败: " + err.Error())
	}
	//再删除书籍元数据
	//因为先删除书籍元数据后，无法通过书籍的id删除关联关系
	tx := BookDAO.db.Delete(book)
	if tx.Error != nil {
		log.Printf("删除书籍元数据失败%s", tx.Error.Error())
		return errors.New("删除书籍元数据失败" + tx.Error.Error())
	}
	return nil
}
