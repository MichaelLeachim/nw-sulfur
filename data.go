package main
import "runtime"
import "path"

var PACKAGE_JSON_FNAME string = "package.json"
var INET_ADDR   string   = "127.0.0.1"
var MD5_SUM_URL string = "http://dl.nwjs.io/MD5SUMS"
var CURRENTOS   string = runtime.GOOS     //"linux", "darwin", "windows" "dragonfly", "netbsd", "openbsd", "plan9", "solaris"
var ARCH        string = runtime.GOARCH   //386, amd64, or arm.
var SULFUR_PATH string = getNWdir()
var SULFUR_FOLDER_NAME_LINUX   string = ".nw-sulfur"
var SULFUR_FOLDER_NAME_WINDOWS string = "nw-sulfur"
var SULFUR_FOLDER_NAME_OSX     string = "nw-sulfur"
var SULFUR_MD5_DATA_STORAGE    string = path.Join(getNWdir(),"SULFURMD5.json")
var PERCENT_DOWNLOADED_NW      int64












































