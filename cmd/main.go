package main

import (
	"fmt"
	"gotickets/internal/config"
	"gotickets/internal/user"
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

config:=config.LoadEnv()

	dsn := config.Dsn
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

	user.RegisterRoutes(e, db)

	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
