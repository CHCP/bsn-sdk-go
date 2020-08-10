package http

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"strings"

	"github.com/chcp/bsn-sdk-go/pkg/common/errors"
	"github.com/wonderivan/logger"
)

func SendPost(dataBytes []byte, url string, cert string) ([]byte, error) {

	var client *http.Client

	isHttps := strings.Contains(url, "https://")

	if isHttps {
		logger.Debug("cert:", cert)
		if cert == "" {
			return nil, errors.New("HTTPS certificate not set")
		}

		//dirPath, err := filepath.Abs(".")
		//if err != nil {
		//	logger.Error("get current directory failed：", err.Error())
		//	return nil, err
		//}
		/*
			//read the content of http cert
			caCert, err := ioutil.ReadFile(cert)
			if err != nil {
				logger.Error("read HTTPS certificate content failed：", err.Error())
				return nil, err
			}
		*/
		//20200810: convert cert string to []byte
		caCert := []byte(cert)

		//build a cert pool
		caCertPool := x509.NewCertPool()
		//add the loaded https cert to the cert pool
		caCertPool.AppendCertsFromPEM(caCert)
		//Http request client
		client = &http.Client{
			//define the mechanism for a single Http request
			Transport: &http.Transport{
				//define TLS client configuration
				TLSClientConfig: &tls.Config{
					//add RootCA cert pool（add public key cert of https to RootCA cert pool）
					RootCAs: caCertPool,
				},
			},
		}
	} else {
		logger.Debug("Http")
		tr := new(http.Transport)
		client = &http.Client{
			//define the mechanism for a single HTTP request
			Transport: tr,
		}
	}
	//invoke interface
	logger.Debug("request message：", string(dataBytes))
	response, err := client.Post(url, "application/json", bytes.NewReader(dataBytes))
	if err != nil {
		logger.Error("request failed：", err.Error())
		return nil, err
	}
	//Get the response message data from the response object and read it
	allBytes := []byte{}
	//buffer
	bytes := make([]byte, response.ContentLength)
	i, err := response.Body.Read(bytes)
	allBytes = append(allBytes, bytes[:i]...)

	for {
		i, err = response.Body.Read(bytes)
		if i == 0 {
			break
		}
		allBytes = append(allBytes, bytes[:i]...)
	}
	response.Body.Close()
	logger.Debug("response message：", string(allBytes))
	return allBytes, nil
}
