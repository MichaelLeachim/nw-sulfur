package main
import (
	"path"
//	"strings"
	"os/exec"
	"log"
	"os"
//	"runtime"
//	"github.com/gpmgo/gopm/cmd"
//	"syscall"
)
//import "os"
//import "github.com/bitly/go-simplejson"

func workOutDownloading() {
	log.Println(SULFUR_MD5_DATA_STORAGE,"md5 storage")
	log.Println(SULFUR_PATH,"sulfur path")

	downloadURL := getURLForOS(getUrlsForSystem())
	_,fname     := path.Split(downloadURL)
	savePath    := path.Join(getNWdir(),fname)
	log.Println("Fname of URL: ",fname)

	// nwExecutable exits on fs
	checkAgain:
	if isExistsOnFs(savePath) {
		log.Println("file exists")
		nwFileCheckSum := readMD5toString(savePath)

		if checkMD5(nwFileCheckSum, fname) {
			nwFolder := path.Join(SULFUR_PATH, ArchiveExtStripper(fname))

			if !isExistsOnFs(nwFolder){
				log.Println("No extracted folder. Let us extract: ")
				extract(path.Join(getNWdir(),fname),getNWdir())
				goto checkAgain
			}
			log.Println("md5 matches")
			// this file is md5 checkable
			log.Println(savePath)

			// exec nw based on data
			switch CURRENTOS{
				case "windows":
				    log.Println("Start running: ",path.Join(nwFolder, "nw.exe"), " ",getScriptDir())
    				exec.Command(path.Join(nwFolder, "nw.exe"), getScriptDir()).Start()
				case "linux":
				log.Println("Start running: ",path.Join(nwFolder, "nw")," ", getScriptDir())
    				exec.Command(path.Join(nwFolder, "nw"), getScriptDir()).Start()
				default:
    				log.Panic("OS ", CURRENTOS, " is not supported right now. Sorry")
			}
		}else{
			os.Remove(savePath)
			goto checkAgain

		}
	}else {
		// open downloading ui
		go startServer()
		// no such file. Download it
		downloadURL = downloadFromUrl(downloadURL,getNWdir())

		_,fname    := path.Split(downloadURL)
		savePath = path.Join(getNWdir(),fname)
		// set md5 after downloading it
		setMD5(readMD5toString(savePath),fname)
		// extract archive.
		extract(savePath,getNWdir())
		goto checkAgain
	}
	os.Exit(1)
}

func main(){
//	if CURRENTOS == "windows"{
//		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
//	}
//	runtime.GOMAXPROCS(4)
	workOutDownloading()
//	simpleUi()l
//	startServer()

}







// Right work:
//   Open package.json
//   Check name
//   Check md5
//   If allright
//     open nw
//   else:
//     download file from url. Extract
//     save it md5 to md5_file

//

// Algorithm is as follows:
//  1. Get node url from package json.
//  2. Check if this version exists in folder. <folder>

//  3. If no such folder: create it
//       1. on Windows: ProgramFiles/sulfur-nw
//       2. on Linux:   <cur-dir>/.sulfur-nw
//     else:
//       check for file version in dir: <url,stripped,just name>
//         if file exists:
//           download md5 hashes:
//             check file against md5 hash
//               if error:
//                 delete that file
//
//
//


//        download from source:
//        download md5
//          if error:
//            print error
//
//
//  3.

// getFolder. Check environment for folders(windows) os.Environ
// getCurUserDir (linux)
// osx ?
