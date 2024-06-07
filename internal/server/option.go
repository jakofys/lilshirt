package server

import (
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"
)

type Option func(*http.Server) error

func WithTLS(certFile string) Option {
	return func(s *http.Server) error {
		if certFile == "" {
			return ErrCertificatePathEmpty
		}
		cert, err := os.ReadFile(certFile)
		if err != nil {
			return err
		}
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(cert)
		s.TLSConfig = &tls.Config{
			ClientCAs:  caCertPool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		}
		return nil
	}
}
