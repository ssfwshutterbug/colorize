package main

import (
	"colorize/colorstr"
	"flag"
	"fmt"
	"os"
)

func main() {
	fg, bg, contents, err := getArgs()
	if err != nil {
		fmt.Println(err)
	}
	pair := colorstr.ColorPair{
		Fg: fg,
		Bg: bg,
	}
	text := colorstr.RenderText(&pair, contents)
	fmt.Println(text)
}

func getArgs() (string, string, string, error) {

	flag.BoolFunc("help", "show help message", func(string) error {
		help()
		os.Exit(0)
		return nil
	})

	fgcolor := flag.String("fg", "nil", "foreground color")
	bgcolor := flag.String("bg", "nil", "background color")
	flag.Parse()

	content := flag.Args()
	var contents string
	for _, j := range content {
		contents = contents + " " + j + " "
	}

	var err error
	if contents == "" || (*fgcolor == "nil" && *bgcolor == "nil") {
		err = fmt.Errorf("error: no color or text")
		return "nil", "nil", "nil", err
	}

	return *fgcolor, *bgcolor, contents, nil
}

func help() {
	fmt.Println(`colorize [-fg <color>] [-bg <color>] text
<color> can be rgb color code or predefined colorname`)

	fmt.Printf("\n available colorname:\n")

	for key := range colorstr.Color {
		fmt.Printf("          %s\n", key)
	}
}
