package kustomize

import (
	"sigs.k8s.io/kustomize/kyaml/filesys"
)

var Kustomize Kustomizer

type Kustomizer interface {
	Run(fs filesys.FileSystem, manifestPath string) ([]byte, error)
}
