package main

import (
	"fmt"

	"github.com/RivellionCS/gator/internal/config"
)

func main() {
	new_config := config.Read()
	new_config.SetUser("RivellionCS")
	new_config = config.Read()
	fmt.Printf("%v\n", new_config)
}