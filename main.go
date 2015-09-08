package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/mongolar/frontend/environment"
)

func main() {
	spew.Dump(environment.Env)

}
