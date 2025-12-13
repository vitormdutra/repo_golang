package contas

import "bank-system/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0
	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Valor Depositado", c.saldo
	} else {
		return "Valor insuficiente", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDeTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDeTransferencia < c.saldo && valorDeTransferencia > 0 {
		c.saldo -= valorDeTransferencia
		contaDestino.Depositar(valorDeTransferencia)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
