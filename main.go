package main

import (
	"log"

	stdinhandler "github.com/MikelGV/WCTool/StdinHandler"
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

        wcI = stdinhandler.StdInHandler{}

    } else {
        wcI = filehandler.FileHandler{
            FileName: cli.FileName,
        }
    }

    if err := utils.PrintOutput(cli, wcI); err != nil {
        log.Fatal(err)
    }
}
