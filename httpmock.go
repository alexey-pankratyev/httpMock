// This module with simple http server to procces and logging to file
// example: 
// get - curl -si "http://localhost:4000/?foo=1&bar=2"
// post - curl -si -X POST -d "test for you" http://localhost:4000/

package main

import (
    "os"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
	logPath := "development.log"
	httpPort := 4000

	openLogFile(logPath)

    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    
    http.HandleFunc("/", reqHand)

	fmt.Printf("listening on %v\n", httpPort)
	fmt.Printf("Logging to %v\n", logPath)

	err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), logRequest(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
    }
    
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        switch r.Method {
        case "GET":
                for k, v := range r.URL.Query() {
                       fmt.Printf("%s: %s\n", k, v)
                       log.Printf("%s - %s: %s\n", r.Method, k, v)
                }
        case "POST":
                 reqBody, err := ioutil.ReadAll(r.Body)
                 if err != nil {
                          log.Fatal(err)
                 }
                 fmt.Printf("%s - %s\n", r.Method, reqBody)
                 log.Printf("%s - %s\n", r.Method, reqBody)
        default:
                 w.WriteHeader(http.StatusNotImplemented)
                 w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
        }
		handler.ServeHTTP(w, r)
	})
}

func reqHand(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
            http.NotFound(w, r)
            return
    }
}

func openLogFile(logfile string) {
	if logfile != "" {
		lf, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)

		if err != nil {
			log.Fatal("OpenLogfile: os.OpenFile:", err)
		}

		log.SetOutput(lf)
	}
}



