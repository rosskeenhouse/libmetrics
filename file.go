package libmetrics

type FileReader interface {
  ReadFile(filename string) ([]byte, error)
}

// func wrapper
// E.g. ioFileReader(func(filename string) { return ioutil.ReadFile() })
type ioFileReader func(string) ([]byte, error)

func (i ioFileReader) ReadFile(filename string) ([]byte, error) {
  return i(filename)
}
