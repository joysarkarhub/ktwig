package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"net/http"
	"log"
	"io"
)

func InstallPkg() {
	cmdName := "apt-get"
	installsupportPkg := []string{
		"install",
		"-y",
		"apt-transport-https",
		"ca-certificates",
		"curl",
		"gnupg-agent",
		"software-properties-common",
	}

	if _, err := exec.Command(cmdName, installsupportPkg...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "Unbale to download required packages:", err)
		os.Exit(1)
	}

	gpgkeyurl := "https://download.docker.com/linux/ubuntu/gpg"
	
	gpgresp, err := http.Get(gpgkeyurl)
	
	if err != nil {
		log.Fatal("Not able to download GPG Key", err)
	}

	defer gpgresp.Body.Close()

	gpgfilename, err := os.Create("gpg")
	if err != nil {
		log.Fatal("Not able to create file", err)
	}

	defer gpgfilename.Close()

	_, err = io.Copy(gpgfilename, gpgresp.Body)
	
	if err != nil {
		log.Fatal("Not able to create file", err)
	}

	keyaddcmd := "apt-key"
	aptkeyaddarg := []string{
		"add",
		"gpg",
	}

	if _, err := exec.Command(keyaddcmd, aptkeyaddarg...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "Unable to add GPG key:", err)
		os.Exit(1)
	}

	repoaddcmd := "add-apt-repository"
	repoaddcmdarg := []string {
		"deb",
		"[arch=amd64]",
		"https://download.docker.com/linux/ubuntu",
		"$(lsb_release -cs)",
		"stable",
	}

	if _, err := exec.Command(repoaddcmd, repoaddcmdarg...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "Unable to add Docker repo:", err)
		os.Exit(1)

	}

}