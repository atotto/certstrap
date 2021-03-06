/*-
 * Copyright 2015 Square Inc.
 * Copyright 2014 CoreOS
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package pkix

import (
	"testing"
)

func TestCreateCertificateAuthority(t *testing.T) {
	key, err := CreateRSAKey(rsaBits)
	if err != nil {
		t.Fatal("Failed creating rsa key:", err)
	}

	crt, err := CreateCertificateAuthority(key, "OU", TimeNow().AddDate(5, 0, 0), "test", "US", "California", "San Francisco", "CA Name", []string{".example.com"}, []string{"www.example.com"})
	if err != nil {
		t.Fatal("Failed creating certificate authority:", err)
	}
	rawCrt, err := crt.GetRawCertificate()
	if err != nil {
		t.Fatal("Failed to get x509.Certificate:", err)
	}

	if err = rawCrt.CheckSignatureFrom(rawCrt); err != nil {
		t.Fatal("Failed to check signature:", err)
	}

	if rawCrt.Subject.OrganizationalUnit[0] != "OU" {
		t.Fatal("Failed to verify hostname:", err)
	}

	if !TimeNow().After(rawCrt.NotBefore) {
		t.Fatal("Failed to be after NotBefore")
	}

	if !TimeNow().Before(rawCrt.NotAfter) {
		t.Fatal("Failed to be before NotAfter")
	}

	if crt.crt.PermittedDNSDomainsCritical != true {
		t.Fatal("Permitted DNS Domains is not set to critical")
	}

	if len(crt.crt.PermittedDNSDomains) != 1 {
		t.Fatal("More than one entry found in list of permitted DNS domains")
	}

	if crt.crt.PermittedDNSDomains[0] != ".example.com" {
		t.Fatalf("Wrong permitted DNS domain, want %q, got %q", ".example.com", crt.crt.PermittedDNSDomains[0])
	}

	if crt.crt.DNSNames[0] != "www.example.com" {
		t.Fatalf("Wrong DNS names, want %q, got %q", "www.example.com", crt.crt.DNSNames[0])
	}
}
