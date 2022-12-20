package main;

import(
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound);
		return;
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound);
		return;
	}

	fmt.Fprintf(w, "Hello!");
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	// Parse html form
	err := r.ParseForm();

	if err != nil {
		fmt.Fprint(w, "ParseError: %w", err);
		return;
	}

	fmt.Fprintf(w, "Post request successful");

	// Get form fields
	name := r.FormValue("name");
	address := r.FormValue("address");

	fmt.Fprint(w, "Name - %s\n", name);
	fmt.Fprintf(w, "Address - %s\n", address);
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"));

	http.Handle("/", fileserver);

	http.HandleFunc("/form", formHandler);

	http.HandleFunc("/hello", helloHandler);

	fmt.Println("Starting server at port 8080");

	err := http.ListenAndServe(":8080", nil); 
	 
	if err != nil {
		log.Fatal(err);
	}
}