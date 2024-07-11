package main

import (
	"github.com/gustavlsouz/trucks-service/pkg"
)

func main() {
	pkg.Start(make(chan<- bool), "../.env.local", "../deployments/migrations")
}
