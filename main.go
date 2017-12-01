package main	
import (
	"net/http"
	"fmt"
    "io/ioutil"
)
type Page struct {
    Body  []byte
}
func main() {
	http.HandleFunc("/", indexHandler)    
    http.ListenAndServe(":9090", nil)
}
func indexHandler(w http.ResponseWriter, r *http.Request){
    p, _ := loadPage()
    fmt.Fprintf(w, "%s", p.Body)
}
func loadPage() (*Page, error) {
    filename := "./myApp/www/index.html"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Body: body}, nil
}