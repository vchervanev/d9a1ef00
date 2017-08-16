package main

import ("net/http"; "fmt"; "runtime")

func DumbHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello, I'm Web Server!"))
	fmt.Println("OS: %s\nArchitecture: %s\n", runtime.GOOS, runtime.GOARCH)
}

func main() {
	http.HandleFunc("/", DumbHandler)
	http.ListenAndServe(":80", nil)
}
