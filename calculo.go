package main

import (
	"math"
	"sync"
	"time"
)

func calculoPiComThreads(numeroThreads int) (float64, time.Duration) {
	somasParciais := make([]float64, numeroThreads)
	tentativas := math.Pow(10, 6)
	intervaloPorBloco := int(tentativas) / numeroThreads
	tempoInicial := time.Now()
	posicaoInicialDoBloco := 0
	posicaoFinalDoBloco := 0

	var wg sync.WaitGroup
	wg.Add(numeroThreads)
	for i := 0; i < numeroThreads; i++ {
		if i < numeroThreads-1 {
			posicaoFinalDoBloco += intervaloPorBloco
		} else {
			posicaoFinalDoBloco = int(tentativas) - 1
		}
		go somaPorBloco(posicaoInicialDoBloco, posicaoFinalDoBloco, i, somasParciais, &wg)
		posicaoInicialDoBloco = posicaoFinalDoBloco + 1
	}
	wg.Wait()
	var soma float64
	for _, value := range somasParciais {
		soma += value
	}

	pi := 4 * soma

	tempoQueLevou := time.Now().Sub(tempoInicial)

	return pi, tempoQueLevou
}

func somaPorBloco(posicaoInicial, posicaoFinal, posicao int, somasParciais []float64, wg *sync.WaitGroup) {
	var soma float64
	for ; posicaoInicial <= posicaoFinal; posicaoInicial++ {
		soma += math.Pow(-1, float64(posicaoInicial)) / float64(2*posicaoInicial+1)
	}
	somasParciais[posicao] = soma
	wg.Done()
}
