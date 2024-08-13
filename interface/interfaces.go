package interfaces

type WcInterface interface {
    ByteCount() (int, error)
    LineCount() (int, error)
    WordCount() (int, error)
}
