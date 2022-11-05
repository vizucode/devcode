package activitymodel

import (
	activitycore "devcode/domains/activity/core"

	"gorm.io/gorm"
)

type Activity struct {
	gorm.Model
	Title string
	Email string
}

func ToCore(model Activity) activitycore.Core {
	return activitycore.Core{
		Id:        model.ID,
		Email:     model.Email,
		Title:     model.Title,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		DeletedAt: model.DeletedAt.Time,
	}
}

func ToModel(Core activitycore.Core) Activity {
	return Activity{
		Title: Core.Title,
		Email: Core.Email,
	}
}
