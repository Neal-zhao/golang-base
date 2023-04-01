package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayGood(w http.ResponseWriter,r *http.Request)  {
	date := []byte
	t,_ := template.ParseFiles("./net.go")
	_ =t.Execute(w,date)
	abc := r.Form.Get("abc")
	fmt.Println("abc: ",abc)
	fmt.Fprintf(w,"%s %s","abc back",abc)
}
func main()  {
	http.HandleFunc("/",sayGood)
	err := http.ListenAndServe(":900",nil)
	if err != nil {
		fmt.Println("err",err)
	}

}

