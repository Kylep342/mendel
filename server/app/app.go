package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "github.com/jackc/pgx/stdlib"

	"github.com/kylep342/mendel/db"
	"github.com/kylep342/mendel/handlers"
	"github.com/kylep342/mendel/models"
	"github.com/kylep342/mendel/responses"
)

// global config struct holding database connection info
type config struct {
	sqlUrl string
	// redisPassword string
	// redisHost     string
	// redisPort     string
	// redisDb       int
}

// method to initialize config struct from environment variables
func (conf *config) Configure() {
	conf.sqlUrl = os.Getenv("DATABASE_URL")
	// conf.redisPassword = os.Getenv("REDIS_PASSWORD")
	// conf.redisHost = os.Getenv("REDIS_HOST")
	// conf.redisPort = os.Getenv("REDIS_PORT")
	// conf.redisDb, _ = strconv.Atoi(os.Getenv("REDIS_DB"))
}

var conf = config{}

// App contains necessary components to run the webserver
// Router is a pointer to a chi router
// Logger is an http handler
// DB is a pointer to a db
// Redis is a pointer to a redis client
type App struct {
	Router chi.Router
	DB     *sql.DB
	// Redis  *redis.Client
}

// InitializeRoutes creates all endpoints for the api
func (a *App) InitializeRoutes() {
	a.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		responses.RespondWithData(w, http.StatusOK, "mendel is alive!")
	})
	a.Router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		responses.RespondWithData(w, http.StatusOK, "ok")
	})

	psRepo := db.NewPlantSpeciesTable(a.DB)

	psHandler := handlers.NewPlantSpeciesHandler(psRepo)

	a.Router.Route("/plant-species", func(r chi.Router) {
		r.Post("/", psHandler.Create)
		// r.Get("/", psHandler.GetAll)
		r.Get("/{id}", psHandler.GetByID)
		r.Put("/{id}", psHandler.Update)
		r.Delete("/{id}", psHandler.Delete)
	})

	plantHandler := &handlers.CRUDHandler[models.Plant]{
		Table: &db.PlantTable{DB: a.DB},
		New:   func() *models.Plant { return &models.Plant{} },
	}
	plantHandler.RegisterRoutes(a.Router, "/plant")
}

// Initialize creates the application as a whole
func (a *App) Initialize() {
	var err error
	conf.Configure()
	a.DB, err = sql.Open("pgx", conf.sqlUrl)

	if err != nil {
		log.Fatal(err)
	}
	_, err = a.DB.Exec("SET search_path TO mendel_core")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to set search_path: %w", err))
	}

	// a.Redis = redis.NewClient(&redis.Options{
	// 	Addr:     fmt.Sprintf("%s:%s", conf.redisHost, conf.redisPort),
	// 	Password: conf.redisPassword,
	// 	DB:       conf.redisDb,
	// })

	a.Router = chi.NewRouter()
	a.Router.Use(middleware.Logger)

	a.InitializeRoutes()
}

// Run starts the app to listen on the port specitied by the env variable SERVER_PORT
func (a *App) Run() {
	port := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	log.Fatal(http.ListenAndServe(port, a.Router))
}
