package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/square/certstrap/pkix"
)

var nowFunc = time.Now

func init() {
	timeNow := os.Getenv("CERTSTRAP_CURRENT_TIME")
	if timeNow != "" {
		d, err := time.Parse(time.RFC3339, timeNow)
		if err == nil {
			fmt.Printf("use %s for current time\n", d)
			nowFunc = func() time.Time {
				return d
			}
		}
	}
	pkix.TimeNow = nowFunc
}
