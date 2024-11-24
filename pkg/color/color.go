package color

import "fmt"

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

func PurpleString(str string) string {
	return fmt.Sprintf("%s%s%s", Purple, str, Reset)
}

func GreenString(str string) string {
	return fmt.Sprintf("%s%s%s", Green, str, Reset)
}

func RedString(str string) string {
	return fmt.Sprintf("%s%s%s", Red, str, Reset)
}
