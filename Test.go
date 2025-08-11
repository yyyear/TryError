package TryError

import (
	"github.com/yyyear/YY"
	"os"
)

func Test() {
	j, e := os.ReadFile("config.json")
	j2 := Try(j, e).Do()
	YY.Info(j2)
}
