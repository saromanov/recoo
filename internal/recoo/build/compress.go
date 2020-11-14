package build

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// GZip provides compress of the target directory
func GZip(path string) error {
	var buf bytes.Buffer
	if err := compress(path, &buf); err != nil {
		return fmt.Errorf("unable to compress file: %v", err)
	}

	fileToWrite, err := os.OpenFile("./exanple.tar.gzip", os.O_CREATE|os.O_RDWR, os.FileMode(600))
	if err != nil {
		return fmt.Errorf("unable to write to file: %v", err)
	}
	if _, err := io.Copy(fileToWrite, &buf); err != nil {
		return fmt.Errorf("unable to copy to file: %v", err)
	}
	return nil
}

func compress(src string, buf io.Writer) error {
	zr := gzip.NewWriter(buf)
	tw := tar.NewWriter(zr)

	// walk through every file in the folder
	filepath.Walk(src, func(file string, fi os.FileInfo, err error) error {
		// generate tar header
		header, err := tar.FileInfoHeader(fi, file)
		if err != nil {
			return err
		}

		header.Name = filepath.ToSlash(file)

		// write header
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		// if not a dir, write file content
		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return err
			}
			if _, err := io.Copy(tw, data); err != nil {
				return err
			}
		}
		return nil
	})

	// produce tar
	if err := tw.Close(); err != nil {
		return err
	}
	// produce gzip
	if err := zr.Close(); err != nil {
		return err
	}
	//
	return nil
}

// check for path traversal and correct forward slashes
func validRelPath(p string) bool {
	if p == "" || strings.Contains(p, `\`) || strings.HasPrefix(p, "/") || strings.Contains(p, "../") {
		return false
	}
	return true
}
