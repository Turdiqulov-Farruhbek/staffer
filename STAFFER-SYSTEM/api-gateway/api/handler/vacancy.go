package handler

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	cp "gateway/genprotos"
)

// VacancyCreate handles the creation of a new vacancy.
// @Summary Create vacancy
// @Description Create a new vacancy
// @Tags vacancy
// @Accept json
// @Produce json
// @Param vacancy body cp.VacancyCreate true "Vacancy data"
// @Success 200 {object} string "Vacancy created"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /vacancies [post]
func (h *Handler) CreateVacancyHandler(c *gin.Context) {
	var req cp.VacancyCreate

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	_, err := h.srvs.VacancyService.Create(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, "Vacancy created")
}

// GetByIdHandler handles the retrieval of a vacancy by ID.
// @Summary Get vacancy by ID
// @Description Get a vacancy by its ID
// @Tags vacancy
// @Accept json
// @Produce json
// @Param id path string true "Vacancy ID"
// @Success 200 {object} cp.VacancyGetResUpdateReq
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /vacancies/{id} [get]
func (h *Handler) GetByIdVacancyHandler(c *gin.Context) {
	id := &cp.Byid{Id: c.Param("id")}
	res, err := h.srvs.VacancyService.GetByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get vacancy", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetAllHandler handles getting all vacancies.
// @Summary Get all vacancies
// @Description Get all vacancies
// @Tags vacancy
// @Accept json
// @Produce json
// @Param title query string false "Title"
// @Param description query string false "Description"
// @Param department_id query string false "Department ID"
// @Param position_id query string false "Position ID"
// @Param status query string false "Status"
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} cp.VacancyGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /vacancies/all [get]
func (h *Handler) GetAllVacancyHandler(c *gin.Context) {
	req := cp.VacancyGetAllReq{
		Title:        c.Query("title"),
		Description:  c.Query("description"),
		DepartmentId: c.Query("department_id"),
		PositionId:   c.Query("position_id"),
		Status:       c.Query("status"),
		Pagination:   &cp.Pagination{},
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

	req.Pagination.Limit = int64(limit)
	req.Pagination.Offset = int64(offset)
	
	res, err := h.srvs.VacancyService.GetAll(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get vacancies", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateHandler handles updating an existing vacancy.
// @Summary Update vacancy
// @Description Update an existing vacancy
// @Tags vacancy
// @Accept json
// @Produce json
// @Param id path string true "Vacancy ID"
// @Param vacancy body cp.VacancyGetResUpdateReq true "Updated vacancy data"
// @Success 200 {object} string "Vacancy updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Vacancy not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /vacancies/{id} [put]
func (h *Handler) UpdateVacancyHandler(c *gin.Context) {
	var req cp.VacancyGetResUpdateReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.Id = c.Param("id")
	_, err := h.srvs.VacancyService.Update(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update vacancy", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Vacancy updated")
}

// DeleteHandler handles deleting a vacancy by ID.
// @Summary Delete vacancy
// @Description Delete a vacancy by ID
// @Tags vacancy
// @Accept json
// @Produce json
// @Param id path string true "Vacancy ID"
// @Success 200 {object} string "Vacancy deleted"
// @Failure 400 {object} string "Invalid vacancy ID"
// @Failure 404 {object} string "Vacancy not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /vacancies/{id} [delete]
func (h *Handler) DeleteVacancyHandler(c *gin.Context) {
	id := &cp.Byid{Id: c.Param("id")}
	_, err := h.srvs.VacancyService.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete vacancy", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Vacancy deleted"})
}

// VacancyApplicationHandler handles getting applications for a specific vacancy.
// @Summary Get vacancy applications
// @Description Get applications for a specific vacancy
// @Tags vacancy
// @Accept json
// @Produce json
// @Param id path string true "Vacancy ID"
// @Success 200 {object} cp.VacancyApplicationsRes
// @Failure 400 {object} string "Invalid vacancy ID"
// @Failure 404 {object} string "Vacancy not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /vacancies/{id}/applications [get]
func (h *Handler) VacancyApplicatHundler(c *gin.Context) {
	id := &cp.Byid{Id: c.Param("id")}
	res, err := h.srvs.VacancyService.GetApplications(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get applications", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// VacancyOffersHandler handles getting offers for a specific vacancy.
// @Summary Get vacancy offers
// @Description Get offers for a specific vacancy
// @Tags vacancy
// @Accept json
// @Produce json
// @Param id path string true "Vacancy ID"
// @Success 200 {object} cp.VacancyOffersRes
// @Failure 400 {object} string "Invalid vacancy ID"
// @Failure 404 {object} string "Vacancy not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /vacancies/{id}/offers [get]
func (h *Handler) VacanyOffersHundler(c *gin.Context) {
	id := &cp.Byid{Id: c.Param("id")}
	res, err := h.srvs.VacancyService.GetOffers(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get offers", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
