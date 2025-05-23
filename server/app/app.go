package app

import (
	"context"
	"database/sql"
	"fmt"

	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	// "github.com/rs/zerolog"
	// "github.com/rs/zerolog/log"

	_ "github.com/jackc/pgx/stdlib"

	"github.com/kylep342/mendel/constants"
	"github.com/kylep342/mendel/db"
	"github.com/kylep342/mendel/handlers"
	"github.com/kylep342/mendel/models/plants"
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
func (conf *config) Configure(ctx context.Context, env constants.EnvConfig) {
	conf.sqlUrl = fmt.Sprintf("%s://%s:%s@%s:%d/%s",
		env.Database.Dialect,
		env.Database.User,
		env.Database.Password,
		env.Database.Host,
		env.Database.Port,
		env.Database.Name)

	// conf.redisPassword = os.Getenv("REDIS_PASSWORD")
	// conf.redisHost = os.Getenv("REDIS_HOST")
	// conf.redisPort = os.Getenv("REDIS_PORT")
	// conf.redisDb, _ = strconv.Atoi(os.Getenv("REDIS_DB"))
}

var conf = config{}

// App is the singleton struct with components to run mendel
// Router is a pointer to a chi router
// Logger is an http handler
// DB is a pointer to a db
// Redis is a pointer to a redis client
type App struct {
	Context context.Context
	DB      *sql.DB
	Router  chi.Router
	// Redis  *redis.Client
}

// InitializeRoutes creates all endpoints for the api
func (a *App) InitializeRoutes() {
	a.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		responses.RespondWithData(w, http.StatusOK, "ok")
	})

	internalHandler := handlers.NewInternalHandler(
		a.DB,
		constants.GetEnv(),
	)

	internalHandler.RegisterRoutes(a.Router, constants.RouteIndex)

	plantSpeciesHandler := handlers.NewCRUDHandler(
		a.DB,
		func() *plants.PlantSpecies { return &plants.PlantSpecies{} },
		func(d *sql.DB) db.CRUDTable[plants.PlantSpecies] {
			return &db.PlantSpeciesTable{DB: d}
		},
	)
	plantSpeciesHandler.RegisterRoutes(a.Router, constants.RoutePlantSpecies)

	plantCultivarHandler := handlers.NewCRUDHandler(
		a.DB,
		func() *plants.PlantCultivar { return &plants.PlantCultivar{} },
		func(d *sql.DB) db.CRUDTable[plants.PlantCultivar] {
			return &db.PlantCultivarTable{DB: d}
		},
	)
	plantCultivarHandler.RegisterRoutes(a.Router, constants.RoutePlantCultivar)

	plantHandler := handlers.NewCRUDHandler(
		a.DB,
		func() *plants.Plant { return &plants.Plant{} },
		func(d *sql.DB) db.CRUDTable[plants.Plant] {
			return &db.PlantTable{DB: d}
		},
	)
	plantHandler.RegisterRoutes(a.Router, constants.RoutePlant)
}

// Initialize creates the application as a whole
func (a *App) Initialize() {
	ctx := context.Background()
	var err error
	env := constants.LoadEnv()
	conf.Configure(ctx, *env)
	a.Context = ctx

	// Postgres setup
	a.DB, err = sql.Open("pgx", conf.sqlUrl)

	if err != nil {
		log.Fatal(err)
	}
	_, err = a.DB.Exec(constants.DBInitQuery)
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
	a.Router.Use(middleware.Recoverer)

	a.InitializeRoutes()
}

// Run starts the app to listen on the port specitied by the env variable SERVER_PORT
func (a *App) Run() {
	port := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	log.Fatal(http.ListenAndServe(port, a.Router))
}
