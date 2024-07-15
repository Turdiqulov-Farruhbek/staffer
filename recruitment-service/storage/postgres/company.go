package postgres

import (
	"context"
	"strconv"
	"time"

	gen "recruitment/genprotos"

	"github.com/jackc/pgx/v5"
)

type Company struct {
	db *pgx.Conn
}

func NewCompany(db *pgx.Conn) *Company {
	return &Company{db: db}
}

// CreateCompany inserts a new company into the database.
func (s *Company) Create(ctx context.Context, req *gen.CreateCompany) (*gen.Void, error) {
	// Prepare the SQL statement for inserting a new company
	query := `INSERT INTO companies (name, address, industry, website) VALUES ($1, $2, $3, $4)`

	// Execute the query with the parameters from the CreateCompany request
	_, err := s.db.Exec(ctx, query, req.Name, req.Address, req.Industry, req.Website)
	if err != nil {
		return nil, err
	}

	// Return an empty response (Void) if successful
	return &gen.Void{}, nil
}

// GetCompanyById retrieves a company's details based on the provided ID, excluding deleted records.
func (s *Company) GetByID(ctx context.Context, req *gen.Byid) (*gen.GetByID, error) {
	// Prepare the SQL statement for selecting a company by ID
	query := `SELECT name, address, industry, website, company_id FROM companies WHERE company_id = $1 AND deleted_at = 0`

	// Execute the query
	row := s.db.QueryRow(ctx, query, req.Id)

	// Parse the result into a GetByID response
	var company gen.GetByID
	err := row.Scan(&company.Name, &company.Address, &company.Industry, &company.Website, &company.CompanyId)
	if err != nil {
		return nil, err
	}

	// Set the success field and return the response
	company.Success = true
	return &company, nil
}

// UpdateCompany updates the details of an existing company in the database.
func (s *Company) Update(ctx context.Context, req *gen.CompanyUpdate) (*gen.Void, error) {
	// First, retrieve the existing company details
	var existing gen.CompanyUpdate
	query := `SELECT name, address, industry, website FROM companies WHERE company_id = $1 AND deleted_at = 0`
	err := s.db.QueryRow(ctx, query, req.Id).Scan(&existing.Name, &existing.Address, &existing.Industry, &existing.Website)
	if err != nil {
		return nil, err
	}

	// Use existing values if the new values are empty
	if req.Name == "" {
		req.Name = existing.Name
	}
	if req.Address == "" {
		req.Address = existing.Address
	}
	if req.Industry == "" {
		req.Industry = existing.Industry
	}
	if req.Website == "" {
		req.Website = existing.Website
	}

	// Prepare the SQL statement for updating a company
	updateQuery := `UPDATE companies SET name = $1, address = $2, industry = $3, website = $4 WHERE company_id = $5 AND deleted_at = 0`

	// Execute the update query with the parameters from the CompanyUpdate request
	_, err = s.db.Exec(ctx, updateQuery, req.Name, req.Address, req.Industry, req.Website, req.Id)
	if err != nil {
		return nil, err
	}

	// Return an empty response (Void) if successful
	return &gen.Void{}, nil
}

// DeleteCompany marks a company as deleted by setting the deleted_at timestamp.
func (s *Company) Delete(ctx context.Context, req *gen.Byid) (*gen.Void, error) {
	// Prepare the SQL statement for soft deleting a company
	query := `UPDATE companies SET deleted_at = $1 WHERE company_id = $2 AND deleted_at = 0`

	// Get the current UNIX timestamp
	deletedAt := time.Now().Unix()

	// Execute the query with the parameters
	_, err := s.db.Exec(ctx, query, deletedAt, req.Id)
	if err != nil {
		return nil, err
	}

	// Return an empty response (Void) if successful
	return &gen.Void{}, nil
}

// GetAllCompanies retrieves all companies from the database with optional filtering.
func (s *Company) GetAll(ctx context.Context, req *gen.GetAllCompanyRequest) (*gen.GetAllCompanyResponse, error) {
	// Prepare the base SQL query
	query := `SELECT name, address, industry, website, company_id FROM companies WHERE 1=1`
	args := []interface{}{}
	argIndex := 1

	// Add filters to the query if they are provided in the request
	if req.Name != "" {
		query += ` AND name = $` + strconv.Itoa(argIndex)
		args = append(args, req.Name)
		argIndex++
	}
	if req.Address != "" {
		query += ` AND address = $` + strconv.Itoa(argIndex)
		args = append(args, req.Address)
		argIndex++
	}
	if req.Industry != "" {
		query += ` AND industry = $` + strconv.Itoa(argIndex)
		args = append(args, req.Industry)
		argIndex++
	}
	if req.Website != "" {
		query += ` AND website = $` + strconv.Itoa(argIndex)
		args = append(args, req.Website)
		argIndex++
	}
	if req.CompanyId != "" {
		query += ` AND company_id = $` + strconv.Itoa(argIndex)
		args = append(args, req.CompanyId)
		argIndex++
	}

	// Execute the query
	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Parse the results into a GetAllCompanyResponse
	var companies []*gen.GetAllCompanyRequest
	for rows.Next() {
		var company gen.GetAllCompanyRequest
		err := rows.Scan(&company.Name, &company.Address, &company.Industry, &company.Website, &company.CompanyId)
		if err != nil {
			return nil, err
		}
		companies = append(companies, &company)
	}
	// Return the response with the list of companies
	// company gen.GetAllCompanyRequest{}

	return &gen.GetAllCompanyResponse{
		Companies: companies,
		Success:   true,
	}, nil
}
