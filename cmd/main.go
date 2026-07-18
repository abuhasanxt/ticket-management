package main

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name     string `json:"name" validate:"required" gorm:"type:varchar(100); not null"`
	Email    string `json:"email" validate:"required,email" gorm:"type:varchar(250); uniqueIndex; not null"`
	Password string `json:"password" validate:"required,min=6" gorm:"type:varchar(100); not null"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally return the error to let each route control the status code.
		return echo.ErrBadRequest.Wrap(err)
	}
	return nil
}

func main() {

	dsn := "host=localhost user=postgres password=abu##228 dbname=gotickets port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{

		TranslateError: true,
	})
	
	if err != nil {
		panic("failed to connect database")
		} else {
			fmt.Println("Database connect successfully!")
		}
		
		db.AutoMigrate(&Users{})

	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
	})

	e.Validator = &CustomValidator{validator: validator.New()}
	e.POST("/users", func(c *echo.Context) error {
		newUser := new(Users)

		//binding the user data
		if err := c.Bind(newUser); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{"error": err.Error()})

		}
		//validating the user data
		if err := c.Validate(newUser); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
		//save the database
result:=db.Create(&newUser)

if result.Error!=nil {
  return c.JSON(http.StatusInternalServerError,map[string]any{"error":result.Error.Error()})
}
		return c.JSON(http.StatusCreated, newUser)
	})

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
