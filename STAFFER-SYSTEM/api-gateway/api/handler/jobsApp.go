package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	cp "gateway/genprotos"

	"github.com/gin-gonic/gin"
)

// CreateJobAppHandler creates a new job application.
// @Summary Create job application
// @Description Create a new job application
// @Tags job_applications
// @Accept json
// @Produce json
// @Param jobApplication body cp.CreateJobApplicationRequest true "Job Application Data"
// @Success 200 {object} cp.JobApplicationResponse "Job Application created"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /applications [post]
func (h *Handler) CreateJobAppHandler(c *gin.Context) {
	var req cp.CreateJobApplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.srvs.JobsAppService.CreateJobApplication(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, "CreateJobApplication created")
}

// CreateJobStepHandler creates a new job step for a job application.
// @Summary Create job step
// @Description Create a new job step for a job application
// @Tags job_applications
// @Accept json
// @Produce json
// @Param id path string true "Job Application ID"
// @Param jobStep body cp.CreateJobStepRequest true "Job Step Data"
// @Success 200 {object} cp.JobStepResponse "Job Step created"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /applications/{id}/steps [post]
func (h *Handler) CreateJobStepHandler(c *gin.Context) {
	var req cp.CreateJobStepRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.ApplicationId = c.Param("id")
	_, err := h.srvs.JobsAppService.CreateJobStep(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, "CreateJobStep created")
}

// GetByIdJobAppHandler retrieves a job application by ID.
// @Summary Get job application by ID
// @Description Get a job application by its ID
// @Tags job_applications
// @Accept json
// @Produce json
// @Param id path string true "Job Application ID"
// @Success 200 {object} cp.JobApplicationResponse "Job Application data"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /applications/{id} [get]
func (h *Handler) GetByIdJobAppHandler(c *gin.Context) {
	req := &cp.GetJobApplicationByIdRequest{Id: c.Param("id")}
	res, err := h.srvs.JobsAppService.GetJobApplicationById(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetJobStepHandler retrieves job steps for a job application.
// @Summary Get job steps
// @Description Get job steps for a job application
// @Tags job_applications
// @Accept json
// @Produce json
// @Param id path string true "Job Application ID"
// @Success 200 {object} cp.JobStepResponse "Job Steps data"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /applications/{id}/steps [get]
func (h *Handler) GetJobStepHandler(c *gin.Context) {
	req := &cp.GetJobStepByIdRequest{ApplicationId: c.Param("id")}
	res, err := h.srvs.JobsAppService.GetJobStepById(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetAllJobAppsHandler retrieves all job applications.
// @Summary Get all job applications
// @Description Get all job applications
// @Tags job_applications
// @Accept json
// @Produce json
// @Param candidate_id query string false "Candidate ID"
// @Param vacancy_id query string false "Vacancy ID"
// @Param resume_id query string false "Resume ID"
// @Param status query string false "Status"
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} cp.GetJobApplicationsRequest "Job Applications data"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Router /applications/all [get]
func (h *Handler) GetAllJobAppsHandler(c *gin.Context) {
	req := &cp.GetJobApplicationsRequest{
		CandidateId: c.Query("candidate_id"),
		VacancyId:   c.Query("vacancy_id"),
		ResumeId:    c.Query("resume_id"),
		Status:      c.Query("status"),
		Pagination:  &cp.Pagination{},
	}

	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	var limit, offset int
	var err error
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
			return
		}
	}
	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
			return
		}
	}

	req.Pagination.Limit = int64(limit)
	req.Pagination.Offset = int64(offset)

	res, err := h.srvs.JobsAppService.GetAllJobApplications(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateJobAppHandler updates a job application.
// @Summary Update job application
// @Description Update an existing job application
// @Tags job_applications
// @Accept json
// @Produce json
// @Param id path string true "Job Application ID"
// @Param jobApplication body cp.UpdateJobApplicationRequest true "Job Application Data"
// @Success 200 {object} string "Job Application updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Job Application not found"
// @Failure 500 {object} string "Server error"
// @Router /applications/{id}/confirm [put]
func (h *Handler) UpdateJobAppHandler(c *gin.Context) {
	var req cp.UpdateJobApplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Id = c.Param("id")
	_, err := h.srvs.JobsAppService.UpdateJobApplication(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, "Job application updated")
}

// UpdateJobStepHandler updates a job step for a job application.
// @Summary Update job step
// @Description Update an existing job step for a job application
// @Tags job_applications
// @Accept json
// @Produce json
// @Param applicationId path string true "Job Application ID"
// @Param stepId path string true "Job Step ID"
// @Param jobStep body cp.UpdateJobStepRequest true "Job Step Data"
// @Success 200 {object} string "Job Step updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Job Step not found"
// @Failure 500 {object} string "Server error"
// @Router /applications/{applicationId}/steps/{stepId} [put]
func (h *Handler) UpdateJobStepHandler(c *gin.Context) {
	var req cp.UpdateJobStepRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.StepId = c.Param("stepId")
	_, err := h.srvs.JobsAppService.UpdateJobStep(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, "Job step updated")
}

// DeleteJobAppHandler deletes a job application.
// @Summary Delete job application
// @Description Delete a job application by its ID
// @Tags job_applications
// @Accept json
// @Produce json
// @Param id path string true "Job Application ID"
// @Success 200 {object} string "Job Application deleted"
// @Failure 400 {object} string "Invalid job application ID"
// @Failure 404 {object} string "Job Application not found"
// @Failure 500 {object} string "Server error"
// @Router /applications/{id} [delete]
func (h *Handler) DeleteJobAppHandler(c *gin.Context) {
	req := &cp.DeleteJobApplicationRequest{Id: c.Param("id")}
	_, err := h.srvs.JobsAppService.DeleteJobApplication(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, "Job application deleted")
}

// DeletedJobStepHandler deletes a job step for a job application.
// @Summary Delete job step
// @Description Delete a job step for a job application by its ID
// @Tags job_applications
// @Accept json
// @Produce json
// @Param applicationId path string true "Job Application ID"
// @Param stepId path string true "Job Step ID"
// @Success 200 {object} string "Job Step deleted"
// @Failure 400 {object} string "Invalid job step ID"
// @Failure 404 {object} string "Job Step not found"
// @Failure 500 {object} string "Server error"
// @Router /applications/{applicationId}/steps/{stepId} [delete]
func (h *Handler) DeletedJobStepHandler(c *gin.Context) {
	req := &cp.DeleteJobStepRequest{
		ApplicationId: c.Param("applicationId"),
		StepId:        c.Param("stepId"),
	}
	_, err := h.srvs.JobsAppService.DeleteJobStep(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, "Job step deleted")
}
