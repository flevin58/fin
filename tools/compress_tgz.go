package tools

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
)

// Compress the source folder to the target .tgz archive file
func TgzCompress(source, target string) error {

	// Create or truncate the target file
	out, err := os.Create(target)
	if err != nil {
		return err
	}
	defer out.Close()

	// Create the gzip and tar writers
	gw := gzip.NewWriter(out)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	// Loop through all the files and folders of the source folder
	return filepath.Walk(source,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Create the TAR header
			header, err := tar.FileInfoHeader(info, info.Name())
			if err != nil {
				return err
			}

			// Set the relative path of the file as the header name
			header.Name, err = filepath.Rel(filepath.Dir(source), path)
			if err != nil {
				return err
			}

			if info.IsDir() {
				header.Name += "/"
			}

			// Write the header
			err = tw.WriteHeader(header)
			if err != nil {
				return err
			}

			// Write the file if not a folder
			if !info.IsDir() {
				// Open the source file
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()

				// Copy its contents to the TAR archive
				_, err = io.Copy(tw, file)
				if err != nil {
					return err
				}
			}
			return nil
		})
}

func TgzList(source string) error {
	source = NormalizePath(source, "")
	f, err := os.Open(source)
	if err != nil {
		return err
	}
	defer f.Close()
	gr, err := gzip.NewReader(f)
	if err != nil {
		return err
	}

	tr := tar.NewReader(gr)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		// Process header
		switch header.Typeflag {
		case tar.TypeDir:
			fmt.Println(header.Name)
		case tar.TypeReg:
			fmt.Println(header.Name)
		default:
			return fmt.Errorf(
				"bad tar header: uknown type: %v in %s",
				header.Typeflag,
				header.Name)
		}
	}
	return nil
}

func TgzExtract(source, target string) error {

	source = NormalizePath(source, "")
	f, err := os.Open(source)
	if err != nil {
		return err
	}
	defer f.Close()

	gr, err := gzip.NewReader(f)
	if err != nil {
		return err
	}

	tr := tar.NewReader(gr)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		// If we define a target dir, make all paths relative to it
		if target != "" {
			header.Name = path.Join(target, header.Name)
		}

		// Create folder / extract file
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(header.Name, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			outFile, err := os.Create(header.Name)
			if err != nil {
				return err
			}
			defer outFile.Close()

			if _, err := io.Copy(outFile, tr); err != nil {
				return err
			}

		default:
			return fmt.Errorf(
				"bad tar header: uknown type: %v in %s",
				header.Typeflag,
				header.Name)
		}
	}
	return nil
}
