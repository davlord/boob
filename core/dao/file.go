package dao

import (
	"bufio"
	"io"
	"math"
	"os"
)

const BUFFER_SIZE = 1024

type byteRange struct {
	start int64
	end   int64
}

func (br *byteRange) splitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	br.start = br.end
	advance, token, err = bufio.ScanLines(data, atEOF)
	br.end += int64(advance)
	return advance, token, err
}

func (br *byteRange) len() int {
	return int(br.end - br.start)
}

func updateFileInPlace(file *os.File, byteRangeToRemove byteRange, replaceWith []byte) error {
	fileSizeDiff := (len(replaceWith) - byteRangeToRemove.len())
	if fileSizeDiff < 0 {
		return reduceFileInPlace(file, byteRangeToRemove, replaceWith)
	} else {
		return increaseFileInPlace(file, byteRangeToRemove, replaceWith)
	}
}

func reduceFileInPlace(file *os.File, byteRangeToRemove byteRange, replaceWith []byte) error {
	buf := make([]byte, 1024)
	writeAt := byteRangeToRemove.start
	readAt := byteRangeToRemove.end
	for {
		n, err := file.ReadAt(buf, readAt)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := file.WriteAt(buf[:n], writeAt); err != nil {
			return err
		}
		writeAt += int64(n)
		readAt += int64(n)
	}

	file.Truncate(writeAt)

	return nil
}

func increaseFileInPlace(file *os.File, byteRangeToRemove byteRange, replaceWith []byte) error {

	newContentLen := len(replaceWith)

	bufferSize := nextPowerOf2Int(newContentLen)
	if bufferSize < BUFFER_SIZE {
		bufferSize = BUFFER_SIZE
	}

	readBuffer := make([]byte, bufferSize)
	writeBuffer := make([]byte, bufferSize)
	readBufferLen := 0
	writeBufferLen := 0

	copy(writeBuffer, replaceWith)
	writeBufferLen = newContentLen

	writeAt := byteRangeToRemove.start
	readAt := byteRangeToRemove.end

	for {
		var err error

		// file -> read buffer
		readBufferLen, err = file.ReadAt(readBuffer, readAt)
		if err != nil && err != io.EOF {
			return err
		}

		// write buffer -> file
		writeBufferLen, err = file.WriteAt(writeBuffer[:writeBufferLen], writeAt)
		if err != nil {
			return err
		}

		// move read/write cursors forward
		writeAt += int64(writeBufferLen)
		readAt += int64(readBufferLen)

		// swap read/write buffers
		readBuffer, writeBuffer = writeBuffer, readBuffer
		readBufferLen, writeBufferLen = writeBufferLen, readBufferLen

		// exit loop if there is nothing more to write
		if writeBufferLen == 0 {
			break
		}
	}

	file.Truncate(writeAt)

	return nil
}

func nextPowerOf2(num float64) float64 {
	return math.Pow(2, math.Ceil(math.Log2(num)))
}

func nextPowerOf2Int(num int) int {
	return int(nextPowerOf2(float64(num)))
}
