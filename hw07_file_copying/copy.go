package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fromFile, err := os.Open(fromPath)
	if err != nil {
		return fmt.Errorf("open from file: %w", err)
	}
	defer fromFile.Close()
	fromFileInfo, err := fromFile.Stat()
	if err != nil {
		return fmt.Errorf("from file stat: %w", err)
	}
	if fromFileInfo.Size() == 0 {
		return ErrUnsupportedFile
	}

	if offset > fromFileInfo.Size() {
		return ErrOffsetExceedsFileSize
	}
	toFile, err := os.Create(toPath)
	if err != nil {
		return fmt.Errorf("create to file: %w", err)
	}
	defer toFile.Close()

	if offset != 0 {
		_, err := io.CopyN(io.Discard, fromFile, offset)
		if err != nil {
			return fmt.Errorf("skip offset in from file: %w", err)
		}
	}

	var copySize int64
	if limit == 0 {
		copySize = fromFileInfo.Size() - offset
	} else {
		copySize = min(limit, fromFileInfo.Size()-offset)
	}
	count := 100
	copyChunk := copySize / int64(100)
	progress := 0
	bar := pb.StartNew(count)

	var counter int64
	for counter < copySize {
		var c int64
		progress += 1
		if progress == 100 {
			c, err = io.CopyN(toFile, fromFile, copySize-counter)
		} else {
			c, err = io.CopyN(toFile, fromFile, copyChunk)
		}
		if err != nil {
			if err != io.EOF {
				return fmt.Errorf("copy from file")
			}
			break
		}
		counter += c
		bar.Increment()
	}
	bar.FinishPrint("The End!")

	return nil
}
