package encryption

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	tests := []struct {
		algorithm int
		input     string
		output    string
	}{
		{Md5, "hello world", "5eb63bbbe01eeed093cb22bb8f5acdc3"},
		{Sha224, "hello world", "2f05477fc24bb4faefd86517156dafdecec45b8ad3cf2522a563582b"},
		{Sha256, "hello world", "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"},
		{Sha384, "hello world", "fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd"},
		{Sha512, "hello world", "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f"},
		{Sha512224, "hello world", "22e0d52336f64a998085078b05a6e37b26f8120f43bf4db4c43a64ee"},
		{Sha512256, "hello world", "0ac561fac838104e3f2e4ad107b4bee3e938bf15f2b15f009ccccd61a913f017"},
		{Hmac512, "hello world", "3OQUyxrE59QA6+dfQ3upCtpBwzmHQnawgHt6jZ1ztW2954mOmcTtkmWfMMzUDHEu5Rf8IpASz/zXmNnvfjV92A=="},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			hash := Hash{
				Input:     test.input,
				Algorithm: test.algorithm,
			}
			hash.Hash()

			if hash.Output != test.output {
				t.Fatalf("expected '%s', got '%s'", test.output, hash.Output)
			}
		})
	}
}
