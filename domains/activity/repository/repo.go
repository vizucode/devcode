package activityrepo

import (
	activitycore "devcode/domains/activity/core"
	activityModel "devcode/domains/activity/model"

	"gorm.io/gorm"
)

type activityRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) *activityRepo {
	return &activityRepo{
		db: db,
	}
}

func (r *activityRepo) Insert(activityCore activitycore.Core) (activitycore.Core, error) {
	model := activityModel.ToModel(activityCore)
	tx := r.db.Create(&model)

	if tx.Error != nil {
		return activitycore.Core{}, tx.Error
	}

	return activityModel.ToCore(model), nil
}

func (r *activityRepo) Update(activityCore activitycore.Core) (activitycore.Core, error) {
	model := activityModel.ToModel(activityCore)
	tx := r.db.Where("id", model.ID).Updates(&model)

	if tx.Error != nil {
		return activitycore.Core{}, tx.Error
	}

	return activityModel.ToCore(model), nil
}

func (r *activityRepo) Delete(activityCore activitycore.Core) (activitycore.Core, error) {
	model := activityModel.ToModel(activityCore)
	tx := r.db.Delete(&model)

	if tx.Error != nil {
		return activitycore.Core{}, tx.Error
	}

	return activityModel.ToCore(model), nil
}

func (r *activityRepo) GetAll() ([]activitycore.Core, error) {
	model := []activityModel.Activity{}
	tx := r.db.Find(&model)

	if tx.Error != nil {
		return []activitycore.Core{}, tx.Error
	}

	coreList := []activitycore.Core{}
	for _, data := range model {
		coreList = append(coreList, activityModel.ToCore(data))
	}

	return coreList, nil
}

func (r *activityRepo) GetSingle(activityCore activitycore.Core) (activitycore.Core, error) {
	model := activityModel.ToModel(activityCore)
	tx := r.db.First(&model)

	if tx.Error != nil {
		return activitycore.Core{}, tx.Error
	}

	return activityModel.ToCore(model), nil
}
