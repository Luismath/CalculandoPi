package main

import (
	"fmt"
	"math"
	"time"
)

var NUMERO_DE_VERIFICACOES = 5

type Execucao struct {
	Pi            float64
	TempoExecucao time.Duration
}

func relatorio(numeroThreads int) {
	var execucao []Execucao
	var somaDosValoresDePi float64
	var somaDosTemposDeExecucao time.Duration

	for i := 0; i < NUMERO_DE_VERIFICACOES; i++ {
		pi, tempoQueLevou := calculoPiComThreads(numeroThreads)
		execucaoAtual := Execucao{
			Pi:            pi,
			TempoExecucao: tempoQueLevou,
		}
		execucao = append(execucao, execucaoAtual)
		somaDosValoresDePi += pi
		somaDosTemposDeExecucao += tempoQueLevou
	}

	valorMedioDePi := somaDosValoresDePi / float64(NUMERO_DE_VERIFICACOES)
	tempoMedioDeExecucao := int(somaDosTemposDeExecucao.Nanoseconds()) / NUMERO_DE_VERIFICACOES
	tempoMedioDeExecucaoEmMilissegundos := float64(tempoMedioDeExecucao) / 1000000

	var numeradorDoDesvioPadraoDePi float64
	var numeradorDoDesvioPadraoDoTempoDeExecucao float64
	for _, execucaoAtual := range execucao {
		numeradorDoDesvioPadraoDePi += math.Pow(execucaoAtual.Pi-valorMedioDePi, 2)
		numeradorDoDesvioPadraoDoTempoDeExecucao += math.Pow(float64(int(execucaoAtual.TempoExecucao.Nanoseconds())-tempoMedioDeExecucao), 2)
	}

	desvioPadraoDosValoresDePi := math.Sqrt(numeradorDoDesvioPadraoDePi / float64(NUMERO_DE_VERIFICACOES))
	desvioPadraoDosTemposDeExecucao := math.Sqrt(numeradorDoDesvioPadraoDoTempoDeExecucao / float64(NUMERO_DE_VERIFICACOES))
	desvioPadraoDosTemposDeExecucaoEmMilissegundos := float64(desvioPadraoDosTemposDeExecucao) / 1000000

	fmt.Printf("Número de Threads: %d\n", numeroThreads)
	fmt.Printf("Média do valor de PI: %.50f\n", valorMedioDePi)
	fmt.Printf("Desvio Padrão do valor de PI: %f\n", desvioPadraoDosValoresDePi)
	fmt.Printf("Média do tempo de Execução: %fms\n", tempoMedioDeExecucaoEmMilissegundos)
	fmt.Printf("Desvio Padrão do tempo de Execução: %fms\n", desvioPadraoDosTemposDeExecucaoEmMilissegundos)
	fmt.Printf("Por execução [{<Valor de PI> <Tempo de Execução do cálculo>}]: %v\n", execucao)
}
