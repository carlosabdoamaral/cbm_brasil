package common

import "log"

func LogInfo(s string) {
	log.Printf("[*] %s", s)
}

func LogWarning(s string) {
	log.Printf("[!] %s", s)
}

func LogError(s string) {
	log.Printf("[!] %s", s)
}
