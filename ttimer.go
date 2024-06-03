package main

import (
	"fmt"
	"github.com/drgrib/ttimer/agent"
	"github.com/drgrib/ttimer/parse"
	"os"
	"path/filepath"
	"runtime"
)

//////////////////////////////////////////////
/// flags
//////////////////////////////////////////////

var args struct {
	t string
	q bool
}

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func init() {
	if len(os.Args) > 1 {

		if os.Args[1] == "--install" || os.Args[1] == "--silent" || os.Args[1] == "--silentWithProgress" {
			ex, err := os.Executable()
			if err != nil {
				panic(err)
			}
			exPath := filepath.Dir(ex)
			path := UserHomeDir()
			fmt.Println("Adding ttimer to WindowsApp...")
			err = os.Link(exPath+"\\ttimer.exe", path+"\\AppData\\Local\\Microsoft\\WindowsApps\\ttimer.exe")
			if err != nil {
				fmt.Println("Error adding ttimer to user path. You will have to do it manually :(")
			}
			fmt.Println("Success adding ttimer to path. You can now use it in cmd/powershell via ttimer :)")
			os.Exit(0)
		}
	}

	switch len(os.Args) {
	case 3:

		if os.Args[1] == "-q" {
			args.q = true
			args.t = os.Args[2]
		}
		if os.Args[2] == "-q" {
			args.q = true
			args.t = os.Args[1]
		}
	case 2:
		args.t = os.Args[1]
	default:
		args.t = "10s"
	}
}

//////////////////////////////////////////////
/// main
//////////////////////////////////////////////

func main() {
	// parse
	d, title, err := parse.Args(args.t)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("\nPlease refer to https://github.com/MaxCoteLeKing/ttimer for usage instructions.")
		return
	}

	// start timer
	t := agent.Timer{Title: title}
	t.AutoQuit = true
	t.Start(d)

	// run UI
	t.CountDown()
	os.Exit(0)
}
