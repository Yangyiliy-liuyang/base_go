package wire

import "fmt"

func useUserRepository() {
	repository := InitUserRepository()
	fmt.Print(repository)
}
