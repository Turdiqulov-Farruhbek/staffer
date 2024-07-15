package postgres

import (
	"context"

	gen "recruitment/genprotos"

	"github.com/jackc/pgx/v5"
)

type JobOffer struct {
	db *pgx.Conn
}

func NewJobOffer(db *pgx.Conn) *JobOffer {
	return &JobOffer{db: db}
}

// CreateJobOffer inserts a new job offer into the database.
func (s *JobOffer) CreateJobOffer(ctx context.Context, req *gen.CreateJobOfferRequest) (*gen.JobOfferResponse, error) {
	query := `INSERT INTO joboffers (title, description, company, location, status) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	row := s.db.QueryRow(ctx, query, req.JobOffer.Title, req.JobOffer.Description, req.JobOffer.Company, req.JobOffer.Location, req.JobOffer.Status)

	var id string
	err := row.Scan(&id)
	if err != nil {
		return nil, err	
	}

	req.JobOffer.Id = id
	return &gen.JobOfferResponse{JobOffer: req.JobOffer}, nil
}

// GetByIdJobOffer retrieves a job offer by ID.
func (s *JobOffer) GetByIdJobOffer(ctx context.Context, req *gen.GetByIdJobOfferRequest) (*gen.JobOfferResponse, error) {
	query := `SELECT id, title, description, company, location, status FROM joboffers WHERE id = $1`
	row := s.db.QueryRow(ctx, query, req.Id)

	var jobOffer gen.JobOffer
	err := row.Scan(&jobOffer.Id, &jobOffer.Title, &jobOffer.Description, &jobOffer.Company, &jobOffer.Location, &jobOffer.Status)
	if err != nil {
		return nil, err
	}

	return &gen.JobOfferResponse{JobOffer: &jobOffer}, nil
}

// GetAllJobOffers retrieves all job offers from the database.
func (s *JobOffer) GetAllJobOffers(ctx context.Context, req *gen.GetAllJobOffersRequest) (*gen.JobOffersResponse, error) {
	query := `SELECT id, title, description, company, location, status FROM joboffers`
	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobOffers []*gen.JobOffer
	for rows.Next() {
		var jobOffer gen.JobOffer
		err := rows.Scan(&jobOffer.Id, &jobOffer.Title, &jobOffer.Description, &jobOffer.Company, &jobOffer.Location, &jobOffer.Status)
		if err != nil {
			return nil, err
		}
		jobOffers = append(jobOffers, &jobOffer)
	}

	return &gen.JobOffersResponse{JobOffers: jobOffers}, nil
}

// UpdateJobOffer updates an existing job offer in the database.
func (s *JobOffer) UpdateJobOffer(ctx context.Context, req *gen.UpdateJobOfferRequest) (*gen.JobOfferResponse, error) {
	query := `UPDATE joboffers SET title = $1, description = $2, company = $3, location = $4, status = $5 WHERE id = $6`
	_, err := s.db.Exec(ctx, query, req.JobOffer.Title, req.JobOffer.Description, req.JobOffer.Company, req.JobOffer.Location, req.JobOffer.Status, req.Id)
	if err != nil {
		return nil, err
	}

	return s.GetByIdJobOffer(ctx, &gen.GetByIdJobOfferRequest{Id: req.Id})
}

// DeleteJobOffer deletes a job offer by ID.
func (s *JobOffer) DeleteJobOffer(ctx context.Context, req *gen.DeleteJobOfferRequest) (*gen.JobOfferResponse, error) {
	query := `DELETE FROM joboffers WHERE id = $1`
	_, err := s.db.Exec(ctx, query, req.Id)
	if err != nil {
		return nil, err
	}

	return &gen.JobOfferResponse{}, nil
}
