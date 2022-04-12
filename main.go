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
    log.Println("info: starting")
    dir := os.Args[1]
    if(dir == "" || dir == " "){
     log.Println("fatal: no directory provided")
     os.Exit(1)
    }  
    log.Println("info running from dir:",dir)
    http.Handle("/", http.FileServer(http.Dir(string(dir))))
    if os.Args[2] == "" {
      port = ":5000"
    } else{
      port = ":"+ os.Args[2]
    }
    log.Println("info: starting on port",port)
    log.Fatal(http.ListenAndServe(port, nil))
}
