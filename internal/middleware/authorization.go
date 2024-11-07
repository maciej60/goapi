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
var ErrUnAuthorizedUser = fmt.Errorf("invalid username or token/password")
var ErrUnAuthorizedApiKey error = fmt.Errorf("invalid Api key or secret")

func Authorization(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        var username string = r.URL.Query().Get("username")
        var password = r.Header.Get("Authorization")
        var apikey = r.Header.Get("apikey")
        var err error

        if username == "" {
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

        loginDetails := (*database).GetUserLoginDetails(username)

        if (loginDetails == nil || (password != (*loginDetails).Password)) {
            log.Error(ErrUnAuthorizedUser)
            api.RequestErrorHandler(w, ErrUnAuthorizedUser)
            return
        }

        next.ServeHTTP(w, r)

    })
}