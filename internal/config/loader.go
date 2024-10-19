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
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// LoadConfig lê o arquivo de configuração YAML e carrega as informações no QueueStore.
func LoadConfig() error {
	// Em produção, esperamos que o arquivo config.yaml esteja na raiz
	executable, err := os.Executable()
	if err != nil {
		log.Fatal("Erro ao obter o caminho do binário:", err)
	}

	path := filepath.Join(filepath.Dir(executable), "config.yaml")

	// Ler o conteúdo do arquivo YAML
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Erro ao ler o arquivo %s: %v", path, err)
	}

	var configs []Config
	err = yaml.Unmarshal(yamlFile, &configs)
	if err != nil {
		return fmt.Errorf("Erro ao fazer unmarshal do YAML: %v", err)
	}

	// Obtém a instância do QueueStore e adiciona as configurações
	store := GetQueueStore()

	for _, config := range configs {
		// Adiciona a configuração diretamente ao store
		store.AddConfig(config)
	}

	return nil
}
