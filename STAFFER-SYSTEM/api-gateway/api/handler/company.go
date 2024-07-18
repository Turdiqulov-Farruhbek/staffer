package handler

import (
	"context"
	cp "gateway/genprotos"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CompanyCreate handles the creation of a new company.
// @Summary Create company
// @Description Create a new company
// @Tags company
// @Accept json
// @Produce json
// @Param company body cp.CompanyCreateReq true "Company data"
// @Success 200 {object} cp.CompanyRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /company [post]
func (h *Handler) CreateCompanyHandler(c *gin.Context) {
	var req cp.CompanyCreateReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	res, err := h.srvs.Company.Create(context.Background(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println("error: ", err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// CompanyGetById handles the get a company.
// @Summary Get company
// @Description Get a company
// @Tags company
// @Accept json
// @Produce json
// @Param id path string true "Company ID"
// @Success 200 {object} cp.CompanyGetByIdRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /company/{id} [get]
func (h *Handler) GetByIdCompanyHandler(c *gin.Context) {
	id := &cp.Byid{Id: c.Param("id")}
	res, err := h.srvs.Company.GetById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get company", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// CompanyGetAll handles getting all companies.
// @Summary Get all companies
// @Description Get all companies
// @Tags company
// @Accept json
// @Produce json
// @Param name query string false "Name"
// @Param address query string false "Address"
// @Param industry query string false "Industry"
// @Param website query string false "website"
// @Param limit query integer false "Limit"
// @Param offset query integer false "Offset"
// @Success 200 {object} cp.CompanyGetAllRes
// @Failure 400 {object} string "Invalid parameters"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /company/all [get]
func (h *Handler) GetAllCompaniesHandler(c *gin.Context) {
	req := cp.CompanyGetAllReq{
		Name: c.Query("name"),
		Address: c.Query("address"),
		Industry: c.Query("industry"),
		Website: c.Query("website"),
		Filter: &cp.Pagination{},
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

	filter := cp.Pagination{
		Limit:  int64(limit),
		Offset: int64(offset),
	}

	req.Filter.Limit = filter.Limit
	req.Filter.Offset = filter.Offset
	
	res, err := h.srvs.Company.GetAll(context.Background(), &req)


	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't get categories", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// CompanyUpdate handles updating an existing company.
// @Summary Update company
// @Description Update an existing company
// @Tags company
// @Accept json
// @Produce json
// @Param id path string true "Company ID"
// @Param company body cp.CompanyCreateReq true "Updated company data"
// @Success 200 {object} cp.CompanyRes
// @Failure 400 {object} string "Invalid request payload"
// @Failure 404 {object} string "Company not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /company/{id} [put]
func (h *Handler) UpdateCompanyHandler(c *gin.Context) {
	id := c.Param("id")
	var req cp.CompanyUpdateReq
	var company cp.CompanyCreateReq

	if err := c.BindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.Id = id
	req.Name = company.Name
	req.Address = company.Address
	req.Industry = company.Industry
	req.Website = company.Website
	res, err := h.srvs.Company.Update(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update company", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// CompanyDelete handles deleting a company by ID.
// @Summary Delete company
// @Description Delete a company by ID
// @Tags company
// @Accept json
// @Produce json
// @Param id path string true "Company ID"
// @Success 200 {object} string "Company deleted"
// @Failure 400 {object} string "Invalid company ID"
// @Failure 404 {object} string "Company not found"
// @Failure 500 {object} string "Server error"
// @Security BearerAuth
// @Router /company/{id} [delete]
func (h *Handler) DeleteCompanyHandler(c *gin.Context) {
	id := &cp.Byid{Id: c.Param("id")}
	_, err := h.srvs.Company.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete company", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Company deleted"})
}
