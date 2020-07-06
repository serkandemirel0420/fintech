package helpers

import (
	"regexp"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/serkandemirel0420/fintech/interfaces"
	"golang.org/x/crypto/bcrypt"
)

//HandleErr handler
func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

//HashAndSalt func
func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandleErr(err)

	return string(hashed)
}

//ConnectDB connect
func ConnectDB() *gorm.DB {
	//docker run --name docker_postgres -e POSTGRES_PASSWORD=123456 -d -p 5432:5432  postgres
	//docker exec -it docker_postgres psql -U postgres
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres password=123456 sslmode=disable")
	HandleErr(err)
	return db
}

//Validation vald
func Validation(values []interfaces.Validation) bool {

	username := regexp.MustCompile(`^([A-Za-z0-9]{5,})+$`)
	email := regexp.MustCompile(`^[A-Za-z0-9]+[@]+[A-Za-z0-9]+[.]+[A-Za-z]+$`)

	for i := 0; i < len(values); i++ {
		switch values[i].Valid {
		case "username":
			if !username.MatchString(values[i].Value) {
				return false
			}
		case "email":
			if !email.MatchString(values[i].Value) {
				return false
			}
		case "password":
			if len(values[i].Value) < 5 {
				return false
			}
		}
	}
	return true
}
