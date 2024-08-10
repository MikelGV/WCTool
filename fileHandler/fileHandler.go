package filehandler 

import (
	"bufio"
	"fmt"
	"os"
)


type FileHandler struct {
    FileName string
    Count int
}

// TODO:: Build CountLines function


func (fh *FileHandler) OpenFile() (*os.File, error) {
    f, err := os.Open(fh.FileName)
    if err != nil {
        return nil, fmt.Errorf("Couldn't open File: %w", err)
    }

    return f, nil
}


func (fh FileHandler) ByteCount() (int, error) {
    f, err := fh.OpenFile()

    if err != nil {
        return 0, err
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanBytes)

    if err := fh.scan(scanner); err != nil {
        return 0, nil
    }

    return fh.Count, nil
}

func (fh FileHandler) CountLines() (int, error) {
    f, err := fh.OpenFile()

    if err != nil {
        return 0, err
    }

    defer f.Close()

    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanLines)

    if err := fh.scan(scanner); err != nil {
        return 0, nil
    }


    return fh.Count, nil
}

func (fh *FileHandler) scan(scanner *bufio.Scanner) error {
    for scanner.Scan() {
        fh.Count++
    }

    if err := scanner.Err(); err != nil {
        return err
    }

    return nil
}
