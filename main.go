package main

import (
	"auditor/attacks"
	"auditor/utils"
	"fmt"
	"strings"
)

func main() {
	var targetHash string
	var hashType string

	fmt.Print("Digite o Hash da senha: ")
	_, err := fmt.Scanln(&targetHash)
	if err != nil {
		fmt.Println("Erro ao ler o hash:", err)
		return
	}

	fmt.Print("Digite o Tipo de Criptografia (Ex: md5, sha1): ")
	_, err = fmt.Scanln(&hashType)
	if err != nil {
		fmt.Println("Erro ao ler o tipo:", err)
		return
	}

	targetHash = strings.TrimSpace(targetHash)
	hashType = strings.ToLower(strings.TrimSpace(hashType))

	fmt.Printf("\n--- Dados Recebidos ---\n")
	fmt.Printf("Hash: %s\n", targetHash)
	fmt.Printf("Tipo: %s\n", hashType)
	
	utils.Info("Carregando wordlist...")

	words, err := utils.LoadWordlist("wordlist.txt")
	if err != nil {
		utils.Error("Erro ao carregar wordlist")
		return
	}

	utils.Info("Iniciando ataque de dicionário...")

	result := attacks.DictionaryAttack(words, targetHash, hashType)

	if result != "" {
		utils.Success("Senha encontrada: " + result)
	} else {
		utils.Error("Senha não encontrada")
	}

	fmt.Println("Finalizado.")
}