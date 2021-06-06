package main

import(
	"AuthService/server"
	"os"
)
func main(){
	os.Setenv("ENVIRONMENT","TEST")
	
	ServiceServer:=server.NewServer()

	ServiceServer.Start()

}