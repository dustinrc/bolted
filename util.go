package bolted

import (
	"fmt"
	"os"
)

func Die(format string, v ...interface{}) {
	fmt.Println(os.Stderr, fmt.Sprintf(format, v...))
	os.Exit(1)
}
