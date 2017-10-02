/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"io"

	"github.com/spf13/cobra"
	"k8s.io/kops/cmd/kops/util"
	"k8s.io/kubernetes/pkg/kubectl/cmd/templates"
	"k8s.io/kubernetes/pkg/kubectl/util/i18n"
)

var (
	edit_long = templates.LongDesc(i18n.T(`Edit a resource configuration.
	This command changes the desired configuration in the registry.

    	To set your preferred editor, you can define the EDITOR environment variable.
    	When you have done this, kops will use the editor that you have set.

    	kops edit does not update the cloud resources, to apply the changes use "kops update cluster".
    	`))

	edit_example = templates.Examples(i18n.T(`
	# Edit a cluster configuration.
	kops edit cluster k8s-cluster.example.com --state=s3://kops-state-1234

	# Edit a instance group configuration.
	kops edit ig --name k8s-cluster.example.com \
	  --state=s3://kops-state-1234 nodes

	# Edit a cluster federation configuration.
	kops edit federation k8s-cluster.example.com --state=s3://kops-state-1234
	`))
)

func NewCmdEdit(f *util.Factory, out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "edit",
		Short:   i18n.T("Edit clusters and other resources."),
		Long:    edit_long,
		Example: edit_example,
	}

	// create subcommands
	cmd.AddCommand(NewCmdEditCluster(f, out))
	cmd.AddCommand(NewCmdEditInstanceGroup(f, out))
	cmd.AddCommand(NewCmdEditFederation(f, out))

	return cmd
}
