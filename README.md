# pipeline042

![CI/CD Status](https://github.com/vinidows/pipeline042/actions/workflows/main.yml/badge.svg)
[![Cluster K8s](https://img.shields.io/badge/DigitalOcean-Kubernetes-0080FF?style=for-the-badge&logo=digitalocean)](https://cloud.digitalocean.com/kubernetes/clusters)

> **NOTA:** Substitua `main.yml` no *badge* acima pelo nome real do seu arquivo de *workflow* (se for diferente) para que ele funcione corretamente.

Visão Geral
Este repositório é dedicado à automação do processo de build e deploy (CI/CD) de uma aplicação, orquestrada via Kubernetes (K8s) e hospedada no DigitalOcean Kubernetes Service (DOKS).

O pipeline de CI/CD é gerenciado pelo GitHub Actions, garantindo que qualquer alteração na branch principal seja automaticamente construída, testada e implantada no cluster DOKS.

Infraestrutura: DigitalOcean Kubernetes Service (DOKS)

Contêinerização: Docker

Orquestração: Kubernetes (K8s)

CI/CD: GitHub Actions

Configs K8s: Arquivos de manifesto YAML

📦 Estrutura do Repositório
.
├── .github/workflows/
│   └── main.yml        # Workflow do GitHub Actions (Pipeline CI/CD)
├── k8s/
│   ├── deployment.yaml # Configuração do Deployment da aplicação
│   ├── service.yaml    # Configuração do Service (ex: LoadBalancer)
│   └── [Opcional] ingress.yaml
├── [Opcional] Dockerfile   # Arquivo de construção da imagem da aplicação
└── README.md

⚙️ Pré-requisitosPara que o pipeline funcione e para gerenciar o cluster manualmente, os seguintes itens são necessários:DigitalOcean Account e Kubernetes Cluster (DOKS) criado.Registro de Contêiner (ex: Docker Hub ou DigitalOcean Container Registry).Ferramentas de linha de comando: kubectl e doctl.🔑 Configuração de Segredos (GitHub Secrets)O pipeline do GitHub Actions depende dos seguintes segredos (Secrets) configurados nas Settings do seu repositório para acessar a DigitalOcean e o Kubernetes:Nome do SegredoDescriçãoDO_ACCESS_TOKENToken de Acesso Pessoal da DigitalOcean com permissão para gerenciar K8s.DOCKER_USERNAMEUsuário do seu Registro de Contêiner.DOCKER_PASSWORDSenha ou Token de Acesso do seu Registro de Contêiner.KUBE_CONFIG_BASE64O conteúdo do seu arquivo kubeconfig do DOKS, codificado em Base64.

📦 Deploy (CI/CD)
O processo de deploy é disparado por meio do GitHub Actions a cada push na branch principal (main).

🔄 Fluxo do Pipeline
Build: A imagem Docker é construída a partir do Dockerfile.

Tagging: A imagem recebe uma tag única, geralmente baseada no SHA do commit.

Push: A imagem é enviada para o Registro de Contêiner.

Deploy (K8s):

O workflow se autentica no cluster DOKS usando o KUBE_CONFIG_BASE64.

O arquivo k8s/deployment.yaml é atualizado programaticamente (via script no workflow) para referenciar a nova tag da imagem.

O comando kubectl apply -f k8s/ é executado, iniciando um Rolling Update no cluster.

🔗 Acessando a Aplicação
A aplicação está exposta através do Service configurado em k8s/service.yaml.

Para obter o endereço IP público (LoadBalancer):

Bash

kubectl get services -n [SEU_NAMESPACE]
Procure pelo EXTERNAL-IP do seu Service.

URL da Aplicação: http://[EXTERNAL-IP]

Mantenha sempre os seus manifestos e o pipeline atualizados para refletir o estado desejado da sua infraestrutura.
