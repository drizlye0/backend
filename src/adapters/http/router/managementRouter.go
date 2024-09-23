package router

import (
	"github.com/Dialosoft/src/adapters/http/controller"
	"github.com/Dialosoft/src/adapters/http/middleware"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type ManagementRouter struct {
	ManagementController *controller.ManagementController
}

func NewManagementRouter(managementController *controller.ManagementController) *ManagementRouter {
	return &ManagementRouter{ManagementController: managementController}
}

func (r *ManagementRouter) SetupManagementRoutes(api fiber.Router, middlewares *middleware.SecurityMiddleware, defaultRoles map[string]uuid.UUID) {
	managementGroup := api.Group("/management")

	{
		managementGroup.Post("/change-user-role/:id", r.ManagementController.ChangeUserRole,
			middlewares.GetAndVerifyAccesToken(),
			middlewares.VerifyRefreshToken(),
			middlewares.RoleRequiredByID(defaultRoles["administrator"].String()),
		)
	}
}
