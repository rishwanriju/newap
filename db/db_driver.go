package db

import (
	"newsapp/entity"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type NewsRepository interface {
	Save(news entity.News)
	Update(news entity.News)
	Delete(news entity.News)
	FindAll() []entity.News
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewNewsRepository() NewsRepository {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&entity.News{})
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

func (db *database) Save(news entity.News) {
	db.connection.Create(&news)
}

func (db *database) Update(news entity.News) {
	db.connection.Save(&news)
}

func (db *database) Delete(news entity.News) {
	db.connection.Delete(&news)
}

func (db *database) FindAll() []entity.News {
	var news []entity.News
	db.connection.Set("gorm:auto_preload", true).Find(&news)
	return news
}
