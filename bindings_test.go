package vimtea

import (
	"testing"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBindingRegistryBasics(t *testing.T) {
	registry := newBindingRegistry()

	registry.RegisterKey(
		key.NewBinding(key.WithKeys("a"), key.WithHelp("a", "Test binding A")),
		func(m *editorModel) tea.Cmd { return nil },
		ModeNormal,
	)

	registry.RegisterKey(
		key.NewBinding(key.WithKeys("b"), key.WithHelp("b", "Test binding B")),
		func(m *editorModel) tea.Cmd { return nil },
		ModeInsert,
	)

	binding := registry.FindExact("a", ModeNormal)
	require.NotNil(t, binding, "Binding for 'a' in normal mode not found")

	// Test that binding exists
	assert.Equal(t, "a", binding.Key, "Expected binding key 'a'")

	nonExistingBinding := registry.FindExact("nonexistent", ModeNormal)
	assert.Nil(t, nonExistingBinding, "Binding for 'nonexistent' should not exist")

	wrongModeBinding := registry.FindExact("a", ModeInsert)
	assert.Nil(t, wrongModeBinding, "Binding for 'a' should not exist in insert mode")
}

func TestBindingRegistryPrefix(t *testing.T) {
	registry := newBindingRegistry()

	registry.RegisterKey(
		key.NewBinding(key.WithKeys("dd"), key.WithHelp("dd", "Delete line")),
		func(m *editorModel) tea.Cmd { return nil },
		ModeNormal,
	)

	registry.RegisterKey(
		key.NewBinding(key.WithKeys("d$"), key.WithHelp("d$", "Delete to end of line")),
		func(m *editorModel) tea.Cmd { return nil },
		ModeNormal,
	)

	registry.RegisterKey(
		key.NewBinding(key.WithKeys("dw"), key.WithHelp("dw", "Delete word")),
		func(m *editorModel) tea.Cmd { return nil },
		ModeNormal,
	)

	// Test prefix detection
	isPrefix := registry.IsPrefix("d", ModeNormal)
	assert.True(t, isPrefix, "'d' should be detected as a prefix")

	notPrefix := registry.IsPrefix("x", ModeNormal)
	assert.False(t, notPrefix, "'x' should not be detected as a prefix")

	// Test complete key sequence
	binding := registry.FindExact("dd", ModeNormal)
	require.NotNil(t, binding, "Binding for 'dd' not found")

	// Test that binding exists
	assert.Equal(t, "dd", binding.Key, "Expected binding key 'dd'")
}

func TestBindingRegistryGetForMode(t *testing.T) {
	registry := newBindingRegistry()

	registry.RegisterKey(
		key.NewBinding(key.WithKeys("a"), key.WithHelp("a", "Test Normal A")),
		func(m *editorModel) tea.Cmd { return nil },
		ModeNormal,
	)

	registry.RegisterKey(
		key.NewBinding(key.WithKeys("b"), key.WithHelp("b", "Test Normal B")),
		func(m *editorModel) tea.Cmd { return nil },
		ModeNormal,
	)

	registry.RegisterKey(
		key.NewBinding(key.WithKeys("c"), key.WithHelp("c", "Test Insert C")),
		func(m *editorModel) tea.Cmd { return nil },
		ModeInsert,
	)

	normalBindings := registry.GetForMode(ModeNormal)
	assert.Len(t, normalBindings, 2, "Expected 2 bindings for normal mode")

	insertBindings := registry.GetForMode(ModeInsert)
	assert.Len(t, insertBindings, 1, "Expected 1 binding for insert mode")

	visualBindings := registry.GetForMode(ModeVisual)
	assert.Empty(t, visualBindings, "Expected 0 bindings for visual mode")
}

func TestCommandRegistry(t *testing.T) {
	registry := newCommandRegistry()

	commandCalled := false
	model := &editorModel{} // Minimal model for testing

	registry.Register("test", func(m *editorModel) tea.Cmd {
		commandCalled = true
		return nil
	})

	cmd := registry.Get("test")
	assert.NotNil(t, cmd, "Command 'test' not found")

	nonExistentCmd := registry.Get("nonexistent")
	assert.Nil(t, nonExistentCmd, "Command 'nonexistent' should not exist")

	// Test command execution
	cmdFunc := registry.Get("test")
	_ = cmdFunc(model)

	assert.True(t, commandCalled, "Command function was not called")
}
