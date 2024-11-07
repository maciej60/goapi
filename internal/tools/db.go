package tools

import (
	"time"
	"database/sql"
    "fmt"
    "log"
    "os"
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

func (d *dB) GetUserLoginDetails(username string) *LoginDetails {
    time.Sleep(time.Second * 1)

    var clientData LoginDetails
    query := fmt.Sprintf("SELECT password FROM users WHERE username = '%s'", username)
    row := db.QueryRow(query)
    err := row.Scan(&clientData.Password)
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