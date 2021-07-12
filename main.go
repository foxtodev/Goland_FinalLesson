package main

import (
	"fmt"
	"github/foxtodev/Goland_FinalLesson/cache"
	"github/foxtodev/Goland_FinalLesson/closure"
	"github/foxtodev/Goland_FinalLesson/recursion"
	"github/foxtodev/Goland_FinalLesson/web"
)

func main() {
	fibImplementations := []web.Fiber{recursion.NewRecursion(), cache.NewCache(), closure.NewClosure()}
	fmt.Printf("Starting server http://localhost:8080/. Use after ?n=10&n=20")
	web.Serve(fibImplementations)
}
