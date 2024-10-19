package docker

import (
	"log"
	"sync"
	"time"

	"github.com/caiomarcatti12/rabbitmq-container-autoscaler/internal/config"
	"github.com/caiomarcatti12/rabbitmq-container-autoscaler/internal/docker/container_store"
)

var (
	containerMonitorMutex sync.Mutex
)

// CheckContainersToStop inicia o processo contínuo de monitoramento e parada de containers inativos.
func CheckContainersToStop() {
	for {
		monitorAndStopContainers()
		time.Sleep(5 * time.Second)
	}
}

// monitorAndStopContainers monitora e para containers que estão inativos além do tempo limite.
func monitorAndStopContainers() {
	containerMonitorMutex.Lock()
	defer containerMonitorMutex.Unlock()

	now := time.Now()

	// Obtém todas as configurações do QueueStore
	store := config.GetQueueStore()
	configs := store.GetAllConfigs()

	for _, cfg := range configs {
		for _, queue := range cfg.Queues {
			container, _ := container_store.GetByContainerName(queue.ContainerName)

			if container != nil {
				log.Printf("Verificando contêiner '%s' associado à fila '%s'. Último acesso: %v\n", queue.ContainerName, queue.Name, container.LastAccess)
				checkAndStopContainer(*container, now, queue)
			} else {
				log.Printf("Contêiner associado à fila '%s' não encontrado.\n", queue.Name)
			}
		}
	}
}

// checkAndStopContainer verifica se o container deve ser parado com base no TTL.
func checkAndStopContainer(container container_store.Container, now time.Time, queue config.Queue) {
	if isContainerExpired(container, now) {
		log.Printf("Contêiner '%s' inativo por mais de 10s. Iniciando parada.\n", container.ID)
		stopAndRemoveContainer(container)
	} else {
		log.Printf("Contêiner '%s' ainda está ativo. Nenhuma ação necessária.\n", container.ID)
	}
}

// isContainerExpired verifica se o container excedeu o tempo de inatividade permitido.
func isContainerExpired(container container_store.Container, now time.Time) bool {
	ttl := 10 * time.Second
	return now.Sub(container.LastAccess) > ttl && container.IsActive
}

// stopAndRemoveContainer para e remove o container da store.
func stopAndRemoveContainer(container container_store.Container) {
	log.Printf("Parando contêiner '%s'.", container.ID)
	StopContainer(container.ID)

	container.IsActive = false
	container_store.Update(container)
	log.Printf("Contêiner '%s' parado e marcado como inativo.\n", container.ID)
}
