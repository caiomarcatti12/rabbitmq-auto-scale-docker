# Configurações do RabbitMQ Auto-Scale Docker

Este documento explica como configurar a autenticação e filas para o projeto RabbitMQ Auto-Scale Docker. Todas as configurações são definidas em um arquivo YAML e permitem personalizar o comportamento do escalonador e das conexões RabbitMQ.

## Estrutura de Configuração

### Exemplo de Configuração:

```yaml
auth:
  protocol: http
  host: host.docker.internal
  port: 15672 
  username: root
  password: password
  vhost: /
queues:
  - name: teste
    containerName: hello-container
```

### Seção de Autenticação (`auth`)

- **protocol**: Define o protocolo de comunicação utilizado para acessar o RabbitMQ (ex.: `http`, `https`).
- **host**: O endereço do servidor RabbitMQ. `host.docker.internal` é frequentemente usado para acessar o host da máquina a partir de contêineres Docker.
- **port**: A porta na qual o RabbitMQ está disponível (ex.: `15672` para a interface de gerenciamento).
- **username**: Nome de usuário para autenticação no RabbitMQ.
- **password**: Senha associada ao usuário para autenticação.
- **vhost**: O Virtual Host do RabbitMQ que será utilizado (ex.: `/` para o virtual host padrão).

### Seção de Filas (`queues`)

Esta seção define uma lista de filas que serão monitoradas e associadas aos contêineres Docker. 

- **name**: O nome da fila RabbitMQ.
- **containerName**: O nome do contêiner Docker que será escalado automaticamente com base na atividade na fila definida.

Cada item na lista de `queues` representa uma associação entre uma fila RabbitMQ e um contêiner específico, permitindo o gerenciamento automatizado do escalonamento com base na demanda de mensagens.
