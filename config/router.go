package config

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"lion-test/controllers"
	"lion-test/handlers"
	"lion-test/middleware"
	"lion-test/repositories"
)

func NewRouter(r *gin.Engine, db *gorm.DB) {
	//#region Library

	movieRepository := repositories.NewMovieRepository(db)
	movieController := controllers.NewMovieController(movieRepository)
	movieHandler := handlers.NewMovieHandler(movieController)

	genreRepository := repositories.NewGenreRepository(db)
	genreController := controllers.NewGenreController(genreRepository)
	genreHandler := handlers.NewGenreHandler(genreController)

	userRepository := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepository)
	userHandler := handlers.NewUserHandler(userController)

	authHandler := handlers.NewAuthHandler(userRepository)

	voteRepository := repositories.NewVoteRepository(db)
	voteController := controllers.NewVoteController(voteRepository)
	voteHandler := handlers.NewVoteHandler(voteController, movieController)

	//#region Login

	r.POST("/login", authHandler.Login)

	//#endregion

	//#region RouteMovie

	r.POST("/movie/insert-movie", middleware.AuthMiddleware([]string{"admin"}), movieHandler.InsertMovieHandler)
	r.POST("/movie/update-movie", middleware.AuthMiddleware([]string{"admin"}), movieHandler.UpdateMovieHandler)
	r.POST("/movie/get-movie", middleware.AuthMiddleware([]string{"admin", "user"}), movieHandler.GetMovieHandler)
	r.POST("/movie/get-movie-detail", middleware.AuthMiddleware([]string{"admin", "user"}), movieHandler.GetMovieDetailHandler)
	r.POST("/movie/get-most-viewed", middleware.AuthMiddleware([]string{"admin", "user"}), movieHandler.GetMostViewedHandler)

	//#endregion

	//#region RouteGenre

	r.POST("/genre/insert-genre", middleware.AuthMiddleware([]string{"admin"}), genreHandler.InsertGenreHandler)
	r.POST("/genre/get-genre", middleware.AuthMiddleware([]string{"admin"}), genreHandler.GetGenreHandler)
	r.POST("/genre/get-genre-detail", middleware.AuthMiddleware([]string{"admin"}), genreHandler.GetGenreDetailHandler)
	r.POST("/genre/get-most-viewed", middleware.AuthMiddleware([]string{"admin"}), genreHandler.GetMostViewedHandler)

	//#endregion

	//#region RouteUser

	r.POST("/user/insert-user", middleware.AuthMiddleware([]string{"admin"}), userHandler.InsertUserHandler)
	r.POST("/user/generate-admin", middleware.AuthMiddleware([]string{"admin"}), userHandler.GenerateAdminHandler)

	//#endregion

	//#region RouteTransaction

	r.POST("/vote/vote", middleware.AuthMiddleware([]string{"user"}), voteHandler.UpdateVoteHandler)
	r.POST("/vote/check-vote", middleware.AuthMiddleware([]string{"admin"}), voteHandler.CheckVoteHandler)

	//#endregion
}
