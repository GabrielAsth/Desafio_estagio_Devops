package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Response representa a estrutura do JSON exigida pelo desafio
type Response struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

// Criando as métricas obrigatórias exigidas pelo desafio
var (
	// Volume de requisições (Contador)
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Volume total de requisicoes recebidas no endpoint /projeto-korp",
		},
		[]string{"path", "status"},
	)
)

func init() {
	// Registra a métrica no coletor padrão do Prometheus
	prometheus.MustRegister(httpRequestsTotal)
}

func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httpRequestsTotal.WithLabelValues("/projeto-korp", "450").Inc()
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	currentUTC := time.Now().UTC().Format(time.RFC3339)

	response := Response{
		Nome:    "Projeto Korp",
		Horario: currentUTC,
	}

	// Incrementa a métrica de volume de requisições com sucesso (Status 200)
	httpRequestsTotal.WithLabelValues("/projeto-korp", "200").Inc()

	json.NewEncoder(w).Encode(response)
}

func main() {
	// Endpoint principal do desafio
	http.HandleFunc("/projeto-korp", projetoKorpHandler)

	// Endpoint obrigatório do Prometheus para coletar as métricas (Disponibilidade e Volume)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Servidor iniciado na porta 8080 com suporte ao Prometheus...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}
