package common

import "log"

func LogError(message string) {
	log.Printf("\n[!!] %s", message)
}

func LogInfo(message string) {
	log.Printf("\n[*] %s", message)
}

func LogSuccess(message string) {
	log.Printf("\n[*] %s", message)
}

func LogWarning(message string) {
	log.Printf("\n[!] %s", message)
}
