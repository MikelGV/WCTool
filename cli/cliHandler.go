package cli

import "flag"

type Cli struct {
    ByteCount bool
    FileName string
}

func New() *Cli {
    byteCount := flag.Bool("c", false, "Gives the byte count")

    flag.Parse()

    return &Cli{
        ByteCount: *byteCount,
        FileName: determineFlName(),
    }

}


func determineFlName() string {
    if flag.NArg() > 0 {
        return flag.Arg(0)
    }

    return ""
}
