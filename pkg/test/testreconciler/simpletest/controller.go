/*
Copyright 2020 The Kubernetes Authors.

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

package simpletest

import (
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"github.com/yuwenma/kubebuilder-declarative-pattern/pkg/patterns/addon"
	"github.com/yuwenma/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/status"
	"github.com/yuwenma/kubebuilder-declarative-pattern/pkg/patterns/declarative"
	"github.com/yuwenma/kubebuilder-declarative-pattern/pkg/patterns/declarative/pkg/applier"

	api "github.com/yuwenma/kubebuilder-declarative-pattern/pkg/test/testreconciler/simpletest/v1alpha1"
)

var _ reconcile.Reconciler = &SimpleTestReconciler{}

// SimpleTestReconciler reconciles a SimpleTest object
type SimpleTestReconciler struct {
	declarative.Reconciler
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme

	watchLabels declarative.LabelMaker

	manifestController declarative.ManifestController

	applier applier.Applier
}

func (r *SimpleTestReconciler) setupReconciler(mgr ctrl.Manager) error {
	labels := map[string]string{
		"example-app": "simpletest",
	}

	r.watchLabels = declarative.SourceLabel(mgr.GetScheme())

	return r.Reconciler.Init(mgr, &api.SimpleTest{},
		declarative.WithObjectTransform(declarative.AddLabels(labels)),
		declarative.WithOwner(declarative.SourceAsOwner),
		declarative.WithLabels(r.watchLabels),
		declarative.WithStatus(status.NewBasic(mgr.GetClient())),

		// TODO: Readd prune
		//declarative.WithApplyPrune(),

		declarative.WithObjectTransform(addon.ApplyPatches),

		// Add other options for testing
		//declarative.WithApplyValidation(),

		// Don't turn on metrics, they create another watch and cause unpredictable requests
		// declarative.WithReconcileMetrics(0, nil),

		declarative.WithManifestController(r.manifestController),
		declarative.WithApplier(r.applier),
	)
}

func (r *SimpleTestReconciler) SetupWithManager(mgr ctrl.Manager) error {
	if err := r.setupReconciler(mgr); err != nil {
		return err
	}

	c, err := controller.New("simpletest-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to SimpleTest objects
	err = c.Watch(&source.Kind{Type: &api.SimpleTest{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to deployed objects
	_, err = declarative.WatchChildren(declarative.WatchChildrenOptions{Manager: mgr, Controller: c, Reconciler: r, LabelMaker: r.watchLabels})
	if err != nil {
		return err
	}

	return nil
}
