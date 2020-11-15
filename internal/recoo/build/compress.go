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

// comperss provides compress of the target directory
func compress(path, archivePath string) error {
	var buf bytes.Buffer
	if err := archive(path, &buf); err != nil {
		return fmt.Errorf("unable to archive file: %v", err)
	}

	fileToWrite, err := os.OpenFile(fmt.Sprintf("%s.tar.gzip", archivePath), os.O_CREATE|os.O_RDWR, os.FileMode(777))
	if err != nil {
		return fmt.Errorf("unable to write to file: %v", err)
	}
	if _, err := io.Copy(fileToWrite, &buf); err != nil {
		return fmt.Errorf("unable to copy to file: %v", err)
	}
	return nil
}

func archive(src string, buf io.Writer) error {
	zr := gzip.NewWriter(buf)
	tw := tar.NewWriter(zr)

	if err := filepath.Walk(src, func(file string, fi os.FileInfo, err error) error {
		header, err := tar.FileInfoHeader(fi, file)
		if err != nil {
			return err
		}

		header.Name = filepath.ToSlash(file)

		if err := tw.WriteHeader(header); err != nil {
			return err
		}
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
	}); err != nil {
		return err
	}

	// produce tar
	if err := tw.Close(); err != nil {
		return fmt.Errorf("unable to produce tar: %v", err)
	}
	// produce gzip
	if err := zr.Close(); err != nil {
		return fmt.Errorf("unable to produce gzip: %v", err)
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
