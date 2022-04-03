package main
 
import (
    "log"
    "net/http"
    "os"
)
var(
  port string
)
func main() {
    log.Println("[INIT] started")
    dir := os.Args[1]
    log.Println("[INFO] running from static dir:",dir)
    http.Handle("/", http.FileServer(http.Dir(string(dir))))
    if os.Args[2] == "" {
      port = ":5000"
    } else{
      port = ":"+ os.Args[2]
    }
    log.Println("[INFO]running on port",port)

    log.Fatal(http.ListenAndServe(port, nil))
}
