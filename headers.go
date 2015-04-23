package main
import "io"


type packageJson struct {
	Sulfur struct {
		Osx32 string `json:"osx32"`
		Osx64 string `json:"osx64"`
		Win32 string `json:"win32"`
		Win64 string `json:"win64"`
		Lin32 string `json:"lin32"`
		Lin64 string `json:"lin64"`
    } `json:"sulfur-url"`
}

type PassThru struct {
	io.Reader
	transfered int64    // Total # of bytes transferred
	total      int64    // Size Downloaded
	percentage int64  // Percents

}





