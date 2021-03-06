// Copyright Jetstack Ltd. See LICENSE for details.
package main

import (
	"fmt"
	"os"

	"github.com/jetstack/kube-oidc-proxy/cmd"
	"github.com/jetstack/kube-oidc-proxy/pkg/util"
)

func main() {
	stopCh := util.SignalHandler()
	cmd := cmd.NewRunCommand(stopCh)

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
