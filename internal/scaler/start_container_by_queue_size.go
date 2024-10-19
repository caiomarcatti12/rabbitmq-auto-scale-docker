package scaler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/caiomarcatti12/rabbitmq-container-autoscaler/internal/config"
	"github.com/caiomarcatti12/rabbitmq-container-autoscaler/internal/docker"
	"github.com/caiomarcatti12/rabbitmq-container-autoscaler/internal/docker/container_store"
)

// Estrutura da resposta da API de filas do RabbitMQ
type QueueInfo struct {
	Name      string `json:"name"`
	Messages  int    `json:"messages"`
	Consumers int    `json:"consumers"`
	State     string `json:"state"`
}

// UpdateQueuesStatus inicia o processo contínuo de verificação dos containers.
func UpdateQueuesStatus() {
	store := config.GetQueueStore()
	configs := store.GetAllConfigs()

	for _, cfg := range configs {
		queues, err := fetchRabbitMQQueues(cfg.Auth)
		if err != nil {
			log.Printf("Erro ao buscar filas do RabbitMQ para '%s': %v\n", cfg.Auth.Host, err)
			continue
		}
		processQueues(cfg, queues)
	}
}

// Função para buscar as filas do RabbitMQ via API
func fetchRabbitMQQueues(auth config.Auth) ([]QueueInfo, error) {
	url := fmt.Sprintf("%s://%s:%s@%s:%d%s/api/queues", auth.Protocol, auth.Username, auth.Password, auth.Host, auth.Port, auth.Path)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição HTTP: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta HTTP: %v", err)
	}

	var queues []QueueInfo
	err = json.Unmarshal(body, &queues)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer unmarshal do JSON: %v", err)
	}

	return queues, nil
}

// Função para processar as filas e atualizar o status
func processQueues(cfg config.Config, queues []QueueInfo) {
	for _, q := range queues {
		for _, queue := range cfg.Queues {
			if q.Name == queue.Name {
				containerName := queue.ContainerName
				containerService, exists := container_store.GetByContainerName(containerName)

				if !exists {
					continue
				}

				if !containerService.IsActive && q.Messages > 0 {
					_, err := docker.StartContainer(containerName)
					if err != nil {
						log.Printf("Erro ao iniciar container '%s': %v", containerName, err)
						return
					}

					containerService.IsActive = true
					container_store.Update(*containerService)
				}

				if q.Messages > 0 {
					container_store.UpdateAccessTime(containerService.ID)
				}
			}
		}
	}
}
