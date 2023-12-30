package main

import (
	machines_management_validaor_backend "github.com/bsydorenko7/machines-management-validator-backend"
	"github.com/bsydorenko7/machines-management-validator-backend/pkg/handler"
	"github.com/bsydorenko7/machines-management-validator-backend/pkg/service"
	"log"
)

func main() {
	services := service.NewService()
	handlers := handler.NewHandler(services)

	srv := new(machines_management_validaor_backend.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
