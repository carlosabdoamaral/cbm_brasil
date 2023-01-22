package common

import "log"

func LogError(message string) {
	log.Printf("[!!] %s", message)
}

func LogInfo(message string) {
	log.Printf("[*] %s", message)
}

func LogSuccess(message string) {
	log.Printf("[*] %s", message)
}

func LogWarning(message string) {
	log.Printf("[!] %s", message)
}
