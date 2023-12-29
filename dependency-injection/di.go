package di

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// this works because io.Writer is the interface of both os.Stdout and bytes.Buffer
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Tin")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
