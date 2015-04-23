
package main
import (
	"testing"
	"log"
	"os"
)
//
//func  TestParseMD5s(t *testing.T) {
//
//
//
//
//}

func TestGetUrlsForSystem(t *testing.T){
//	data := getUrlsForSystem()
//
//	if data.Sulfur.Lin32 != "http://dl.nwjs.io/v0.11.0/node-webkit-v0.11.0-linux-ia32.tar.gz" {
//		t.Error("Expected","http://dl.nwjs.io/v0.11.0/node-webkit-v0.11.0-linux-ia32.tar.gz","got",data.Sulfur.Lin32)
//	}
}





func TestParseMD5s(t *testing.T){

	result := parseMD5s("./test_data/MD5SUMS")
	if result["node-webkit-v0.8.3-linux-ia32.tar.gz"] != "b76890570306aec25697f445d594eb18" {
		log.Println(result)
		t.Error("got checksum: ",result["node-webkit-v0.8.3-linux-ia32.tar.gz"])

	}
}

func TestDownloadFromUrl(t *testing.T){
	downloadFromUrl("http://dl.nwjs.io/MD5SUMS","./test_data/ho/")
	if !isExistsOnFs("./test_data/ho/MD5SUMS"){
		t.Error("download file error")
	}
	os.Remove("./test_data/ho/MD5SUMS")
}


func TestTrimSuffix(t *testing.T){
	if ArchiveExtStripper("bla.bla/gogo/hohol.tar.gz") != "bla.bla/gogo/hohol" {
		t.Error("archive .tar.gz stripped in: ",ArchiveExtStripper("bla.bla/gogo/hohol.tar.gz"))
	}
	if ArchiveExtStripper("bla.bla/gogo/hohol.zip") != "bla.bla/gogo/hohol" {
		t.Error("archive .zip stripped in: ",ArchiveExtStripper("bla.bla/gogo/hohol.zip"))
	}
}




func TestGetSetMd5(t *testing.T){
	setMD5("checksum","filename.doc")
	setMD5("checksum2","filename2.doc")
	setMD5("checksum3","filename3.doc")
	if checkMD5("checksum","filename.doc") != true {
		t.Error("checksum checking not working")
	}
	if checkMD5("checksum","gogogog") != false {
		t.Error("not existant checksum checking not working")
	}
	os.Remove(SULFUR_MD5_DATA_STORAGE)
}




func TestReadMD5ofFile(t *testing.T){
	// Go fuck these packers. We do not use md5 checker(as md5 on site is something custom
//	md5hash     := readMD5toString("./test_data/MD5SUMS")
//	md5RealHash := parseMD5s("./test_data/MD5SUMS")["MD5SUMS"]
//	if md5hash != md5RealHash{
//		t.Error(md5hash,"!=",md5RealHash)
//	}

}



