package main

import (
	"fmt"
	"github.com/HashDataInc/packer-plugin-qingcloud/builder/qingcloud"
	qingcloudVersion "github.com/HashDataInc/packer-plugin-qingcloud/version"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder(plugin.DEFAULT_NAME, new(qingcloud.Builder))
	pps.SetVersion(qingcloudVersion.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
