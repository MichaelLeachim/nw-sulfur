package main

import (
	"archive/tar"
	"compress/gzip"
	"archive/zip"
	"path/filepath"
	"fmt"
	"io"
	"os"
	"strings"
	"log"
	"path"
)

func extractTar(sourcefile string,extractDir string) {

	if sourcefile == "" {
		log.Panic("No such file")
	}

	file, err := os.Open(sourcefile)

	if err != nil {
		log.Panic(err)
	}

	defer file.Close()

	var fileReader io.ReadCloser = file

	// just in case we are reading a tar.gz file, add a filter to handle gzipped file
	if strings.HasSuffix(sourcefile, ".gz") {
		if fileReader, err = gzip.NewReader(file); err != nil {
			log.Panic(err)
		}
		defer fileReader.Close()
	}

	tarBallReader := tar.NewReader(fileReader)

	// Extracting tarred files

	for {
		header, err := tarBallReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		// get the individual filename and extract to the current directory
		filename := path.Join(extractDir,header.Name)

		switch header.Typeflag {
			case tar.TypeDir:
			// handle directory
			fmt.Println("Creating directory :", filename)
			err = os.MkdirAll(filename, os.FileMode(header.Mode)) // or use 0755 if you prefer

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			case tar.TypeReg:
			// handle normal file
			fmt.Println("Untarring :", filename)
			writer, err := os.Create(filename)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			io.Copy(writer, tarBallReader)

			err = os.Chmod(filename, os.FileMode(header.Mode))

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			writer.Close()
			default:
			fmt.Printf("Unable to untar type : %c in file %s", header.Typeflag, filename)
		}
	}
}


func extractZip(zipfile string,extractDir string) {

	if zipfile == "" {
		log.Panic("no such file")
	}

	reader, err := zip.OpenReader(zipfile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer reader.Close()

	for _, f := range reader.Reader.File {

		zipped, err := f.Open()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		defer zipped.Close()
		// get the individual file name and extract the current directory
		path := filepath.Join(extractDir, f.Name)
		// make dir all
		curDIRpath,_ := filepath.Split(path)
		os.MkdirAll(curDIRpath,0777)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
			fmt.Println("Creating directory", path)
		} else {
			writer, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, f.Mode())

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			defer writer.Close()

			if _, err = io.Copy(writer, zipped); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("Decompressing : ", path)
		}
	}
}


func extract(fpath string,extractDir string){
	log.Println("extracting: ",fpath)
	log.Println("extracting Dir:",extractDir)
	if strings.HasSuffix(fpath,"bz2") {

		extractTar(fpath,extractDir)
		return
	}
	if strings.HasSuffix(fpath,"gz") {
		extractTar(fpath,extractDir)
		return
	}
	if strings.HasSuffix(fpath,"zip"){
		log.Println("Extracting zip",fpath)
		extractZip(fpath,extractDir)
		return
	}
	log.Panic("No such file: ",fpath)
}



