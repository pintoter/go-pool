package tools

import (
	"Day02/internal/entity"
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var a bool

func init() {
	flag.BoolVar(&a, "a", false, "output directory all files from path")
}

func RunRotate() {
	flag.Parse()

	var files []string
	var directory string = "."

	if a {
		directory = os.Args[2]
		files = os.Args[3:]
	} else {
		files = os.Args[1:]
	}

	if len(files) == 0 {
		log.Fatal(entity.EmptyFiles)
	}

	if _, err := isCorrectDir(directory); err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go archive(directory, file, &wg)
	}
	wg.Wait()
}

func isCorrectDir(dir string) (bool, error) {
	fi, err := os.Stat(dir)
	if err != nil && (os.IsNotExist(err) || os.IsPermission(err)) {
		return false, err
	}

	if !fi.IsDir() {
		return false, entity.PathIsNotDirectory
	}

	return true, nil
}

func archive(directory, filename string, wg *sync.WaitGroup) error {
	defer wg.Done()

	newFilename := getNewFilename(directory, filename)

	zippedFile, err := os.Create(newFilename)
	if err != nil {
		return err
	}
	defer zippedFile.Close()

	err = writeCompressedFile(zippedFile, filename)
	if err != nil {
		return err
	}

	return nil
}

func getNewFilename(directory, filename string) string {
	dir := strings.TrimSuffix(directory, "/")
	name := strings.Split(filename, filepath.Ext(filename))
	name = strings.Split(name[0], "/")

	return fmt.Sprintf("%s/%s_%d.tar.gz", dir, name[len(name)-1], time.Now().Unix())
}

func writeCompressedFile(f *os.File, filename string) error {
	gw := gzip.NewWriter(f)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := os.Stat(filename)
	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}

	header.Name = filename
	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(tw, file)
	if err != nil {
		return err
	}

	return nil
}
