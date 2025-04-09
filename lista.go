package main

import (
	"errors"
	"fmt"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
)

type Node struct {
	coche Coche
	Next  *Node
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

func (l *List) InsertarNodo(n *Node) {
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
<<<<<<< HEAD

func (l *List) EliminarCoche(coche Coche, vendidos *List) bool {
	n := l.Head
	index := 0
	encontrado := false
	for n != nil {
		//if coche.año==n.coche.año && coche.color==n.coche.color && coche.marca==n.coche.marca
		if coche == n.coche {
			encontrado = true
			break
		}
		n = n.Next
		index += 1
		if index >= 104 {
			break
		}
	}
	if encontrado {
		vendidos.InsertarNodo(l.Obtener(index))
		return true
	} else {
		return false
	}
}

=======
>>>>>>> 3f2ca0e (format)
func (l *List) Ultimo() *Node {
	curr := l.Head
	if l.Head == nil {
		return curr
	} else if l.Head.Next == nil {
		return nil
	}
	prev := curr
	curr = curr.Next
	for curr.Next != nil {
		prev = curr
		curr = curr.Next
	}
	prev.Next = nil
	return curr
}

func (l *List) Obtener(index int) *Node {
	curr := l.Head
	if l.Head == nil {
		return curr
	} else if l.Head.Next == nil {
		return nil
	}
	prev := curr
	curr = curr.Next
	i := 0
	for curr.Next != nil {
		if index == i {
			break
		}
		prev = curr
		curr = curr.Next
		i++
	}
	prev.Next = curr.Next
	curr.Next = nil
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
	marca  string
	modelo string
	año    int
	color  string
	tamaño string
}

<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 3f2ca0e (format)
func (c *Coche) Random() {
	c.año = 1900 + rand.Int()%123
	c.color = RandStringRunes()
	c.marca = RandStringRunes()
	c.modelo = RandStringRunes()
	tamaños := []string{"Grande", "Mediano", "Pequeño"}
	c.tamaño = tamaños[rand.Intn(3)]
<<<<<<< HEAD
=======
func (c *Coche)Random(){
	c.año=1900+rand.Int()%123
	c.color=RandStringRunes()
	c.marca=RandStringRunes()
	c.modelo=RandStringRunes()
	tamaños:=[]string{"Grande", "Mediano", "Pequeño"}
	c.tamaño=tamaños[rand.Intn(3)]
>>>>>>> ba0016e (update)
=======
>>>>>>> 3f2ca0e (format)
}

func RandStringRunes() string {
	n := 5
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
<<<<<<< HEAD
=======
	}
	return string(b)
}

<<<<<<< HEAD
func not_main() {
	disponiblesVacio := new(Coche)
	compradosVacio := new(Coche)
	vendidosVacio := new(Coche)

	disponiblesNodoVacio := Node{*disponiblesVacio, nil}
	compradosNodoVacio := Node{*compradosVacio, nil}
	vendidosNodoVacio := Node{*vendidosVacio, nil}

	disponibles := List{&disponiblesNodoVacio}
	comprados := List{&compradosNodoVacio}
	vendidos := List{&vendidosNodoVacio}

	for i := 0; i < 6; i++ {
		cochevacio := new(Coche)
		cochevacio.año = i + 1
		disponibles.InsertarCoche(*cochevacio)
>>>>>>> 3f2ca0e (format)
	}
	return string(b)
}

=======
>>>>>>> 2b27c6c (cambios)
func Comprar(año int, color string, marca string,
	modelo string, tamaño string, disponibles *List, comprados *List) error {
	nuevo := Coche{año: año, color: color, marca: marca, modelo: modelo, tamaño: tamaño}
	ultimo := disponibles.Ultimo()
	if ultimo == nil {
		return errors.New("No hay espacio disponible")
	}
	ultimo.coche = nuevo
	comprados.InsertarNodo(ultimo)
	return nil
}

func ComprarCoche(nuevo *Coche, disponibles *List, comprados *List) error {
	ultimo := disponibles.Ultimo()
	if ultimo == nil {
		return errors.New("No hay espacio disponible")
	}
	ultimo.coche = *nuevo
	comprados.InsertarNodo(ultimo)
	return nil
}

func (l *List) bubbleSort() {
	if l.Head == nil || l.Head.Next == nil {
		return
	}

	swapped := true
	for swapped {
		swapped = false
		current := l.Head
		for current.Next != nil {
			if current.coche.marca > current.Next.coche.marca {
				current.coche, current.Next.coche = current.Next.coche, current.coche
				swapped = true
			}
			current = current.Next
		}
	}
}

func swap(a, b *Node) {
	a.coche, b.coche = b.coche, a.coche
}

func partition(low, high *Node) *Node {
	pivot := high
	i := low

	for j := low; j != high; j = j.Next {
		if j.coche.marca < pivot.coche.marca {
			swap(i, j)
			i = i.Next
		}
	}
	swap(i, high)
	return i
}

// Función principal de QuickSort
func quickSort(low, high *Node) {
	if low != nil && high != nil && low != high && low != high.Next {
		pivot := partition(low, high)
		quickSort(low, pivot)
		quickSort(pivot.Next, high)
	}
}

func agregarLista(p *Node, ui *fyne.Container) {
	for p != nil {
		zona := container.New(layout.NewHBoxLayout())
		zona.Add(widget.NewLabel("Modelo: " + p.coche.modelo))
		zona.Add(widget.NewLabel("Marca: " + p.coche.marca))
		zona.Add(widget.NewLabel("Color: " + p.coche.color))
		zona.Add(widget.NewLabel("Tamaño: " + p.coche.tamaño))
		zona.Add(widget.NewLabel(fmt.Sprintf("Año: %d", p.coche.año)))
		line := canvas.NewLine(color.White)
		line.StrokeWidth = 2
		ui.Add(zona)
		ui.Add(line)
		p = p.Next
	}

}

// merge is a helper function to merge two sorted linked lists
func merge(left, right *Node) *Node {
	var result *Node

	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	if left.coche.marca <= right.coche.marca {
		result = left
		result.Next = merge(left.Next, right)
	} else {
		result = right
		result.Next = merge(left, right.Next)
	}

	return result
}

// mergeSort sorts the linked list using merge sort algorithm
func mergeSort(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	// Split the list into two halves
	var left, right *Node
	slow, fast := head, head.Next

	for fast != nil {
		fast = fast.Next
		if fast != nil {
			slow = slow.Next
			fast = fast.Next
		}
	}

	left = head
	right = slow.Next
	slow.Next = nil

	// Recursively sort the two halves
	left = mergeSort(left)
	right = mergeSort(right)

	// Merge the sorted halves
	return merge(left, right)
}
