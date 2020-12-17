package service

import (
	"newsapp/db"
	"newsapp/entity"
)

type NewsService interface {
	Save(entity.News) entity.News
	FindAll() []entity.News
	Update(news entity.News)
	Delete(news entity.News)
}

type newsService struct {
	newsRepository db.NewsRepository
}

func New(repo db.NewsRepository) NewsService {

	return &newsService{

		newsRepository: repo,
	}
}

func (service *newsService) Save(newss entity.News) entity.News {

	service.newsRepository.Save(newss)

	return newss

}
func (service *newsService) FindAll() []entity.News {

	return service.newsRepository.FindAll()
}

func (service *newsService) Update(news entity.News) {
	service.newsRepository.Update(news)
}
func (service *newsService) Delete(news entity.News) {
	service.newsRepository.Delete(news)
}
