// Dup2 exibe a contagem e o texto das linhas que aparecem mais de uma
// vez na entrada. Ele lê de stdin ou de uma lista de arquivos nomeados.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts) // função chamada antes de sua declaração. Funções no mesmo pacote não tem ordem para declaração
	} else {
		for _, arg := range files {  // itera sobre a lista dos nomes passados
			f, err := os.Open(arg) // abre o arquivo com base no nome
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// counts passado por referencia (cópia) pode ser alterado e é visivel pela refenrência ao map de quem chamou. (main)
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// Nota ignorando erros em potencial de input.Err()
}