package controllers_test

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"github.com/DmytroKha/underwater/config"
	"github.com/DmytroKha/underwater/internal/app"
	"github.com/DmytroKha/underwater/internal/infra/database"
	httpInt "github.com/DmytroKha/underwater/internal/infra/http"
	"github.com/DmytroKha/underwater/internal/infra/http/controllers"
	"github.com/stretchr/testify/require"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type requestTest struct {
	name           string
	url            string
	method         string
	bodyData       string
	expectedCode   int
	responseRegexg string
	msg            string // Test error message
}

func TestMain(m *testing.M) {
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestControllers(t *testing.T) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var conf = &config.Configuration{
		DatabaseName:      "halo",
		DatabaseHost:      "localhost:5433",
		DatabaseUser:      "db_user",
		DatabasePassword:  "db_password",
		MigrateToVersion:  "latest",
		MigrationLocation: "../../database/migrations",
		RedisDatabase:     "0",
		RedisHost:         "localhost:6364",
	}

	config.CreateRedisClient(ctx, *conf)

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		conf.DatabaseUser,
		conf.DatabasePassword,
		conf.DatabaseHost,
		conf.DatabaseName,
	)

	internalSQLDriver, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("Unable to open DB connection: %q\n", err)
	}

	sess, err := postgresql.New(internalSQLDriver)
	if err != nil {
		log.Fatalf("Unable to create new DB session: %q\n", err)
	}

	sensorRepository := database.NewSensorRepository(sess)
	readingRepository := database.NewReadingRepository(sess)
	fishSpeciesRepository := database.NewFishSpeciesRepository(sess)
	groupRepository := database.NewGroupRepository(sess)

	var readingService app.ReadingService
	sensorService := app.NewSensorService(sensorRepository, &readingService)
	fishSpeciesService := app.NewFishSpeciesService(fishSpeciesRepository)
	readingService = app.NewReadingService(readingRepository, sensorService, fishSpeciesService)
	groupService := app.NewGroupService(groupRepository, sensorService, readingService, fishSpeciesService)

	sensorController := controllers.NewSensorController(sensorService)
	groupController := controllers.NewGroupController(ctx, groupService)

	router := httpInt.Router(
		sensorController,
		groupController,
	)

	iterateOverTests(t, sensorControllerTests, router)
	iterateOverTests(t, groupControllerTests, router)
}

func iterateOverTests(t *testing.T, tests []*requestTest, router http.Handler) {
	for _, tt := range tests {
		t.Run("Controller "+tt.name, func(t *testing.T) {
			fmt.Printf("[%-6s] %-35s", tt.method, tt.url)
			bodyData := tt.bodyData
			req, err := http.NewRequest(tt.method, tt.url, bytes.NewBufferString(bodyData))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			body := w.Body.String()

			fmt.Printf("[%d]\n", w.Code)
			require.Equal(t, tt.expectedCode, w.Code, "Response Status - "+tt.msg+"\nBody:\n"+body)
			require.Regexp(t, tt.responseRegexg, body, "Response Content - "+tt.msg)
		})
	}
}
