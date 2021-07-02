package hasher_test

import (
	"testing"

	"github.com/don-nv/pkg/hasher"
)

func Test_MakeSHA256_FromString(t *testing.T) {
	hasher := hasher.New()

	hash, _ := hasher.MakeSHA256_FromString("string")
	println(hash)

}
