package ngrok

import (
	"fmt"
	"github.com/atotto/clipboard"
	"os/exec"
	"regexp"
)

var reg = regexp.MustCompile("(https://[a-z0-9]+(\\.)([a-z]{2}\\.)?(ngrok.io))")

type w struct {
	filename string
	key      string
}

func (ww w) Write(p []byte) (n int, err error) {
	s := string(p)
	u := reg.FindString(s)
	if len(u) > 0 {
		full := fmt.Sprintf("%s/?key=%s", u, ww.key)
		curl := fmt.Sprintf("curl -o \"%s\" \"%s/download?key=%s\"", ww.filename, u, ww.key)
		fmt.Println("Avalible through Ngrok at \n ", full)
		fmt.Println(" ", curl)
		fmt.Println("and url is copied to clipboard")
		clipboard.WriteAll(full)
	}

	return len(p), nil
}

func Start(filename string, key string, port int) {
	go start(filename, key, port)
}

func start(filename string, key string, port int) {
	cmd := exec.Command("ngrok", "http", "--log", "stdout", fmt.Sprintf("%d", port))
	cmd.Stdout = w{key: key, filename: filename}
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
}
