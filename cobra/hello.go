package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Hello() *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:   "hello [name]",
		Short: "retorna Olá + name passado",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Olá %s\n", name)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "Mundo", "flag para concatenar com Olá")

	return cmd
}
