package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const filemode = os.FileMode(0700)

var files = []string{
	"D877F783D5D3EF8C",
	"D877F783D5D3EF8C0",
	"settings1",
}

func exitWithError(msg string) {
	fmt.Println(msg)
	os.Exit(0)
}

func check(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func dirExist(path string) error {
	if s, err := os.Stat(path); err != nil || !s.IsDir() {
		if err == nil {
			return errors.New(path + " is regular file! Folder expected!")
		}

		return err
	}

	return nil
}

func clearDir(dirpath string) error {
	dir, err := os.Open(dirpath)
	if err != nil {
		return err
	}

	defer dir.Close()

	files, err := dir.Readdirnames(0)
	if err != nil {
		return err
	}

	for _, f := range files {
		err = os.RemoveAll(filepath.Join(dirpath, f))
		if err != nil {
			return err
		}
	}

	return nil
}

func copy(dest, src string) error {
	srcStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}

	defer srcFile.Close()

	if srcStat.IsDir() {
		files, err := srcFile.Readdir(0)
		if err != nil {
			return err
		}

		err = os.Mkdir(dest, filemode)
		if err != nil {
			return err
		}

		for _, f := range files {
			n := f.Name()
			err = copy(dest+"/"+n, src+"/"+n)
			if err != nil {
				return err
			}
		}

		return nil
	}

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destFile.Close()

	if _, err = io.Copy(destFile, srcFile); err != nil {
		return err
	}

	return nil
}

func neededFilesExist(dir string) (bool, error) {
	for _, f := range files {
		if _, err := os.Stat(filepath.Join(dir, f)); err != nil {
			if os.IsNotExist(err) {
				return false, nil
			}

			return false, err
		}
	}

	return true, nil
}

func copyRequiredFiles(dest, src string) error {
	for _, f := range files {
		destPath := filepath.Join(dest, f)
		srcPath := filepath.Join(src, f)

		if err := copy(destPath, srcPath); err != nil {
			return err
		}
	}

	return nil
}
