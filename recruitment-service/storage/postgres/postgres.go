package postgres

import (
	"context"
	"fmt"
	"log/slog"
	"recruitment/config"
	stg "recruitment/storage"

	"github.com/jackc/pgx/v5"
)

type StorageStruct struct {
	Db        *pgx.Conn
	Company_S stg.CompanyInterface
	Job_S     stg.JobInterface
	Vacancy_S stg.VacancyInterface
	JobOffer_S stg.JobOfferInterface
	JobApp_S  stg.JobAppInterface
}

func DBCon() (*StorageStruct, error) {
	var (
		db  *pgx.Conn
		err error
	)
	cfg := config.Load()
	dbcon := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase)

	db, err = pgx.Connect(context.Background(), dbcon)
	if err != nil {
		slog.Error("Unable to connect to database:", "erroRRR:", err)
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		slog.Warn("unable to ping database:", "erroRRR:", err)
		return nil, err
	}

	return &StorageStruct{
		Db: db,
	}, nil
}


func (s *StorageStruct) Company() stg.CompanyInterface {
    if s.Company_S == nil {
		s.Company_S = NewCompany(s.Db)
	}
	return s.Company_S
}

func (s *StorageStruct) Job() stg.JobInterface {
	if s.Job_S == nil {
        s.Job_S = NewJob(s.Db)
    }
    return s.Job_S
}

func (s *StorageStruct) Vacancy() stg.VacancyInterface {
	if s.Vacancy_S == nil {
        s.Vacancy_S = NewVacancy(s.Db)
    }
    return s.Vacancy_S
}

func (s *StorageStruct) JobOffer() stg.JobOfferInterface {
	if s.JobOffer_S == nil {
        s.JobOffer_S = NewJobOffer(s.Db)
    }
    return s.JobOffer_S
}

func (s *StorageStruct) JobApp() stg.JobAppInterface {
	if s.JobApp_S == nil {
        s.JobApp_S = NewJobApp(s.Db)
    }
    return s.JobApp_S
}