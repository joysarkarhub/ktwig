package cmd

import (
	"io/ioutil"
	"log"
	"net/http"
)



func FindInternalHostName() string {

	ec2MetaDataUrl := "http://169.254.169.254/latest/meta-data/local-hostname"

	req, err := http.NewRequest("GET", ec2MetaDataUrl, nil)

	if err != nil {
		log.Fatal("Not able to access EC2 Metadata", err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal("Issue with the response", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal("Issue to Read the Response Body", err)
	}

	return string(body)

}
func FindInternalIp() string {

	ec2MetaDataUrl := "http://169.254.169.254/latest/meta-data/local-ipv4"

	req, err := http.NewRequest("GET", ec2MetaDataUrl, nil)

	if err != nil {
		log.Fatal("Not able to access EC2 Metadata", err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal("Issue with the response", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal("Issue to Read the Response Body", err)
	}

	return string(body)

}
