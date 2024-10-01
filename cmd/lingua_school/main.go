package lingua_school

import (
	"fmt"
	"lingua_school/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

}
