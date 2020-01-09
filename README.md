# QS - Quick Share

```bash
$ go get github.com/crholm/qs

$ qs --help  
NAME:
   qs - Quick Share creates a local web server in order to share text from stdin or a file

USAGE:
   qs [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --nokey                 no access key is needed to view content  (default: false)
   --noclip                do not copy url to clipboard (default: false)
   --file value, -f value  the file specified to be shared
   --ascii, -a             will treat a file as text input and display it on the web, just as for std in (default: false)
   --help, -h              show help (default: false)

$ echo A string to share | qs
Avalible at http://192.168.1.10:37381?key=Ys1NZfzTTjHGMmaE

$ qs -f file.txt
Avalible at http://192.168.1.10:37381?key=Ys1NZfzTTjHGMmaE

```