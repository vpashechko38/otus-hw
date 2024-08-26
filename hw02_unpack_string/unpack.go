package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	char := ""
	sb := strings.Builder{}

	for _, ch := range input {
		// проверяем, число это или нет
		dig, err := strconv.Atoi(string(ch))
		if len(char) == 0 && err == nil {
			return "", ErrInvalidString
		}
		// нашли число после символа
		if len(char) != 0 && err == nil {
			for i := 0; i < dig; i++ {
				sb.WriteString(char)
			}
			char = ""
			continue
		}
		// одинокий символ без числа
		if len(char) != 0 && err != nil {
			sb.WriteString(char)
			char = ""
		}
		// нашли символ, но пока не нашли число
		if len(char) == 0 && err != nil {
			char = string(ch)
		}
	}
	// одинокий символ без числа в конце
	if char != "" {
		sb.WriteString(char)
	}
	return sb.String(), nil
}
