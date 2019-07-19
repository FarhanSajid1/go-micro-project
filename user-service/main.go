package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	k8s "github.com/micro/examples/kubernetes/go/micro"

	"github.com/micro/go-micro/errors"

	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"

	user "github.com/farhansajid1/go-micro-project/user-service/proto/user"
	"github.com/micro/go-micro"
)

// for handling the server and all methods
type serv struct{}

func GenerateUUID() string {
	uuid := uuid.NewV4()
	if err != nil {
		log.Printf("could not create uuid %v", err)
	}
	return uuid.String()
}

func JWToken() string {
	mySigningKey := []byte("AllYourBase")
	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "go.micro.srv.user",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Printf("could not send the token %v", err)
	}
	return t

}

func (s *serv) Create(ctx context.Context, req *user.User, resp *user.Response) error {
	// create a new user, we are going to hash the password as well..
	req.Id = GenerateUUID()
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	log.Printf("hashed password : %v", hashedPass)
	if err != nil {
		log.Printf("could not encrypt the string %v", err)
	}
	// this is now going to hash the password, when we store them in our database
	req.Password = string(hashedPass)
	if req.Email == "" {
		req.Email = req.Name + "." + req.Company + "@gmail.com"
	}
	db.Create(&req)
	resp.User = req

	// before returning we are going to publish this event.
	// we are going to publish the request object as well.
	if err = publisher.Publish(ctx, req); err != nil {
		log.Printf("could not publish the event... %v", err)
	}

	return nil
}

func (s *serv) Auth(ctx context.Context, req *user.User, resp *user.Token) error {
	// used to query the database for a particular user and check if the passwords match
	log.Printf("logging in with %s", req.Email)

	// we will get my email
	u := &user.User{}
	// .First() wants a struct to unpack into
	db.Where(&user.User{Email: req.Email}).First(&u)

	if u.Password == "" {
		return errors.New("go.micro.srv.user", "user could not be found", 400)
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password))
	if err != nil {
		log.Printf("that password does not match the one associated with our registry %v", err)
		resp.Token = "failed"
		return err
	}

	resp.Token = JWToken()
	resp.Valid = true

	return nil

}

/*string id = 1;
string name = 2;
string company = 3;
string email = 4;
string password = 5;

*/

var DB_USER = getEnv("PGUSER", "postgres")
var DB_PASSWORD = getEnv("PGPASSWORD", "postgres")
var DB_NAME = getEnv("PGDATABASE", "postgres")
var DB_HOST = getEnv("PGHOST", "localhost")
var K8S = getEnv("K8S", "false")

var topic = "user.created"

var db *gorm.DB
var err error
var publisher micro.Publisher

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

// a type satisfies an interface if it possess all of the methods
// once it does we are able to reassign the variable to another interface with the same methods

func main() {
	var service micro.Service
	if K8S == "false" {
		service = micro.NewService(
			micro.Name("go.micro.srv.user"),
			micro.RegisterTTL(time.Second*30),
			micro.RegisterInterval(time.Second*10),
		)
	} else {
		service = k8s.NewService(
			micro.Name("go.micro.srv.user"),
			micro.RegisterTTL(time.Second*30),
			micro.RegisterInterval(time.Second*10),
		)
	}

	// database
	dbinfo := fmt.Sprintf("host=%s port=5432 user=%s dbname=%s password=%s sslmode=disable",
		DB_HOST, DB_USER, DB_NAME, DB_PASSWORD)
	db, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		log.Printf("could not connect %v", err)
	}
	log.Print("connected to database")
	db.AutoMigrate(&user.User{})

	service.Init()

	user.RegisterUserServiceHandler(service.Server(), new(serv))

	// set up publisher
	// public takes in a topic to publish to
	publisher = micro.NewPublisher(topic, service.Client())

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
