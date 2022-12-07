package main

import (
	"bytes"
	"fmt"
	"os"
)

const input = "./6/input.txt"
const bufferLen = 4

func newBuffer(data []byte) *charBuffer {
	b := &charBuffer{
		data:     data,
		readHead: 0,
		buffer:   [bufferLen]byte{},
	}
	if len(data) < bufferLen {
		panic("input string shorter than 4")
	}
	b.newbuf()
	return b
}

type charBuffer struct {
	data     []byte
	buffer   [bufferLen]byte
	readHead int
}

func (b *charBuffer) display() string {
	return string(b.buffer[:])
}

func (b *charBuffer) newbuf() {
	b.buffer = [4]byte{}
	for i := b.readHead; i < b.readHead+bufferLen; i++ {
		b.buffer[i-b.readHead] = b.data[i]
	}
}
func (b *charBuffer) next() error {
	if b.readHead > len(b.data)-bufferLen-1 {
		return fmt.Errorf("end of buffer")
	}
	b.readHead++
	b.newbuf()
	return nil
}

func (b *charBuffer) isUnique() (int, bool) {
	unique := true
	for i := 0; i < bufferLen; i++ {
		for j := i + 1; j < bufferLen; j++ {
			if b.buffer[i] == b.buffer[j] {
				unique = false
			}
		}
	}
	return b.readHead + bufferLen, unique
}

func main() {
	// Open Input file
	data, err := os.ReadFile(input)
	if err != nil {
		panic("Could not open file: " + err.Error())
	}
	inputString := bytes.NewBuffer(data).Bytes()
	b := newBuffer(inputString)
	for {
		err := b.next()
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		if i, ok := b.isUnique(); ok {
			fmt.Printf("first unique set of %d bytes :: %#v\n", bufferLen, b.display())
			fmt.Printf("%d bytes processed\n", i)
			break
		}
	}

}

/*
NOTES
your subroutine needs to identify the first position where the four most recently received characters were all different. Specifically, it needs to report the number of characters from the beginning of the buffer to the end of the first such four-character marker.
*/
