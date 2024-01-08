package rest

import (
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"net/http"
	"net/http/httptest"
	"test_task/business/service"
	"test_task/business/storage"
	"test_task/utils/config"
	"testing"
)

type RestSuite struct {
	suite.Suite
	router *gin.Engine
	db     *sql.DB
}

func (suite *RestSuite) SetupSuite() {
	rawJSON := []byte(`{
	 "level": "debug",
	 "encoding": "json",
	 "outputPaths": ["/tmp/logs"],
	 "errorOutputPaths": [],
	 "initialFields": {"foo": "bar"},
	 "encoderConfig": {
	  "messageKey": "message",
	  "levelKey": "level",
	  "levelEncoder": "lowercase"
	 }
	}`)
	var cf zap.Config
	err := json.Unmarshal(rawJSON, &cf)
	suite.NoError(err)
	l := zap.Must(cf.Build())
	defer l.Sync()

	var cfg config.Config
	if err := godotenv.Load("../../.env"); err != nil {
		suite.Fail("can't load .env file ", err)
	}
	if err := envconfig.Process("", &cfg); err != nil {
		suite.Fail("can't read OS env", err)
	}
	db, err := sql.Open("mysql", cfg.MysqlTestDsn)
	if err != nil {
		suite.Fail("can't establish database connection", err)
	}
	err = db.Ping()
	if err != nil {
		suite.Fail("database connection credentials are not valid", err)
	}
	suite.db = db
	st := storage.NewStorage(storage.New(db), l)
	s := service.New(l, st)
	server := New(l, s, cfg)
	suite.router = server.setupRouter()
}

func (suite *RestSuite) TearDownSuite() {
	suite.db.Close()
}

func (suite *RestSuite) TestRequestAuthFailed() {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/profile", nil)
	suite.NoError(err)
	req.Header.Set("api-key", "wrongkey")
	suite.router.ServeHTTP(w, req)
	suite.Equal(403, w.Code)
	suite.JSONEq(`{"error": "access forbidden"}`, w.Body.String())
}

func (suite *RestSuite) TestGetAllProfiles() {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/profile", nil)
	suite.NoError(err)
	req.Header.Set("Api-key", "www-dfq92-sqfwf")
	suite.router.ServeHTTP(w, req)
	suite.Equal(200, w.Code)
}

func (suite *RestSuite) TestGetUsernameProfile() {
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/profile", nil)
	suite.NoError(err)
	req.Header.Set("Api-key", "www-dfq92-sqfwf")
	q := req.URL.Query()
	q.Add("username", "guest")
	req.URL.RawQuery = q.Encode()
	suite.router.ServeHTTP(w, req)
	suite.Equal(200, w.Code)
	suite.JSONEq(`{"Id":3,"Username":"guest","FirstName":"Василий","LastName":"Шпак","City":"Житомир","School":"Медична гімназія №33 міста Києва"}`, w.Body.String())
}

func TestRestSuite(t *testing.T) {
	suite.Run(t, new(RestSuite))
}
