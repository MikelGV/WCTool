package main

import (
	"errors"
	"flag"
	"os"
)


func main() {
    // Need to take a file from cli input. ej wc foo bar
    fs := flag.NewFlagSet("cFunc", flag.ContinueOnError)
    fs.SetOutput(os.Stdout)

    fs.Func("c", "`txt file` to parse", func(s string) error {
        cnt, err := os.ReadFile(s)
        if err != nil {
            return errors.New("could not parse file")
        }
        return nil 
    })

    flag.Parse()
}
