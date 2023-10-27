package tools

import (
	"Day02/internal/entity"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type FindFlags struct {
	sl, d, f bool
	ext      string
}

var fl FindFlags

func init() {
	flag.BoolVar(&fl.sl, "sl", false, "symlinks")
	flag.BoolVar(&fl.d, "d", false, "directories")
	flag.BoolVar(&fl.f, "f", false, "files")
	flag.StringVar(&fl.ext, "ext", "", "extension")
}

func RunFind() {
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		log.Fatal(entity.ArgumentIsEmpty)
	}

	if !fl.f && fl.ext != "" {
		log.Fatal(entity.UncorrectExt)
	}

	if !fl.sl && !fl.d && !fl.f {
		fl.sl, fl.d, fl.f = true, true, true
	}

	for _, arg := range args {
		err := filepath.Walk(arg, walkFunc)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if os.IsPermission(err) {
		return filepath.SkipDir
	} else if err != nil {
		return err
	}

	mode := info.Mode()

	if mode.IsRegular() && fl.ext != "" {
		extension := filepath.Ext(path)
		if extension == ("." + fl.ext) {
			fmt.Println("./" + path)
		}
	} else {

		if fl.sl && mode&fs.ModeSymlink != 0 {
			newPath, _ := filepath.EvalSymlinks(path)
			if _, err := os.Stat(newPath); os.IsNotExist(err) {
				fmt.Println("./"+newPath, "->", "[broken]")
			} else {
				fmt.Println("./"+path, "->", newPath)
			}
		}

		if fl.f && mode.IsRegular() {
			fmt.Println("./" + path)
		}

		if fl.d && mode.IsDir() {
			fmt.Println("./" + path)
		}
	}

	return nil
}
