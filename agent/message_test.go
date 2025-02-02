package agent_test

import (
	"testing"

	"github.com/kcarretto/paragon/agent"
	"github.com/kcarretto/paragon/api/codec"
	"github.com/stretchr/testify/require"
)

func TestMessageWrite(t *testing.T) {
	expected := "test message"

	msg := &agent.Message{}
	require.True(t, msg.IsEmpty())

	n, err := msg.Write([]byte(expected))
	require.NoError(t, err)
	require.Equal(t, len(expected), n)
	require.Equal(t, 1, len(msg.Logs))
	require.Equal(t, expected, msg.Logs[0])
}

func TestMessageWriteResult(t *testing.T) {
	expected := &codec.Result{}

	msg := &agent.Message{}
	msg.WriteResult(expected)
	require.Equal(t, 1, len(msg.Results))
	require.Equal(t, expected, msg.Results[0])
}

func TestServerMessageWrite(t *testing.T) {
	expectedTask := &codec.Task{}
	expected := &agent.ServerMessage{}

	msg := &agent.ServerMessage{
		Tasks: []*codec.Task{expectedTask},
	}
	expected.WriteServerMessage(msg)

	require.NotNil(t, expected)
	require.Equal(t, 1, len(expected.Tasks))
	require.Equal(t, expectedTask, expected.Tasks[0])
}
