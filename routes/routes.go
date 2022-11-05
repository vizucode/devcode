package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	activityhandler "devcode/domains/activity/handler"
	activityrepo "devcode/domains/activity/repository"
	activityservice "devcode/domains/activity/service"

	todohandler "devcode/domains/todo/handler"
	todorepo "devcode/domains/todo/repository"
	todoservice "devcode/domains/todo/service"
)

func InitRoutes(ctx *fiber.App, db *gorm.DB) {
	/*
		Dependency Injection
	*/

	activityRepo := activityrepo.New(db)
	activityService := activityservice.New(activityRepo)
	activityHandler := activityhandler.New(activityService)

	todoRepo := todorepo.New(db)
	todoService := todoservice.New(todoRepo)
	todoHandler := todohandler.New(todoService)

	/*
		Routes
	*/

	ctx.Post("/activity-groups", activityHandler.Create)
	ctx.Patch("/activity-groups/:id", activityHandler.Update)
	ctx.Delete("/activity-groups/:id", activityHandler.Delete)
	ctx.Get("/activity-groups", activityHandler.FindAll)
	ctx.Get("/activity-groups/:id", activityHandler.FindSingle)

	ctx.Post("/todo-items", todoHandler.Create)
	ctx.Patch("/todo-items/:id", todoHandler.Update)
	ctx.Delete("/todo-items/:id", todoHandler.Delete)
	ctx.Get("/todo-items", todoHandler.FindAll)
	ctx.Get("/todo-items/:id", todoHandler.FindById)
}
