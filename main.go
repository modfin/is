package main

import (
	"crypto/rand"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/crholm/qs/assets"
	"github.com/crholm/qs/ngrok"
	"github.com/phayes/freeport"
	"github.com/urfave/cli/v2"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func outboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func randString(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	getint := func() int {
		b := make([]byte, 4)
		_, err := rand.Read(b)
		check(err)
		v := int(b[0])<<24 | int(b[1])<<16 | int(b[2])<<8 | int(b[3])
		if v > 0 {
			return v
		}
		return -v
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[getint()%len(letterRunes)]
	}
	return string(b)
}

func main() {

	var ascii = true
	var nokey bool
	var once bool
	var ungrok bool
	var oncemux sync.Mutex
	var noclip bool
	var filepath string
	var filename string

	app := &cli.App{
		Name:  "qs",
		Usage: "Quick Share creates a local web server in order to share text snippets and files on your local network through http",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "nokey",
				Usage: "no access key is needed to view content ",
			},
			&cli.BoolFlag{
				Name:  "noclip",
				Usage: "do not copy url to clipboard",
			},
			&cli.BoolFlag{
				Name:    "ascii",
				Aliases: []string{"a"},
				Usage:   "will treat a file as text input and display it on the web, just as for std in",
			},
			&cli.BoolFlag{
				Name:    "once",
				Aliases: []string{"o"},
				Usage:   "terminates the server after first page load, so things dont hang around (copies curl command to clipboard)",
			},
			&cli.BoolFlag{
				Name:    "ngrok",
				Aliases: []string{"n"},
				Usage:   "starts an ngrok instance linking the snippet to it, see https://ngrok.com/. Make sure to have it in your $PATH",
			},
		},
		UsageText: "echo \"foo bar\" | qs [global options]\n   qs [global options] [filename]",
		Action: func(c *cli.Context) error {

			nokey = c.Bool("nokey")
			noclip = c.Bool("noclip")
			once = c.Bool("once")
			ungrok = c.Bool("ngrok")
			filepath = c.Args().First()
			if filepath != "" {
				ascii = c.Bool("ascii")
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	check(err)

	for _, a := range os.Args {
		if a == "help" || a == "--help" || a == "-h" {
			os.Exit(0)
		}
	}

	var key string
	if !nokey {
		key = randString(16)
	}

	var text string
	if ascii {
		var file io.Reader = os.Stdin
		var closer = func() {}
		if len(filepath) > 0 {
			f, err := os.Open(filepath)
			check(err)
			closer = func() { _ = f.Close() }
			file = f
		}
		b, err := ioutil.ReadAll(file)
		check(err)
		closer()
		text = string(b)
	}

	filename = fmt.Sprintf("qs_%s.txt", time.Now().Format(time.RFC3339))
	if filepath != "" {
		parts := strings.Split(filepath, "/")
		filename = parts[len(parts)-1]
	}

	if !ascii {
		text = filename
	}

	ip := outboundIP()
	port, err := freeport.GetFreePort()
	check(err)

	u := fmt.Sprintf("http://%s:%d?key=%s", ip, port, key)
	curl := fmt.Sprintf("curl -o \"%s\" \"http://%s:%d/download?key=%s\"", filename, ip, port, key)
	if key == "" {
		u = fmt.Sprintf("http://%s:%d", ip, port)
		curl = fmt.Sprintf("curl -o \"%s\" \"http://%s:%d/download\"", filename, ip, port)
	}

	htmltmpl, err := template.New("name").Parse(assets.Template)
	check(err)

	fmt.Print("Avalible ")
	if once {
		fmt.Print("once ")
	}
	fmt.Println("at")
	fmt.Println(" ", u)
	fmt.Println(" ", curl)
	if !noclip {
		if once {
			fmt.Print("and curl ")
			_ = clipboard.WriteAll(curl)
		} else {
			fmt.Print("and url ")
			_ = clipboard.WriteAll(u)
		}
		fmt.Println("is copied to clipboard")
	}
	fmt.Println()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if once {
			oncemux.Lock()
			go func() {
				<-r.Context().Done()
				fmt.Println("terminating after first request")
				os.Exit(0)
			}()
		}

		fmt.Println("Got request from", r.RemoteAddr, r.URL.Path)

		vals := r.URL.Query()
		if vals.Get("key") != key {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("bad key"))
			return
		}

		if r.URL.Path == "/download" {
			fmt.Println("Download", filename)
			w.Header().Set("content-type", "application/octet-stream")
			w.Header().Set("content-disposition", fmt.Sprintf("filename=\"%s\"", filename))

			if filepath == "" {
				_, err = w.Write([]byte(text))
				check(err)
				return
			}

			f, err := os.Open(filepath)
			check(err)

			_, err = io.Copy(w, f)
			check(err)
			return
		}

		if r.URL.Path == "/raw" {
			_, _ = w.Write([]byte(text))
			return

		}

		_ = htmltmpl.Execute(w, map[string]interface{}{
			"text":      text,
			"ascii":     ascii,
			"key":       key,
			"bootstrap": template.CSS(assets.CSSBootstrap),
			"clipboard": template.JS(assets.JSClipboard),
		})
	})

	if ungrok {
		ngrok.Start(filename, key, port)
	}

	check(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}
