package tls

import (
	"crypto"
	"crypto/x509"
)

// TlsCertificateRetriever is the interface used by
// both windows cert store implementation, linux and cert from file retriever.
type TlsCertificateRetriever interface {
	GetCertificate() (*x509.Certificate, error)
	GetPrivateKey() (crypto.PrivateKey, error)
}