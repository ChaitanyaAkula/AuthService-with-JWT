package server

import(
 "os"
 "github.com/gorilla/mux"
 "AuthService/api/v1/handlers"
 "AuthService/api/v1/services"
 "net/http"
 "AuthService/config"
 "log"
 "fmt"
)

type ServerInterface interface{
	Start()
}
type Server struct{}

func NewServer() ServerInterface {
	return Server{}
}

func(s Server)Start(){
	os.Setenv("ENV","TEST")

	serverconfigs:=config.GetServerConfig()
	fmt.Println(serverconfigs)
	router:=mux.NewRouter()
	authService:=services.NewAuthService()
	handlers:=handlers.NewAuthHandler(authService)
	router.HandleFunc("/",handlers.Homepage)


	authRouter:=router.Methods(http.MethodPost).Subrouter()
	authRouter.HandleFunc("/auth",handlers.Authentication)
	
	srv:=&http.Server{
		Handler: router,
		Addr : "localhost:9000",

	}
	log.Fatal(srv.ListenAndServe())
}
