package handler

import (
	"context"
	"log"
	"net/http"

	cp "gateway/genprotos"

	"github.com/gin-gonic/gin"
)

// CreateJobOfferHandler handles the creation of a new job offer.
// @Summary Create job offer
// @Description Create a new job offer
// @Tags joboffer
// @Accept json
// @Produce json
// @Param job_offer body cp.CreateJobOfferRequest true "Job Offer data"
// @Success 200 {object} string "Job offer created"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /offers [post]
func (h *Handler) CreateJobOfferHandler(c *gin.Context) {
	var req cp.CreateJobOfferRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	_, err := h.srvs.JobsOfferService.CreateJobOffer(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, "Job offer created")
}

// GetByIdJobOfferHandler handles the retrieval of a job offer by ID.
// @Summary Get job offer by ID
// @Description Get a job offer by its ID
// @Tags joboffer
// @Accept json
// @Produce json
// @Param id path string true "Job Offer ID"
// @Success 200 {object} cp.JobOfferResponse
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /offers/{id} [get]
func (h *Handler) GetByIdJobOfferHandler(c *gin.Context) {
	id := &cp.GetByIdJobOfferRequest{Id: c.Param("id")}
	res, err := h.srvs.JobsOfferService.GetByIdJobOffer(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get job offer", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetAllJobOffersHandler handles getting all job offers.
// @Summary Get all job offers
// @Description Get all job offers
// @Tags joboffer
// @Accept json
// @Produce json
// @Success 200 {object} cp.JobOffersResponse
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /offers/all [get]
func (h *Handler) GetAllJobOffersHandler(c *gin.Context) {
	req := &cp.GetAllJobOffersRequest{}

	res, err := h.srvs.JobsOfferService.GetAllJobOffers(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get job offers", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateJobOfferHandler handles updating an existing job offer.
// @Summary Update job offer
// @Description Update an existing job offer
// @Tags joboffer
// @Accept json
// @Produce json
// @Param id path string true "Job Offer ID"
// @Param job_offer body cp.UpdateJobOfferRequest true "Updated job offer data"
// @Success 200 {object} string "Job offer updated"
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Job offer not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /offers/{id}/confirm [put]
func (h *Handler) UpdateJobOfferHandler(c *gin.Context) {
	var req cp.UpdateJobOfferRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.Id = c.Param("id")
	_, err := h.srvs.JobsOfferService.UpdateJobOffer(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update job offer", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Job offer updated")
}

// DeleteJobOfferHandler handles deleting a job offer by ID.
// @Summary Delete job offer
// @Description Delete a job offer by ID
// @Tags joboffer
// @Accept json
// @Produce json
// @Param id path string true "Job Offer ID"
// @Success 200 {object} string "Job offer deleted"
// @Failure 400 {object} string "Invalid job offer ID"
// @Failure 404 {object} string "Job offer not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /offers/{id} [delete]
func (h *Handler) DeleteJobOfferHandler(c *gin.Context) {
	id := &cp.DeleteJobOfferRequest{Id: c.Param("id")}
	_, err := h.srvs.JobsOfferService.DeleteJobOffer(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete job offer", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Job offer deleted"})
}
