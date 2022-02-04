package types

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
)

type TLSConfig struct {
	CertFile      string
	KeyFile       string
	CAFile        string
	ServerAddress string
	Server        bool
}

func SetupTLSConfig(config *TLSConfig) (*tls.Config, error) {
	var err error
	tlsConfig := &tls.Config{}
	if config == nil {
		return nil, errors.New("SetupTLSConfig: nothing to do")
	}
	if config.CertFile != "" && config.KeyFile != "" {
		tlsConfig.Certificates = make([]tls.Certificate, 1)
		tlsConfig.Certificates[0], err = tls.LoadX509KeyPair(config.CertFile, config.KeyFile)
		if err != nil {
			return nil, err
		}
	}
	if config.CAFile != "" {
		caFile, err := ioutil.ReadFile(config.CAFile)
		if err != nil {
			return nil, err
		}
		ca := x509.NewCertPool()
		ok := ca.AppendCertsFromPEM([]byte(caFile))
		if !ok {
			return nil, fmt.Errorf("failed to parse root certificate: %q", config.CAFile)
		}
		if config.Server {
			tlsConfig.ClientCAs = ca
			tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
		} else {
			tlsConfig.RootCAs = ca
		}
		tlsConfig.ServerName = config.ServerAddress
	}
	return tlsConfig, nil
}
