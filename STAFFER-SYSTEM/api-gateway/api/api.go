package api

import (
	"gateway/api/handler"
	"gateway/grpc/clients"

	_ "gateway/docs"
	_ "gateway/genprotos"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "gateway/docs"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(connC *clients.GrpcClients) *gin.Engine {
	h := handler.NewHandler(connC)

	router := gin.Default()

	router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	company := router.Group("/companies")
	{
		company.POST("/", h.CreateCompanyHandler)
		company.GET("/:id", h.GetByIdCompanyHandler)
		company.PUT("/:id", h.UpdateCompanyHandler)
		company.DELETE("/:id", h.DeleteCompanyHandler)
		company.GET("/all", h.GetAllCompaniesHandler)
	}


	jobApplications := router.Group("/applications")
	{
		jobApplications.POST("/", h.CreateJobAppHandler)
		jobApplications.POST("/:id/steps", h.CreateJobStepHandler)
		
		jobApplications.GET("/:id", h.GetByIdJobAppHandler)
		jobApplications.GET("/:id/steps", h.GetJobStepHandler)
		jobApplications.GET("/all", h.GetAllJobAppsHandler)

		jobApplications.PUT("/:applicationId/steps/:stepId", h.UpdateJobStepHandler)
		jobApplications.PUT("/jobApp/:id/confirm", h.UpdateJobAppHandler)

		jobApplications.DELETE("/:applicationId/steps/:stepId", h.DeletedJobStepHandler)
		jobApplications.DELETE("/jobApp/:id", h.DeleteJobAppHandler)
	}
	

	jobOffer := router.Group("/offers")
	{
		jobOffer.POST("/", h.CreateJobOfferHandler)
		jobOffer.GET("/:id", h.GetByIdJobOfferHandler)
		jobOffer.GET("/all", h.GetAllJobOffersHandler)
		jobOffer.PUT("/:id/confirm", h.UpdateJobOfferHandler)
		jobOffer.DELETE("/:id", h.DeleteJobOfferHandler)

	}
	
	
	vacancy := router.Group("/vacancies")
	{
		vacancy.POST("/", h.CreateVacancyHandler)
		vacancy.GET("/:id", h.GetByIdVacancyHandler)
		vacancy.GET("/all", h.GetAllVacancyHandler)
		vacancy.PUT("/:id", h.UpdateVacancyHandler)
		vacancy.DELETE("/:id", h.DeleteVacancyHandler)
		vacancy.GET("/Vacancy/:id/applications", h.VacancyApplicatHundler)
		vacancy.GET("/offers/:id", h.VacanyOffersHundler)
	}

	return router
}
