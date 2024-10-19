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
package config

import (
	"errors"
	"sync"
)

type QueueStore struct {
	configs []Config
}

var (
	once     sync.Once
	instance *QueueStore
)

// GetQueueStore retorna a instância Singleton de QueueStore.
func GetQueueStore() *QueueStore {
	once.Do(func() {
		instance = &QueueStore{
			configs: []Config{},
		}
	})
	return instance
}

// AddConfig adiciona uma nova configuração de autenticação e filas ao QueueStore.
func (qs *QueueStore) AddConfig(config Config) {
	qs.configs = append(qs.configs, config)
}

// UpdateConfig atualiza uma configuração existente baseada no Host da Auth.
func (qs *QueueStore) UpdateConfig(config Config) error {
	for i, c := range qs.configs {
		if c.Auth.Host == config.Auth.Host {
			qs.configs[i] = config
			return nil
		}
	}
	return errors.New("config not found")
}

// GetConfig retorna uma configuração de autenticação e filas com base no Host.
func (qs *QueueStore) GetConfig(authHost string) (Config, error) {
	for _, config := range qs.configs {
		if config.Auth.Host == authHost {
			return config, nil
		}
	}
	return Config{}, errors.New("config not found")
}

// GetAllConfigs retorna todas as configurações armazenadas no QueueStore.
func (qs *QueueStore) GetAllConfigs() []Config {

	return qs.configs
}
