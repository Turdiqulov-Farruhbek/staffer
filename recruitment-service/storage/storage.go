package storage

import (
	"context"
	gen "recruitment/genprotos"
)

type StorageInterface interface {
	Company() CompanyInterface
	Job() JobInterface
	Vacancy() VacancyInterface
	JobOffer() JobOfferInterface
	JobApp() JobAppInterface
}


type CompanyInterface interface {
	Create(ctx context.Context, req *gen.CreateCompany) (*gen.Void, error)
	GetByID(ctx context.Context, req *gen.Byid) (*gen.GetByID, error)
	Update(ctx context.Context, req *gen.CompanyUpdate) (*gen.Void, error)
	Delete(ctx context.Context, req *gen.Byid) (*gen.Void, error)
	GetAll(ctx context.Context, req *gen.GetAllCompanyRequest) (*gen.GetAllCompanyResponse, error)
}


type JobInterface interface {
	Create(ctx context.Context, req *gen.JobApplicationCreate) (*gen.Void, error)
	GetByID(ctx context.Context, req *gen.Byid) (*gen.JobApplitcationGetRes, error)
	Update(ctx context.Context, req *gen.JobApplicationUpdate) (*gen.Void, error)
	Delete(ctx context.Context, req *gen.Byid) (*gen.Void, error)
	GetAll(ctx context.Context, req *gen.JobApplicationGetAll) (*gen.JobApplicationGetAllRes, error)
}


type VacancyInterface interface {
	Create(ctx context.Context, req *gen.VacancyCreate) (*gen.Void, error)
	GetByID(ctx context.Context, req *gen.Byid) (*gen.VacancyGetResUpdateReq, error)
	Update(ctx context.Context, req *gen.VacancyGetResUpdateReq) (*gen.Void, error)
	Delete(ctx context.Context, req *gen.Byid) (*gen.Void, error)
	GetAll(ctx context.Context, req *gen.VacancyGetAllReq) (*gen.VacancyGetAllRes, error)
}


type JobOfferInterface interface {
	CreateJobOffer(ctx context.Context, req *gen.CreateJobOfferRequest) (*gen.JobOfferResponse, error)
	GetByIdJobOffer(ctx context.Context, req *gen.GetByIdJobOfferRequest) (*gen.JobOfferResponse, error)
	GetAllJobOffers(ctx context.Context, req *gen.GetAllJobOffersRequest) (*gen.JobOffersResponse, error)
	UpdateJobOffer(ctx context.Context, req *gen.UpdateJobOfferRequest) (*gen.JobOfferResponse, error)
	DeleteJobOffer(ctx context.Context, req *gen.DeleteJobOfferRequest) (*gen.JobOfferResponse, error)
}


type JobAppInterface interface {
	CreateJobApplication(ctx context.Context, req *gen.CreateJobApplicationRequest) (*gen.JobApplicationResponse, error)
	CreateJobStep(ctx context.Context, req *gen.CreateJobStepRequest) (*gen.JobStepResponse, error)
	GetJobApplicationById(ctx context.Context, req *gen.GetJobApplicationByIdRequest) (*gen.JobApplicationResponse, error)
	GetJobStepById(ctx context.Context, req *gen.GetJobStepByIdRequest) (*gen.JobStepResponse, error)
	GetAllJobApplications(ctx context.Context, req *gen.GetJobApplicationsRequest) (*gen.JobApplicationsResponse, error)
	UpdateJobApplication(ctx context.Context, req *gen.UpdateJobApplicationRequest) (*gen.JobApplicationResponse, error)
	UpdateJobStep(ctx context.Context, req *gen.UpdateJobStepRequest) (*gen.JobStepResponse, error)
	DeleteJobApplication(ctx context.Context, req *gen.DeleteJobApplicationRequest) (*gen.JobApplicationResponse, error)
	DeleteJobStep(ctx context.Context, req *gen.DeleteJobStepRequest) (*gen.JobStepResponse, error)
}