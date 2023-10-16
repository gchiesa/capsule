package podlabels

import (
	"context"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
)

type PodLabelsReconciler struct {
	abstractPodLabelsReconciler

	Log logr.Logger
}

func (r *PodLabelsReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager) error {
	r.abstractPodLabelsReconciler = abstractPodLabelsReconciler{
		obj:    &corev1.Pod{},
		client: mgr.GetClient(),
		log:    r.Log,
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(r.abstractPodLabelsReconciler.obj, r.abstractPodLabelsReconciler.forOptionPerInstanceName(ctx)).
		Complete(r)
}
