package utils

import (
	"fmt"
	"strings"
)

func RespTemplate(desc string) {
	fmt.Println(strings.Repeat("=~=", 10))
	fmt.Println(desc)
	fmt.Println(strings.Repeat("=~=", 10))
}
