package activityservice

import (
	activitycore "devcode/domains/activity/core"
	"devcode/exceptions"
	"fmt"

	"gorm.io/gorm"
)

type activityService struct {
	repo activitycore.IRepoActivity
}

func New(repo activitycore.IRepoActivity) *activityService {
	return &activityService{
		repo: repo,
	}
}

func (s *activityService) Create(activityCore activitycore.Core) activitycore.Core {
	result, err := s.repo.Insert(activityCore)

	if err != nil {
		panic(exceptions.NewInternalServerError(err.Error()))
	}

	return result
}

func (s *activityService) Update(activityCore activitycore.Core) activitycore.Core {
	result, err := s.repo.Update(activityCore)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(exceptions.NewNotFoundError(err.Error()))

		} else {
			msg := fmt.Sprintf("Activity with ID %d Not Found", activityCore.Id)
			panic(exceptions.NewInternalServerError(msg))
		}
	}

	return result
}

func (s *activityService) Delete(activityCore activitycore.Core) activitycore.Core {
	result, err := s.repo.Delete(activityCore)

	if err != nil {
		panic(exceptions.NewInternalServerError(err.Error()))
	}

	return result
}

func (s *activityService) FindAll() []activitycore.Core {
	result, err := s.repo.GetAll()

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(exceptions.NewNotFoundError(err.Error()))

		} else {
			panic(exceptions.NewInternalServerError(err.Error()))
		}
	}

	return result
}

func (s *activityService) FindSingle(activityCore activitycore.Core) activitycore.Core {
	result, err := s.repo.GetSingle(activityCore)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(exceptions.NewNotFoundError(err.Error()))

		} else {
			msg := fmt.Sprintf("Activity with ID %d Not Found", activityCore.Id)
			panic(exceptions.NewInternalServerError(msg))
		}
	}

	return result
}
