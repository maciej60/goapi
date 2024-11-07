# simple golang api - By maciej the magician

- import the sql file goapi.sql
- update the .env file with your params
- run > go get github.com/joho/godotenv
- run > go run cmd/api/main.go, to start the api service on port 8000 (change the port)
- middleware/authorization.go uses: 1) apikey and 2) username/password or username/token auth policy
    the password is applied on the header with key Authorization 
    the apikey is also added to the header