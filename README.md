# RESULTADOS

## Desenvolvedores:

- Lucas Rodrigues Santos
- Luis Carlos da Silva

## Descrição

Este repositório contém um algoritmo em Go responsável por calcular o valor de PI utilizando múltiplas threads.

Um fato legal para se observar é a variação no valor de PI com a mudança no número de threads, isso
ocorre por conta da inconsistência presente nos números decimais na computação e por consequência acaba nos fazendo
observar uma certa inconsistência nos valores, como você pode observar no exemplo abaixo após `3.141591653589…`.

Outra coisa também muito interessante é o tempo de execução para realizar os cálculos (irá variar de computador para
computador e até mesmo a depender da quantidade de processos executando na sua máquina naquele momento). Vale observar
que até certo ponto quanto mais threads você colocar para o programa executar mais rápido ele será executado, porém como
citado anteriormente, em certo ponto o programa começa a demorar mais a executar com o aumento do número de threads
ocorrendo isso por conta da quantidade de trocas de contexto necessárias para o processador finalizar a aplicação,
como você pode observar no exemplo abaixo onde o tempo de execução em média vai diminuindo até 8 threads, mas aumenta
a partir daí.

## Verificando os dados

Para executar o programa você precisará apenas ter o Go instalado na sua máquina, baixar esse repositório e por fim, na
pasta do projeto, executar:

```shell
$ go run .
```

E verá um relatório bastante parecido com esse:

```
Número de Threads: 2
Média do valor de PI: 3.14159165358969216796936052560340613126754760742188
Desvio Padrão do valor de PI: 0.000000
Média do tempo de Execução: 48.028687ms
Desvio Padrão do tempo de Execução: 0.132874ms
Por execução [{<Valor de PI> <Tempo de Execução do cálculo>}]: [{3.141591653589692 47.917891ms} {3.141591653589692 47.973273ms} {3.141591653589692 47.88845ms} {3.141591653589692 48.12575ms} {3.141591653589692 48.238074ms}]



Número de Threads: 4
Média do valor de PI: 3.14159165358978098581133053812664002180099487304688
Desvio Padrão do valor de PI: 0.000000
Média do tempo de Execução: 25.364401ms
Desvio Padrão do tempo de Execução: 0.234133ms
Por execução [{<Valor de PI> <Tempo de Execução do cálculo>}]: [{3.141591653589781 25.63285ms} {3.141591653589781 25.20203ms} {3.141591653589781 25.044445ms} {3.141591653589781 25.630113ms} {3.141591653589781 25.312567ms}]



Número de Threads: 8
Média do valor de PI: 3.14159165358972769510614853061269968748092651367188
Desvio Padrão do valor de PI: 0.000000
Média do tempo de Execução: 19.910608ms
Desvio Padrão do tempo de Execução: 1.259495ms
Por execução [{<Valor de PI> <Tempo de Execução do cálculo>}]: [{3.1415916535897277 22.396846ms} {3.1415916535897277 19.573149ms} {3.1415916535897277 18.9932ms} {3.1415916535897277 19.159631ms} {3.1415916535897277 19.430217ms}]



Número de Threads: 16
Média do valor de PI: 3.14159165358976455451056608580984175205230712890625
Desvio Padrão do valor de PI: 0.000000
Média do tempo de Execução: 20.334142ms
Desvio Padrão do tempo de Execução: 1.353678ms
Por execução [{<Valor de PI> <Tempo de Execução do cálculo>}]: [{3.1415916535897646 20.070359ms} {3.1415916535897646 19.18091ms} {3.1415916535897646 19.45155ms} {3.1415916535897646 22.956976ms} {3.1415916535897646 20.010919ms}]



Número de Threads: 32
Média do valor de PI: 3.14159165358979342030920633987989276647567749023438
Desvio Padrão do valor de PI: 0.000000
Média do tempo de Execução: 22.464246ms
Desvio Padrão do tempo de Execução: 1.486721ms
Por execução [{<Valor de PI> <Tempo de Execução do cálculo>}]: [{3.1415916535897934 25.344607ms} {3.1415916535897934 22.150697ms} {3.1415916535897934 21.784516ms} {3.1415916535897934 21.975857ms} {3.1415916535897934 21.065553ms}]
```