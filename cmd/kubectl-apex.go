package main

import (
	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"apexauto/pkg/cmd"
	"os"
)

func main() {
	flags := pflag.NewFlagSet("kubectl-apex", pflag.ExitOnError)
	pflag.CommandLine = flags

	root := cmd.NewCmdApex(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}

