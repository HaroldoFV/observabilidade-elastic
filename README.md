# Projeto Go com Elastic Stack, APM e Kibana

## Visão Geral

Este projeto é uma aplicação simples em Go que inclui um servidor HTML básico. Ele está integrado com o Elastic Stack para coletar e visualizar logs e métricas de desempenho da aplicação usando Filebeat, Elasticsearch, APM e Kibana.
 

## Pré-requisitos

- Docker
- Docker Compose

## Serviços

- **Aplicação Go**: Um servidor HTTP simples escrito em Go.
- **Nginx**: Atua como um proxy reverso para a aplicação Go.
- **Elasticsearch**: Armazena logs e dados de APM.
- **Kibana**: Fornece uma interface web para pesquisa e visualização de logs e dados de APM.
- **APM Server**: Coleta dados de desempenho da aplicação Go.
- **Filebeat**: Encaminha dados de logs para o Elasticsearch.
- **Metricbeat**: Coleta e encaminha métricas do sistema para o Elasticsearch.
- **Heartbeat**: Monitora a disponibilidade da aplicação Go.

 
