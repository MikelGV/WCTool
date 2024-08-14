package stdinhandler

import (
	"bufio"
	"os"
)

type StdInHandler struct {
    Count int
}

func (sh StdInHandler) ByteCount() (int, error) {


    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanBytes)

    if err := sh.scan(scanner); err != nil {
        return 0, err
    }

    return sh.Count, nil
}

func (sh StdInHandler) LineCount() (int, error) {

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanLines)

    if err := sh.scan(scanner); err != nil {
        return 0, err
    }

    return sh.Count, nil
}

func (sh StdInHandler) WordCount() (int, error) {
    
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)

    if err := sh.scan(scanner); err != nil {
        return 0, err
    }

    return sh.Count, nil
}

func (sh StdInHandler) CharCount() (int, error) {

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanRunes)

    if err := sh.scan(scanner); err != nil {
        return 0, err
    }

    return sh.Count, nil
}

func (sh *StdInHandler) scan(scanner *bufio.Scanner) error {
    for scanner.Scan() {
        sh.Count++
    }

    if err := scanner.Err(); err != nil {
        return err
    }

    return nil
}

