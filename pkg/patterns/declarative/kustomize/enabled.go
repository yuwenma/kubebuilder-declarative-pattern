//go:build !without_kustomize
// +build !without_kustomize

package kustomize

import (
	"fmt"

	"sigs.k8s.io/kustomize/api/krusty"
	"sigs.k8s.io/kustomize/kyaml/filesys"
)

func init() {
	Kustomize = &EnabledKustomize{}
}

var _ Kustomizer = &EnabledKustomize{}

type EnabledKustomize struct{}

func (k *EnabledKustomize) Run(fs filesys.FileSystem, manifestPath string) ([]byte, error) {
	// run kustomize to create final manifest
	opts := krusty.MakeDefaultOptions()
	kustomizer := krusty.MakeKustomizer(opts)
	m, err := kustomizer.Run(fs, manifestPath)
	if err != nil {
		return nil, fmt.Errorf("error running kustomize: %v", err)
	}

	manifestYaml, err := m.AsYaml()
	if err != nil {
		return nil, fmt.Errorf("error converting kustomize output to yaml: %v", err)
	}
	return manifestYaml, fmt.Errorf("!!!!!EnabledKustomize!!!!!")
}
