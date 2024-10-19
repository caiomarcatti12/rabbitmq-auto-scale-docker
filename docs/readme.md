# RabbitMQ Auto-Scale Docker

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](license) ![Static Badge](https://img.shields.io/badge/N%C3%A3o%20pronto%20para%20produ%C3%A7%C3%A3o-red)

## Introdução

O RabbitMQ Auto-Scale Docker é uma solução projetada para gerenciar o escalonamento automático de contêineres RabbitMQ em um ambiente Docker. Inspirado na capacidade de sistemas distribuídos de se adaptarem dinamicamente às demandas, este projeto visa fornecer uma alternativa escalável e eficiente para operações de RabbitMQ em ambientes diversos, incluindo produção, homologação e desenvolvimento.

## Características Principais

- **Escalonamento Automático de RabbitMQ**: Inicia e para contêineres RabbitMQ dinamicamente com base na carga e nas mensagens na fila.
- **Eficiência de Recursos**: Otimiza o uso de recursos ao desligar contêineres ociosos após um período de inatividade.
- **Configuração Flexível**: Permite a configuração detalhada para ajustes personalizados de escalonamento e gerenciamento de filas.

## Configuração do RabbitMQ Auto-Scale

O projeto utiliza um arquivo de configuração YAML para definir as regras de escalonamento, rotas e outras opções específicas. Consulte o guia detalhado disponível em [Configuração](./config.md) para entender como personalizar e configurar seu ambiente de auto-escalonamento de RabbitMQ.

## Como Contribuir

Estamos sempre abertos a contribuições! Se você deseja ajudar a melhorar o projeto, seja através de correções de bugs, melhorias ou novas funcionalidades, siga nosso [Guia de Contribuição](./contributing.md) para entender o processo e garantir que sua contribuição seja integrada da melhor forma possível.

## Código de Conduta

Estamos comprometidos em proporcionar uma comunidade acolhedora e inclusiva para todos. Esperamos que todos os participantes do projeto sigam nosso [Código de Conduta](./code_of_conduct.md). Pedimos que leia e siga estas diretrizes para garantir um ambiente respeitoso e produtivo para todos os colaboradores.

## Licença

Este projeto está licenciado sob a licença Apache 2.0. Consulte o arquivo [LICENSE](./license) para obter detalhes.
