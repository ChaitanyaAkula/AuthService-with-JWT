package services

import(
	"testing"
	"github.com/stretchr/testify/assert"

)
var(
	AuthService AuthInterface
)
func TestAuth(t *testing.T){
   
	tokenString:="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.Dk2f4_2nIMho23sow-H7eO9dOp1NiSq7JkPBtN1RtI0"
	var requestData interface{}
	Expectedresponse:=AuthService.Authenticate(tokenString,requestData)
	  // assert equality
	  assert.Equal(t,200,Expectedresponse.Responsecode, "they should be equal")
	
}