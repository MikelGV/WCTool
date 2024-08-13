package main

import (
	"fmt"
	"log"

	"github.com/MikelGV/WCTool/cli"
	filehandler "github.com/MikelGV/WCTool/fileHandler"
	interfaces "github.com/MikelGV/WCTool/interface"
	"github.com/MikelGV/WCTool/utils"
)


var wcI interfaces.WcInterface

func main() {
    // Here we handle logic
    cli := cli.New()

    if cli.FileName == "" {
        val, err := utils.IsEmpty()
        
        if err != nil {
            log.Fatal(err)
        }

        if val {
            log.Fatal("there is no filename")
        }

        fmt.Println("nothing to see here")

    } else {
        wcI = filehandler.FileHandler{
            FileName: cli.FileName,
        }
    }


    // I need to access the output somehow
    if err := utils.PrintOutput(cli, wcI); err != nil {
        log.Fatal(err)
    }
}
