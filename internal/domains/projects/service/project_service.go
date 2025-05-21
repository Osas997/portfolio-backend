package service

type ProjectService interface {
	FindAll()
	FindById(userId string)
	Create()
	Update(userId string)
	Delete(userId string)
}
