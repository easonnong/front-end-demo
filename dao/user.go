package dao

import (
	"log"

	"github.com/easonnong/front-end-demo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Mgr *Manager

type IManager interface {
	AddUser(user *models.User)
	ListUser() []models.User
}

type Manager struct {
	db *gorm.DB
}

//初始化
func init() {
	dsn := "root:2222@tcp(127.0.0.1:3306)/front-end-demo?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to init db:", err)
	}

	Mgr = &Manager{
		db: db,
	}
	db.AutoMigrate(&models.User{})
}

//增加用户
func (mgr *Manager) AddUser(user *models.User) {
	mgr.db.Create(user)
}

//查找用户
func (mgr *Manager) ListUser() []models.User {
	var users []models.User
	mgr.db.Find(&users)
	return users
}
