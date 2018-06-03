package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
)

var (
	flagPort = flag.String("port", ":1313", "change the port of the server")
)

func main() {
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s := formatVar(*r)
		w.Write([]byte(s))
	})
	fmt.Printf("serving at http://localhost%s\n", *flagPort)
	err := http.ListenAndServe(*flagPort, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func formatVar(i interface{}) string {
	s := fmt.Sprintf("%#v", i)
	s = strings.Replace(s, ", ", ",\n\n", -1)
	s = strings.Replace(s, ":", ":\n", -1)
	return s
}
