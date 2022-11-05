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
	tx := r.db.Model(activityModel.Activity{}).Where("id", activityCore.Id).Updates(&model)
	if tx.Error != nil {
		return activitycore.Core{}, tx.Error
	}

	if tx.RowsAffected < 1 {
		return activitycore.Core{}, gorm.ErrRecordNotFound
	}

	tx = r.db.Model(activityModel.Activity{}).First(&model)
	if tx.Error != nil {
		return activitycore.Core{}, tx.Error
	}

	return activityModel.ToCore(model), nil
}

func (r *activityRepo) Delete(activityCore activitycore.Core) (activitycore.Core, error) {
	model := activityModel.ToModel(activityCore)
	tx := r.db.Where("id", activityCore.Id).Delete(&model)

	if tx.Error != nil {
		return activitycore.Core{}, tx.Error
	}

	if tx.RowsAffected < 1 {
		return activitycore.Core{}, gorm.ErrRecordNotFound
	}

	return activityModel.ToCore(model), nil
}

func (r *activityRepo) GetAll() ([]activitycore.Core, error) {
	model := []activityModel.Activity{}
	tx := r.db.Unscoped().Find(&model)

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
	tx := r.db.Unscoped().First(&model, activityCore.Id)

	if tx.Error != nil {
		return activitycore.Core{}, tx.Error
	}

	return activityModel.ToCore(model), nil
}
