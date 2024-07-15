package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	cp "gateway/genprotos"

	"github.com/gin-gonic/gin"
)

// CreateJobApplicationHandler handles the creation of a new job application.
// @Summary Create job application
// @Description Create a new job application
// @Tags jobapplication
// @Accept json
// @Produce json
// @Param job_application body cp.JobApplicationCreate true "Job Application data"
// @Success 200 {object} string "Job application created"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /job-applications [post]
func (h *Handler) CreateJobApplicationHandler(c *gin.Context) {
	var req cp.JobApplicationCreate

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	_, err := h.srvs.JobsService.Create(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, "Job application created")
}

// GetJobApplicationHandler handles the retrieval of a job application by ID.
// @Summary Get job application by ID
// @Description Get a job application by its ID
// @Tags jobapplication
// @Accept json
// @Produce json
// @Param id path string true "Job Application ID"
// @Success 200 {object} cp.JobApplitcationGetRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /job-applications/{id} [get]
func (h *Handler) GetJobApplicationHandler(c *gin.Context) {
	id := &cp.Byid{Id: c.Param("id")}
	res, err := h.srvs.JobsService.Get(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get job application", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetAllJobApplicationsHandler handles getting all job applications.
// @Summary Get all job applications
// @Description Get all job applications
// @Tags jobapplication
// @Accept json
// @Produce json
// @Param candidate_id query string false "Candidate ID"
// @Param vacancy_id query string false "Vacancy ID"
// @Param resume_id query string false "Resume ID"
// @Param status query string false "Status"
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} cp.JobApplicationGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /job-applications/all [get]
func (h *Handler) GetAllJobApplicationsHandler(c *gin.Context) {
	req := cp.JobApplicationGetAll{
		CandidateId: c.Query("candidate_id"),
		VacancyId:   c.Query("vacancy_id"),
		ResumeId:    c.Query("resume_id"),
		Status:      c.Query("status"),
		Filter:      &cp.Pagination{},
	}

	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	var limit, offset int
	var err error
	if limitStr == "" {
		limit = 0
	} else {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
			return
		}
	}
	if offsetStr == "" {
		offset = 0
	} else {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset parameter"})
			return
		}
	}

	req.Filter.Limit = int64(limit)
	req.Filter.Offset = int64(offset)

	res, err := h.srvs.JobsService.GetAll(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get job applications", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateJobApplicationHandler handles updating an existing job application.
// @Summary Update job application
// @Description Update an existing job application
// @Tags jobapplication
// @Accept json
// @Produce json
// @Param id path string true "Job Application ID"
// @Param job_application body cp.JobApplicationUpdate true "Updated job application data"
// @Success 200 {object} string "Job application updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Job application not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /job-applications/{id} [put]
func (h *Handler) UpdateJobApplicationHandler(c *gin.Context) {
	var req cp.JobApplicationUpdate

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.Id = c.Param("id")
	_, err := h.srvs.JobsService.Update(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update job application", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Job application updated")
}

// DeleteJobApplicationHandler handles deleting a job application by ID.
// @Summary Delete job application
// @Description Delete a job application by ID
// @Tags jobapplication
// @Accept json
// @Produce json
// @Param id path string true "Job Application ID"
// @Success 200 {object} string "Job application deleted"
// @Failure 400 {object} string "Invalid job application ID"
// @Failure 404 {object} string "Job application not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /job-applications/{id} [delete]
func (h *Handler) DeleteJobApplicationHandler(c *gin.Context) {
	id := &cp.Byid{Id: c.Param("id")}
	_, err := h.srvs.JobsService.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete job application", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Job application deleted"})
}
