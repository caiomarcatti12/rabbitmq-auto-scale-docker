/*
 * Copyright 2023 Caio Matheus Marcatti Calimério
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package container_store

import (
	"time"
)

var containers = make(map[string]Container)
var containersBySvc = make(map[string]Container)

func Add(container Container) {
	containers[container.ID] = container
	containersBySvc[container.ContainerName] = container
}
func Update(container Container) {
	containers[container.ID] = container
	containersBySvc[container.ContainerName] = container
}

func Remove(containerID string) {
	if container, exists := containers[containerID]; exists {
		delete(containersBySvc, container.ContainerName)
		delete(containers, containerID)
	}
}

func GetByID(containerID string) (Container, bool) {
	container, exists := containers[containerID]
	return container, exists
}

func GetByContainerName(serviceName string) (*Container, bool) {
	container, exists := containersBySvc[serviceName]
	if !exists {
		return nil, false
	}
	return &container, true
}

func UpdateAccessTime(containerID string) {
	if container, exists := containers[containerID]; exists {
		container.LastAccess = time.Now()
		containers[containerID] = container
		containersBySvc[container.ContainerName] = container
	}
}

func GetAll() map[string]Container {
	return containers
}
