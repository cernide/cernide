/*
Copyright 2018-2020 Polyaxon, Inc.

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
	"strings"
	"time"

	operationv1 "github.com/polyaxon/polyaxon/operator/api/v1"
	"github.com/polyaxon/polyaxon/operator/controllers/config"
	"github.com/polyaxon/polyaxon/operator/controllers/plugins"
)

const (
	apiServerDefaultTimeout = 35 * time.Second
)

func (r *OperationReconciler) instanceSyncStatus(instance *operationv1.Operation) error {
	lastCond := instance.Status.Conditions[len(instance.Status.Conditions)-1]
	return r.syncStatus(instance, lastCond)
}

func (r *OperationReconciler) syncStatus(instance *operationv1.Operation, statusCond operationv1.OperationCondition) error {
	if !config.GetBoolEnv("POLYAXON_AGENT_ENABLED", true) || !instance.SyncStatuses {
		return nil
	}

	log := r.Log

	log.Info("Operation sync status", "Syncing", instance.GetName())

	instanceName, ok := instance.ObjectMeta.Labels["app.kubernetes.io/instance"]
	if !ok || instanceName == "" {
		log.Info("Operation cannot be synced", "Instance", instance.Name, "Does not exist", instance.GetName())
		return nil
	}
	jobName := strings.Split(instanceName, ".")
	if len(jobName) != 4 {
		log.Info("Operation cannot be synced", "Instance", instance.Name, "Job name is not valid", instanceName)
		return nil
	}

	return plugins.LogPolyaxonRunStatus(jobName[0], jobName[1], jobName[3], statusCond, r.Log)
}

func (r *OperationReconciler) notify(instance *operationv1.Operation) error {

	if !config.GetBoolEnv("POLYAXON_AGENT_ENABLED", true) || len(instance.Notifications) == 0 {
		return nil
	}

	log := r.Log

	log.Info("Operation notify status", "Notifying", instance.GetName())

	instanceName, ok := instance.ObjectMeta.Labels["app.kubernetes.io/instance"]
	if !ok || instanceName == "" {
		log.Info("Operation cannot be notified", "Instance", instance.Name, "Does not exist", instance.GetName())
		return nil
	}
	jobName := strings.Split(instanceName, ".")
	if len(jobName) != 4 {
		log.Info("Operation cannot be notified", "Instance", instance.Name, "Job name is not valid", instanceName)
		return nil
	}

	name, ok := instance.ObjectMeta.Labels["app.kubernetes.io/name"]
	if !ok || instanceName == "" {
		name = ""
	}

	if len(instance.Status.Conditions) == 0 {
		log.Info("Operation cannot be notified", "Instance", instance.Name, "No conditions", instance.GetName())
		return nil
	}
	lastCond := instance.Status.Conditions[len(instance.Status.Conditions)-1]

	connections := []string{}
	for _, notification := range instance.Notifications {
		if notification.Trigger == operationv1.OperationDoneTrigger || operationv1.OperationConditionType(notification.Trigger) == lastCond.Type {
			connections = append(connections, notification.Connection)
		}
	}

	if len(connections) == 0 {
		log.Info("Operation cannot be notified", "Instance", instance.Name, "No connections", instance.Notifications)
		return nil
	}

	log.Info("Operation notify status", "Status", lastCond.Type, "Instance", instance.GetName())
	return plugins.NotifyPolyaxonRunStatus(instance.Namespace, name, jobName[0], jobName[1], jobName[3], lastCond, connections, r.Log)
}

func (r *OperationReconciler) collectLogs(instance *operationv1.Operation) error {

	if !config.GetBoolEnv("POLYAXON_AGENT_ENABLED", true) || !instance.CollectLogs {
		return nil
	}

	log := r.Log

	instanceName, ok := instance.ObjectMeta.Labels["app.kubernetes.io/instance"]
	if !ok || instanceName == "" {
		log.Info("Operation cannot be synced", "Instance", instance.Name, "Does not exist", instance.GetName())
		return nil
	}
	jobName := strings.Split(instanceName, ".")
	if len(jobName) != 4 {
		log.Info("Operation cannot be synced", "Instance", instance.Name, "Job name is not valid", instanceName)
		return nil
	}

	log.Info("Operation collect logs", "Instance", instance.GetName())
	return plugins.CollectPolyaxonRunLogs(instance.Namespace, jobName[0], jobName[1], jobName[3], r.Log)
}
