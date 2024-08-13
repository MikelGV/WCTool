package cli

import "flag"

type Cli struct {
    ByteCount bool
    LineCount bool
    FileName string
}

func New() *Cli {
    byteCount := flag.Bool("c", false, "Gives the byte count")
    lineCount := flag.Bool("l", false, "Gives the line count")

    flag.Parse()

    return &Cli{
        ByteCount: *byteCount,
        LineCount: *lineCount,
        FileName: determineFlName(),
    }

}


func determineFlName() string {
    if flag.NArg() > 0 {
        return flag.Arg(0)
    }

    return ""
}
