package item

import (
	"midterm-api/internal/constant"
	"midterm-api/internal/model"

	"gorm.io/gorm"
)

type Service struct {
	Repository Repository
	Validate Validate
}

func NewService(db *gorm.DB) Service {
	return Service{
		Repository: NewRepository(db),
	}
}

func (service Service) Create(req model.RequestItem) (model.Item, error) {
	item := model.Item {
		Title: req.Title,
		Amount: req.Amount,
		Quantity: req.Quantity,
		Status: constant.ItemPendingStatus,
	}

	if err := service.Repository.Create(&item); err != nil {
		return model.Item{}, err
	}

	return item, nil
}

func (service Service) Find(query model.RequestFindItem) ([]model.Item, error) {
	return service.Repository.Find(query)
}

func (service Service) UpdateStatus(id uint, status constant.ItemStatus) (model.Item, error) {
    
    item, err := service.Repository.FindByID(id)
    if err != nil {
        return model.Item{}, err
    }
    
    item.Status = status
    
    if err := service.Repository.Replace(item); err != nil {
        return model.Item{}, err
    }
    return item, nil
}

func (service Service) FindbyId(id uint) (model.Item, error) {
	query, err := service.Repository.FindByID(id)
    if err != nil {
        return model.Item{}, err
    }
	return query, nil
}

func (service Service) DeleteByID(id uint) (model.Item, error) {
    
    item, err := service.Repository.FindByID(id)
    if err != nil {
        return model.Item{}, err
    }
       
    if err := service.Repository.Delete(item); err != nil {
        return model.Item{}, err
    }
    return item, nil
}

func (service Service) UpdateItem(id uint, req model.RequestItem) (model.Item, error) {
	item, err := service.Repository.FindByID(id)
    if err != nil {
        return model.Item{}, err
    }
    
    if err := service.Validate.UpdateItem(item.Status); err != nil {
		return model.Item{}, err
	}

	data := model.Item {
		ID: item.ID,
		Title:    req.Title,
		Amount:   req.Amount,
		Quantity: req.Quantity,
		Status: item.Status,
	}

	if err := service.Repository.Replace(data); err != nil {
		return model.Item{}, err
	}
    return data, nil
}
