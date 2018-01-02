package main

import (
	"io"
	"os"
)

func copyFile(src, dst string) error {
	if src == "" {
		dir, err := os.Getwd()
		if err != nil { return err }
		src = dir + "/notify"
	}
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}
