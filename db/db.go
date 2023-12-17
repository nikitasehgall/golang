package db

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

// Your database-related code goes here

var db *gorm.DB
var err error

type Movie struct {
	ID          string `json:"id" gorm:"primarykey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func InitPostgresDB() {
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		dbUser   = os.Getenv("DB_USER")
		dbName   = os.Getenv("DB_NAME")
		password = os.Getenv("DB_PASSWORD")
	)
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)

	db, err = gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(Movie{})
}
func CreateMovie(movie *Movie) (*Movie, error) {
	movie.ID = uuid.New().String()
	res := db.Create(&movie)
	if res.Error != nil {
		return nil, res.Error
	}
	return movie, nil
}
func GetMovie(id string) (*Movie, error) {
	var movie Movie
	res := db.First(&movie, "id = ?", id)
	if res.RowsAffected == 0 {
		return nil, errors.New(fmt.Sprintf("movie of id %s not found", id))
	}
	return &movie, nil
}
func GetMovies() ([]*Movie, error) {
	var movies []*Movie
	res := db.Find(&movies)
	if res.Error != nil {
		return nil, errors.New("no movies found")
	}
	return movies, nil
}
func UpdateMovie(movie *Movie) (*Movie, error) {
	var movieToUpdate Movie
	result := db.Model(&movieToUpdate).Where("id = ?", movie.ID).Updates(movie)
	if result.RowsAffected == 0 {
		return &movieToUpdate, errors.New("movie not updated")
	}
	return movie, nil
}
func DeleteMovie(id string) error {
	var deletedMovie Movie
	result := db.Where("id = ?", id).Delete(&deletedMovie)
	if result.RowsAffected == 0 {
		return errors.New("movie not deleted")
	}
	return nil
}
func postMovie(ctx *gin.Context) {
	var movie db.Movie
	err := ctx.Bind(&movie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := db.CreateMovie(&movie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"movie": res,
	})
}
func getMovies(ctx *gin.Context) {
	res, err := db.GetMovies()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"movies": res,
	})
}

func getMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := db.GetMovie(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"movie": res,
	})
}
func putMovie(ctx *gin.Context) {
	var updatedMovie db.Movie
	err := ctx.Bind(&updatedMovie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := ctx.Param("id")
	dbMovie, err := db.GetMovie(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	dbMovie.Name = updatedMovie.Name
	dbMovie.Description = updatedMovie.Description

	res, err := db.UpdateMovie(dbMovie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"task": res,
	})
}
func deleteMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	err := db.DeleteMovie(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "task deleted successfully",
	})
}
