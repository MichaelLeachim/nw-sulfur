package main
import (
	"github.com/skratchdot/open-golang/open"
	"fmt"
//	"regexp"
	"encoding/json"
	"io/ioutil"
    "path"
	"os"
	"crypto/md5"
	"log"
	"net/http"
	"strings"
	"io"
//	"regexp"
	"bufio"
//	"go/format"
	"encoding/hex"
	"github.com/kardianos/osext"

)
func openLoader (){
	open.Run(INET_ADDR)
}
func getUserHomeDir () string {
	if CURRENTOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}else{
		return os.Getenv("HOME")
	}
}

func isExistsOnFs (fpath string) bool{
	_,err := os.Stat(fpath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err){
		return false
	}
	return false
}
func getNWdir() string{
	//	//check if dir exists
	//	// return it
	//	// else: make it
	//	//  return it
	//	return "blblabla"
	// determine nw-home
	homeDir := getUserHomeDir()
	var fullPath string
	switch CURRENTOS {
		case "linux":
    		fullPath = path.Join(homeDir, SULFUR_FOLDER_NAME_LINUX)
		case "windows":
            fullPath = path.Join(homeDir, SULFUR_FOLDER_NAME_WINDOWS)
		case "osx":
		    fullPath = path.Join(homeDir, SULFUR_FOLDER_NAME_OSX)
	}
	// if not exists: create one
	if !isExistsOnFs(fullPath) {
		os.MkdirAll(fullPath,0777)
	}
	return fullPath
}

func getURLForOS (data packageJson)string{
	switch {
		case ARCH == "amd64" && CURRENTOS == "windows":
		  return data.Sulfur.Win64
		case ARCH == "386"   && CURRENTOS == "windows":
		  return data.Sulfur.Win32
		case ARCH == "amd64" && CURRENTOS == "linux":
		  return data.Sulfur.Lin64
		case ARCH == "386"   && CURRENTOS == "linux":
		  return data.Sulfur.Lin32
		case ARCH == "amd64" && CURRENTOS == "darwin":
		  return data.Sulfur.Osx64
		case ARCH == "386"   && CURRENTOS == "darwin":
		  return data.Sulfur.Osx32
		default:
		  log.Panic("Architecture: " + ARCH + " is not supported on Platform: ",CURRENTOS)
	}
	return ""
}

//
//}
//func readMD5ofFile(fpath string)string{
//	fileHandler,err := ioutil.ReadFile(fpath)
//	if err != nil {
//		log.Panic(err)
//	}
//	checkSum:= md5.Sum(fileHandler)
//	return hex.EncodeToString(checkSum[:])
//}

func readMD5toString(path string)string{
	result,err := readMD5ofFile(path)
	if err != nil {
		log.Panic(err)
	}
	return hex.EncodeToString(result[:])
}

func readMD5ofFile(filePath string) ([16]byte, error) {
	var result []byte
	var errorResult [16]byte
	file, err := os.Open(filePath)
	if err != nil {
		return errorResult, err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return errorResult, err
	}
	return md5.Sum(result),nil
}

func getUrlsForSystem() packageJson {
	packageBytes,error := ioutil.ReadFile(path.Join(getScriptDir(),PACKAGE_JSON_FNAME))
	var result packageJson
	if error != nil {
		log.Printf("Error:",error)
	}

	parseError := json.Unmarshal(packageBytes,&result)
	if parseError!= nil {
		log.Printf("Error:",error)
	}
	return result
}

func ArchiveExtStripper(filepath string)string{
	if strings.HasSuffix(filepath,".tar.gz"){
	    return strings.TrimSuffix(filepath,".tar.gz")
	}
	if strings.HasSuffix(filepath,".zip"){
		return strings.TrimSuffix(filepath,".zip")
	}
	log.Println("unrecognizeable archive: ",filepath)
	return filepath
}


// Do it later
//func buildMD5OfDir(root string, saveTo string) {
//	result := make(map[string]string)
//	visitFunc := func (curPath string, f os.FileInfo, err error) error {
//		if err != nil {
//			return err
//		}
//		if !f.IsDir() {
//			_,splittedPath := path.Split(curPath)
//			md5ofFile  := readMD5toString(curPath)
//			result[splittedPath] = md5ofFile
//		}
//		return nil
//    }
//	json.Marshal(result,result)
//}


// TESTED
func parseMD5s (md5Path string)map[string]string{
	//  file struct is as follows:
    //	b138c9bbd04540f90113b296fb860121  ./v0.8.6/node-webkit-v0.8.6-win-ia32.zip
    //	0cf5505a3b0d7e20db66c7175c187bd4  ./v0.8.6/node-webkit-v0.8.6-linux-x64.tar.gz
	data := make(map[string]string)
	file, err := os.Open(md5Path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parsedData := strings.Fields(scanner.Text())
//		log.Println(parsedData[1])
		_,fname := path.Split(parsedData[1])
		data[fname] = parsedData[0]
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return data
}

func checkMD5(checksum string,fname string) bool{
	var applicationData = make(map[string]string)
	fileData,err := ioutil.ReadFile(SULFUR_MD5_DATA_STORAGE)
	if err != nil{
		log.Println("File cannot be read in checkMD5",err)
		return false
		// no file. thus, no checksum. thus return false
	}else{
		json.Unmarshal(fileData,&applicationData)
		return applicationData[fname] == checksum
	}
}


func setMD5(checksum string,fname string) {
	var applicationData = make(map[string]string)
	// Read json data
	var fileData []byte
	fileData,err := ioutil.ReadFile(SULFUR_MD5_DATA_STORAGE)
	if err != nil{
		log.Println("File cannot be read in setMD5",err)
	}else{
		json.Unmarshal(fileData,&applicationData)
	}
//	log.Println(applicationData)
	applicationData[fname] = checksum

	byteResult,err := json.Marshal(applicationData)
	if err != nil {
		log.Panic(err)
	}
//	log.Println(byteResult)

	ioutil.WriteFile(SULFUR_MD5_DATA_STORAGE,byteResult,0777)
	return
}


func getScriptDir() string {
	folderPath, err := osext.ExecutableFolder()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(folderPath)
	return folderPath
}

func (pt *PassThru) Read(p []byte) (int, error) {
	n, err := pt.Reader.Read(p)
	pt.transfered += int64(n)

	pt.percentage = int64((float64(pt.transfered)/float64(pt.total))*100)
	PERCENT_DOWNLOADED_NW = pt.percentage

    log.Println("downloading data",pt.transfered,"  ",pt.total)
	log.Println("Downloaded: ",pt.percentage,"%")
	return n, err
}


func downloadFromUrl(furl string,folder string) string {
	_,fname := path.Split(furl)

	fileName := path.Join(folder,fname)
	fmt.Println("Downloading ", furl, " to ", fileName)
    log.Println("GOGOGOGOGOG")
	if isExistsOnFs(fileName){
		err := os.Remove(fileName)
		if err != nil {
			log.Println(err)
		}
	}
	output, err := os.Create(fileName)
	if err != nil {
		log.Println("Error while creating", fileName, "-", err)

	}

	defer output.Close()

	response, err := http.Get(furl)
	if err != nil {
		log.Println("Error while downloading", furl, "-", err)
	}

//	log.Println(response.ContentLength)
	defer response.Body.Close()
	src := &PassThru{Reader:response.Body,total:response.ContentLength}

	n, err := io.Copy(output, src)
	if err != nil {
		log.Println("Error while downloading", furl, "-", err)
	}

	log.Println(n, "bytes downloaded.")
	return fileName
}



