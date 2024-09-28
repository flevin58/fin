package tools

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// ZIP

func ZipCompress(source, target string) error {

	// Create or truncate the target file
	out, err := os.Create(target)
	if err != nil {
		return err
	}
	defer out.Close()

	// Create the zip writer
	zw := zip.NewWriter(out)
	defer zw.Close()

	// Loop through all the files and folders of the source folder
	return filepath.Walk(source,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Create the ZIP header
			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}

			// Set the default compression
			header.Method = zip.Deflate

			// Set the relative path of the file as the header name
			header.Name, err = filepath.Rel(filepath.Dir(source), path)
			if err != nil {
				return err
			}
			if info.IsDir() {
				header.Name += "/"
			}

			// Create the header writer (hw) for the file header
			hw, err := zw.CreateHeader(header)
			if err != nil {
				return err
			}

			// Don't process a folder
			if info.IsDir() {
				return nil
			}

			// Write file contents to zip archive
			in, err := os.Open(path)
			if err != nil {
				return err
			}
			defer in.Close()

			_, err = io.Copy(hw, in)
			return err
		})
}

func ZipList(source string) error {

	// Get a zip reader (zr) from the source zip archive
	zr, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer zr.Close()

	// Loop through all the entries
	for _, f := range zr.File {
		fmt.Println(f.Name)
	}
	return nil
}

func ZipExtract(source, target string) error {
	// Get a zip reader (zr) from the source zip archive
	zr, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer zr.Close()

	// Loop through all the entries
	for _, f := range zr.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		path := filepath.Join(target, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
