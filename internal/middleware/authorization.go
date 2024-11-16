package middleware

import (
    "fmt"
    "net/http"
    "os"
    "github.com/joho/godotenv"
    "github.com/maciej60/goapi/api"
    "github.com/maciej60/goapi/internal/tools"
    log "github.com/sirupsen/logrus"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
}
var ErrUnAuthorizedUser error = fmt.Errorf("invalid user or token")
var ErrUnAuthorizedApiKey error = fmt.Errorf("invalid Api key")

func Authorization(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // var appId string = r.URL.Query().Get("appId")
        var token = r.Header.Get("Authorization")
        var apikey = r.Header.Get("apikey")
        var err error
        if token == "" {
            api.RequestErrorHandler(w, ErrUnAuthorizedUser)
            return
        }
        if apikey == "" || apikey != os.Getenv("APIKEY") {
            api.RequestErrorHandler(w, ErrUnAuthorizedApiKey)
            return
        }
        var database *tools.DatabaseInterface
        database, err = tools.NewDatabase()
        if err != nil {
            api.InternalErrorHandler(w)
            return
        }
        authDetails := (*database).GetAuth(token)
        if (authDetails == nil) {
            log.Error(ErrUnAuthorizedUser)
            api.RequestErrorHandler(w, ErrUnAuthorizedUser)
            return
        }
        next.ServeHTTP(w, r)
    })
}