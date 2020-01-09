# QS - Quick Share

```bash
$ go get github.com/crholm/qs

$ qs --help  
NAME:
   qs - Quick creates a local web url to share text from stdin or file

USAGE:
   qs [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --nokey                 no access key is needed to view content  (default: false)
   --noclip                do not copy url to clipboard (default: false)
   --file value, -f value  read data from file
   --help, -h              show help (default: false)

$ echo A string to share | qs
Avalible at http://192.168.1.10:37381?key=Ys1NZfzTTjHGMmaE

$ qs -f file.txt
Avalible at http://192.168.1.10:37381?key=Ys1NZfzTTjHGMmaE

```