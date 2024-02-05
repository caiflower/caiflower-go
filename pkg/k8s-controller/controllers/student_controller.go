package controllers

import (
	v1 "caiflower.com/caiflower-go/pkg/k8s-controller/api/v1"
	"context"
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type StudentReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

func (sr *StudentReconciler) Reconcile(ctx context.Context, req reconcile.Request) (res reconcile.Result, err error) {
	// TODO 业务逻辑: 设置一个简单的业务逻辑，如果stu的svc不存在那么创建
	sr.Log.Info("accept", req.Namespace, req.Name)

	student := v1.Student{}
	err = sr.Get(ctx, req.NamespacedName, &student)
	if err != nil {
		sr.Log.Error(err, err.Error())
	}

	service := corev1.Service{}
	err = sr.Get(ctx, req.NamespacedName, &service)
	if err != nil {
		if apierrs.IsNotFound(err) {
			// 创建svc
			generateServiceFromStudent(student, &service)
			// 设置reference
			if err := ctrl.SetControllerReference(&student, &service, sr.Scheme); err != nil {
				return ctrl.Result{}, err
			}

			err = sr.Create(ctx, &service)
			if err != nil {
				sr.Log.Error(err, err.Error())
			}
		} else {
			sr.Log.Error(err, err.Error())
		}
	}

	return
}

func generateServiceFromStudent(student v1.Student, svcList ...*corev1.Service) (svc *corev1.Service) {
	if len(svcList) > 0 {
		svc = svcList[0]
	}
	svc.ObjectMeta.Name = student.ObjectMeta.Name
	svc.ObjectMeta.Namespace = student.ObjectMeta.Namespace

	svc.Spec.Type = corev1.ServiceTypeClusterIP
	svc.Spec.Ports = append(svc.Spec.Ports, corev1.ServicePort{
		Name:       "http",
		Port:       80,
		Protocol:   corev1.ProtocolTCP,
		TargetPort: intstr.IntOrString{IntVal: 8080},
	})
	m := make(map[string]string)
	m["app"] = "healthz"
	svc.Spec.Selector = m

	return svc
}

// SetupWithManager 设置需要监控的对象
func (sr *StudentReconciler) SetupWithManager(manager ctrl.Manager) error {
	err := ctrl.NewControllerManagedBy(manager).
		For(&v1.Student{}).
		Owns(&corev1.Service{}). // 监听ownerReferences、gkv属于Student的service，且controller：true那种
		Complete(sr)

	if err != nil {
		return err
	}

	return nil
}
