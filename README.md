# QS - Quick Share

```bash
$ go get github.com/crholm/qs

$ qs --help  
NAME:
   qs - Quick Share creates a local web server in order to share text from stdin or a file

USAGE:
   echo "foo bar" | qs [global options]
   qs [global options] [filename]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --nokey      no access key is needed to view content  (default: false)
   --noclip     do not copy url to clipboard (default: false)
   --ascii, -a  will treat a file as text input and display it on the web, just as for std in (default: false)
   --once, -o   terminates the server after first page load, so things dont hang around (copies curl command to clipboard) (default: false)
   --help, -h   show help (default: false
```


**Usage**
```bash
$ echo A string to share | qs
Avalible at http://192.168.1.10:37381?key=Ys1NZfzTTjHGMmaE

$ qs file.txt
Avalible at http://192.168.1.10:37381?key=Ys1NZfzTTjHGMmaE

```