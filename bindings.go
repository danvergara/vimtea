// Package vimtea provides a Vim-like text editor component for terminal applications
package vimtea

import (
	"reflect"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
)

type NormalKeyMap struct {
	EnterInsertMode          key.Binding
	BeginVisualSelection     key.Binding
	BeginVisualLineSelection key.Binding
	DeleteCharAtCursor       key.Binding
	EnterModeCommand         key.Binding
	AppendAfterCursor        key.Binding
	AppendAtEndOfLine        key.Binding
	InsertatStartOfLine      key.Binding
	OpenLineBelow            key.Binding
	OpenLineAbove            key.Binding
	YankLine                 key.Binding
	DeleteLine               key.Binding
	DeleteToEndOfLine        key.Binding
	PasteAfter               key.Binding
	PasteBefore              key.Binding
	Undo                     key.Binding
	Redo                     key.Binding
	DeleteInnerWord          key.Binding
	YankInnerWord            key.Binding
	ChangeInnerWord          key.Binding

	MoveCursorLeft            key.Binding
	MoveCursorDown            key.Binding
	MoveCursorUp              key.Binding
	MoveCursorRight           key.Binding
	MoveToNextWordStart       key.Binding
	MoveToPrevWordStart       key.Binding
	MoveCursorRightOrNextLine key.Binding
	MoveToStartOfLine         key.Binding
	MoveToFirstNonWhitespace  key.Binding
	MoveToEndOfLine           key.Binding
	MoveToStartOfDocument     key.Binding
	MoveToEndOfDocument       key.Binding
}

type VisualKeyMap struct {
	MoveCursorLeft            key.Binding
	MoveCursorDown            key.Binding
	MoveCursorUp              key.Binding
	MoveCursorRight           key.Binding
	MoveToNextWordStart       key.Binding
	MoveToPrevWordStart       key.Binding
	MoveCursorRightOrNextLine key.Binding
	MoveToStartOfLine         key.Binding
	MoveToFirstNonWhitespace  key.Binding
	MoveToEndOfLine           key.Binding
	MoveToStartOfDocument     key.Binding
	MoveToEndOfDocument       key.Binding

	ExitModeVisual                 key.Binding
	EnterModeCommand               key.Binding
	YankVisualSelection            key.Binding
	DeleteVisualSelection          key.Binding
	ReplaceVisualSelectionWithYank key.Binding
}

type InsertKeyMap struct {
	ExitModeInsert        key.Binding
	HandleInsertBackspace key.Binding
	HandleInsertTab       key.Binding
	HandleInsertEnterKey  key.Binding
	MoveUp                key.Binding
	MoveDown              key.Binding
	MoveLeft              key.Binding
	MoveRight             key.Binding
}

type CommandKeyMap struct {
	ExitModeCommand    key.Binding
	ExecuteModeCommand key.Binding
	CommandBackspace   key.Binding
}

type KeyMap struct {
	Normal  NormalKeyMap
	Visual  VisualKeyMap
	Insert  InsertKeyMap
	Command CommandKeyMap
}

func (k KeyMap) IsZero() bool {
	return reflect.ValueOf(k).IsZero()
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Normal: NormalKeyMap{
			EnterInsertMode:          key.NewBinding(key.WithKeys("i"), key.WithHelp("i", "Enter insert mode")),
			BeginVisualSelection:     key.NewBinding(key.WithKeys("v"), key.WithHelp("v", "Enter visual mode")),
			BeginVisualLineSelection: key.NewBinding(key.WithKeys("V"), key.WithHelp("V", "Enter visual line mode")),
			DeleteCharAtCursor:       key.NewBinding(key.WithKeys("x"), key.WithHelp("x", "Delete character at cursor")),
			EnterModeCommand:         key.NewBinding(key.WithKeys(":"), key.WithHelp(":", "Enter command mode")),

			AppendAfterCursor:   key.NewBinding(key.WithKeys("a"), key.WithHelp("a", "Append after cursor")),
			AppendAtEndOfLine:   key.NewBinding(key.WithKeys("A"), key.WithHelp("A", "Append at end of line")),
			InsertatStartOfLine: key.NewBinding(key.WithKeys("I"), key.WithHelp("I", "Insert at start of line")),
			OpenLineBelow:       key.NewBinding(key.WithKeys("o"), key.WithHelp("o", "Open line below")),
			OpenLineAbove:       key.NewBinding(key.WithKeys("O"), key.WithHelp("O", "Open line above")),

			YankLine:          key.NewBinding(key.WithKeys("yy"), key.WithHelp("yy", "Yank line")),
			DeleteLine:        key.NewBinding(key.WithKeys("dd"), key.WithHelp("dd", "Delete line")),
			DeleteToEndOfLine: key.NewBinding(key.WithKeys("D"), key.WithHelp("D", "Delete to end of line")),
			PasteAfter:        key.NewBinding(key.WithKeys("p"), key.WithHelp("p", "Paste after cursor")),
			PasteBefore:       key.NewBinding(key.WithKeys("P"), key.WithHelp("P", "Paste before cursor")),

			Undo:            key.NewBinding(key.WithKeys("u"), key.WithHelp("u", "Undo")),
			Redo:            key.NewBinding(key.WithKeys("ctrl+r"), key.WithHelp("ctrl+r", "Redo")),
			DeleteInnerWord: key.NewBinding(key.WithKeys("diw"), key.WithHelp("diw", "Delete inner word")),
			YankInnerWord:   key.NewBinding(key.WithKeys("yiw"), key.WithHelp("yiw", "Yank inner word")),
			ChangeInnerWord: key.NewBinding(key.WithKeys("ciw"), key.WithHelp("ciw", "Change inner word")),

			MoveCursorLeft:      key.NewBinding(key.WithKeys("h", "left"), key.WithHelp("h", "Move cursor left")),
			MoveCursorDown:      key.NewBinding(key.WithKeys("j", "down"), key.WithHelp("j", "Move cursor down")),
			MoveCursorUp:        key.NewBinding(key.WithKeys("k", "up"), key.WithHelp("k", "Move cursor up")),
			MoveCursorRight:     key.NewBinding(key.WithKeys("l", "right"), key.WithHelp("l", "Move cursor right")),
			MoveToNextWordStart: key.NewBinding(key.WithKeys("w"), key.WithHelp("w", "Move to next word")),
			MoveToPrevWordStart: key.NewBinding(key.WithKeys("b"), key.WithHelp("b", "Move to previous word")),

			MoveCursorRightOrNextLine: key.NewBinding(key.WithKeys("space"), key.WithHelp("space", "Move cursor right")),
			MoveToStartOfLine:         key.NewBinding(key.WithKeys("0"), key.WithHelp("0", "Move to start of line")),
			MoveToFirstNonWhitespace:  key.NewBinding(key.WithKeys("^"), key.WithHelp("^", "Move to first non-whitespace character")),
			MoveToEndOfLine:           key.NewBinding(key.WithKeys("$"), key.WithHelp("$", "Move to end of line")),
			MoveToStartOfDocument:     key.NewBinding(key.WithKeys("gg"), key.WithHelp("gg", "Move to document start")),
			MoveToEndOfDocument:       key.NewBinding(key.WithKeys("G"), key.WithHelp("G", "Move to document end")),
		},
		Visual: VisualKeyMap{
			ExitModeVisual:                 key.NewBinding(key.WithKeys("esc", "v", "V"), key.WithHelp("esc", "Exit visual mode")),
			EnterModeCommand:               key.NewBinding(key.WithKeys(":"), key.WithHelp(":", "Enter command mode")),
			YankVisualSelection:            key.NewBinding(key.WithKeys("y"), key.WithHelp("y", "Yank selection")),
			DeleteVisualSelection:          key.NewBinding(key.WithKeys("d", "x"), key.WithHelp("d", "Delete selection")),
			ReplaceVisualSelectionWithYank: key.NewBinding(key.WithKeys("p"), key.WithHelp("p", "Replace with yanked text")),

			MoveCursorLeft:            key.NewBinding(key.WithKeys("h", "left"), key.WithHelp("h", "Move cursor left")),
			MoveCursorDown:            key.NewBinding(key.WithKeys("j", "down"), key.WithHelp("j", "Move cursor down")),
			MoveCursorUp:              key.NewBinding(key.WithKeys("k", "up"), key.WithHelp("k", "Move cursor up")),
			MoveCursorRight:           key.NewBinding(key.WithKeys("l", "right"), key.WithHelp("l", "Move cursor right")),
			MoveToNextWordStart:       key.NewBinding(key.WithKeys("w"), key.WithHelp("w", "Move to next word")),
			MoveToPrevWordStart:       key.NewBinding(key.WithKeys("b"), key.WithHelp("b", "Move to previous word")),
			MoveCursorRightOrNextLine: key.NewBinding(key.WithKeys("space"), key.WithHelp("space", "Move cursor right")),
			MoveToStartOfLine:         key.NewBinding(key.WithKeys("0"), key.WithHelp("0", "Move to start of line")),
			MoveToFirstNonWhitespace:  key.NewBinding(key.WithKeys("^"), key.WithHelp("^", "Move to first non-whitespace character")),
			MoveToEndOfLine:           key.NewBinding(key.WithKeys("$"), key.WithHelp("$", "Move to end of line")),
			MoveToStartOfDocument:     key.NewBinding(key.WithKeys("gg"), key.WithHelp("gg", "Move to document start")),
			MoveToEndOfDocument:       key.NewBinding(key.WithKeys("G"), key.WithHelp("G", "Move to document end")),
		},
		Insert: InsertKeyMap{
			ExitModeInsert:        key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "Exit insert mode")),
			HandleInsertBackspace: key.NewBinding(key.WithKeys("backspace"), key.WithHelp("Backspace", "Backspace")),
			HandleInsertTab:       key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "Tab")),
			HandleInsertEnterKey:  key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "Enter")),
			MoveUp:                key.NewBinding(key.WithKeys("up"), key.WithHelp("up", "Move cursor up")),
			MoveDown:              key.NewBinding(key.WithKeys("down"), key.WithHelp("down", "Move cursor down")),
			MoveLeft:              key.NewBinding(key.WithKeys("left"), key.WithHelp("left", "Move cursor left")),
			MoveRight:             key.NewBinding(key.WithKeys("right"), key.WithHelp("right", "Move cursor right")),
		},
		Command: CommandKeyMap{
			ExitModeCommand:    key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "Exit command mode")),
			ExecuteModeCommand: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "Execute command mode")),
			CommandBackspace:   key.NewBinding(key.WithKeys("backspace"), key.WithHelp("backspace", "Backspace")),
		},
	}
}

// Command is a function that performs an action on the editor model
// and returns a bubbletea command
type Command func(m *editorModel) tea.Cmd

// KeyBinding represents a key binding that can be registered with the editor
// This is the public API for adding key bindings
type KeyBinding struct {
	Key         string               // The key sequence to bind (e.g. "j", "dd", "ctrl+f")
	Description string               // Human-readable description for help screens
	Mode        EditorMode           // Which editor mode this binding is active in
	Handler     func(Buffer) tea.Cmd // Function to execute when the key is pressed
}

// UndoRedoMsg is sent when an undo or redo operation is performed
// It contains the new cursor position and operation status
type UndoRedoMsg struct {
	NewCursor Cursor // New cursor position after undo/redo
	Success   bool   // Whether the operation succeeded
	IsUndo    bool   // True for undo, false for redo
}

// internalKeyBinding is the internal representation of a key binding
// used by the binding registry
type internalKeyBinding struct {
	Key     string     // The key sequence
	Help    string     // Help text describing the binding
	Command Command    // The command function to execute
	Mode    EditorMode // The editor mode this binding is active in
}

// CommandRegistry stores and manages commands that can be executed in command mode
// Commands are invoked by typing ":command" in command mode
type CommandRegistry struct {
	commands map[string]Command // Map of command names to command functions
}

// BindingRegistry manages key bindings for the editor
// It supports exact matches and prefix detection for multi-key sequences
type BindingRegistry struct {
	// Maps EditorMode -> key sequence -> binding
	exactBindings map[EditorMode]map[string]internalKeyBinding

	// Maps EditorMode -> key prefix -> true
	// Used to detect if a key sequence could be a prefix of a longer binding
	prefixBindings map[EditorMode]map[string]bool

	// List of all bindings for help display
	allBindings []internalKeyBinding
}

// newBindingRegistry creates a new empty binding registry
func newBindingRegistry() *BindingRegistry {
	return &BindingRegistry{
		exactBindings:  make(map[EditorMode]map[string]internalKeyBinding),
		prefixBindings: make(map[EditorMode]map[string]bool),
		allBindings:    []internalKeyBinding{},
	}
}

// newCommandRegistry creates a new empty command registry
func newCommandRegistry() *CommandRegistry {
	return &CommandRegistry{
		commands: make(map[string]Command),
	}
}

// RegisterKey registers a new key binding with the registry
// It automatically builds prefix maps for multi-key sequences
func (r *BindingRegistry) RegisterKey(b key.Binding, cmd Command, mode EditorMode) {
	binding := internalKeyBinding{
		Command: cmd,
		Mode:    mode,
		Help:    b.Help().Desc,
	}

	// Initialize mode map if needed.
	if r.exactBindings[mode] == nil {
		r.exactBindings[mode] = make(map[string]internalKeyBinding)
	}

	// Initialize prefix map if needed.
	if r.prefixBindings[mode] == nil {
		r.prefixBindings[mode] = make(map[string]bool)
	}

	// Map all physical string keys in a binding to a single Internal Key Binding.
	for _, k := range b.Keys() {
		bk := binding
		bk.Key = k
		r.exactBindings[mode][k] = bk

		// Register all prefixes of the key sequence.
		// For example, for "dw", register "d" as a prefix.
		for i := 1; i < len(k); i++ {
			prefix := k[:i]
			r.prefixBindings[mode][prefix] = true
		}

		// Add to the list of all bindings.
		r.allBindings = append(r.allBindings, bk)
	}
}

// FindExact looks for an exact match for the given key sequence in the specified mode
// It can handle numeric prefixes by ignoring them when looking for the command
func (r *BindingRegistry) FindExact(keySeq string, mode EditorMode) *internalKeyBinding {
	// Find where the numeric prefix ends (if any)
	nonDigitStart := 0
	for i, c := range keySeq {
		if c < '0' || c > '9' {
			nonDigitStart = i
			break
		}
	}

	// If the sequence is all digits, it's not a command
	if nonDigitStart == len(keySeq) {
		return nil
	}

	// Try to match without the numeric prefix
	cmdPart := keySeq[nonDigitStart:]
	if modeBindings, ok := r.exactBindings[mode]; ok {
		if binding, ok := modeBindings[cmdPart]; ok {
			return &binding
		}
	}

	// Try to match the full sequence (including any numeric prefix)
	if modeBindings, ok := r.exactBindings[mode]; ok {
		if binding, ok := modeBindings[keySeq]; ok {
			return &binding
		}
	}

	return nil
}

// IsPrefix checks if the key sequence is a prefix of any registered binding
// This is used to determine if we should wait for more input
func (r *BindingRegistry) IsPrefix(keySeq string, mode EditorMode) bool {
	if prefixes, ok := r.prefixBindings[mode]; ok {
		return prefixes[keySeq]
	}
	return false
}

// GetAll returns all registered key bindings
func (r *BindingRegistry) GetAll() []internalKeyBinding {
	return r.allBindings
}

// GetForMode returns all key bindings for the specified mode
func (r *BindingRegistry) GetForMode(mode EditorMode) []internalKeyBinding {
	var result []internalKeyBinding
	for _, binding := range r.allBindings {
		if binding.Mode == mode {
			result = append(result, binding)
		}
	}
	return result
}

// Register adds a command to the registry with the given name
func (r *CommandRegistry) Register(name string, cmd Command) {
	r.commands[name] = cmd
}

// Get retrieves a command by name, returning nil if not found
func (r *CommandRegistry) Get(name string) Command {
	cmd, ok := r.commands[name]
	if !ok {
		return nil
	}
	return cmd
}

// GetAll returns all registered commands as a map
func (r *CommandRegistry) GetAll() map[string]Command {
	return r.commands
}
