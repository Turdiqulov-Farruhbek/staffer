package services

import (
	"context"

	gen "recruitment/genprotos"
	"recruitment/storage"
)

// JobService provides methods to interact with Job Application entities.
type JobService struct {
	store storage.StorageInterface
	gen.UnimplementedJobServiceServer
}

// NewJobService creates a new instance of JobService.
func NewJobService(store storage.StorageInterface) *JobService {
	return &JobService{store: store}
}

// Create inserts a new job application into the database.
func (s *JobService) Create(ctx context.Context, req *gen.JobApplicationCreate) (*gen.Void, error) {
	return s.store.Job().Create(ctx, req)
}

// GetByID retrieves a job application by ID.
func (s *JobService) Get(ctx context.Context, req *gen.Byid) (*gen.JobApplitcationGetRes, error) {
	return s.store.Job().GetByID(ctx, req)
}

// Update modifies an existing job application in the database.
func (s *JobService) Update(ctx context.Context, req *gen.JobApplicationUpdate) (*gen.Void, error) {
	return s.store.Job().Update(ctx, req)
}

// Delete marks a job application as deleted by setting the deleted_at timestamp.
func (s *JobService) Delete(ctx context.Context, req *gen.Byid) (*gen.Void, error) {
	return s.store.Job().Delete(ctx, req)
}

// GetAll retrieves all job applications from the database with optional filtering.
func (s *JobService) GetAll(ctx context.Context, req *gen.JobApplicationGetAll) (*gen.JobApplicationGetAllRes, error) {
	return s.store.Job().GetAll(ctx, req)
}
