package oidc

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"net/url"
	"testing"
	"time"

	"github.com/hashicorp/boundary/internal/db"
	wrapping "github.com/hashicorp/go-kms-wrapping"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
)

// TestAuthMethod creates an oidc auth method
func TestAuthMethod(
	t *testing.T,
	conn *gorm.DB,
	databaseWrapper wrapping.Wrapper,
	scopeId string,
	state AuthMethodState,
	discoveryUrl *url.URL,
	clientId string,
	clientSecret ClientSecret,
	opt ...Option) *AuthMethod {
	t.Helper()
	opts := getOpts(opt...)
	require := require.New(t)
	rw := db.New(conn)
	ctx := context.Background()

	authMethod, err := NewAuthMethod(scopeId, discoveryUrl, clientId, clientSecret, opt...)
	require.NoError(err)
	id, err := newAuthMethodId()
	require.NoError(err)
	authMethod.PublicId = id
	err = authMethod.encrypt(ctx, databaseWrapper)
	require.NoError(err)
	err = rw.Create(ctx, authMethod)
	require.NoError(err)

	if len(opts.withCallbackUrls) > 0 {
		newCallbacks := make([]interface{}, 0, len(opts.withCallbackUrls))
		for _, c := range opts.withCallbackUrls {
			callback, err := NewCallbackUrl(authMethod.PublicId, c)
			require.NoError(err)
			newCallbacks = append(newCallbacks, callback)
		}
		err := rw.CreateItems(ctx, newCallbacks)
		require.NoError(err)
	}
	if len(opts.withAudClaims) > 0 {
		newAudClaims := make([]interface{}, 0, len(opts.withAudClaims))
		for _, a := range opts.withAudClaims {
			aud, err := NewAudClaim(authMethod.PublicId, a)
			require.NoError(err)
			newAudClaims = append(newAudClaims, aud)
		}
		err := rw.CreateItems(ctx, newAudClaims)
		require.NoError(err)
	}
	if len(opts.withCertificates) > 0 {
		newCerts := make([]interface{}, 0, len(opts.withCertificates))
		for _, c := range opts.withCertificates {
			pem := TestEncodeCertificates(t, c)
			cert, err := NewCertificate(authMethod.PublicId, pem[0])
			require.NoError(err)
			newCerts = append(newCerts, cert)
		}
		err := rw.CreateItems(ctx, newCerts)
		require.NoError(err)
	}
	return authMethod
}

// TestEncodeCertificates will encode a number of x509 certificates to PEMs.
func TestEncodeCertificates(t *testing.T, certs ...*x509.Certificate) []string {
	t.Helper()
	require := require.New(t)
	require.NotEmpty(certs)

	var pems []string
	for _, cert := range certs {
		var buffer bytes.Buffer
		require.NotNil(cert)
		err := pem.Encode(&buffer, &pem.Block{
			Type:  "CERTIFICATE",
			Bytes: cert.Raw,
		})
		require.NoError(err)
		pems = append(pems, buffer.String())
	}
	return pems
}

// TestConvertToUrls will convert URL string representations to a slice of
// *url.URL
func TestConvertToUrls(t *testing.T, urls ...string) []*url.URL {
	t.Helper()
	require := require.New(t)
	require.NotEmpty(urls)
	var convertedUrls []*url.URL
	for _, u := range urls {
		parsed, err := url.Parse(u)
		require.NoError(err)
		require.Contains([]string{"http", "https"}, parsed.Scheme)
		convertedUrls = append(convertedUrls, parsed)
	}
	return convertedUrls
}

// testGenerateCA will generate a test x509 CA cert, along with it encoded in a
// PEM format.
func testGenerateCA(t *testing.T, hosts ...string) (*x509.Certificate, string) {
	t.Helper()
	require := require.New(t)

	priv, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	require.NoError(err)

	// ECDSA, ED25519 and RSA subject keys should have the DigitalSignature
	// KeyUsage bits set in the x509.Certificate template
	keyUsage := x509.KeyUsageDigitalSignature

	validFor := 2 * time.Minute
	notBefore := time.Now()
	notAfter := notBefore.Add(validFor)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	require.NoError(err)

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage:              keyUsage,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		} else {
			template.DNSNames = append(template.DNSNames, h)
		}
	}

	template.IsCA = true
	template.KeyUsage |= x509.KeyUsageCertSign

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	require.NoError(err)

	c, err := x509.ParseCertificate(derBytes)
	require.NoError(err)

	return c, string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes}))
}