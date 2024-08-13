package filehandler

import (
	"bufio"
	"log"
	"os"
)

type FileHandler struct {
    // In here i have to add linesCount and others
    FileName string
    Count int
}

func (fh FileHandler) OpenFile() (*os.File, error) {
    file, err := os.Open(fh.FileName)

    if err != nil {
        return nil, err
    }

    return file, nil
}

func (fh FileHandler) ByteCount() (int, error) {
    file, err := fh.OpenFile() 

   if  err != nil {
       log.Fatal(err)
    }

    defer file.Close()


    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanBytes)

    if err := fh.scan(scanner); err != nil {
        return 0, err
    }

    return fh.Count, nil
}

func (fh FileHandler) LineCount() (int, error) {
    file, err := fh.OpenFile() 

   if  err != nil {
       log.Fatal(err)
    }

    defer file.Close()


    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    if err := fh.scan(scanner); err != nil {
        return 0, err
    }

    return fh.Count, nil
}

func (fh FileHandler) WordCount() (int, error) {
    file, err := fh.OpenFile() 

   if  err != nil {
       log.Fatal(err)
    }

    defer file.Close()


    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)

    if err := fh.scan(scanner); err != nil {
        return 0, err
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

