package main

import (
	"context"
	"modularis/config"
	"modularis/repository"
	"modularis/repository/mongoUserRepo"
	"modularis/repository/ramPostRepo"
	"modularis/server"
	"modularis/service"
	"modularis/service/basicAuthService"
	"modularis/service/basicPostService"
)

func main() {
	cfg := config.NewConfig("config.yaml")
	svc := ServiceConfig(cfg)

	srv := server.NewServer(struct {
		Host string
		Port string
	}(cfg.Server), svc)
	srv.Run()
}

func RepoConfig(cfg *config.Config) *repository.Repository {
	postRepo := ramPostRepo.NewRepository()
	userRepo, _ := mongoUserRepo.NewRepository(cfg.MongoDatabase, context.TODO())
	return repository.NewRepository(userRepo, postRepo)
}

func ServiceConfig(cfg *config.Config) *service.Service {
	repo := RepoConfig(cfg)
	postService := basicPostService.NewService(repo)
	authService := basicAuthService.NewService(repo)
	return service.NewService(authService, postService)
}
