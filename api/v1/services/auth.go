package services

import(
	"AuthService/models"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
)
var(
	AuthService AuthInterface
)
func init(){
	AuthService=&Authenticate{}
}
var mySecretKey= []byte("My Secret key is not secret@dont reveal nkln")
type Authenticate struct{}

type AuthInterface interface{
	Authenticate(string,interface{})*models.AuthResponse
}
func NewAuthService()*Authenticate{
	auth:=&Authenticate{}
	return auth
}

func (aus *Authenticate)Authenticate(tokenString string,requestData interface{})*models.AuthResponse{

	token,err:=jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
				
		if _,ok:=token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("we got some error: %v",token.Header["alg"])
		}
		return mySecretKey,nil
	})
	if err!=nil{
		return aus.GenerateResponse(requestData,false,190,err)
		fmt.Println(err.Error())
	}
	if !token.Valid{
		return aus.GenerateResponse(requestData,false,190,err)
	}
	return aus.GenerateResponse(requestData,true,200,nil)
}

func (aus *Authenticate)GenerateResponse(requestData interface{},success bool,responseCode int64,err error)*models.AuthResponse{

	response:=new(models.AuthResponse)

	if err!=nil{
		fmt.Println(err.Error())
	}
	response.Status="fail"
	if success && responseCode==200{
		response.Responsecode=200
		response.Status="success"
	}else if success && responseCode ==299{
		response.Responsecode=299
		response.Status="processing"
	} else{
		response.Responsecode=190
		response.Status="failed"
	}

	return response
}