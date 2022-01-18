package main

import (
	"fmt"

	"github.com/RBosio/lista-enlazada/models"
)

func main() {
	list := models.Lista{}

	list.InsertNodeFinal(&models.User{ID: 1, Name: "Fiido", Age: 40})
	list.InsertNodeFinal(&models.User{ID: 2, Name: "Marcelo", Age: 30})
	list.InsertNodeFinal(&models.User{ID: 3, Name: "Maria", Age: 12})
	list.InsertNodeFinal(&models.User{ID: 4, Name: "Carlos", Age: 63})
	list.InsertNodeMedium(&models.User{ID: 5, Name: "Juan", Age: 84}, 2)

	if nodo := list.GetNode(3); nodo != nil {
		fmt.Println(nodo.User)
	}
	fmt.Println("Cantidad de nodos:", list.Count)

	list.DeleteNode(3)
	if nodo := list.GetNode(3); nodo != nil {
		fmt.Println(nodo.User)
	}

	fmt.Println("Cantidad de nodos:", list.Count)
}
