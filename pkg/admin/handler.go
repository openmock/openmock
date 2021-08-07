package admin

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/openmock/openmock"
	"github.com/openmock/openmock/swagger_gen/models"
	"github.com/openmock/openmock/swagger_gen/restapi/operations"
	"github.com/openmock/openmock/swagger_gen/restapi/operations/evaluate"
	"github.com/openmock/openmock/swagger_gen/restapi/operations/health"
	"github.com/openmock/openmock/swagger_gen/restapi/operations/template"
	"github.com/openmock/openmock/swagger_gen/restapi/operations/template_set"
)

// Setup initialize all the handler functions, returns whether admin HTTP
// should be enabled
func Setup(api *operations.OpenMockAPI, customOpenmock *openmock.OpenMock) bool {
	om := customOpenmock
	if om == nil {
		// Start openmock
		om = &openmock.OpenMock{}
		om.ParseEnv()
	}

	defer om.Stop()
	om.Start()

	// Setup swagger API
	setupHealth(api)
	setupCRUD(api, om)

	return om.AdminHTTPEnabled
}

func setupHealth(api *operations.OpenMockAPI) {
	api.HealthGetHealthHandler = health.GetHealthHandlerFunc(
		func(health.GetHealthParams) middleware.Responder {
			return health.NewGetHealthOK().WithPayload(&models.Health{Status: "OK"})
		},
	)
}

func setupCRUD(api *operations.OpenMockAPI, om *openmock.OpenMock) {
	c := NewCRUD(om)

	// Templates
	api.TemplateGetTemplatesHandler = template.GetTemplatesHandlerFunc(c.GetTemplates)
	api.TemplatePostTemplatesHandler = template.PostTemplatesHandlerFunc(c.PostTemplates)
	api.TemplateDeleteTemplatesHandler = template.DeleteTemplatesHandlerFunc(c.DeleteTemplates)
	api.TemplateDeleteTemplateHandler = template.DeleteTemplateHandlerFunc(c.DeleteTemplate)

	// Template sets
	api.TemplateSetDeleteTemplateSetHandler = template_set.DeleteTemplateSetHandlerFunc(c.DeleteTemplateSet)
	api.TemplateSetPostTemplateSetHandler = template_set.PostTemplateSetHandlerFunc(c.PostTemplateSet)

	// evaluate
	api.EvaluateEvaluateHandler = evaluate.EvaluateHandlerFunc(c.PostEvaluate)
}
