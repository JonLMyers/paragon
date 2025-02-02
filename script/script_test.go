package script_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/kcarretto/paragon/script"

	"github.com/stretchr/testify/require"
)

const testScriptContent = `
print("loading")

def count():
	nums = [str(x) for x in range(10)]
	print(",".join(nums))

def main():
    count()
`
const testScriptOutput = "[test_execute] loading\n[test_execute] 0,1,2,3,4,5,6,7,8,9\n"

func TestNewScript(t *testing.T) {
	code := script.New("test_execute", bytes.NewBufferString(testScriptContent))

	val := make([]byte, len(testScriptContent))
	n, err := code.Read(val)
	require.NoError(t, err)
	require.Equal(t, len(testScriptContent), n)
	require.Equal(t, testScriptContent, string(val))
}

func TestExecOutput(t *testing.T) {
	code := script.New("test_execute", bytes.NewBufferString(testScriptContent))

	err := code.Exec(context.Background())
	require.NoError(t, err)

	// TODO: Check output of logger

	// result, err := ioutil.ReadAll(reader)
	// require.NoError(t, err)

	// require.Equal(t, testScriptOutput, string(result))
}
