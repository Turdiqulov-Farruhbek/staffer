package clients

import (
	"gateway/config"
	cp "gateway/genprotos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClients struct {
	Company    cp.CompanyServiceClient
	VacancyService cp.VacancyServiceClient
	JobsService cp.JobServiceClient
	JobsAppService cp.RecruitmentServiceClient
	JobsOfferService cp.JobOfferServiceClient
}

func NewGrpcClients(cfg *config.Config) (*GrpcClients, error) {
	connC, err := grpc.NewClient("recruitment-service"+cfg.RECRUITMENT_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &GrpcClients{
		Company:    cp.NewCompanyServiceClient(connC),
		VacancyService: cp.NewVacancyServiceClient(connC),
        JobsService: cp.NewJobServiceClient(connC),
		JobsAppService: cp.NewRecruitmentServiceClient(connC),
		JobsOfferService: cp.NewJobOfferServiceClient(connC),
	}, nil
}
