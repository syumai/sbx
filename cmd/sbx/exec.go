package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"syscall"
)

func sandboxExec(ctx context.Context, policy string, command string, args ...string) error {
	if command == "" {
		return errors.New("command is required")
	}
	sandboxArgs := []string{"-p", policy, command}
	sandboxArgs = append(sandboxArgs, args...)

	cmd := exec.CommandContext(ctx, "sandbox-exec", sandboxArgs...)

	cmd.Stdin = os.Stdin

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create stdout pipe: %v\n", err)
		os.Exit(1)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create stderr pipe: %v\n", err)
		os.Exit(1)
	}

	if err := cmd.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to start command: %v\n", err)
		os.Exit(1)
	}

	go func() {
		io.Copy(os.Stdout, stdout)
	}()

	go func() {
		io.Copy(os.Stderr, stderr)
	}()

	err = cmd.Wait()

	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				exitCode = status.ExitStatus()
			} else {
				exitCode = 1
			}
		} else {
			fmt.Fprintf(os.Stderr, "command execution failed: %v\n", err)
			exitCode = 1
		}
	}

	os.Exit(exitCode)
	return nil
}
