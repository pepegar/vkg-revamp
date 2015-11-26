package cmd

import "testing"

func TestFormat(t *testing.T) {
	plugins := []string{
		"gundo",
		"vim-airline",
		"ack-vim",
	}

	expected := `* gundo
* vim-airline
* ack-vim`

	if formatOutput(plugins) != expected {
		t.Error(formatOutput(plugins))
	}
}
