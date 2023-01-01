package common

import (
	"fmt"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

func LogInfo(str string) {
	str = fmt.Sprintf("[*] %s - %s", TimeToString(nil), str)
	fmt.Println(Green + str + Reset)
}

func LogWarning(str string) {
	str = fmt.Sprintf("[!] %s - %s", TimeToString(nil), str)
	fmt.Println(Yellow + str + Reset)
}

func LogError(str string) {
	str = fmt.Sprintf("[!!] %s - %s", TimeToString(nil), str)
	fmt.Println(Red + str + Reset)
}
