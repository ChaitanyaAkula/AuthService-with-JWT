package main

import(
	"AuthService/server"
	"os"
)
func main(){
	//Starting Server
	os.Setenv("ENVIRONMENT","TEST")
	
	ServiceServer:=server.NewServer()

	ServiceServer.Start()

}