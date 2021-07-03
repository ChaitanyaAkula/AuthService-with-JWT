package services

import(
	"testing"
	"github.com/stretchr/testify/assert"
	"AuthService/models"
)
var(
	AuthenticateFunction func(tokenString string,requestData interface{})*models.AuthResponse
)
type ServiceMock struct{}

func(s *ServiceMock)Authenticate(tokenString string,requestData interface{})*models.AuthResponse{
	return AuthenticateFunction(tokenString,requestData)
}
func TestAuthSuccess(t *testing.T){
	AuthService=&ServiceMock{}
	AuthenticateFunction=func(tokenString string,requestData interface{})*models.AuthResponse{
		response:=new(models.AuthResponse)
		response.Responsecode=200
		response.Status="Success"
		return response
	}

	tokenString:="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.Dk2f4_2nIMho23sow-H7eO9dOp1NiSq7JkPBtN1RtI0"
	var requestData interface{}
	Actualresponse:=AuthService.Authenticate(tokenString,requestData)
	  // assert equality
	  assert.Equal(t,200,Actualresponse.Responsecode, "they should be equal")
	
}
func TestAuthFail(t *testing.T){
	AuthService=&ServiceMock{}
	AuthenticateFunction=func(tokenString string,requestData interface{})*models.AuthResponse{
		response:=new(models.AuthResponse)
		response.Responsecode=199
		response.Status="Success"
		return response
	}

	tokenString:="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.Dk2f4_2nIMho23sow-H7eO9dOp1NiSq7JkPBtN1RtI0"
	var requestData interface{}
	Actualresponse:=AuthService.Authenticate(tokenString,requestData)
	  // assert equality
	  assert.Equal(t,200,Actualresponse.Responsecode, "they should be equal")
	
}