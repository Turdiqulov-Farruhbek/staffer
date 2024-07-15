package services

import (
	"context"

	gen "recruitment/genprotos"
	"recruitment/storage"
)

// CompanyService provides methods to interact with Company entities.
type CompanyService struct {
	store storage.StorageInterface
	gen.UnimplementedCompanyServiceServer
}

// NewCompanyService creates a new instance of CompanyService.
func NewCompanyService(store storage.StorageInterface) *CompanyService {
	return &CompanyService{store: store}
}

// Create creates a new company entry.
func (s *CompanyService) Create(ctx context.Context, req *gen.CreateCompany) (*gen.Void, error) {
	return s.store.Company().Create(ctx, req)
}

// Get retrieves a company entry by ID.
func (s *CompanyService) GetByID(ctx context.Context, req *gen.Byid) (*gen.GetByID, error) {
	return s.store.Company().GetByID(ctx, req)
}

// Update updates an existing company entry.
func (s *CompanyService) Update(ctx context.Context, req *gen.CompanyUpdate) (*gen.Void, error) {
	return s.store.Company().Update(ctx, req)
}

// Delete deletes a company entry by ID.
func (s *CompanyService) Delete(ctx context.Context, req *gen.Byid) (*gen.Void, error) {
	return s.store.Company().Delete(ctx, req)
}

// GetAll retrieves all company entries based on optional filters.
func (s *CompanyService) GetAll(ctx context.Context, req *gen.GetAllCompanyRequest) (*gen.GetAllCompanyResponse, error) {
	return s.store.Company().GetAll(ctx, req)
}
