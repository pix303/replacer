package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func execChangeExtension(rootDir, from, to string) {
	if !strings.HasPrefix(from, ".") {
		from = "." + from
	}

	if !strings.HasPrefix(to, ".") {
		to = "." + to
	}

	err := filepath.Walk(rootDir, func(filename string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(info.Name()) == from {
			src := filename
			dst := strings.TrimSuffix(src, from)
			dst += to
			if err := os.Rename(src, dst); err != nil {
				fmt.Println(err)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("error walking on ", rootDir)
	}
}

func execChangeContains(rootDir, from, to string) {
	err := filepath.Walk(rootDir, func(filename string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.Contains(filepath.Base(info.Name()), from) {
			src := filename
			dst := strings.ReplaceAll(src, from, to)
			if err := os.Rename(src, dst); err != nil {
				fmt.Println(err)
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("error walking on ", rootDir)
	}
}

func execSnakeCase(rootDir string) {
	err := filepath.Walk(rootDir, func(filename string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		fInfo, err := os.Stat(rootDir)
		if err != nil {
			return err
		}

		basePath := rootDir
		if !fInfo.IsDir() {
			basePath = filepath.Dir(rootDir)
		}

		newName := ""
		for _, v := range info.Name() {
			if !unicode.IsUpper(v) {
				newName += string(v)
			} else {
				newName += "_" + string(unicode.ToLower(v))
			}
		}

		err = os.Rename(filename, basePath+string(os.PathSeparator)+newName)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Println("error walking on ", rootDir)
	}
}

func execCamelCase(rootDir string) error {
	err := filepath.Walk(rootDir, func(filename string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fInfo, err := os.Stat(rootDir)
		if err != nil {
			return err
		}

		basePath := rootDir
		if !fInfo.IsDir() {
			basePath = filepath.Dir(rootDir)
		}

		newName := ""
		nextUpper := false
		for _, v := range info.Name() {
			if v == '_' {
				nextUpper = true
				continue
			}

			if nextUpper {
				newName += string(unicode.ToUpper(v))
				nextUpper = false
			} else {
				newName += string(v)
			}
		}

		err = os.Rename(filename, basePath+string(os.PathSeparator)+newName)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
