package config

import(
	"os"
	"fmt"
)

func getstringEmpty(key string)string{

	value:=os.Getenv(key)
	fmt.Println("Value:",value)
	return value
}
