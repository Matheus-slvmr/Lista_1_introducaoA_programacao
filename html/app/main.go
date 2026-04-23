package main
import (
	"fmt"
	"log"
	"net/http"
)

func main (){
	//cria um servidor de arquivos que serve os arquivos da pasta ./static.
	fileserver := http.FileServer(http.Dir("./static"))

	//associa o servidor de arquivos a rota raiz
	http.Handle("/", fileserver)

	fmt.Printf("port running on http://localhost:8081/\n")

	if err := http.ListenAndServe(":8081",nil); err != nil {
		log.Fatal(err)
	}


}