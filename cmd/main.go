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
package main

import (
	"log"
	"time"

	"github.com/caiomarcatti12/rabbitmq-container-autoscaler/internal/config"
	"github.com/caiomarcatti12/rabbitmq-container-autoscaler/internal/docker"
	"github.com/caiomarcatti12/rabbitmq-container-autoscaler/internal/scaler"
)

func main() {
	err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	go docker.CheckContainersActive()
	go docker.CheckContainersToStop()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	go scaler.UpdateQueuesStatus()

	for range ticker.C {
		go scaler.UpdateQueuesStatus()
	}
}
