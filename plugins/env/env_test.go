package env

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/apex/apex/function"
	"github.com/apex/log"
	"github.com/apex/log/handlers/discard"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.SetHandler(discard.New())
}

func TestPlugin_Run_buildHook(t *testing.T) {
	p := &Plugin{}

	f := &function.Function{
		Log:  log.Log,
		Path: os.TempDir(),
		Config: function.Config{
			Environment: map[string]string{
				"foo": "bar",
				"bar": "baz",
			},
		},
	}

	err := p.Run(function.BuildHook, f)
	assert.NoError(t, err)

	b, err := ioutil.ReadFile(filepath.Join(os.TempDir(), ".env.json"))
	assert.NoError(t, err)

	result := `{"bar":"baz","foo":"bar"}
`

	assert.Equal(t, result, string(b))
}
