package main  
  
import (  
	"encoding/json"  
	"fmt"  
	"net/http"  
)  
  
type MetaData struct {  
	Host string `json:"host"`  
	Method string `json:"method"`  
}  
  
func handler(w http.ResponseWriter, r *http.Request) {  
	MetaData := &MetaData{  
		Host:     r.URL.Host,  
		Method:   r.Method,  
	}  
  
	response, _ := json.Marshal(MetaData)  
	w.Header().Set("Content-Type", "application/json")  
	w.Write(response)  
  
	fmt.Fprint(w, "\nMetadata:\n")  
	fmt.Fprintf(w, "%v\n", MetaData)  
}  
  
func main() {  
	http.HandleFunc("/", handler)  
	http.ListenAndServe(":8080", nil)  
} 
