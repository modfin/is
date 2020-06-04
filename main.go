package main

import (
	"archive/tar"
	"compress/gzip"
	"crypto/rand"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/c2h5oh/datasize"
	"github.com/itsy-sh/is/assets"
	"github.com/itsy-sh/is/ngrok"
	"github.com/phayes/freeport"
	"github.com/urfave/cli/v2"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
	"unicode"
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
func isAsciiPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) && !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

func filename(path string) string {
	_, name := filepath.Split(path)
	return name
}

func main() {
	var stdin bool

	var size string
	var nokey bool
	var once bool
	var ascii bool
	var ungrok bool
	var oncemux sync.Mutex
	var noclip bool
	var filepaths []string
	var filesizes []string

	app := &cli.App{
		Name:  "is",
		Usage: "Itsy Share creates a local web server in order to share text snippets and files on your local network through http",
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
		UsageText: "echo \"foo bar\" | is [global options]\n   is [global options] [filename]",
		Action: func(c *cli.Context) error {

			nokey = c.Bool("nokey")
			noclip = c.Bool("noclip")
			once = c.Bool("once")
			ungrok = c.Bool("ngrok")
			filepaths = c.Args().Slice()

			if len(filepaths) == 0 {
				stdin = true
				ascii = true
			}
			if len(filepaths) == 1{
				ascii = true
			}
			if len(filepaths) > 1{
				ascii = false
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

	if stdin {
		var file io.Reader = os.Stdin
		var closer = func() {}
		b, err := ioutil.ReadAll(file)
		check(err)
		closer()
		f, err := ioutil.TempFile("", "itsy_share_"+time.Now().Format(time.RFC3339))
		check(err)
		defer func() {
			os.Remove(f.Name())
		}()
		err = ioutil.WriteFile(f.Name(), b, 660)
		check(err)

		filepaths = append(filepaths, f.Name())
	}

	tmppaths := filepaths
	filepaths = []string{}
	for i := range tmppaths {
		path := tmppaths[i]
		f, err := os.Open(path)
		check(err)
		i, err := f.Stat()
		check(err)
		if i.IsDir() {
			continue
		}
		filepaths = append(filepaths, path)
		filesizes = append(filesizes, datasize.ByteSize(uint64(i.Size())).HumanReadable())
		check(f.Close())
	}

	if ascii{
		f, err := os.Open(filepaths[0])
		check(err)
		i, err := f.Stat()
		check(err)
		if i.Size() > 40000 { // 40kb
			ascii = false
		}
		check(f.Close())
	}
	if ascii {
		b, err := ioutil.ReadFile(filepaths[0])
		check(err)
		ascii = isAsciiPrintable(string(b))
	}

	ip := outboundIP()
	port, err := freeport.GetFreePort()
	check(err)

	dlname := filename(filepaths[0])
	if len(filepaths) > 1 {
		dlname = "itsy_share_bundle" + time.Now().Format(time.RFC3339) + ".tar.gz"
	}

	u := fmt.Sprintf("http://%s:%d?key=%s", ip, port, key)
	curl := fmt.Sprintf("curl -o \"%s\" \"http://%s:%d/download?key=%s\"", dlname, ip, port, key)
	if key == "" {
		u = fmt.Sprintf("http://%s:%d", ip, port)
		curl = fmt.Sprintf("curl -o \"%s\" \"http://%s:%d/download\"", dlname, ip, port)
	}

	htmltmpl, err := template.New("name").Funcs(template.FuncMap{
		"filename" : filename,
	}).Parse(assets.Template)
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
			fp := filepaths
			name := dlname

			specific := r.URL.Query().Get("file")
			if specific != "" {
				for _, path := range filepaths {
					if path == specific {
						fp = []string{path}
						name = filename(path)
					}
				}
			}

			if len(fp) == 1 {
				path := fp[0]

				fmt.Println("Download", name)
				w.Header().Set("content-type", "application/octet-stream")
				w.Header().Set("content-disposition", fmt.Sprintf("filename=\"%s\"", name))

				f, err := os.Open(path)
				check(err)
				_, err = io.Copy(w, f)
				check(err)
				check(f.Close())
			}

			if len(fp) > 1 {
				fmt.Println("Download", name)
				w.Header().Set("content-type", "application/octet-stream")
				w.Header().Set("content-disposition", fmt.Sprintf("filename=\"%s\"", name))
				zw := gzip.NewWriter(w)
				zw.Name = dlname
				zw.ModTime = time.Now()
				tw := tar.NewWriter(zw)
				for _, path := range fp {
					f, err := os.Open(path)
					check(err)
					s, err := f.Stat()
					check(err)

					hdr := &tar.Header{
						Name: filename(path),
						Mode: 0600,
						Size: s.Size(),
					}
					err = tw.WriteHeader(hdr)
					check(err)

					_, err = io.Copy(tw, f)
					check(err)
					check(f.Close())
				}
				err = tw.Close()
				check(err)
				err = zw.Close()
				check(err)
			}

			return
		}

		var content string
		if ascii {
			b, err := ioutil.ReadFile(filepaths[0])
			check(err)
			content = string(b)
		}

		_ = htmltmpl.Execute(w, map[string]interface{}{
			"filepaths": filepaths,
			"filesizes": filesizes,
			"ascii":     ascii,
			"content": content,
			"key":       key,
			"size":      size,
			"clipboard": template.JS(assets.JSClipboard),
		})
	})

	if ungrok {
		ngrok.Start(dlname, key, port)
	}

	// TODO listen to termination signals...
	check(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))

}
