package main

import (
	"context"
	"fmt"
	"job-portal-api/config"
	"job-portal-api/internal/authentication"
	"job-portal-api/internal/cache"
	"job-portal-api/internal/database"
	"job-portal-api/internal/handler"
	"job-portal-api/internal/repository"
	"job-portal-api/internal/service"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
)

func main() {
	err := StartApp()
	if err != nil {
		log.Panic().Err(err).Send()
	}
	log.Info().Msg("Heloo this is our job-portal-app")
}

func StartApp() error {

	cfg := config.GetConfig()

	// log.Info().Interface("cfg", cfg).Msg("config")

	//initializing authentication support
	log.Info().Msg("main started : initializing with the authentication support")

	//reading private key file
	// privatePemFile, err := os.ReadFile(`private.pem`)
	// if err != nil {
	// 	log.Info().Msg("Error in reading private Key file")
	// 	return fmt.Errorf("error in reading private key file : %w", err)
	// }

	//parsing private pem filr content to rsa private key
	privatePemFile := []byte(cfg.AuthConfig.PrivateKey)
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePemFile)
	if err != nil {
		log.Info().Msg("Error in reading private Key")
		return fmt.Errorf("error in parsing private key : %w", err)
	}

	//reading public key file
	// publicPemFile, err := os.ReadFile(`pubkey.pem`)
	// if err != nil {
	// 	log.Info().Msg("Error in reading public Key filer")
	// 	return fmt.Errorf("error in reading public key file : %w", err)
	// }

	//parsing public pem file content to rsa public key
	publicPemFile := []byte(cfg.AuthConfig.PublicKey)
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicPemFile)
	if err != nil {
		log.Info().Msg("error in parsing public key")
		return fmt.Errorf("error in paring public key")
	}

	//callig factory function to get authentication
	auth, err := authentication.NewAuth(privateKey, publicKey)
	if err != nil {
		log.Info().Msg("error in auth function ")
		return fmt.Errorf("error in auth function : %w", err)
	}

	//starting with dtatbase connection
	log.Info().Msg("main started : initializing the database")

	db, err := database.DatabaseConnection()
	if err != nil {
		log.Info().Msg("error while opening data base connection")
		return fmt.Errorf("error while opening data base connection : %w", err)
	}

	redis := database.ConnectToRedis()

	rdb, err := cache.NewRDBLayer(redis)
	if err != nil {
		log.Info().Msg("error while initializing redis service")
		return fmt.Errorf("error while initializing redis service : %w", err)
	}

	//initializing the repo layer
	userRepo, err := repository.NewUserRepo(db)
	if err != nil {
		log.Info().Msg("error while initializing the user repository")
		return err
	}

	companyRepo, err := repository.NewCompanyRepo(db)
	if err != nil {
		log.Info().Msg("error while initializing the company repository")
		return err
	}

	jobRepo, err := repository.NewJobRepo(db)
	if err != nil {
		log.Info().Msg("error while initializing the job repository")
		return err
	}

	userService, err := service.NewUserService(userRepo, auth, rdb)
	if err != nil {
		log.Info().Msg("error while initializing user service")
		return fmt.Errorf("error while initializing uservservice : %w", err)
	}

	companyService, err := service.NewCompanyService(companyRepo)
	if err != nil {
		log.Info().Msg("error while initializing company service")
		return fmt.Errorf("error while initializing company service : %w", err)
	}

	jobService, err := service.NewJobService(jobRepo, rdb)
	if err != nil {
		log.Info().Msg("error while initializing job service")
		return fmt.Errorf("error while initializing job service : %w", err)
	}

	//initilazing http server
	api := http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.AppConfig.Port),
		ReadTimeout:  time.Duration(cfg.AppConfig.ReadTimeOut) * time.Second,
		WriteTimeout: time.Duration(cfg.AppConfig.WriteTimeOut) * time.Second,
		IdleTimeout:  time.Duration(cfg.AppConfig.IdleTimeout) * time.Second,
		Handler:      handler.SetupApi(auth, userService, companyService, jobService),
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Info().Str("port", api.Addr).Msg("main started : api is listening")
		serverErrors <- api.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt)

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error : %w", err)

	case sig := <-shutdown:
		log.Info().Msgf("main: start shutdown %s", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		err := api.Shutdown(ctx)
		if err != nil {
			err := api.Close()
			return fmt.Errorf("could not stop server gracefully : %w", err)
		}

	}

	return nil

}
