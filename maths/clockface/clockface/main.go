package main

import (
	"os"
	"time"

	"github.com/Popoliku/cursoGo/maths/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
