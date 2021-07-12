package web

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func Serve(fiber []Fiber) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		nSlice, ok := request.URL.Query()["n"] //http://localhost:8080/?n=10&n=20

		rand.Seed(time.Now().UnixNano())
		fiberRnd := rand.Intn(len(fiber))

		data := fmt.Sprintf("%s\n", fiber[fiberRnd].Name())
		if ok {
			for _, nValue := range nSlice {
				n, err := strconv.Atoi(nValue)
				if err == nil {
					res := fiber[fiberRnd].Fib(n)
					if res != -1 {
						data += fmt.Sprintf("Fib(%d) = %d\n", n, res)
					} else {
						data += fmt.Sprintf("Fib(%d) = integer overflow\n", n)
					}
				}
			}
		}

		_, err := writer.Write([]byte(data))
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
