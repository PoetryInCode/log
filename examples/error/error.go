package main

import (
	"fmt"

	"github.com/PoetryInCode/log"
)

func main() {
	err := fmt.Errorf("too much sugar")
	log.Error("failed to bake cookies", "err", err)
}
