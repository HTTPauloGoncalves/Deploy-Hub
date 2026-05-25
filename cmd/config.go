/*
Copyright (c) 2026
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/HTTPauloGoncalves/Deploy-Hub/internal/config"
	"github.com/kballard/go-shellquote"
	"github.com/spf13/cobra"
)

// implementar mais funções
var configCmd = &cobra.Command{
	Use:     "config",
	Aliases: []string{"confg"},
	Short:   "Gerencia o arquivo de configuracao",
}

var configEditCmd = &cobra.Command{
	Use:   "edit",
	Short: "Abre o arquivo de configuracao",
	Run: func(cmd *cobra.Command, args []string) {
		err := openConfigFile()
		if err != nil {
			fmt.Println("Erro:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configEditCmd)
}

func openConfigFile() error {
	if err := os.MkdirAll(filepath.Dir(config.DefaultPath), 0755); err != nil {
		return err
	}

	command, args := editorCommand(config.DefaultPath)
	editor := exec.Command(command, args...)
	editor.Stdin = os.Stdin
	editor.Stdout = os.Stdout
	editor.Stderr = os.Stderr

	return editor.Run()
}

func editorCommand(path string) (string, []string) {
	if runtime.GOOS == "windows" {
		return "notepad", []string{path}
	}

	if editor := os.Getenv("EDITOR"); editor != "" {
		parts, err := shellquote.Split(editor)
		if err == nil && len(parts) > 0 {
			return parts[0], append(parts[1:], path)
		}

		return editor, []string{path}
	}

	if _, err := exec.LookPath("nano"); err == nil {
		return "nano", []string{path}
	}

	return "vi", []string{path}
}
