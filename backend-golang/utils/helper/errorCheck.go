package helper

import (
	"fmt"
	"log"
)

//FatalError is a function for handle fatal error
func FatalError(err error, message string) {
	if err != nil {
		log.Println(message)
		LogApp(message)
		log.Fatal(err)
	}
}

//GetKeyError is a function for handle keyname error
func GetKeyError(err error, keyName string, defaultValue string, message string) (string, bool) {
	if err != nil {
		log.Println(message)
		log.Printf("Keyname : %v, not found, default key value : %v, has been loaded", keyName, defaultValue)
		LogApp(message)
		LogApp("Keyname : " + keyName + ", not found, default key value : " + defaultValue + ", has been loaded")
		return defaultValue, true
	}
	return "", false
}

//GetValueError is a function for handle value from a keyname error
func GetValueError(keyName string, err bool, message string) {
	if !err {
		log.Println(message)
		fmt.Println(keyName)
		LogApp(message)
		LogApp(keyName)
		log.Fatalf("Invalid type assertion")
	}
}

//ErrorNotNil is a function for check create new user error
func ErrorNotNil(err error) bool {
	if err != nil {
		log.Println(err)
		log.Print(" | Error is not nil")
		LogApp("Error not nil")
		return true
	}
	return false
}
