package services

import (
	"context"

	gen "recruitment/genprotos"
	"recruitment/storage"
)

// VacancyService provides methods to interact with Vacancy entities.
type VacancyService struct {
	store storage.StorageInterface
	gen.UnimplementedVacancyServiceServer
}

// NewVacancyService creates a new instance of VacancyService.
func NewVacancyService(store storage.StorageInterface) *VacancyService {
	return &VacancyService{store: store}
}

// Create inserts a new vacancy into the database.
func (s *VacancyService) Create(ctx context.Context, req *gen.VacancyCreate) (*gen.Void, error) {
	return s.store.Vacancy().Create(ctx, req)
}

// GetByID retrieves a vacancy by ID.
func (s *VacancyService) GetByID(ctx context.Context, req *gen.Byid) (*gen.VacancyGetResUpdateReq, error) {
	return s.store.Vacancy().GetByID(ctx, req)
}

// Update modifies an existing vacancy in the database.
func (s *VacancyService) Update(ctx context.Context, req *gen.VacancyGetResUpdateReq) (*gen.Void, error) {
	return s.store.Vacancy().Update(ctx, req)
}

// Delete marks a vacancy as deleted by setting the deleted_at timestamp.
func (s *VacancyService) Delete(ctx context.Context, req *gen.Byid) (*gen.Void, error) {
	return s.store.Vacancy().Delete(ctx, req)
}

// GetAll retrieves all vacancies from the database with optional filtering.
func (s *VacancyService) GetAll(ctx context.Context, req *gen.VacancyGetAllReq) (*gen.VacancyGetAllRes, error) {
	return s.store.Vacancy().GetAll(ctx, req)
}
