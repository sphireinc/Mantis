package mantis

import "log"

// Handle error
func HandleError(message string, err error) {
	if err != nil {
		log.Println(message, err)
	}
}

func HandleFatalError(err error) {
	if err != nil {
		panic(err)
	}
}
