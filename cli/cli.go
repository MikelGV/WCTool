package cli

import "flag"

type Cli struct {
    CheckByte bool
    FileName string
    // here goes all the other checks
}


func New() *Cli {
    checkBytes := flag.Bool("c", false, "checks for the file byte count")


    flag.Parse()

    return &Cli{
        CheckByte: *checkBytes,
        FileName: determineFile(),
    }
}


func determineFile() string {
    if flag.NArg() > 0 {
        return flag.Arg(0)
    }
    return ""
}
