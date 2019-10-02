/*

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

package controller

import (
	"context"

	"github.com/go-logr/logr"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	ir "github.com/pivotal/image-relocation/pkg/api/v1alpha1"
	webhookv1alpha1 "github.com/pivotal/image-relocation/pkg/api/v1alpha1"
)

// ImageMapReconciler reconciles a ImageMap object
type ImageMapReconciler struct {
	client.Client
	Log logr.Logger
}

func (r *ImageMapReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("imagemap", req.NamespacedName)

	log.Info("Reconcile", "req", req)

	var imageMap ir.ImageMap
	if err := r.Get(ctx, req.NamespacedName, &imageMap); err != nil {
		if apierrs.IsNotFound(err) {
			log.Info("deleting", "ImageMap", req.NamespacedName)
			// TODO: delete the image map
			return ctrl.Result{}, nil
		}
		log.Error(err, "unable to fetch ImageMap")
		return ctrl.Result{}, err
	}

	log.Info("adding", "ImageMap", req.NamespacedName)
	// TODO: add the image map

	return ctrl.Result{}, nil
}

func (r *ImageMapReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&webhookv1alpha1.ImageMap{}).
		Complete(r)
}
