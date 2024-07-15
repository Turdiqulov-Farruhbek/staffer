package services

import (
	"context"
	"errors"

	gen "recruitment/genprotos"
	"recruitment/storage"
)

// JobAppService provides methods to interact with JobApplication and JobStep entities.
type JobAppService struct {
	store storage.StorageInterface
	gen.UnimplementedRecruitmentServiceServer
}

// NewJobAppService creates a new instance of JobAppService.
func NewJobAppService(store storage.StorageInterface) *JobAppService {
	return &JobAppService{store: store}
}

// CreateJobApplication inserts a new job application into the database.
func (s *JobAppService) CreateJobApplication(ctx context.Context, req *gen.CreateJobApplicationRequest) (*gen.JobApplicationResponse, error) {
	response, err := s.store.JobApp().CreateJobApplication(ctx, req)
	if err != nil {
		return nil, errors.New("failed to create job application: " + err.Error())
	}
	return response, nil
}

// CreateJobStep inserts a new job step into the database.
func (s *JobAppService) CreateJobStep(ctx context.Context, req *gen.CreateJobStepRequest) (*gen.JobStepResponse, error) {
	response, err := s.store.JobApp().CreateJobStep(ctx, req)
	if err != nil {
		return nil, errors.New("failed to create job step: " + err.Error())
	}
	return response, nil
}

// UpdateJobStep updates an existing job step in the database.
func (s *JobAppService) UpdateJobStep(ctx context.Context, req *gen.UpdateJobStepRequest) (*gen.JobStepResponse, error) {
	response, err := s.store.JobApp().UpdateJobStep(ctx, req)
	if err != nil {
		return nil, errors.New("failed to update job step: " + err.Error())
	}
	return response, nil
}

// UpdateJobApplication updates an existing job application in the database.
func (s *JobAppService) UpdateJobApplication(ctx context.Context, req *gen.UpdateJobApplicationRequest) (*gen.JobApplicationResponse, error) {
	response, err := s.store.JobApp().UpdateJobApplication(ctx, req)
	if err != nil {
		return nil, errors.New("failed to update job application: " + err.Error())
	}
	return response, nil
}

// GetJobApplicationById retrieves a job application by ID.
func (s *JobAppService) GetJobApplicationById(ctx context.Context, req *gen.GetJobApplicationByIdRequest) (*gen.JobApplicationResponse, error) {
	response, err := s.store.JobApp().GetJobApplicationById(ctx, req)
	if err != nil {
		return nil, errors.New("failed to get job application by ID: " + err.Error())
	}
	return response, nil
}

// GetJobStepById retrieves a job step by ID.
func (s *JobAppService) GetJobStepById(ctx context.Context, req *gen.GetJobStepByIdRequest) (*gen.JobStepResponse, error) {
	response, err := s.store.JobApp().GetJobStepById(ctx, req)
	if err != nil {
		return nil, errors.New("failed to get job step by ID: " + err.Error())
	}
	return response, nil
}

// GetAllJobApplications retrieves all job applications with optional filters.
func (s *JobAppService) GetAllJobApplications(ctx context.Context, req *gen.GetJobApplicationsRequest) (*gen.JobApplicationsResponse, error) {
	response, err := s.store.JobApp().GetAllJobApplications(ctx, req)
	if err != nil {
		return nil, errors.New("failed to get all job applications: " + err.Error())
	}
	return response, nil
}

// DeleteJobApplication deletes a job application by ID.
func (s *JobAppService) DeleteJobApplication(ctx context.Context, req *gen.DeleteJobApplicationRequest) (*gen.JobApplicationResponse, error) {
	response, err := s.store.JobApp().DeleteJobApplication(ctx, req)
	if err != nil {
		return nil, errors.New("failed to delete job application: " + err.Error())
	}
	return response, nil
}

// DeleteJobStep deletes a job step by ID.
func (s *JobAppService) DeleteJobStep(ctx context.Context, req *gen.DeleteJobStepRequest) (*gen.JobStepResponse, error) {
	response, err := s.store.JobApp().DeleteJobStep(ctx, req)
	if err != nil {
		return nil, errors.New("failed to delete job step: " + err.Error())
	}
	return response, nil
}
