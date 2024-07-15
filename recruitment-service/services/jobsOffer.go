package services

import (
	"context"

	gen "recruitment/genprotos"
	"recruitment/storage"
)

// JobOfferService provides methods to interact with JobOffer entities.
type JobOfferService struct {
	store storage.StorageInterface
	gen.UnimplementedJobOfferServiceServer
}

// NewJobOfferService creates a new instance of JobOfferService.
func NewJobOfferService(store storage.StorageInterface) *JobOfferService {
	return &JobOfferService{store: store}
}

// CreateJobOffer inserts a new job offer into the database.
func (s *JobOfferService) CreateJobOffer(ctx context.Context, req *gen.CreateJobOfferRequest) (*gen.JobOfferResponse, error) {
	return s.store.JobOffer().CreateJobOffer(ctx, req)
}

// GetByIdJobOffer retrieves a job offer by ID.
func (s *JobOfferService) GetByIdJobOffer(ctx context.Context, req *gen.GetByIdJobOfferRequest) (*gen.JobOfferResponse, error) {
	return s.store.JobOffer().GetByIdJobOffer(ctx, req)
}

// GetAllJobOffers retrieves all job offers from the database.
func (s *JobOfferService) GetAllJobOffers(ctx context.Context, req *gen.GetAllJobOffersRequest) (*gen.JobOffersResponse, error) {
	return s.store.JobOffer().GetAllJobOffers(ctx, req)
}

// UpdateJobOffer updates an existing job offer in the database.
func (s *JobOfferService) UpdateJobOffer(ctx context.Context, req *gen.UpdateJobOfferRequest) (*gen.JobOfferResponse, error) {
	return s.store.JobOffer().UpdateJobOffer(ctx, req)
}

// DeleteJobOffer deletes a job offer by ID.
func (s *JobOfferService) DeleteJobOffer(ctx context.Context, req *gen.DeleteJobOfferRequest) (*gen.JobOfferResponse, error) {
	return s.store.JobOffer().DeleteJobOffer(ctx, req)
}
