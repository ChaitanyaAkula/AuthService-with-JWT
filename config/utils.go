package config

import(


)

func getstringEmpty(key string)string{

	value:=os.Getenv(key)
	return value
}
