package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/MikelGV/WCTool/cli"
	interfaces "github.com/MikelGV/WCTool/interface"
)


func IsEmpty() (bool, error) {
    fi, err := os.Stdin.Stat()

    if err != nil {
        return false, err
    }

    return fi.Mode()&os.ModeNamedPipe == 0, nil
}

func isFlagSet(ci *cli.Cli) bool {
    return ci.ByteCount || ci.LineCount || ci.WordCount || ci.CharCount 
}


func PrintOutput(ci *cli.Cli, w interfaces.WcInterface) error {
    
    var output int
    var err error

    if !isFlagSet(ci) {
        val, err := printNArg(ci, w)

        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(val)
        return nil
    } else {
        switch {
            case ci.ByteCount:
                output, err = w.ByteCount()
            case ci.LineCount:
                output, err = w.LineCount()
            case ci.WordCount:
                output, err = w.WordCount()
            case ci.CharCount:
                output, err = w.CharCount()
        }
    } 
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%d \t%v\n", output, w)
    return nil
}

func printNArg(ci *cli.Cli, w interfaces.WcInterface) (string, error) {
    qrtyByte, err := w.ByteCount()

    if err != nil {
        return "", err
    }

    qrtyLine, err := w.LineCount()

    if err != nil {
        return "", err
    }

    return fmt.Sprintf("\t%d\n \t%d\n \t%v\n", qrtyLine, qrtyByte, ci.FileName), nil

}
