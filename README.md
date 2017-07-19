# mplater

CLI tool for filling out a Golang-based template.

## Usage

```
# full command
$ mplater -i ./my_template -o ./output.txt -x myvariable=someValue -x anothervariable=appropriatevalue

# Help message
$ mplater -h
NAME:
   mplater - A template filler-upper

USAGE:
   mplater [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --input value, -i value     Input template file
   --output value, -o value    Output file (default: "output.txt")
   --keyvalue value, -x value  Key=Value pair
   --help, -h                  show help
   --version, -v               print the version
```

For more examples, please refer to the examples folder
