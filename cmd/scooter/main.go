package main

import (
	"database/sql"
	"math"
	server "net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/shonjord/e-scooter/cmd/config"
	"github.com/shonjord/e-scooter/internal/pkg/application/event"
	"github.com/shonjord/e-scooter/internal/pkg/application/handler"
	"github.com/shonjord/e-scooter/internal/pkg/infrastructure/location"
	"github.com/shonjord/e-scooter/internal/pkg/infrastructure/mysql"
	"github.com/shonjord/e-scooter/internal/pkg/presentation/http"
	"github.com/shonjord/e-scooter/internal/pkg/presentation/http/action"
	"github.com/shonjord/e-scooter/internal/pkg/presentation/http/middleware"
	log "github.com/sirupsen/logrus"
)

var (
	spec config.Specification
)

const (
	retries = 4
)

func init() {
	if err := config.LoadEnvironmentVariables(&spec); err != nil {
		log.WithError(err).Fatal("error while loading environment variables.")
	}

	level, err := log.ParseLevel(spec.Log.Level)
	if err != nil {
		log.WithError(err).Fatal("log level could not be parsed.")
	}

	log.SetLevel(level)
}

func main() {
	var (
		mysqlConnection *sql.DB
		err             error
	)

	// infrastructure layer
	retry(func() error {
		mysqlConnection, err = sql.Open(spec.Database.Driver, spec.Database.DSN)
		if err != nil {
			return err
		}

		if err = mysqlConnection.Ping(); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.WithError(err).Fatal("connection to DB was not possible.")
	}

	dbConnection := mysql.NewConnection(mysqlConnection)
	geographicLocator := location.NewGeographicLocator()
	scooterRepository := mysql.NewScooterRepository(dbConnection)
	mobileRepository := mysql.NewMobileRepository(dbConnection)
	tripRepository := mysql.NewTripRepository(dbConnection)

	// application layer
	tripStartedEventHandler := event.NewTripStartedEventHandler(
		tripRepository,
		scooterRepository,
		geographicLocator,
	)
	domainEventDispatcher, err := event.NewDomainEventDispatcher(
		tripStartedEventHandler,
	)
	if err != nil {
		log.WithError(err).Fatal("error while registering event handlers")
	}
	findAvailableScooters := handler.NewFindAvailableScooters(scooterRepository)
	scooterDisconnector := handler.NewDisconnectScooter(tripRepository, scooterRepository)
	scooterConnector := handler.NewConnectScooter(
		scooterRepository,
		mobileRepository,
		tripRepository,
		domainEventDispatcher,
	)

	// presentation layer
	mobileXApiKeyAuthenticator := middleware.NewMobileXApiKeyAuthenticator(spec.XApiKeys.Mobile)
	scooterXApiKeyAuthenticator := middleware.NewScooterXApiKeyAuthenticator(spec.XApiKeys.Scooter)
	getAvailableScooters := action.NewGetAvailableScootersAction(findAvailableScooters)
	postScooterConnection := action.NewConnectScooterAction(scooterConnector)
	postScooterDisConnection := action.NewDisconnectScooterAction(scooterDisconnector)

	// REST endpoints
	router := chi.NewRouter()

	// mobile clients endpoints
	router.Group(func(r chi.Router) {
		r.Use(http.Bridge(mobileXApiKeyAuthenticator))

		r.Get("/clients/scooters/available", http.Handler(getAvailableScooters))
	})

	// scooter clients endpoints
	router.Group(func(r chi.Router) {
		r.Use(http.Bridge(scooterXApiKeyAuthenticator))

		r.Route("/scooters", func(r chi.Router) {
			r.Post("/connect", http.Handler(postScooterConnection))
			r.Post("/disconnect", http.Handler(postScooterDisConnection))
		})
	})

	if err = server.ListenAndServe(spec.Server.Port, router); err != nil {
		log.WithError(err).Fatal("error initializing server.")
	}
}

// retry receives a callback and executes it, if an errors is encountered, it retries after the $retries attempts.
func retry(f func() error) {
	count := 0

	for {
		if err := f(); err == nil {
			break
		}

		count++

		if retries == count {
			break
		}

		value := time.Duration(math.Pow(3, float64(count)))

		time.Sleep(value * time.Second)
	}
}
