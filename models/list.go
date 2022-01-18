package models

import "fmt"

type Lista struct {
	Root  *Nodo
	Count int
}

func createNode(user *User) *Nodo {
	return &Nodo{user, nil}
}

func (lista *Lista) InsertNode(user *User) {
	node := createNode(user)
	if lista.Root == nil {
		lista.Root = node
		lista.Count++
		return
	}
	node.Next = lista.Root
	lista.Root = node
	lista.Count++
}

func (lista *Lista) InsertNodeFinal(user *User) {
	node := createNode(user)
	if lista.Root == nil {
		lista.Root = node
		lista.Count++
		return
	}
	puntero := lista.Root
	for puntero.Next != nil {
		puntero = puntero.Next
	}
	puntero.Next = node
	lista.Count++
}

func (lista *Lista) InsertNodeMedium(user *User, n int) {
	node := createNode(user)
	if lista.Root == nil {
		lista.Root = node
		lista.Count++
		return
	} else {
		puntero := lista.Root
		count := 0
		for count < n {
			puntero = puntero.Next
			count++
			if puntero == nil {
				fmt.Println("Nodo inexistente")
				return
			}
		}
		node.Next = puntero.Next
		puntero.Next = node
		lista.Count++
	}
}

func (lista *Lista) GetNode(n int) *Nodo {
	if lista.Root == nil {
		fmt.Println("Lista vacia")
		return nil
	} else {
		puntero := lista.Root
		count := 0
		for count < n {
			puntero = puntero.Next
			count++
			if puntero == nil {
				fmt.Println("Nodo inexistente")
				return nil
			}
		}
		return puntero
	}
}

func (lista *Lista) DeleteNode(n int) {
	if lista.Root == nil {
		fmt.Println("Lista vacia")
	} else if n == 0 {
		fmt.Println("Eliminado:", lista.Root.User)
		lista.Root = lista.Root.Next
		lista.Count--
	} else {
		puntero := lista.Root
		count := 0
		for count < n-1 {
			puntero = puntero.Next
			count++
			if puntero == nil {
				fmt.Println("Nodo inexistente")
				return
			}
		}
		fmt.Println("Eliminado:", puntero.Next.User)
		puntero.Next = puntero.Next.Next
		lista.Count--
	}
}
