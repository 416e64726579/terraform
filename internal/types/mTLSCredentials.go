package types

import "crypto/tls"

type TLSConfigProvider interface {
	GetTLSConfig() (*tls.Config, error)
}
