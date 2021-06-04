package server

import(
 "os"
 "github.com/gorilla/mux"
 "AuthService/api/v1/handlers"
 "net/http"
 "AuthService/config"
)

type ServerInterface interface{
	Start()
}
type Server struct{}

type NewServer() ServerInterface {
	return Server{}
}

func(s Server)Start(){
	os.Setenv("ENV","TEST")

	router:=mux.NewRouter()

	router.HandleFunc("/",Homepage)

	serverConfigs:=config.GetServerConfig()
	authRouter:=router.Methods(http.MethodPost).Subrouter()
	router.HandleFunc("/auth",handlers.Authentication)
	
	srv:=&http.Server{
		Handler: router,
		Addr : serverConfigs.Host+":"+serverConfigs.Port,

	}
	log.Fatal(srv.ListenAndServer())
}
