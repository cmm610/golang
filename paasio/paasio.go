package paasio

import (
	"io"
)
type rCounter struct {
	io.Reader
	count int
}
type wCounter struct {
	io.Writer
	count int
}
type readWriter struct {
	io.Reader
	io.Writer
}

// ReadCount returns the total number of bytes successfully read along
// with the total number of calls to Read().
func (rc rCounter) ReadCount() (n int64, nops int) {
	var bs []byte
	i, err := rc.Read(bs)
	if err != nil {
		return
	}
	rc.count ++
	return int64(i), rc.count
}
func (wc wCounter) WriteCount() (n int64, c int) {
	var bs []byte
	i, err := wc.Write(bs)
	if err != nil {
		return
	}
	wc.count ++
	return int64(i), wc.count
}
// WriteCount returns the total number of bytes successfully read along
// with the total number of calls to Read().
func NewReadCounter(w io.Reader) (rc ReadCounter) {
	//rc = w.
}

// WriteCount returns the total number of bytes successfully written along
// with the total number of calls to Write().
func NewWriteCounter(w io.Writer) (wc WriteCounter) {

}

func NewReadWriteCounter(rw readWriter) (readWriter) {
	rcc := rCounter{
		rw.Reader,
	}
	wrc := wCounter{
		rw.Writer,
	}
	if rcc != nil {
		return rcc
	}
}
