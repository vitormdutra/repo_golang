package main

import (
	"bank-system/clientes"
	"bank-system/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	clienteDoRobson := clientes.Titular{Nome: "Robson", CPF: "0974263152", Profissao: "Professor"}
	contaDoRobson := contas.ContaCorrente{Titular: clienteDoRobson, NumeroAgencia: 2345, NumeroConta: 44}
	contaDoMarcelo := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Marcelo",
		CPF:       "123456789",
		Profissao: "Padeiro",
	}, NumeroAgencia: 1234,
		NumeroConta: 1241,
	}

	contaDoDenis := contas.ContaPoupanca{Titular: clientes.Titular{
		Nome:      "Denis",
		CPF:       "4124155",
		Profissao: "Motorista",
	}, NumeroAgencia: 4434,
		NumeroConta: 124411,
	}

	fmt.Println(contaDoMarcelo)
	fmt.Println(contaDoRobson)
	fmt.Println(contaDoDenis)

	contaDoDenis.Depositar(100)
	fmt.Println(contaDoDenis.ObterSaldo())
	PagarBoleto(&contaDoDenis, 40)

	fmt.Println(contaDoDenis.ObterSaldo())
}
