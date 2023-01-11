package regexf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var regex = NewRegex()

func TestRegex_Pwd(t *testing.T) {
	assert.Equal(t, regex.Pwd("Light123!@#$"), true)
	assert.Equal(t, regex.Pwd("Light123333"), true)
	assert.Equal(t, regex.Pwd("Li23456789"), true)
	assert.Equal(t, regex.Pwd("123456789l!"), true)
	assert.Equal(t, regex.Pwd("l123456789!"), true)
	assert.Equal(t, regex.Pwd("LightWelcome"), false)
	assert.Equal(t, regex.Pwd("123456789"), false)
	assert.Equal(t, regex.Pwd("L23456789"), false)
	assert.Equal(t, regex.Pwd("Ll2啊56789"), false)
}

func TestRegex_Email(t *testing.T) {
	assert.Equal(t, regex.Email("Light@test.com"), true)
	assert.Equal(t, regex.Email("Light@test"), false)
}
