package zip

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func Zip(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	gw := gzip.NewWriter(d)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	for _, file := range files {
		err := compress(file, "", tw)
		if err != nil {
			return err
		}
	}
	return nil
}

func compress(file *os.File, prefix string, tw *tar.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, tw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := tar.FileInfoHeader(info, "")
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		err = tw.WriteHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(tw, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

//解压 tar.gz
func Unzip(tarFile, dest string) error {
	err := unzipTargz(tarFile, dest)
	if err != nil {
		err = unzipRar(tarFile, dest)
	}

	return err
}

//解压 tar.gz
func unzipTargz(tarFile, dest string) error {
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	fmt.Println(err)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		filename := dest + hdr.Name

		file, err := createFile(filename, os.FileMode(hdr.Mode))
		if err != nil {
			return err
		}
		io.Copy(file, tr)
	}
	return nil
}

func unzipRar(zipFile, dest string) error {
	rootFileName := string([]rune(zipFile)[0:strings.LastIndex(zipFile, ".")])
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()

	isFirstFile := true
	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return err
		}

		if len(file.Name) > 8 {
			s := string([]rune(file.Name)[0:8])
			if s == "__MACOSX" {
				continue
			}
		}
		defer rc.Close()
		if isFirstFile {
			isFirstFile = false
			if rootFileName+"/" == file.Name {
				continue
			}

			dest = dest + string([]rune(zipFile)[0:strings.LastIndex(zipFile, ".")]) + "/"
		}
		filename := dest + file.Name

		if file.FileInfo().IsDir() {
			fmt.Println(filename)
			os.MkdirAll(filename, file.Mode())
			continue
		}

		w, err := createFile(filename, file.Mode())

		if err != nil {
			return err
		}

		_, err = io.Copy(w, rc)
		if err != nil {
			return err
		}
		w.Close()
		rc.Close()
	}
	return nil
}

func getDir(path string) string {
	return subString(path, 0, strings.LastIndex(path, "/"))
}

func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

func createFile(name string, mode os.FileMode) (*os.File, error) {
	err := os.MkdirAll(string([]rune(name)[0:strings.LastIndex(name, "/")]), 0755)
	if err != nil {
		return nil, err
	}

	return os.OpenFile(name, os.O_WRONLY|os.O_CREATE, mode)
}
