package main

import (
	"github.com/"
	"github.com/"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: aws.Provider})
}
