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

	path := filepath.Join(filepath.Dir(executable), "../config.yaml")

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
