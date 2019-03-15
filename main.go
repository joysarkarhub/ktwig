package main

import (
	"github.com/joysarkarhub/ktwig/cmd"

	"os"
	"log"

)

func main () {

	hostname := cmd.FindInternalHostName()
	hostip := cmd.FindInternalIp()

	hostnameip := hostip + " " + hostname

	f, err := os.OpenFile("/etc/hosts", os.O_APPEND|os.O_WRONLY, 6000)
	if err != nil {
		log.Fatal("Issue to open file", err)
	}

	defer f.Close()

	if _, err = f.WriteString(hostnameip); err != nil {
		panic(err)
	}

	cmd.InstallPkg()

}