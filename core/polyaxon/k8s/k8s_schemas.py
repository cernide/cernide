# All schemas used by Polyaxon

from kubernetes import client

V1Affinity = client.V1Affinity
V1SecurityContext = client.V1SecurityContext
V1PodDNSConfig = client.V1PodDNSConfig
V1Toleration = client.V1Toleration
V1HostAlias = client.V1HostAlias
V1Container = client.V1Container
V1EnvVar = client.V1EnvVar
V1VolumeMount = client.V1VolumeMount
V1ContainerPort = client.V1ContainerPort
V1ResourceRequirements = client.V1ResourceRequirements
V1EnvFromSource = client.V1EnvFromSource
V1Volume = client.V1Volume
V1ObjectFieldSelector = client.V1ObjectFieldSelector
V1EnvVarSource = client.V1EnvVarSource
V1ConfigMapKeySelector = client.V1ConfigMapKeySelector
V1SecretKeySelector = client.V1SecretKeySelector
V1PodSpec = client.V1PodSpec
V1ObjectMeta = client.V1ObjectMeta
V1PodTemplateSpec = client.V1PodTemplateSpec
V1HostPathVolumeSource = client.V1HostPathVolumeSource
V1EmptyDirVolumeSource = client.V1EmptyDirVolumeSource
V1PersistentVolumeClaimVolumeSource = client.V1PersistentVolumeClaimVolumeSource
V1SecretVolumeSource = client.V1SecretVolumeSource
V1ConfigMapVolumeSource = client.V1ConfigMapVolumeSource
V1LocalObjectReference = client.V1LocalObjectReference
