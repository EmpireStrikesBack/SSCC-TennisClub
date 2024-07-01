package ext

import (
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
)

func StartServer() {
	http.HandleFunc("/", Home)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("(http://localhost:1994) - Server started on port 1994")

	go func() {
		OpenInBrowser("http://localhost:1994")
	}()

	fmt.Println("starting server at port 1994")
	if err := http.ListenAndServe(":1994", nil); err != nil {
		fmt.Println(err)
	}
}

func OpenInBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	case "windows":
		err = exec.Command("cmd", "/c", "start", url).Start()
	default:
		err = fmt.Errorf("unsupported plateform")
	}

	if err != nil {
		fmt.Println("failed to open browser", err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}
