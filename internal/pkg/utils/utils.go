package utils

import (
	"encoding/json"
	"os"
)

// GetEnv - get the value of an env-var with a fallback
func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}

// DeepCopy copies fields from in to out using json tags out should be address
func DeepCopy(in interface{}, out interface{}) {
	temp, _ := json.Marshal(in)
	json.Unmarshal(temp, out)

}
