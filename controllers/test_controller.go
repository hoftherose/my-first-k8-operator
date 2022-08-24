/*
Copyright 2022.

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

package controllers

import (
	"context"
	"errors"
	"os"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/go-logr/logr"
	testingv1 "github.com/hoftherose/my-first-k8-operator/api/v1"
)

// TestReconciler reconciles a Test object
type TestReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=testing.example.com,resources=tests,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=testing.example.com,resources=tests/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=testing.example.com,resources=tests/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Test object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.12.1/pkg/reconcile
func (r *TestReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)

	test_resource := &testingv1.Test{}
	r.Get(ctx, types.NamespacedName{Name: req.Name, Namespace: req.Namespace}, test_resource)

	l.Info("Reconciling")

	if _, err := os.Stat(test_resource.Status.Foo); errors.Is(err, os.ErrNotExist) {
		l.Info("Creating file")
		os.Create(test_resource.Spec.Foo)
		l.Info("File created")
		// Update created status
		test_resource.Status.Foo = test_resource.Spec.Foo
		r.Status().Update(ctx, test_resource)
	}

	if test_resource.Status.Foo != test_resource.Spec.Foo {
		renameFile(ctx, test_resource, l)
		r.Status().Update(ctx, test_resource)
	}

	return ctrl.Result{}, nil
}

func renameFile(ctx context.Context, test_resource *testingv1.Test, l logr.Logger) {
	err := os.Rename(test_resource.Status.Foo, test_resource.Spec.Foo)
	if err != nil {
		l.Error(err, "Could not rename file.")
	} else {
		l.Info("File changed correctly")
		test_resource.Status.Foo = test_resource.Spec.Foo
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *TestReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&testingv1.Test{}).
		Complete(r)
}
