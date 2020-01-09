package main

import (
	"crypto/rand"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/phayes/freeport"
	"github.com/urfave/cli/v2"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"qs/assets"
	"strings"
	"time"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func RandStringRunes(n int) string {
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
	var noclip bool
	var filepath string
	var filename string

	app := &cli.App{
		Name:  "qs",
		Usage: "Quick Share creates a local web server in order to share text from stdin or a file",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "nokey",
				Usage: "no access key is needed to view content ",
			},
			&cli.BoolFlag{
				Name:  "noclip",
				Usage: "do not copy url to clipboard",
			},
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "the file specified to be shared",
			},
			&cli.BoolFlag{
				Name:    "ascii",
				Aliases: []string{"a"},
				Usage:   "will treat a file as text input and display it on the web, just as for std in",
			},
		},
		Action: func(c *cli.Context) error {
			nokey = c.Bool("nokey")
			noclip = c.Bool("noclip")
			filepath = c.String("file")
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
		key = RandStringRunes(16)
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


	filename = fmt.Sprintf("ql_%v.txt", time.Now())
	if filepath != "" {
		parts := strings.Split(filepath, "/")
		filename = parts[len(parts)-1]
	}

	if !ascii{
		text = filename
	}

	ip := GetOutboundIP()
	port, err := freeport.GetFreePort()
	check(err)

	u := fmt.Sprintf("http://%s:%d?key=%s", ip, port, key)
	if key == "" {
		u = fmt.Sprintf("http://%s:%d", ip, port)
	}


	htmltmpl, err := template.New("name").Parse(tmpl)
	check(err)


	fmt.Println("Avalible at", u)
	if !noclip {
		fmt.Println(" and copied to clipboard")
		_ = clipboard.WriteAll(u)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

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

			if filepath == ""{
				_, err = w.Write([]byte(text))
				check(err)
				return
			}

			f, err := os.Open(filepath)
			check(err)

			buf := make([]byte, 1024)
			for {
				// read a chunk
				n, err := f.Read(buf)
				if err != nil && err != io.EOF {
					check(err)
				}
				if n == 0 {
					break
				}

				// write a chunk
				if _, err := w.Write(buf[:n]); err != nil {
					check(err)
				}
			}
		}

		if r.URL.Path == "/raw" {
			_, _ = w.Write([]byte(text))
			return

		}

		_ = htmltmpl.Execute(w, map[string]interface{}{
			"text": text,
			"ascii": ascii,
			"key": key,
			"bootstrap": template.CSS(assets.CSSBootstrap),
			"clipboard": template.JS(assets.JSClipboard),
			})
	})
	fmt.Println()
	check(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}

const tmpl = `
<html>
<head>
<style type="text/css">
{{ index . "bootstrap" }}
</style>
</head>
	<h1 style="text-align: center; padding: 30px">Quick Share</h1>
	<div style=" padding-top: 30px">
  		<div style="text-align: center">
			<div class="card" style="display: inline-block; text-align: left; min-width: 152px;">
			  <div class="card-body">
				<div style="position: absolute; top: -30px; left:0; min-width: 153px;">
					{{if index . "ascii"}}
						<button id="copy" style="font-size: 10px;" 
						class="btn btn-sm btn-outline-secondary" 
						data-clipboard-action="copy" 
						data-clipboard-target="#content">
							Copy
						</button>
						<a href="/raw?key={{index . "key"}}" class="btn btn-sm btn-outline-secondary"  style="font-size: 10px;">Raw</a>
					{{end}}
					<a href="/download?key={{index . "key"}}" class="btn btn-sm btn-outline-secondary"  style="font-size: 10px;">Download</a>
				</div>
				
				<pre id="content" style="margin: 0">{{index . "text"}}</pre>
				
			
			  </div>
			</div>
		</div> 
    </div>
</div>
<script>
	{{index . "clipboard"}}
</script>
<script>
	new ClipboardJS('#copy');
</script>

</html>
`
