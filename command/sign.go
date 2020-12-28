package command

import (
	"flag"
	"fmt"
)

type Sign struct {
	flagSet *flag.FlagSet

	path string
	key  string
}

func (cmd *Sign) Name() string {
	return "sign"
}

func (cmd *Sign) Init(args []string) error {
	cmd.flagSet = flag.NewFlagSet(cmd.Name(), flag.ExitOnError)

	cmd.flagSet.StringVar(&cmd.path, "path", "", "path to the package (required)")
	cmd.flagSet.StringVar(&cmd.key, "key", "", "path to the private key (required)")

	return cmd.flagSet.Parse(args)
}

func (cmd *Sign) Run() error {
	fmt.Println("path: ", cmd.path, "!")
	fmt.Println("key: ", cmd.key, "!")

	return nil
}