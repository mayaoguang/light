package cryptof

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHash(t *testing.T) {
	var (
		factory = NewHash()
		text    = "那是一道光"
	)

	assert.Equal(t, factory.Md5(text), "3b9a71f0d8e0b48f38fb594ca8da81b0")
	assert.Equal(t, factory.Sha1(text), "f62180d40ee9c68e3e29a2a686662b8bb3265964")
	assert.Equal(t, factory.Sha224(text), "d3ccb705834adbd164764b456f821b07189acaba9d96350119dcaf5d")
	assert.Equal(t, factory.Sha256(text), "429fa35c12afa68091c91255eaca70f99b0308ec8d5dd6057ac0bd16fb5ae761")
	assert.Equal(t, factory.Sha384(text), "97b38341272b22eac4d0913c2dc05ca8423501ff24fb6425750d64c419a793c4b714e083441684d1b3404120a8cfd841")
	assert.Equal(t, factory.Sha512(text), "c0a86be532b67c99a1caec6daf5fe685a76985b4d0f77393afa7f22ad102d0b6595c7f2147d95cc0769405a0881124cfb380ef82c96fe21629736de64d355415")
}
