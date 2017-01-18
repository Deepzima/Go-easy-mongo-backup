package main
import (
	"fmt"
	"os/exec"
  //"encoding/binary"
)

func main() {
	path := "/var/www/"
	wcArgs := []string{
    "find",
    "db_backup/",
    "-type", "f", "|", "wc", "-l",
  }
	cmdBackup := "gesu.sh"
  //var fullCmd string
	var (
		cmdOut []byte
		cmdErr error
	)
  tmpCount := fileCounter(wcArgs, path)
  fmt.Printf("%d", tmpCount);
  fullCmd := getPath(path, cmdOut, cmdBackup)

	cmdOutput := exec.Command(fullCmd)
	cmdOut, cmdErr = cmdOutput.Output()
  if cmdErr != nil {
    fmt.Printf("error ", cmdErr)
    failWith(cmdErr)
  } else {
      fmt.Printf("%s", cmdOut)
  }
}

func failWith(e error) {
	if e != nil {
		fmt.Printf("can't get stout: ", e)
		checkErr(e)
	}
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func getPath(s string, b []byte, c string) string {
  	if string(b[:]) == s {
  		return "./" + c
  	} else {
  		return s + c
  	}
}

func fileCounter(args []string, s string) int {
  /*var (
		cmdOut []byte
		cmdErr error
	)*/
  fmt.Printf("%v", args...)
  /*cmdOut,cmdErr = exec.Command(args[0], s, args[0:]...)
  if cmdErr != nil {
    failWith(cmdErr)
  } else {
    return binary.ReadVarInt(cmdOut)
  }*/

  return 1
}
