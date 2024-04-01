package tests

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const testHost = "http://localhost"

func SetupMinimumInstance(path string) error {

	_ = path
	cnf := viper.New()
	cnf.Set("mode", "test")
	gin.SetMode("debug")

	//Connect to Redis
	// rClient := redis.NewClient(&redis.Options{
	// 	Addr: "localhost:6379",
	// })

	// kv.Init(rClient)

	// Connect to MongoDB
	//mongoURI := "mongodb://localhost:27017"
	// client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	// if err != nil {
	// 	panic(err)
	// }
	// if err := client.Ping(context.Background(), nil); err != nil {
	// 	panic(err)
	// }

	// userRepo := m.NewUserRepoMongosh(client)

	// kv.Init(userRepo)

	// fmt.Println("Successfully connected to the database!")

	//Connect to Postgres
	db, err := sqlx.Open("postgres", "user=postgres dbname=userdb password=1234 host=localhost port=5544 sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}
	// userRepo := p.NewUserRepo(db)
	// kv.Init(userRepo)

	fmt.Println("Successfully connected to the database!")

	return nil
}

func Server(handler gin.HandlerFunc, uri string, req *http.Request, router *gin.Engine) (*httptest.ResponseRecorder, error) {
	resp := httptest.NewRecorder()

	switch req.Method {
	case http.MethodPost:
		router.POST(uri, handler)
	case http.MethodGet:
		router.GET(uri, handler)
	case http.MethodDelete:
		router.DELETE(uri, handler)
	case http.MethodPut:
		router.PUT(uri, handler)
	case http.MethodPatch:
		router.PATCH(uri, handler)

	}
	router.ServeHTTP(resp, req)

	return resp, nil
}

func NewResponse() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
}

func NewRequest(method string, uri string, body []byte) *http.Request {
	req, _ := http.NewRequest(method, testHost+uri, nil)
	req.Header.Set("Host", "localhost")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	req.Header.Set("X-Forwarded-For", "79.104.42.249")
	if body != nil {
		req.Body = io.NopCloser(bytes.NewBuffer(body))
	}

	return req
}

func OpenFile(fileName string) ([]byte, error) {
	return os.ReadFile(fileName)
}
