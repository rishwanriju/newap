package service

import "newsapp/entity"

type NewsService interface {
	Save(entity.News) entity.News
	FindAll() []entity.News
}

type newsService struct {
	news []entity.News
}

func New() NewsService {

	return &newsService{}
}

func (service *newsService) Save(newss entity.News) entity.News {

	service.news = append(service.news, newss)

	return newss

}
func (service *newsService) FindAll() []entity.News {

	return service.news
}
