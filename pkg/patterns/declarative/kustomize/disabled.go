//go:build without_kustomize
// +build without_kustomize

package kustomize

import (
	"sigs.k8s.io/kustomize/kyaml/filesys"
)

func init() {
	Kustomize = &NoopKustomize{}
}

var _ Kustomizer = &NoopKustomize{}

type NoopKustomize struct{}

func (k *NoopKustomize) Run(fs filesys.FileSystem, manifestPath string) ([]byte, error) {
	return nil, nil
}
