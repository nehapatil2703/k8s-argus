// Code generated by go generate; DO NOT EDIT.
package enums

import (
	"fmt"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	batchv1 "k8s.io/api/batch/v1"
	batchv1beta1 "k8s.io/api/batch/v1beta1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	networkingv1beta1 "k8s.io/api/networking/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
)

// ResourceType resource
type ResourceType uint32

// NOTE: RESOURCE_MODIFICATION need to change when adding/deleting resource for monitoring
// START This comment starts the section to be changed when enum added or modified
const (
	Unknown ResourceType = iota
	Pods
	Deployments
	Services
	Hpas
	Nodes
	ETCD
	Namespaces
	DaemonSets
	ReplicaSets
	StatefulSets
	ConfigMaps
	PersistentVolumes
	PersistentVolumeClaims
	CronJobs
	Jobs
	Ingresses
	Secrets
	EndPoints
	NetworkPolicies
)

// ALLResourceTypes All resource type slice
// NOTE: RESOURCE_MODIFICATION need to change when adding/deleting resource for monitoring
var ALLResourceTypes = []ResourceType{
	Pods,
	Deployments,
	Services,
	Hpas,
	Nodes,
	ETCD,
	Namespaces,
	DaemonSets,
	ReplicaSets,
	StatefulSets,
	ConfigMaps,
	PersistentVolumes,
	PersistentVolumeClaims,
	CronJobs,
	Jobs,
	Ingresses,
	Secrets,
	EndPoints,
	NetworkPolicies,
}

// nolint: clyclop
func (resourceType ResourceType) MarshalText() ([]byte, error) {
	switch resourceType {
	case Unknown:
		return []byte("unknown"), nil
	case Pods:
		return []byte("pods"), nil
	case Deployments:
		return []byte("deployments"), nil
	case Services:
		return []byte("services"), nil
	case Hpas:
		return []byte("horizontalpodautoscalers"), nil
	case Nodes:
		return []byte("nodes"), nil
	case ETCD:
		return []byte("etcd"), nil
	case Namespaces:
		return []byte("namespaces"), nil
	case DaemonSets:
		return []byte("daemonsets"), nil
	case ReplicaSets:
		return []byte("replicasets"), nil
	case StatefulSets:
		return []byte("statefulsets"), nil
	case ConfigMaps:
		return []byte("configmaps"), nil
	case PersistentVolumes:
		return []byte("persistentvolumes"), nil
	case PersistentVolumeClaims:
		return []byte("persistentvolumeclaims"), nil
	case CronJobs:
		return []byte("cronjobs"), nil
	case Jobs:
		return []byte("jobs"), nil
	case Ingresses:
		return []byte("ingresses"), nil
	case Secrets:
		return []byte("secrets"), nil
	case EndPoints:
		return []byte("endpoints"), nil
	case NetworkPolicies:
		return []byte("networkpolicies"), nil
	}
	return nil, fmt.Errorf("not a valid ResourceType to marshal %v", uint32(resourceType))
}

// nolint: clyclop
func ParseResourceType(resourceType string) (ResourceType, error) {
	switch strings.ToLower(resourceType) {
	case "pods", "pod", "po":
		return Pods, nil
	case "deployments", "deployment", "deploy":
		return Deployments, nil
	case "services", "service", "svc":
		return Services, nil
	case "horizontalpodautoscalers", "horizontalpodautoscaler", "hpas", "hpa":
		return Hpas, nil
	case "node", "nodes":
		return Nodes, nil
	case "etcd":
		return ETCD, nil
	case "namespaces", "namespace", "ns":
		return Namespaces, nil
	case "daemonsets", "daemonset", "ds":
		return DaemonSets, nil
	case "replicasets", "replicaset", "rs":
		return ReplicaSets, nil
	case "statefulsets", "statefulset", "sts":
		return StatefulSets, nil
	case "configmaps", "configmap", "cm":
		return ConfigMaps, nil
	case "persistentvolumes", "persistentvolume", "pv":
		return PersistentVolumes, nil
	case "persistentvolumeclaims", "persistentvolumeclaim", "pvc":
		return PersistentVolumeClaims, nil
	case "cronjobs", "cronjob", "cj":
		return CronJobs, nil
	case "jobs", "job":
		return Jobs, nil
	case "ingresses", "ingress", "ing":
		return Ingresses, nil
	case "secrets", "secret":
		return Secrets, nil
	case "endpoints", "endpoint":
		return EndPoints, nil
	case "networkpolicies", "networkpolicy", "netpol":
		return NetworkPolicies, nil
	}

	return Unknown, fmt.Errorf("not a valid ResourceType to parse: %s", resourceType)
}

// nolint: clyclop
func (resourceType ShortResourceType) MarshalText() ([]byte, error) {
	switch ResourceType(resourceType) {
	case Unknown:
		return []byte("unknown"), nil
	case Pods:
		return []byte("pod"), nil
	case Deployments:
		return []byte("deploy"), nil
	case Services:
		return []byte("svc"), nil
	case Hpas:
		return []byte("hpa"), nil
	case Nodes:
		return []byte("node"), nil
	case ETCD:
		return []byte("etcd"), nil
	case Namespaces:
		return []byte("ns"), nil
	case DaemonSets:
		return []byte("ds"), nil
	case ReplicaSets:
		return []byte("rs"), nil
	case StatefulSets:
		return []byte("sts"), nil
	case ConfigMaps:
		return []byte("cm"), nil
	case PersistentVolumes:
		return []byte("pv"), nil
	case PersistentVolumeClaims:
		return []byte("pvc"), nil
	case CronJobs:
		return []byte("cj"), nil
	case Jobs:
		return []byte("job"), nil
	case Ingresses:
		return []byte("ing"), nil
	case Secrets:
		return []byte("secret"), nil
	case EndPoints:
		return []byte("ep"), nil
	case NetworkPolicies:
		return []byte("netpol"), nil
	}

	return nil, fmt.Errorf("not a valid ShortResourceType to marshal %d", resourceType)
}

// nolint: clyclop
func ParseShortResourceType(shortResourceType string) (ShortResourceType, error) {
	var l ResourceType
	switch strings.ToLower(shortResourceType) {
	case "unknown":
		l = Unknown
	case "pod":
		l = Pods
	case "deploy":
		l = Deployments
	case "svc":
		l = Services
	case "hpa":
		l = Hpas
	case "node":
		l = Nodes
	case "etcd":
		l = ETCD
	case "ns":
		l = Namespaces
	case "ds":
		l = DaemonSets
	case "rs":
		l = ReplicaSets
	case "sts":
		l = StatefulSets
	case "cm":
		l = ConfigMaps
	case "pv":
		l = PersistentVolumes
	case "pvc":
		l = PersistentVolumeClaims
	case "cj":
		l = CronJobs
	case "job":
		l = Jobs
	case "ing":
		l = Ingresses
	case "secret":
		l = Secrets
	case "ep":
		l = EndPoints
	case "netpol":
		l = NetworkPolicies
	default:

		return ShortResourceType(Unknown), fmt.Errorf("not a valid ShortResourceType to parse: %q", shortResourceType)
	}

	return ShortResourceType(l), nil
}

// nolint: clyclop
func (resourceType *ResourceType) Title() string {
	switch *resourceType {
	case Unknown:
		return "Unknown"
	case Pods:
		return "Pod"
	case Deployments:
		return "Deployment"
	case Services:
		return "Service"
	case Hpas:
		return "HorizontalPodAutoscaler"
	case Nodes:
		return "Node"
	case ETCD:
		return "Etcd"
	case Namespaces:
		return "Namespace"
	case DaemonSets:
		return "DaemonSet"
	case ReplicaSets:
		return "ReplicaSet"
	case StatefulSets:
		return "StatefulSet"
	case ConfigMaps:
		return "ConfigMap"
	case PersistentVolumes:
		return "PersistentVolume"
	case PersistentVolumeClaims:
		return "PersistentVolumeClaim"
	case CronJobs:
		return "CronJob"
	case Jobs:
		return "Job"
	case Ingresses:
		return "Ingress"
	case Secrets:
		return "Secret"
	case EndPoints:
		return "EndPoint"
	case NetworkPolicies:
		return "NetworkPolicy"
	}

	return "Unknown"
}

// nolint: clyclop
func (resourceType *ResourceType) TitlePlural() string {
	switch *resourceType {
	case Unknown:
		return "Unknown"
	case Pods:
		return "Pods"
	case Deployments:
		return "Deployments"
	case Services:
		return "Services"
	case Hpas:
		return "Hpas"
	case Nodes:
		return "Nodes"
	case ETCD:
		return "Etcds"
	case Namespaces:
		return "Namespaces"
	case DaemonSets:
		return "DaemonSets"
	case ReplicaSets:
		return "ReplicaSets"
	case StatefulSets:
		return "StatefulSets"
	case ConfigMaps:
		return "ConfigMaps"
	case PersistentVolumes:
		return "PersistentVolumes"
	case PersistentVolumeClaims:
		return "PersistentVolumeClaims"
	case CronJobs:
		return "CronJobs"
	case Jobs:
		return "Jobs"
	case Ingresses:
		return "Ingresses"
	case Secrets:
		return "Secrets"
	case EndPoints:
		return "EndPoints"
	case NetworkPolicies:
		return "NetworkPolicies"
	}

	return "Unknown"
}

// nolint: clyclop
func (resourceType *ResourceType) K8SObjectType() runtime.Object {
	switch *resourceType {
	case Pods:
		return &corev1.Pod{} // nolint: exhaustivestruct
	case Deployments:
		return &appsv1.Deployment{} // nolint: exhaustivestruct
	case Services:
		return &corev1.Service{} // nolint: exhaustivestruct
	case Hpas:
		return &autoscalingv1.HorizontalPodAutoscaler{} // nolint: exhaustivestruct
	case Nodes:
		return &corev1.Node{} // nolint: exhaustivestruct
	case Namespaces:
		return &corev1.Namespace{} // nolint: exhaustivestruct
	case DaemonSets:
		return &appsv1.DaemonSet{} // nolint: exhaustivestruct
	case ReplicaSets:
		return &appsv1.ReplicaSet{} // nolint: exhaustivestruct
	case StatefulSets:
		return &appsv1.StatefulSet{} // nolint: exhaustivestruct
	case ConfigMaps:
		return &corev1.ConfigMap{} // nolint: exhaustivestruct
	case PersistentVolumes:
		return &corev1.PersistentVolume{} // nolint: exhaustivestruct
	case PersistentVolumeClaims:
		return &corev1.PersistentVolumeClaim{} // nolint: exhaustivestruct
	case CronJobs:
		return &batchv1beta1.CronJob{} // nolint: exhaustivestruct
	case Jobs:
		return &batchv1.Job{} // nolint: exhaustivestruct
	case Ingresses:
		return &networkingv1beta1.Ingress{} // nolint: exhaustivestruct
	case Secrets:
		return &corev1.Secret{} // nolint: exhaustivestruct
	case EndPoints:
		return &corev1.Endpoints{} // nolint: exhaustivestruct
	case NetworkPolicies:
		return &networkingv1.NetworkPolicy{} // nolint: exhaustivestruct
	case Unknown, ETCD:
		return nil
	default:

		return nil
	}
}

// nolint: clyclop
func (resourceType *ResourceType) K8SAPIVersion() string {
	switch *resourceType {
	case Unknown, ETCD:
		return ""
	case Deployments, DaemonSets, ReplicaSets, StatefulSets:
		return "apps/v1"
	case Hpas:
		return "autoscaling/v1"
	case Jobs:
		return "batch/v1"
	case CronJobs:
		return "batch/v1beta1"
	case Pods, Services, Nodes, Namespaces, ConfigMaps, PersistentVolumes, PersistentVolumeClaims, Secrets, EndPoints:
		return "core/v1"
	case NetworkPolicies:
		return "networking/v1"
	case Ingresses:
		return "networking/v1beta1"
	default:
		return ""
	}
}

func (resourceType *ResourceType) IsNamespaceScopedResource() bool {
	switch *resourceType {
	case Pods, Deployments, Services, Hpas, Namespaces, DaemonSets, ReplicaSets, StatefulSets, ConfigMaps, PersistentVolumeClaims, CronJobs, Jobs, Ingresses, Secrets, EndPoints, NetworkPolicies:
		return true
	case Unknown, Nodes, ETCD, PersistentVolumes:
		return false
	default:
		return false
	}
}

func (resourceType *ResourceType) APIGroup() string {
	switch *resourceType {
	case Pods, Services, Nodes, Namespaces, ConfigMaps, PersistentVolumes, PersistentVolumeClaims, Secrets, EndPoints:
		return ""
	case Ingresses, NetworkPolicies:
		return "networking"
	case Unknown, ETCD:
		return ""
	case Deployments, DaemonSets, ReplicaSets, StatefulSets:
		return "apps"
	case Hpas:
		return "autoscaling"
	case CronJobs, Jobs:
		return "batch"
	default:
		return ""
	}
}

func (resourceType *ResourceType) IsK8SPingResource() bool {
	switch *resourceType {
	case Unknown, Pods:
		return true
	case Deployments, Services, Hpas, Nodes, ETCD, Namespaces, DaemonSets, ReplicaSets, StatefulSets, ConfigMaps, PersistentVolumes, PersistentVolumeClaims, CronJobs, Jobs, Ingresses, Secrets, EndPoints, NetworkPolicies:
		return false
	default:
		return true
	}
}
