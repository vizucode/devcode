package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	activityhandler "devcode/domains/activity/handler"
	activityrepo "devcode/domains/activity/repository"
	activityservice "devcode/domains/activity/service"
)

func InitRoutes(ctx *fiber.App, db *gorm.DB) {
	/*
		Dependency Injection
	*/

	activityRepo := activityrepo.New(db)
	activityService := activityservice.New(activityRepo)
	activityHandler := activityhandler.New(activityService)

	/*
		Routes
	*/

	ctx.Post("/activity-groups", activityHandler.Create)
	ctx.Patch("/activity-groups/:id", activityHandler.Update)
	ctx.Delete("/activity-groups/:id", activityHandler.Delete)
	ctx.Get("/activity-groups", activityHandler.FindAll)
	ctx.Get("/activity-groups/:id", activityHandler.FindSingle)

}
