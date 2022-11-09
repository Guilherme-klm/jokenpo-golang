package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestDeveRetornarJogador1ComoVencedorQuandoJogada1ForPedraJogada2ForTesoura(t *testing.T) {

	jogador1 := "Alberto"
	jogador2 := "Josias"
	jogadaJogador1 := "PEDRA"
	jogadaJogador2 := "TESOURA"
	expectResult := jogador1

	ganhador := defineVencedor(jogadaJogador1, jogadaJogador2, jogador1, jogador2)

	if ganhador != expectResult {
		t.Error("Experava", expectResult, "mas retornou", ganhador)
	}

}
func TestDeveRetornarJogador2ComoVencedorQuandoJogada1ForPedraJogada2ForPapel(t *testing.T) {

	jogador1 := "Alberto"
	jogador2 := "Josias"
	jogadaJogador1 := "PEDRA"
	jogadaJogador2 := "PAPEL"
	expectResult := jogador2

	ganhador := defineVencedor(jogadaJogador1, jogadaJogador2, jogador1, jogador2)

	if ganhador != expectResult {
		t.Error("Experava", expectResult, "mas retornou", ganhador)
	}

}

func TestDeveRetornarEmpateQuandoJogada1ForPedraJogada2ForPedra(t *testing.T) {

	jogador1 := "Alberto"
	jogador2 := "Josias"
	jogadaJogador1 := "PEDRA"
	jogadaJogador2 := "PEDRA"
	expectResult := "Empate"

	ganhador := defineVencedor(jogadaJogador1, jogadaJogador2, jogador1, jogador2)

	if ganhador != expectResult {
		t.Error("Experava", expectResult, "mas retornou", ganhador)
	}

}

func TestDeveRetornarJogador1ComoVencedorQuandoJogada1ForPapelJogada2ForPedra(t *testing.T) {

	jogador1 := "Alberto"
	jogador2 := "Josias"
	jogadaJogador1 := "PAPEL"
	jogadaJogador2 := "PEDRA"
	expectResult := jogador1

	ganhador := defineVencedor(jogadaJogador1, jogadaJogador2, jogador1, jogador2)

	if ganhador != expectResult {
		t.Error("Experava", expectResult, "mas retornou", ganhador)
	}

}

func TestDeveRetornarJogador2ComoVencedorQuandoJogada1ForPapelJogada2ForTesoura(t *testing.T) {

	jogador1 := "Alberto"
	jogador2 := "Josias"
	jogadaJogador1 := "PAPEL"
	jogadaJogador2 := "TESOURA"
	expectResult := jogador2

	ganhador := defineVencedor(jogadaJogador1, jogadaJogador2, jogador1, jogador2)

	if ganhador != expectResult {
		t.Error("Experava", expectResult, "mas retornou", ganhador)
	}

}

func TestDeveRetornarEmpateQuandoJogada1ForPapelJogada2ForPapel(t *testing.T) {

	jogador1 := "Alberto"
	jogador2 := "Josias"
	jogadaJogador1 := "PAPEL"
	jogadaJogador2 := "PAPEL"
	expectResult := "Empate"

	ganhador := defineVencedor(jogadaJogador1, jogadaJogador2, jogador1, jogador2)

	if ganhador != expectResult {
		t.Error("Experava", expectResult, "mas retornou", ganhador)
	}

}

func TestDeveRetornarJogador1ComoVencedorQuandoJogada1ForTesouraJogada2ForPapel(t *testing.T) {

	jogador1 := "Alberto"
	jogador2 := "Josias"
	jogadaJogador1 := "TESOURA"
	jogadaJogador2 := "PAPEL"
	expectResult := jogador1

	ganhador := defineVencedor(jogadaJogador1, jogadaJogador2, jogador1, jogador2)

	if ganhador != expectResult {
		t.Error("Experava", expectResult, "mas retornou", ganhador)
	}

}

func TestDeveRetornarJogador2ComoVencedorQuandoJogada1ForTesouraJogada2ForPedra(t *testing.T) {

	jogador1 := "Alberto"
	jogador2 := "Josias"
	jogadaJogador1 := "TESOURA"
	jogadaJogador2 := "PEDRA"
	expectResult := jogador2

	ganhador := defineVencedor(jogadaJogador1, jogadaJogador2, jogador1, jogador2)

	if ganhador != expectResult {
		t.Error("Experava", expectResult, "mas retornou", ganhador)
	}

}

func TestDeveRetornarEmpateQuandoJogada1ForTesouraJogada2ForTesoura(t *testing.T) {

	jogador1 := "Alberto"
	jogador2 := "Josias"
	jogadaJogador1 := "TESOURA"
	jogadaJogador2 := "TESOURA"
	expectResult := "Empate"

	ganhador := defineVencedor(jogadaJogador1, jogadaJogador2, jogador1, jogador2)

	if ganhador != expectResult {
		t.Error("Experava", expectResult, "mas retornou", ganhador)
	}

}

func TestDeveRetornarJogador2ComoVencedorQuandoJogada2ForPedraJogada1ForTesoura(t *testing.T) {

	jogador1 := "Alberto"
	jogador2 := "Josias"
	jogadaJogador1 := "TESOURA"
	jogadaJogador2 := "PEDRA"
	expectResult := jogador2

	ganhador := defineVencedor(jogadaJogador1, jogadaJogador2, jogador1, jogador2)

	if ganhador != expectResult {
		t.Error("Experava", expectResult, "mas retornou", ganhador)
	}

}

func TestDeveRetornarJogador2ComoVencedorQuandoJogada2ForPapelJogada1ForPedra(t *testing.T) {

	jogador1 := "Alberto"
	jogador2 := "Josias"
	jogadaJogador1 := "PEDRA"
	jogadaJogador2 := "PAPEL"
	expectResult := jogador2

	ganhador := defineVencedor(jogadaJogador1, jogadaJogador2, jogador1, jogador2)

	if ganhador != expectResult {
		t.Error("Experava", expectResult, "mas retornou", ganhador)
	}

}

func TestDeveRetornarJogador2ComoVencedorQuandoJogada2ForTesouraJogada1ForPapel(t *testing.T) {

	jogador1 := "Alberto"
	jogador2 := "Josias"
	jogadaJogador1 := "PAPEL"
	jogadaJogador2 := "TESOURA"
	expectResult := jogador2

	ganhador := defineVencedor(jogadaJogador1, jogadaJogador2, jogador1, jogador2)

	if ganhador != expectResult {
		t.Error("Experava", expectResult, "mas retornou", ganhador)
	}

}

func TestDeveRetornarNomeDosDoisJogadores(t *testing.T) {
	expectNomeJogador1 := "Alberto"
	expectNomeJogador2 := "Josias"
	nomesJogadores := expectNomeJogador1 + "\n" + expectNomeJogador2
	tmpfile, _ := ioutil.TempFile("", "example")

	defer os.Remove(tmpfile.Name())

	tmpfile.WriteString(nomesJogadores)
	tmpfile.Seek(0, 0)

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	os.Stdin = tmpfile
	jogador1, jogador2 := cadastrarJogadores()

	if jogador1 != expectNomeJogador1 {
		t.Error("Experava", expectNomeJogador1, "mas retornou", jogador1)
	}

	if jogador2 != expectNomeJogador2 {
		t.Error("Experava", expectNomeJogador2, "mas retornou", jogador2)
	}

	tmpfile.Close()
}

func TestDeveRetornarMensagemDoGanhador(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	expectMessage := "======== Resultado ========\nVencedor da partida: Alberto !\n"
	exibeVencedor("Alberto")

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if string(out) != expectMessage {
		t.Error("Esperava", expectMessage, "mas retornou:", out)
	}
}

func TestDeveRetornarMensagemDeEmpate(t *testing.T) {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	expectMessage := "======== Resultado ========\nO jogo empatou!\n"
	exibeVencedor("Empate")

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	if string(out) != expectMessage {
		t.Error("Esperava", expectMessage, "mas retornou:", out)
	}
}
