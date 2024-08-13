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
    return ci.ByteCount 
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

    return fmt.Sprintf("\t%d\n\t%v\n", qrtyByte, ci.FileName), nil

}
