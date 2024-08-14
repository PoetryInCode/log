package main

import (
	"os"

	"github.com/PoetryInCode/log"
)

func main() {
	logger := log.New(os.Stderr)
	logger.Warn("chewy!", "butter", true)
}
