package handlers

import(
	"net/http"
	"AuthService/api/v1/services"
	"log"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"AuthService/utils"
)

type HandlerInterface interface{
	Authentication(http.ResponseWriter,*http.Request)
}
type AuthHandler struct{
	AuthService services.AuthInterface
}
func NewAuthHandler(authService services.AuthInterface)*AuthHandler{
	Auth:=&AuthHandler{AuthService:authService}
	return Auth
}
func (AUH *AuthHandler)Homepage(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"HomePage")
	utils.ResponseJson(w,`{"Success":true}`)
}

func (AUH *AuthHandler)Authentication(w http.ResponseWriter,r *http.Request){
	
	requestData:=new(interface{})

	body,_:=ioutil.ReadAll(r.Body)

	err:=json.Unmarshal([]byte(body),&requestData)
	if err!=nil{
		log.Fatal("Error while unmarshalling request body",err)
	}

	tokenString:=r.Header.Get("Token")
	fmt.Println(tokenString)
	
	response:=AUH.AuthService.Authenticate(tokenString,requestData)

	utils.ResponseJson(w,response)

}