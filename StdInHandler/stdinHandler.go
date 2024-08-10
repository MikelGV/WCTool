package stdinhandler

import (
	"bufio"
	"log"
	"os"
)


type StdinHandler struct {
    Count int
}


func (st StdinHandler) ByteCount() (int, error) {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanBytes)

    if err := st.scan(scanner); err != nil {
        log.Fatal(err)
    }

    return st.Count, nil
}

func (st StdinHandler) CountLines() (int, error) {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanLines)

    if err := st.scan(scanner); err != nil {
        log.Fatal(err)
    }

    return st.Count, nil
}

func (st *StdinHandler) scan(scanner *bufio.Scanner) error {
    for scanner.Scan() {
        st.Count++
    }

    if err := scanner.Err(); err != nil {
        return err
    }

    return nil
}
