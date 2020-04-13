package caddytls

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
)

// generateTLSConfig the TLS config for https.
// HTTPUsers are the users that can connect to the server.
// HTTPUsers []string
func generateTLSConfig(certFile, keyFile, caFile string, hTTPUsers []string) (*tls.Config, error) {
	//  Load in cert and private key
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, fmt.Errorf("error parsing X509 certificate/key pair (%s, %s): %v", certFile, keyFile, err)
	}
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		return nil, fmt.Errorf("error parsing certificate (%s): %v", certFile, err)
	}
	// Create our TLS configuration
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}

	if hTTPUsers != nil && len(hTTPUsers) > 0 {
		config.ClientAuth = tls.RequireAndVerifyClientCert
	} else {
		config.ClientAuth = tls.NoClientCert
	}

	// Add in CAs if applicable.
	if caFile != "" {
		rootPEM, err := ioutil.ReadFile(caFile)
		if err != nil || rootPEM == nil {
			return nil, fmt.Errorf("failed to load root ca certificate (%s): %v", caFile, err)
		}
		pool := x509.NewCertPool()
		ok := pool.AppendCertsFromPEM(rootPEM)
		if !ok {
			return nil, fmt.Errorf("failed to parse root ca certificate")
		}
		config.ClientCAs = pool
	}
	return config, nil
}


/*
if m.conf.APISSL {
		go func() {
			pool := x509.NewCertPool()
			caCrt, err := ioutil.ReadFile(m.conf.APICaFile)
			if err != nil {
				logrus.Fatal("ReadFile ca err:", err)
				return
			}
			pool.AppendCertsFromPEM(caCrt)
			s := &http.Server{
				Addr:    m.conf.APIAddrSSL,
				Handler: m.r,
				TLSConfig: &tls.Config{
					ClientCAs:  pool,
					ClientAuth: tls.RequireAndVerifyClientCert,
				},
			}
			logrus.Infof("api listen on (HTTPs) %s", m.conf.APIAddrSSL)
			logrus.Fatal(s.ListenAndServeTLS(m.conf.APICertFile, m.conf.APIKeyFile))
		}()
	}
*/