package handlers

import (
	"encoding/json"
	"fmt"
	"os"
    "github.com/joho/godotenv"
	"net/http"
	"github.com/maciej60/goapi/api"
	"github.com/maciej60/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
}

type LoginParams struct {
	Password string
	Username  string
}

type LoginResponse struct {
	Code int // Success Code, Usually 200
	Data *tools.LoginDetails
}

var ErrUnAuthorizedApiKey error = fmt.Errorf("invalid Api key")
var ErrUnAuthorizedUser error = fmt.Errorf("invalid username or password")
var ErrInactiveUser error = fmt.Errorf("user is inactive")

func Login(w http.ResponseWriter, r *http.Request) {
	var apikey = r.Header.Get("apikey")
	var err error
	if apikey == "" || apikey != os.Getenv("APIKEY") {
		api.RequestErrorHandler(w, ErrUnAuthorizedApiKey)
		return
	}
	var params = LoginParams{}
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}
	var lg *tools.LoginDetails = (*database).GetUserLoginDetails(params.Username)
	if lg == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	if (*lg).Password != params.Password {
		api.RequestErrorHandler(w, ErrUnAuthorizedUser)
		return
	}
	if (*lg).Status != 1 {
		api.RequestErrorHandler(w, ErrInactiveUser)
		return
	}
	(*lg).Password = "******"
	var response = LoginResponse{
		Code: http.StatusOK,
		Data: lg,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}else{
		log.Info("Successfully logged in user: ", params.Username)
		fmt.Println(response)
	}
}

