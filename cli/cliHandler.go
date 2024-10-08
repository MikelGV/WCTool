package cli

import "flag"

type Cli struct {
    ByteCount bool
    LineCount bool
    WordCount bool
    CharCount bool
    FileName string
}

func New() *Cli {
    byteCount := flag.Bool("c", false, "Gives the byte count")
    lineCount := flag.Bool("l", false, "Gives the line count")
    wordCount := flag.Bool("w", false, "Gives the word count")
    charCount := flag.Bool("m", false, "Gives the characters count")

    flag.Parse()

    return &Cli{
        ByteCount: *byteCount,
        LineCount: *lineCount,
        WordCount: *wordCount,
        CharCount: *charCount,
        FileName: determineFlName(),
    }

}


func determineFlName() string {
    if flag.NArg() > 0 {
        return flag.Arg(0)
    }

    return ""
}
