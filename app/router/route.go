package router

import (
	cd "github.com/GroupProject2-Kelompok4/BE/features/class/data"
	ch "github.com/GroupProject2-Kelompok4/BE/features/class/handler"
	cs "github.com/GroupProject2-Kelompok4/BE/features/class/service"
	fd "github.com/GroupProject2-Kelompok4/BE/features/feedback/data"
	fh "github.com/GroupProject2-Kelompok4/BE/features/feedback/handler"
	fs "github.com/GroupProject2-Kelompok4/BE/features/feedback/service"
	md "github.com/GroupProject2-Kelompok4/BE/features/mentee/data"
	mh "github.com/GroupProject2-Kelompok4/BE/features/mentee/handler"
	ms "github.com/GroupProject2-Kelompok4/BE/features/mentee/service"
	ud "github.com/GroupProject2-Kelompok4/BE/features/user/data"
	uh "github.com/GroupProject2-Kelompok4/BE/features/user/handler"
	us "github.com/GroupProject2-Kelompok4/BE/features/user/service"
	"github.com/GroupProject2-Kelompok4/BE/utils/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	e.Use(middleware.CORS())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	initUserRouter(db, e)
	initClassRouter(db, e)
	initMenteeRouter(db, e)
	initFeedbackRouter(db, e)
}

func initUserRouter(db *gorm.DB, e *echo.Echo) {
	userData := ud.New(db)
	userService := us.New(userData)
	userHandler := uh.New(userService)

	e.POST("/login", userHandler.Login())
	e.POST("/register", userHandler.Register(), middlewares.JWTMiddleware())          //*** by admin
	e.GET("/users", userHandler.SearchUser(), middlewares.JWTMiddleware())            //*** both
	e.GET("/users/:id", userHandler.ProfileUser(), middlewares.JWTMiddleware())       //*** both
	e.DELETE("/users/:id", userHandler.DeactiveUser(), middlewares.JWTMiddleware())   //*** by admin
	e.PUT("/users", userHandler.UpdateProfile(), middlewares.JWTMiddleware())         //*** by user
	e.PUT("/users/:id", userHandler.UpdateUserProfile(), middlewares.JWTMiddleware()) //*** by admin
}

func initClassRouter(db *gorm.DB, e *echo.Echo) {
	classData := cd.New(db)
	classService := cs.New(classData)
	classHandler := ch.New(classService)

	e.POST("/classes", classHandler.RegisterClass(), middlewares.JWTMiddleware())
	e.GET("/classes", classHandler.ListClasses(), middlewares.JWTMiddleware())
	e.GET("/classes/:id", classHandler.GetClass(), middlewares.JWTMiddleware())
	e.DELETE("/classes/:id", classHandler.DeleteClass(), middlewares.JWTMiddleware())
	e.PUT("/classes/:id", classHandler.UpdateClass(), middlewares.JWTMiddleware())
}

func initMenteeRouter(db *gorm.DB, e *echo.Echo) {
	menteeData := md.New(db)
	menteeService := ms.New(menteeData)
	menteeHandler := mh.New(menteeService)

	e.POST("/mentees", menteeHandler.RegisterMentee(), middlewares.JWTMiddleware())
	e.GET("/mentees", menteeHandler.SearchMentee(), middlewares.JWTMiddleware())
	e.GET("/mentees/:id", menteeHandler.ProfileMentee(), middlewares.JWTMiddleware())
	e.GET("/mentees/:id/feedbacks", menteeHandler.ProfileMenteeAndFeedback(), middlewares.JWTMiddleware())
	e.PUT("/mentees/:id", menteeHandler.UpdateMentee(), middlewares.JWTMiddleware())
	e.DELETE("/mentees/:id", menteeHandler.DeactiveMentee(), middlewares.JWTMiddleware())
}

func initFeedbackRouter(db *gorm.DB, e *echo.Echo) {
	feedbackData := fd.New(db)
	feedbackService := fs.New(feedbackData)
	feedbackHandler := fh.New(feedbackService)

	e.POST("/feedbacks", feedbackHandler.RegisterFeedbackMentee(), middlewares.JWTMiddleware())
	e.PUT("/feedbacks/:id", feedbackHandler.UpdateFeedbackMentee(), middlewares.JWTMiddleware())
}
