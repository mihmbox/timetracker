package main
import (
	"model"
	"fmt"
)



func main() {
	if err:=model.CreateDB(); err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB created")
	}
}