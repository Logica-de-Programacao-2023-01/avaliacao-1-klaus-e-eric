package q4

import "errors"

//Uma loja virtual de roupas recebeu várias listas de produtos vendidos em diferentes dias da semana. O dono da loja
//deseja analisar as listas para entender melhor o comportamento de suas vendas. Para isso, ele precisa classificar cada
//lista como em ordem crescente, decrescente ou aleatória, de acordo com o preço dos produtos.
//
//Você deve implementar uma função que recebe uma lista de preços e retorna um valor inteiro indicando se a lista está em
//ordem crescente, decrescente ou aleatória. A função deve retornar 1 se a lista estiver em ordem crescente, 2 se a lista
//estiver em ordem decrescente e 3 se a lista estiver aleatória. A função deve retornar um erro se a lista estiver vazia.
//Caso a lista possua apenas um elemento, a função deve retornar 3.

func ClassifyPrices(prices []int) (int, error) {
	len_lista := len(prices)

	if len_lista == 0 {
		return 0, errors.New("Lista Vazia")
	}
	if len_lista == 1 {
		return 3, nil
	}

	var cres bool = true
	var dec bool = true

	for i := 1; i < len_lista; i++ {
		if prices[i] > prices[i-1] {
			dec = false
		} else if prices[i] < prices[i-1] {
			cres = false
		}
	}
	if cres {
		return 1, nil
	} else if dec {
		return 2, nil
	} else {
		return 3, nil
	}
}
