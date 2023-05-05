package main

import (
	"fmt"
	"math/rand"
	"errors"
)

type Node struct {
	coche Coche
	Next *Node
}

type List struct {
	Head *Node
}

func (l *List) InsertarCoche(car Coche) {
	list := &Node{coche: car, Next: nil}
	if l.Head == nil {
		l.Head = list
	} else {
		p := l.Head
		for p.Next != nil {
			p = p.Next
		}
		p.Next = list
	}
}

func (l *List)InsertarNodo(n *Node){
	if l.Head == nil {
		l.Head = n
	} else {
		p := l.Head
		for p.Next != nil {
			p = p.Next
		}
		p.Next = n
	}
}
func (l *List)Ultimo() *Node{
	curr := l.Head
	if l.Head==nil {
		return curr
	}else if l.Head.Next==nil{
		return nil
	}
	prev:=curr
	curr=curr.Next
	for curr.Next!=nil {
		prev = curr
		curr=curr.Next
	}
	prev.Next=nil
	return curr
}

func (l *List)Obtener(index int) *Node {
	curr := l.Head
	if l.Head==nil {
		return curr
	}else if l.Head.Next==nil{
		return nil
	}
	prev:=curr
	curr=curr.Next
	i:=0
	for curr.Next!=nil {
		if index==i{
			break
		}
		prev = curr
		curr=curr.Next
		i++
	}
	prev.Next=curr.Next
	curr.Next=nil
	return curr
}

func Show(l *List) {
	p := l.Head
	for p != nil {
		fmt.Printf("-> %v ", p.coche)
		p = p.Next
	}
}


type Coche struct {
	marca string
	modelo string
	año int
	color string
}

func (c *Coche)Random(){
	c.año=1900+rand.Int()%123
	c.color=RandStringRunes()
	c.marca=RandStringRunes()
	c.modelo=RandStringRunes()
}

func RandStringRunes() string {
	n:=5
    letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func not_main() {
	disponiblesVacio:=new(Coche)
	compradosVacio:=new(Coche)
	vendidosVacio:=new(Coche)

	disponiblesNodoVacio:=Node{*disponiblesVacio,nil}
	compradosNodoVacio:=Node{*compradosVacio,nil}
	vendidosNodoVacio:=Node{*vendidosVacio,nil}

	disponibles:=List{&disponiblesNodoVacio}
	comprados:=List{&compradosNodoVacio}
	vendidos:=List{&vendidosNodoVacio}

	for i:=0;i<6;i++{
		cochevacio:=new(Coche)
		cochevacio.año=i+1
		disponibles.InsertarCoche(*cochevacio)
	}
	Show(&comprados)
	fmt.Println()
	Show(&vendidos)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	disponibles.Obtener(3)
	Show(&disponibles)

}

func Comprar(año int, color string, marca string,
modelo string, disponibles* List, comprados* List) error {
	nuevo:=Coche{año:año,color:color,marca:marca,modelo:modelo}
	ultimo:=disponibles.Ultimo()
	if ultimo==nil{
		return errors.New("No hay espacio disponible")
	}
	ultimo.coche=nuevo
	comprados.InsertarNodo(ultimo)
	return nil
}

func ComprarCoche(nuevo* Coche, disponibles* List, comprados* List) error {
	ultimo:=disponibles.Ultimo()
	if ultimo==nil{
		return errors.New("No hay espacio disponible")
	}
	ultimo.coche=*nuevo
	comprados.InsertarNodo(ultimo)
	return nil
}
