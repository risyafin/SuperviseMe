package main

import (
	"fmt"
	"net/http"
	"os"
	"superviseMe/config"
	"superviseMe/core/module"
	"superviseMe/domain"
	"superviseMe/handler"
	cardrepository "superviseMe/repository/card-repository"
	commentrepository "superviseMe/repository/comment-repository"
	goalsrepository "superviseMe/repository/goals-repository"
	listrepository "superviseMe/repository/list-repository"
	notificationrepository "superviseMe/repository/notification-repository"
	userrepository "superviseMe/repository/user-repository"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func init() {
	godotenv.Load()
}

func main() {
	conf := config.GetConfig()
	conn := config.InitDatabaseConnection(conf)
	config.AutoMigration(conn)

	var (
		OAuthConfig = &oauth2.Config{
			ClientID:     os.Getenv("CLIENT_ID"),
			ClientSecret: os.Getenv("CLIENT_SECRET"),
			RedirectURL:  os.Getenv("REDIRECT_URL"),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
			Endpoint:     google.Endpoint,
		}
	)
	domain.InitGoogleConfig()

	GoalsRepo := goalsrepository.NewGoalsRepository(conn)
	notificationRepo := notificationrepository.NewNotificationRepository(conn)
	GoalsUsecase := module.NewGoalsUseCase(GoalsRepo, notificationRepo)
	GoalsHandler := handler.NewGoalsHandler(GoalsUsecase)

	UserRepo := userrepository.NewUserRepository(conn)
	UserUsecase := module.NewUserUseCase(UserRepo)
	UserHandler := handler.NewUserHandler(UserUsecase, OAuthConfig)

	ListRepo := listrepository.NewListRepository(conn)
	ListUsecase := module.NewListUseCase(ListRepo)
	ListHandler := handler.NewListHandler(ListUsecase)

	CardRepo := cardrepository.NewCardRepository(conn)
	CardUsecase := module.NewCardUsecase(CardRepo)
	CardHandler := handler.NewCardHandler(CardUsecase)

	NotificationRepo := notificationrepository.NewNotificationRepository(conn)
	NotificationUsecase := module.NewNotificationUseCase(NotificationRepo)
	NotificationHandler := handler.NewNotificationHandler(NotificationUsecase)

	CommentRepo := commentrepository.NewCommentRepository(conn)
	CommentUsecase := module.NewCommentUsecase(CommentRepo)
	CommentHandler := handler.NewCommentHandler(CommentUsecase)

	const port string = ":8080"
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome, please login first"))
	}).Methods(http.MethodGet)

	r.HandleFunc("/auth/login", UserHandler.LoginGoogle)
	r.HandleFunc("/auth/google-callback", UserHandler.CallbackGoogle)
	r.HandleFunc("/registration", UserHandler.Registration).Methods("POST")
	r.HandleFunc("/login", UserHandler.Login).Methods("POST")
	r.HandleFunc("/profile", jwtMiddleware(UserHandler.GetUserByGmail)).Methods("GET")
	r.HandleFunc("/update-name", jwtMiddleware(UserHandler.UpdateName)).Methods("PUT")

	r.HandleFunc("/home", jwtMiddleware(UserHandler.GetGoalsBySuperviseeUser)).Methods("GET")
	r.HandleFunc("/supervisor", jwtMiddleware(UserHandler.GetGoalSupervisor)).Methods("GET")
	r.HandleFunc("/personal", jwtMiddleware(UserHandler.GetGoalPersonal)).Methods("GET")

	r.HandleFunc("/goals", jwtMiddleware(GoalsHandler.CreateGoals)).Methods("POST")
	r.HandleFunc("/goals/accept-supervisor", jwtMiddleware(GoalsHandler.AcceptSupervisorRequest)).Methods("PUT")
	// r.HandleFunc("/goals/reject-supervisor", jwtMiddleware(GoalsHandler.RejectSupervisorRequest)).Methods("POST")
	// r.HandleFunc("/goals/{id}", GoalsHandler.GetGoalsByID).Methods("GET")

	r.HandleFunc("/list", jwtMiddleware(ListHandler.GetList)).Methods("GET")
	// r.HandleFunc("/goals/{id}", GoalsHandler.DeleteGoals).Methods("PATCH")

	r.HandleFunc("/card", jwtMiddleware(CardHandler.CreateCard)).Methods("POST")

	r.HandleFunc("/notification", jwtMiddleware(NotificationHandler.GetNotification)).Methods("GET")

	r.HandleFunc("/cards/{card_id}/comments", jwtMiddleware(CommentHandler.CreateComment)).Methods("POST")

	fmt.Println("localhost:8080")
	http.ListenAndServe(port, r)
}
