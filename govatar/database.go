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

var connectionString string
var db *gorm.DB

// InitDb Initialise the database
func InitDb() {
	connectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", Conf.Database.Username, Conf.Database.Password, Conf.Database.Database)

	var err error
	db, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic("failed to connect to the database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})
}

// CreateUser Create a new user
func CreateUser(email string, password string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), Conf.App.BcryptCost)
	if err != nil {
		panic(err)
	}

	db.Create(&User{Email: email, Password: string(hashedPassword), Hash: SHA256Hash(email)})
}
