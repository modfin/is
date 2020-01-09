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

	var nokey bool
	var noclip bool
	var filename string

	app := &cli.App{
		Name:  "qs",
		Usage: "Quick creates a local web url to share text from stdin or file",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "nokey",
				Usage: "no access key is needed to view content ",
			},
			&cli.BoolFlag{
				Name:    "noclip",
				Usage:   "do not copy url to clipboard",
			},
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "read data from file",
			},
		},
		Action: func(c *cli.Context) error {
			nokey = c.Bool("nokey")
			noclip = c.Bool("noclip")
			filename = c.String("file")
			return nil
		},
	}
	err := app.Run(os.Args)
	check(err)

	for _, a := range os.Args{
		if a == "help" || a == "--help" || a == "-h"{
			os.Exit(0)
		}
	}


	var key string
	if !nokey {
		key = RandStringRunes(16)
	}

	var file io.Reader = os.Stdin
	var closer = func() {}
	if len(filename) > 0 {
		f, err := os.Open(filename)
		check(err)
		closer =  func(){_ = f.Close()}
		file = f
	}

	b, err := ioutil.ReadAll(file)
	check(err)
	closer()

	text := string(b)

	ip := GetOutboundIP()
	port, err := freeport.GetFreePort()
	check(err)

	htmltmpl, err := template.New("name").Parse(tmpl)
	check(err)

	u := fmt.Sprintf("http://%s:%d?key=%s\n", ip, port, key)
	if key == "" {
		u = fmt.Sprintf("http://%s:%d\n", ip, port)
	}
	fmt.Println("Avalible at", u)
	if !noclip {
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

		if r.URL.Path == "/raw" {
			_, _ = w.Write([]byte(text))
			return
		}

		_ =htmltmpl.Execute(w, text)
	})
	check(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}

var tmpl = `
<html>
<head>
<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
</head>
	<h1 style="text-align: center; padding: 30px">Quick Share</h1>
	<div style=" padding-top: 30px">
  		<div style="text-align: center">
			<div class="card" style="display: inline-block; text-align: left; min-width: 152px;">
			  <div class="card-body">
				<div style="position: absolute; top: -30px; left:0; min-width: 153px;">
					<button id="copy" style="font-size: 10px;" 
					class="btn btn-sm btn-outline-secondary" 
					data-clipboard-action="copy" 
					data-clipboard-target="#content">
						Copy
					</button>
					<a  class="btn btn-sm btn-outline-secondary" href="javascript:raw()" style="font-size: 10px;">Raw</a>
					<a id="download" class="btn btn-sm btn-outline-secondary"  style="font-size: 10px;">Download</a>
				</div>
				
				<pre id="content" style="margin: 0">{{.}}</pre>
				
			
			  </div>
			</div>
		</div> 
    </div>
</div>
<script src="https://cdnjs.cloudflare.com/ajax/libs/clipboard.js/2.0.4/clipboard.min.js"></script>
<script>
	new ClipboardJS('#copy');

	var fileName = 'qs_'+ new Date().toISOString() +'.txt';
	var fileContent = document.getElementById('content').innerText;
	var myFile = new Blob([fileContent], {type: 'text/plain'});
	window.URL = window.URL || window.webkitURL;
	document.getElementById('download').setAttribute('href', window.URL.createObjectURL(myFile));
	document.getElementById('download').setAttribute('download', fileName);
	
	function raw() {
	  window.location.pathname = "/raw"
	}
</script>

</html>
`
