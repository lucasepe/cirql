package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/lucasepe/cirql/internal/cmd/pull"
	"github.com/lucasepe/cirql/internal/cmd/push"
	"github.com/lucasepe/cirql/internal/cmd/rm"
)

var (
	BuildKey = buildKey{}
)

type Action int

const (
	NoAction Action = iota
	Push
	Pull
	Remove
	ShowHelp
	ShowVersion
)

func Run(ctx context.Context) (err error) {
	op := NoAction

	nargs := len(os.Args)
	if nargs > 1 {
		op = chosenAction(os.Args[1:])
	}
	if (op == ShowHelp) || (op == NoAction) {
		usage(os.Stdout)
		return nil
	}

	if op == ShowVersion {
		bld := ctx.Value(BuildKey).(string)
		fmt.Fprintf(os.Stdout, "%s - build: %s\n", appName, bld)
		return nil
	}

	switch op {
	case Remove:
		err = rm.Do(os.Args[2:])
	case Pull:
		err = pull.Do(os.Args[2:])
	case Push:
		err = push.Do(os.Args[2:])
	}

	return err
}

func chosenAction(args []string) Action {
	for _, opt := range args {
		switch opt {
		case "version":
			return ShowVersion
		case "rm":
			return Remove
		case "push":
			return Push
		case "pull":
			return Pull
		}
	}

	return ShowHelp
}

type (
	buildKey struct{}
)
