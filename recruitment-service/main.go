package main



import (
	"log"
	"net"
	cfg "recruitment/config"
	er "recruitment/genprotos"
	src "recruitment/services"
	post "recruitment/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	db, err := post.DBCon()
	if err != nil {
		log.Fatal(err)
	}

	// config := cfg.Load()
	liss, err := net.Listen("tcp", cfg.Load().HTTPPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	er.RegisterCompanyServiceServer(server, src.NewCompanyService(db))
	er.RegisterJobServiceServer(server, src.NewJobService(db))
	er.RegisterVacancyServiceServer(server, src.NewVacancyService(db))

	log.Println("Server is running on port :5050")
	if err := server.Serve(liss); err != nil {
		log.Fatalf("error listening: %v", err)
	}

}



