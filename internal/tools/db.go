package tools

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type dB struct{}
var db *sql.DB

func init() {

	if err := godotenv.Load(); err != nil {
        log.Fatal(err)
    }

    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    os.Getenv("DBNET"),
        Addr:   os.Getenv("DBADDR"),
        DBName: os.Getenv("DBNAME"),
		AllowNativePasswords: true,
    }

    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Database Connected!")
}

func (d *dB) GetAuth(token string) *Auth {
    type Username struct {
        Username string `json:"username"`
    }
    var data = []byte(token)
    var uname Username
    var clientData Auth
    json.Unmarshal(data, &uname)
    query := fmt.Sprintf("SELECT id,name,password FROM users WHERE username = '%s'", uname.Username)
    row := db.QueryRow(query)
    err := row.Scan(&clientData.UserId, &clientData.Name, &clientData.Password)
    if err == sql.ErrNoRows {
        return nil
    } else if err!= nil {
        log.Fatal(err)
    }
    clientData.Username = uname.Username
    return &clientData
}

func (d *dB) GetUserLoginDetails(username string) *LoginDetails {
    time.Sleep(time.Second * 1)
    var clientData LoginDetails
    query := fmt.Sprintf("SELECT id,name,password,status,created_at FROM users WHERE username = '%s'", username)
    row := db.QueryRow(query)
    err := row.Scan(&clientData.UserId, &clientData.Name, &clientData.Password, &clientData.Status, &clientData.CreatedAt)
    if err == sql.ErrNoRows {
        return nil
    } else if err!= nil {
        log.Fatal(err)
    }
    clientData.Username = username
    return &clientData
}

func (d *dB) GetUserCoins(username string) *CoinDetails {
    time.Sleep(time.Second * 1)
    var clientData CoinDetails
    query := fmt.Sprintf("SELECT coin_balance FROM user_coins WHERE username = '%s'", username)
    row := db.QueryRow(query)
    err := row.Scan(&clientData.Coins)
    if err == sql.ErrNoRows {
        return nil
    } else if err!= nil {
        log.Fatal(err)
    }
    clientData.Username = username
    return &clientData
}

func (d *dB) SetupDatabase() error {
	return nil
}