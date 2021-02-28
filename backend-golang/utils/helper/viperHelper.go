package helper

import (
	"github.com/spf13/viper"
)

//GetKeyValue is a function for get a value from a key from .env file
func GetKeyValue(keyName, defaultValue string) string {
	viper.SetConfigName("configvalue")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	defaultKeyValueError, keyError := GetKeyError(err, keyName, defaultValue, " | Router Key Name Error")
	if keyError {
		return defaultKeyValueError
	}
	value, valueErr := viper.Get(keyName).(string)
	GetValueError(keyName, valueErr, "Router Key Name Value Error")

	return value
}
