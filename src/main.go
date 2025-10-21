package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8") // Adicionei charset para melhor compatibilidade
	fmt.Fprintf(w, "No mar da nuvem, onde o código a bailar, Kubernetes surge, para orquestrar. Um mestre timoneiro, com poder e precisão, A gerir contêineres, em vasta imensidão.

Do Pod que nasce, a menor unidade, À Deployment que garante estabilidade. O ReplicaSet zelando pelo número certo, De instâncias ativas, mantendo o concerto.

O Master é a mente, o cérebro a pensar, Com API Server, a todos a escutar. O etcd, memória, chave-valor guardiã, Da frota de nós, na manhã e na manhã.

O Kubelet no Node, vigilante e fiel, Executa o que lhe manda o alto céu. E o Controller Manager, com laços de amor, Corrige desvios, com todo o vigor.

Do YAML que escrevo, a regra e o querer, Ele o traduz em ação, faz o serviço a correr. Service expõe a porta, o tráfego a fluir, Para o mundo externo, o app a se exibir.

Seja stateless, seja o StatefulSet a mandar, Dados persistentes ele sabe guardar. Volumes e Secrets, ConfigMaps no lugar, A essência da app, ele sabe configurar.

Com escala elástica, cresce e pode encolher, Segundo a demanda, para não se exceder. Em cada cluster, um universo a operar, A magia do K8s, a nos fascinar.")
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe("0.0.0.0:5000", nil))
}