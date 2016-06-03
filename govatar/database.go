package govatar

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// We need this to use the mysql dialect
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
)

// User Our user model
type User struct {
	gorm.Model
	Email      string `gorm:"not null;unique;"`
	Password   string `gorm:"not null;size:64"`
	Hash       string `gorm:"not null;index:idx_user_hash;size:64"`
	EmailToken string `gorm:"size:64"`
	Confirmed  bool
	Avatar     []byte `gorm:"type:longblob"`
}

// InitDb Initialise the database
func InitDb() {
	db := connDb()

	// Migrate the schema
	db.AutoMigrate(&User{})
}

func connDb() *gorm.DB {
	var connectionString string

	if Conf.Database.Type == "unix" {
		connectionString = fmt.Sprintf("%s:%s@unix(%s)/%s?charset=utf8&parseTime=True&loc=Local", Conf.Database.Username, Conf.Database.Password, Conf.Database.Server, Conf.Database.Database)
	} else {
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", Conf.Database.Username, Conf.Database.Password, Conf.Database.Server, Conf.Database.Port, Conf.Database.Database)
	}

	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic("failed to connect to the database")
	}

	return db
}

// CreateUser Create a new user
func CreateUser(email string, password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), Conf.App.BcryptCost)
	if err != nil {
		panic(err)
	}

	db := connDb()
	db.Create(&User{Email: email, Password: string(hashedPassword), Hash: SHA256Hash(email)})
}

// GetUserByHash Get a user by their email hash
func GetUserByHash(hash string) User {
	db := connDb()

	var user User
	db.Where("hash = ?", hash).First(&user)

	return user
}
