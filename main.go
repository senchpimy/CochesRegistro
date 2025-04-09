package main

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 3f2ca0e (format)
func createButtonCallback(n int, comprados *List, vendidos *List, disponibles *List) func() {
	return func() {
		nodo := comprados.Obtener(n)
		if nodo != nil {
			vendidos.InsertarNodo(nodo)
			disponibles.InsertarNodo(new(Node))
			compradosUi.RemoveAll()
			content := ElementosComprados(vendidos, disponibles)
<<<<<<< HEAD
			compradosUi.AddObject(content)
		}
=======
func createButtonCallback(n int, comprados* List,vendidos* List,disponibles* List) func() {
    return func() {
	nodo:=comprados.Obtener(n)	
	if nodo!=nil{
		vendidos.InsertarNodo(nodo)
		disponibles.InsertarNodo(new(Node))
		compradosUi.RemoveAll()
		content:=ElementosComprados(vendidos,disponibles)
		compradosUi.Add(content)
>>>>>>> ba0016e (update)
=======
			compradosUi.Add(content)
		}
>>>>>>> 3f2ca0e (format)
	}
}

func ElementosVendidos(l *List) *fyne.Container {
<<<<<<< HEAD
	compradosUi := container.New(layout.NewVBoxLayout())
	marca := widget.NewEntry()
	marca.SetPlaceHolder("Ingrese la Marca")
	modelo := widget.NewEntry()
	modelo.SetPlaceHolder("Ingrese el modelo")
	año := widget.NewEntry()
	año.SetPlaceHolder("Ingrese el año")
	colorW := widget.NewEntry()
	colorW.SetPlaceHolder("Ingrese la color")
	tamaSel := ""
	tamaño := widget.NewSelect([]string{"Grande", "Mediano", "Pequeño"}, func(val string) { tamaSel = val })
	botonAñadir := widget.NewButton("Vender Coche", func() {
		añoInt, err := strconv.Atoi(año.Text)
		if err != nil {
			errorWidget.SetText(err.Error())
			return
		}
		BuscarCoche := Coche{marca: marca.Text, modelo: modelo.Text, año: añoInt, color: colorW.Text, tamaño: tamaSel}
		vendido := false
		if tamaSel == "Grande" {
			vendido = compradosGrandes.EliminarCoche(BuscarCoche, &vendidos)
		} else if tamaSel == "Mediano" {
			vendido = compradosMedianos.EliminarCoche(BuscarCoche, &vendidos)
		} else if tamaSel == "Pequeño" {
			vendido = compradosChicos.EliminarCoche(BuscarCoche, &vendidos)
		} else {
			fmt.Println("No se encontro")
		}

		if !vendido {
			errorWidget.SetText("Coche no Encontrado")
		} else {
			errorWidget.SetText("")
		}

		nuevo := ElementosVendidos(&vendidos)
		nuevoScroll := container.NewVScroll(nuevo)
		nuevoScroll.SetMinSize(fyne.NewSize(500, 500))
		zonaPrincipal.RemoveAll()
		zonaPrincipal.AddObject(nuevoScroll)
	})
	compradosUi.AddObject(marca)
	compradosUi.AddObject(modelo)
	compradosUi.AddObject(año)
	compradosUi.AddObject(colorW)
	compradosUi.AddObject(tamaño)
	compradosUi.AddObject(botonAñadir)
	p := l.Head.Next
<<<<<<< HEAD
	i := 0
	zona := container.New(layout.NewHBoxLayout())
	lugaresVacios := widget.NewLabel("Coches Vendidos")
	zona.AddObject(lugaresVacios)
=======
	compradosUi:=container.New(layout.NewVBoxLayout())
	i:=0
	zona:=container.New(layout.NewHBoxLayout())
	lugaresVacios:=widget.NewLabel("Coches Vendidos")
=======
	p := l.Head.Next
	compradosUi := container.New(layout.NewVBoxLayout())
	i := 0
	zona := container.New(layout.NewHBoxLayout())
	lugaresVacios := widget.NewLabel("Coches Vendidos")
>>>>>>> 3f2ca0e (format)
	zona.Add(lugaresVacios)
>>>>>>> ba0016e (update)
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 10
	compradosUi.Add(zona)
	compradosUi.Add(line)
	for p != nil {
		zona := container.New(layout.NewHBoxLayout())
		line := canvas.NewLine(color.White)
		line.StrokeWidth = 2
<<<<<<< HEAD
<<<<<<< HEAD
		zona.AddObject(widget.NewLabel("Modelo: " + p.coche.modelo))
		zona.AddObject(widget.NewLabel("Marca: " + p.coche.marca))
		zona.AddObject(widget.NewLabel("Color: " + p.coche.color))
		zona.AddObject(widget.NewLabel("Tamaño: " + p.coche.tamaño))
		zona.AddObject(widget.NewLabel(fmt.Sprintf("Año: %d", p.coche.año)))
		compradosUi.AddObject(zona)
		compradosUi.AddObject(line)
=======
		zona.Add(widget.NewLabel("Modelo: "+p.coche.modelo))
		zona.Add(widget.NewLabel("Marca: "+p.coche.marca))
		zona.Add(widget.NewLabel("Color: "+p.coche.color))
		zona.Add(widget.NewLabel("Tamaño: "+p.coche.tamaño))
		zona.Add(widget.NewLabel(fmt.Sprintf("Año: %d",p.coche.año)))
=======
		zona.Add(widget.NewLabel("Modelo: " + p.coche.modelo))
		zona.Add(widget.NewLabel("Marca: " + p.coche.marca))
		zona.Add(widget.NewLabel("Color: " + p.coche.color))
		zona.Add(widget.NewLabel("Tamaño: " + p.coche.tamaño))
		zona.Add(widget.NewLabel(fmt.Sprintf("Año: %d", p.coche.año)))
>>>>>>> 3f2ca0e (format)
		compradosUi.Add(zona)
		compradosUi.Add(line)
>>>>>>> ba0016e (update)
		p = p.Next
		i += 1
	}
	lugaresVacios.SetText(fmt.Sprintf("%d Coches han sido vendidos", i))
	return compradosUi
}

func ElementosDisponibles(l *List) *fyne.Container {
	p := l.Head.Next
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 3f2ca0e (format)
	compradosUi := container.New(layout.NewVBoxLayout())
	i := 0
	zona := container.New(layout.NewHBoxLayout())
	lugaresVacios := widget.NewLabel("")
<<<<<<< HEAD
	zona.AddObject(lugaresVacios)
=======
	compradosUi:=container.New(layout.NewVBoxLayout())
	i:=0
	zona:=container.New(layout.NewHBoxLayout())
	lugaresVacios:=widget.NewLabel("")
=======
>>>>>>> 3f2ca0e (format)
	zona.Add(lugaresVacios)
>>>>>>> ba0016e (update)
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 10
	compradosUi.Add(zona)
	compradosUi.Add(line)
	for p != nil {
<<<<<<< HEAD
<<<<<<< HEAD
		zona := container.New(layout.NewHBoxLayout())
		//zona.AddObject(widget.NewLabel(fmt.Sprintf("Lugar #%d",i)))
		zona.AddObject(widget.NewLabel("Vacio"))
=======
		zona:=container.New(layout.NewHBoxLayout())
=======
		zona := container.New(layout.NewHBoxLayout())
>>>>>>> 3f2ca0e (format)
		//zona.Add(widget.NewLabel(fmt.Sprintf("Lugar #%d",i)))
		zona.Add(widget.NewLabel("Vacio"))
>>>>>>> ba0016e (update)
		line := canvas.NewLine(color.White)
		line.StrokeWidth = 2
		compradosUi.Add(zona)
		compradosUi.Add(line)
		p = p.Next
		i += 1
	}
	lugaresVacios.SetText(fmt.Sprintf("Hay %d lugares vacios", i))
	return compradosUi
}

<<<<<<< HEAD
func ElementosVista() *fyne.Container {
<<<<<<< HEAD
	p := compradosGrandes.Head.Next
	p2 := compradosMedianos.Head.Next
	p3 := compradosChicos.Head.Next
	compradosUi := container.New(layout.NewVBoxLayout())

	zona := container.New(layout.NewHBoxLayout())
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 10
	compradosUi.AddObject(zona)
	compradosUi.AddObject(line)
	compradosUi.AddObject(widget.NewLabel("Grandes"))
	for p != nil {
		zona := container.New(layout.NewHBoxLayout())
		zona.AddObject(widget.NewLabel("Modelo: " + p.coche.modelo))
		zona.AddObject(widget.NewLabel("Marca: " + p.coche.marca))
		zona.AddObject(widget.NewLabel("Color: " + p.coche.color))
		zona.AddObject(widget.NewLabel("Tamaño: " + p.coche.tamaño))
		zona.AddObject(widget.NewLabel(fmt.Sprintf("Año: %d", p.coche.año)))
		line := canvas.NewLine(color.White)
		line.StrokeWidth = 2
		compradosUi.AddObject(zona)
		compradosUi.AddObject(line)
		p = p.Next
	}
	compradosUiMed := container.New(layout.NewVBoxLayout())
	zonaMed := container.New(layout.NewHBoxLayout())
	compradosUiMed.AddObject(zonaMed)
	compradosUiMed.AddObject(line)
	compradosUiMed.AddObject(widget.NewLabel("Medianos"))
	for p2 != nil {
		zonaMed := container.New(layout.NewHBoxLayout())
		zonaMed.AddObject(widget.NewLabel("Modelo: " + p2.coche.modelo))
		zonaMed.AddObject(widget.NewLabel("Marca: " + p2.coche.marca))
		zonaMed.AddObject(widget.NewLabel("Color: " + p2.coche.color))
		zonaMed.AddObject(widget.NewLabel("Tamaño: " + p2.coche.tamaño))
		zonaMed.AddObject(widget.NewLabel(fmt.Sprintf("Año: %d", p2.coche.año)))
		line := canvas.NewLine(color.White)
		line.StrokeWidth = 2
		compradosUiMed.AddObject(zonaMed)
		compradosUiMed.AddObject(line)
		p2 = p2.Next
	}

	compradosUiPeq := container.New(layout.NewVBoxLayout())
	zonaPeq := container.New(layout.NewHBoxLayout())
	compradosUiPeq.AddObject(zonaPeq)
	compradosUiPeq.AddObject(line)
	compradosUiPeq.AddObject(widget.NewLabel("Pequeños"))
	for p3 != nil {
		zonaPeq := container.New(layout.NewHBoxLayout())
		zonaPeq.AddObject(widget.NewLabel("Modelo: " + p3.coche.modelo))
		zonaPeq.AddObject(widget.NewLabel("Marca: " + p3.coche.marca))
		zonaPeq.AddObject(widget.NewLabel("Color: " + p3.coche.color))
		zonaPeq.AddObject(widget.NewLabel("Tamaño: " + p3.coche.tamaño))
		zonaPeq.AddObject(widget.NewLabel(fmt.Sprintf("Año: %d", p3.coche.año)))
		line := canvas.NewLine(color.White)
		line.StrokeWidth = 2
		compradosUiPeq.AddObject(zonaPeq)
		compradosUiPeq.AddObject(line)
		p3 = p3.Next
	}

	grid := container.New(layout.NewGridLayout(3), compradosUi, compradosUiMed, compradosUiPeq)
	return grid
}

func ElementosComprados(vendidos *List, disponibles *List) *fyne.Container {
	p := compradosGrandes.Head.Next
	p2 := compradosMedianos.Head.Next
	p3 := compradosChicos.Head.Next
	compradosUi := container.New(layout.NewVBoxLayout())
	i := 0
	zona := container.New(layout.NewHBoxLayout())
	lugaresVacios := widget.NewLabel("")
	zona.AddObject(lugaresVacios)
=======
=======
func ElementosVista() *fyne.Container { //Muestra los Elementos Y los ordena
	g := container.New(layout.NewVBoxLayout())
>>>>>>> 2b27c6c (cambios)
	p := compradosGrandes.Head.Next
	p2 := compradosMedianos.Head.Next
	p3 := compradosChicos.Head.Next
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 10

	compradosUi := container.New(layout.NewVBoxLayout())
	var ordenarGrand fyne.CanvasObject
	ordenarGrand = widget.NewButton("Ordenar Por Quick Sort (Marca)", func() {
		quickSort(compradosGrandes.Head.Next, nil)
		compradosUi.RemoveAll()
		compradosUi.Add(widget.NewLabel("Grandes"))
		compradosUi.Add(line)
		compradosUi.Add(ordenarGrand)
		agregarLista(compradosGrandes.Head.Next, compradosUi)
	})
	compradosUi.Add(widget.NewLabel("Grandes"))
	compradosUi.Add(line)
	compradosUi.Add(ordenarGrand)
	agregarLista(p, compradosUi)

	compradosUiMed := container.New(layout.NewVBoxLayout())
	var ordenarMed fyne.CanvasObject
	ordenarMed = widget.NewButton("Ordenar Por Merge Sort (Marca)", func() {
		compradosMedianos.Head.Next = mergeSort(compradosMedianos.Head.Next)
		tmp := compradosMedianos.Head.Next
		compradosUiMed.RemoveAll()
		compradosUiMed.Add(widget.NewLabel("Medianos"))
		compradosUiMed.Add(line)
		compradosUiMed.Add(ordenarMed)
		agregarLista(tmp, compradosUiMed)
	})
	compradosUiMed.Add(widget.NewLabel("Medianos"))
	compradosUiMed.Add(line)
	compradosUiMed.Add(ordenarMed)
	agregarLista(p2, compradosUiMed)

	compradosUiPeq := container.New(layout.NewVBoxLayout())
	var ordenarPeq fyne.CanvasObject
	ordenarPeq = widget.NewButton("Ordenar Por Burbuja (Marca)", func() {
		compradosChicos.bubbleSort()
		compradosUiPeq.RemoveAll()
		compradosUiPeq.Add(widget.NewLabel("Pequeños"))
		compradosUiPeq.Add(line)
		compradosUiPeq.Add(ordenarPeq)
		tmp_p := compradosChicos.Head.Next
		agregarLista(tmp_p, compradosUiPeq)
	})
	compradosUiPeq.Add(widget.NewLabel("Pequeños"))
	compradosUiPeq.Add(line)
	compradosUiPeq.Add(ordenarPeq)
	agregarLista(p3, compradosUiPeq)

	grid := container.New(layout.NewGridLayout(3))
	btn1 := widget.NewButton("Coches Grandes", func() {
		l := len(g.Objects)
		if l == 1 {
			g.Add(compradosUi)
		} else {
			g.Objects[1] = compradosUi
		}
	})
	btn2 := widget.NewButton("Coches Medianos", func() {
		l := len(g.Objects)
		if l == 1 {
			g.Add(compradosUiMed)
		} else {
			g.Objects[1] = compradosUiMed
		}
	})
	btn3 := widget.NewButton("Coches Pequeños", func() {
		l := len(g.Objects)
		if l == 1 {
			g.Add(compradosUiPeq)
		} else {
			g.Objects[1] = compradosUiPeq
		}
	})
	grid.Add(btn1)
	grid.Add(btn2)
	grid.Add(btn3)
	g.Add(grid)
	return g
}

func ElementosComprados(vendidos *List, disponibles *List) *fyne.Container { // Interfaz encargada de Comprar coches
	p := compradosGrandes.Head.Next
	p2 := compradosMedianos.Head.Next
	p3 := compradosChicos.Head.Next
	compradosUi := container.New(layout.NewVBoxLayout())
	i := 0
	zona := container.New(layout.NewHBoxLayout())
	lugaresVacios := widget.NewLabel("")
	zona.Add(lugaresVacios)
>>>>>>> ba0016e (update)
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 10
	compradosUi.Add(zona)
	compradosUi.Add(line)
	compradosUi.Add(widget.NewLabel("Grandes"))
	for p != nil {
<<<<<<< HEAD
<<<<<<< HEAD
		zona := container.New(layout.NewHBoxLayout())
		zona.AddObject(widget.NewLabel("Modelo: " + p.coche.modelo))
		zona.AddObject(widget.NewLabel("Marca: " + p.coche.marca))
		zona.AddObject(widget.NewLabel("Color: " + p.coche.color))
		zona.AddObject(widget.NewLabel("Tamaño: " + p.coche.tamaño))
		zona.AddObject(widget.NewLabel(fmt.Sprintf("Año: %d", p.coche.año)))
		zona.AddObject(widget.NewButton("Vender", createButtonCallback(i, &compradosGrandes, vendidos, disponibles)))
=======
		zona:=container.New(layout.NewHBoxLayout())
		zona.Add(widget.NewLabel("Modelo: "+p.coche.modelo))
		zona.Add(widget.NewLabel("Marca: "+p.coche.marca))
		zona.Add(widget.NewLabel("Color: "+p.coche.color))
		zona.Add(widget.NewLabel("Tamaño: "+p.coche.tamaño))
		zona.Add(widget.NewLabel(fmt.Sprintf("Año: %d",p.coche.año)))
		zona.Add(widget.NewButton("Vender",createButtonCallback(i,&compradosGrandes,vendidos,disponibles)))
>>>>>>> ba0016e (update)
=======
		zona := container.New(layout.NewHBoxLayout())
		zona.Add(widget.NewLabel("Modelo: " + p.coche.modelo))
		zona.Add(widget.NewLabel("Marca: " + p.coche.marca))
		zona.Add(widget.NewLabel("Color: " + p.coche.color))
		zona.Add(widget.NewLabel("Tamaño: " + p.coche.tamaño))
		zona.Add(widget.NewLabel(fmt.Sprintf("Año: %d", p.coche.año)))
		zona.Add(widget.NewButton("Vender", createButtonCallback(i, &compradosGrandes, vendidos, disponibles)))
>>>>>>> 3f2ca0e (format)
		line := canvas.NewLine(color.White)
		line.StrokeWidth = 2
		compradosUi.Add(zona)
		compradosUi.Add(line)
		p = p.Next
		i += 1
	}
	compradosUi.Add(widget.NewLabel("Medianos"))
	for p2 != nil {
<<<<<<< HEAD
<<<<<<< HEAD
		zona := container.New(layout.NewHBoxLayout())
		zona.AddObject(widget.NewLabel("Modelo: " + p2.coche.modelo))
		zona.AddObject(widget.NewLabel("Marca: " + p2.coche.marca))
		zona.AddObject(widget.NewLabel("Color: " + p2.coche.color))
		zona.AddObject(widget.NewLabel("Tamaño: " + p2.coche.tamaño))
		zona.AddObject(widget.NewLabel(fmt.Sprintf("Año: %d", p2.coche.año)))
		zona.AddObject(widget.NewButton("Vender", createButtonCallback(i, &compradosMedianos, vendidos, disponibles)))
=======
		zona:=container.New(layout.NewHBoxLayout())
		zona.Add(widget.NewLabel("Modelo: "+p2.coche.modelo))
		zona.Add(widget.NewLabel("Marca: "+p2.coche.marca))
		zona.Add(widget.NewLabel("Color: "+p2.coche.color))
		zona.Add(widget.NewLabel("Tamaño: "+p2.coche.tamaño))
		zona.Add(widget.NewLabel(fmt.Sprintf("Año: %d",p2.coche.año)))
		zona.Add(widget.NewButton("Vender",createButtonCallback(i,&compradosMedianos,vendidos,disponibles)))
>>>>>>> ba0016e (update)
=======
		zona := container.New(layout.NewHBoxLayout())
		zona.Add(widget.NewLabel("Modelo: " + p2.coche.modelo))
		zona.Add(widget.NewLabel("Marca: " + p2.coche.marca))
		zona.Add(widget.NewLabel("Color: " + p2.coche.color))
		zona.Add(widget.NewLabel("Tamaño: " + p2.coche.tamaño))
		zona.Add(widget.NewLabel(fmt.Sprintf("Año: %d", p2.coche.año)))
		zona.Add(widget.NewButton("Vender", createButtonCallback(i, &compradosMedianos, vendidos, disponibles)))
>>>>>>> 3f2ca0e (format)
		line := canvas.NewLine(color.White)
		line.StrokeWidth = 2
		compradosUi.Add(zona)
		compradosUi.Add(line)
		p2 = p2.Next
		i += 1
	}
	compradosUi.Add(widget.NewLabel("Pequeños"))
	for p3 != nil {
<<<<<<< HEAD
<<<<<<< HEAD
		zona := container.New(layout.NewHBoxLayout())
		zona.AddObject(widget.NewLabel("Modelo: " + p3.coche.modelo))
		zona.AddObject(widget.NewLabel("Marca: " + p3.coche.marca))
		zona.AddObject(widget.NewLabel("Color: " + p3.coche.color))
		zona.AddObject(widget.NewLabel("Tamaño: " + p3.coche.tamaño))
		zona.AddObject(widget.NewLabel(fmt.Sprintf("Año: %d", p3.coche.año)))
		zona.AddObject(widget.NewButton("Vender", createButtonCallback(i, &compradosChicos, vendidos, disponibles)))
=======
		zona:=container.New(layout.NewHBoxLayout())
		zona.Add(widget.NewLabel("Modelo: "+p3.coche.modelo))
		zona.Add(widget.NewLabel("Marca: "+p3.coche.marca))
		zona.Add(widget.NewLabel("Color: "+p3.coche.color))
		zona.Add(widget.NewLabel("Tamaño: "+p3.coche.tamaño))
		zona.Add(widget.NewLabel(fmt.Sprintf("Año: %d",p3.coche.año)))
		zona.Add(widget.NewButton("Vender",createButtonCallback(i,&compradosChicos,vendidos,disponibles)))
>>>>>>> ba0016e (update)
=======
		zona := container.New(layout.NewHBoxLayout())
		zona.Add(widget.NewLabel("Modelo: " + p3.coche.modelo))
		zona.Add(widget.NewLabel("Marca: " + p3.coche.marca))
		zona.Add(widget.NewLabel("Color: " + p3.coche.color))
		zona.Add(widget.NewLabel("Tamaño: " + p3.coche.tamaño))
		zona.Add(widget.NewLabel(fmt.Sprintf("Año: %d", p3.coche.año)))
		zona.Add(widget.NewButton("Vender", createButtonCallback(i, &compradosChicos, vendidos, disponibles)))
>>>>>>> 3f2ca0e (format)
		line := canvas.NewLine(color.White)
		line.StrokeWidth = 2
		compradosUi.Add(zona)
		compradosUi.Add(line)
		p3 = p3.Next
		i += 1
	}
	lugaresVacios.SetText(fmt.Sprintf("Hay %d coches comprados", i))

	return compradosUi
}

var zonaPrincipal = container.NewVBox()
var ingresoDatos = container.NewVBox()
var compradosUi = container.NewVBox()

var compradosGrandesVacio = new(Coche)
var compradosGrandesNodoVacio = Node{*compradosGrandesVacio, nil}
var compradosGrandes = List{&compradosGrandesNodoVacio}

var compradosMedianosVacio = new(Coche)
var compradosMedianosNodoVacio = Node{*compradosMedianosVacio, nil}
var compradosMedianos = List{&compradosMedianosNodoVacio}

var compradosChicosVacio = new(Coche)
var compradosChicosNodoVacio = Node{*compradosChicosVacio, nil}
var compradosChicos = List{&compradosChicosNodoVacio}

var tamañoSeleccionado = ""
var errorWidget = widget.NewLabel("")
<<<<<<< HEAD

var vendidosVacio = new(Coche)
var vendidosNodoVacio = Node{*vendidosVacio, nil}
var vendidos = List{&vendidosNodoVacio}

func main() {
	disponiblesVacio := new(Coche)

	disponiblesNodoVacio := Node{*disponiblesVacio, nil}

	disponibles := List{&disponiblesNodoVacio}
=======

func main() {
	disponiblesVacio := new(Coche)
	vendidosVacio := new(Coche)

	disponiblesNodoVacio := Node{*disponiblesVacio, nil}
	vendidosNodoVacio := Node{*vendidosVacio, nil}

	disponibles := List{&disponiblesNodoVacio}
	vendidos := List{&vendidosNodoVacio}
>>>>>>> 3f2ca0e (format)

	for i := 0; i < 100; i++ {
		cochevacio := new(Coche)
		cochevacio.año = i + 1
		disponibles.InsertarCoche(*cochevacio)
	}

	a := app.New()
	w := a.NewWindow("Venta de Autos")

	hello := widget.NewLabel("Venta de Autos")

	//Area de Disponibles
	disponiblesUi := ElementosDisponibles(&disponibles)
	disponiblesScroll := container.NewVScroll(disponiblesUi)
	disponiblesScroll.SetMinSize(fyne.NewSize(500, 500))
	//zonaPrincipal=container.NewVBox(error ,disponiblesScroll)
	zonaPrincipal = container.NewVBox(disponiblesScroll)

	//Area de ingreso
	marca := widget.NewEntry()
	marca.SetPlaceHolder("Ingrese la Marca")
	modelo := widget.NewEntry()
	modelo.SetPlaceHolder("Ingrese el modelo")
	año := widget.NewEntry()
	año.SetPlaceHolder("Ingrese el año")
	color := widget.NewEntry()
	color.SetPlaceHolder("Ingrese la color")
	tamaño := widget.NewSelect([]string{"Grande", "Mediano", "Pequeño"}, func(val string) { tamañoSeleccionado = val })
	botonAñadir := widget.NewButton("Comprar Coche", func() {})
	botonRandom := widget.NewButton("Comprar Coche al azar", func() {})
	botonAñadir = widget.NewButton("Comprar Coche", func() {
		fmt.Println("Boton apretado")
		añoInt, err := strconv.Atoi(año.Text)
		if err != nil {
			errorWidget.SetText("No es un año valido")
			return
		}
		err1 := err
		if tamañoSeleccionado == "Grande" {
			err1 = Comprar(añoInt, color.Text, marca.Text, modelo.Text, tamañoSeleccionado, &disponibles, &compradosGrandes)
		} else if tamañoSeleccionado == "Mediano" {
			err1 = Comprar(añoInt, color.Text, marca.Text, modelo.Text, tamañoSeleccionado, &disponibles, &compradosMedianos)
		} else if tamañoSeleccionado == "Pequeño" {
			err1 = Comprar(añoInt, color.Text, marca.Text, modelo.Text, tamañoSeleccionado, &disponibles, &compradosChicos)
		} else {
			fmt.Printf("Por algun motivo el valor no mathceo $%s$\n", tamañoSeleccionado)
		}
		if err1 != nil {
			errorWidget.SetText(err1.Error())
		}
		marca.SetText("")
		modelo.SetText("")
		año.SetText("")
		color.SetText("")
		tamaño.ClearSelected()
		compradosUi = ElementosComprados(&vendidos, &disponibles)
		compradosScroll := container.NewVScroll(compradosUi)
		compradosScroll.SetMinSize(fyne.NewSize(500, 300))
		zonaPrincipal.Objects[1] = compradosScroll
	})
<<<<<<< HEAD
<<<<<<< HEAD
=======
>>>>>>> 2b27c6c (cambios)

	botonRandom = widget.NewButton("Comprar Coche al azar", func() {
		coche := new(Coche)
		coche.Random()
<<<<<<< HEAD
=======
	botonRandom = widget.NewButton("Comprar Coche al azar", func() {
		coche := new(Coche)
		coche.Random()
>>>>>>> 3f2ca0e (format)
		tamañoSel := coche.tamaño
		var err1 error
		if tamañoSel == "Grande" {
			err1 = ComprarCoche(coche, &disponibles, &compradosGrandes)
		} else if tamañoSel == "Mediano" {
			err1 = ComprarCoche(coche, &disponibles, &compradosMedianos)
		} else if tamañoSel == "Pequeño" {
			err1 = ComprarCoche(coche, &disponibles, &compradosChicos)
		} else {
			fmt.Printf("Por algun motivo el valor no acerto $%s$\n", tamañoSel)
		}
<<<<<<< HEAD
=======
		tamañoSel:=coche.tamaño
		err1:=errors.New("")
		if tamañoSel=="Grande"{
			err1=ComprarCoche(coche,&disponibles,&compradosGrandes)
		}else if tamañoSel=="Mediano"{ 
			err1=ComprarCoche(coche,&disponibles,&compradosMedianos)
		}else if tamañoSel=="Pequeño"{
			err1=ComprarCoche(coche,&disponibles,&compradosChicos)
		}else{fmt.Printf("Por algun motivo el valor no mathceo $%s$\n",tamañoSel)}
>>>>>>> ba0016e (update)
=======
>>>>>>> 3f2ca0e (format)

		if err1 != nil {
			errorWidget.SetText(err1.Error())
		}
		marca = widget.NewEntry()
		marca.SetPlaceHolder("Ingrese la Marca")
		modelo = widget.NewEntry()
		modelo.SetPlaceHolder("Ingrese el modelo")
		año = widget.NewEntry()
		año.SetPlaceHolder("Ingrese el año")
		color = widget.NewEntry()
		color.SetPlaceHolder("Ingrese la color")
		tamaño = widget.NewSelect([]string{"Grande", "Mediano", "Pequeño"}, func(val string) { tamañoSeleccionado = val })
		ingresoDatos := container.New(layout.NewVBoxLayout(),
			marca, modelo, color, año, tamaño, botonAñadir, botonRandom)
		compradosUi = ElementosComprados(&vendidos, &disponibles)
		compradosScroll := container.NewVScroll(compradosUi)
		compradosScroll.SetMinSize(fyne.NewSize(500, 300))
		zonaPrincipal.RemoveAll()
		zonaPrincipal.Add(ingresoDatos)
		zonaPrincipal.Add(compradosScroll)
	})
	ingresoDatos = container.New(layout.NewVBoxLayout(),
		marca, modelo, color, año, tamaño, botonAñadir, botonRandom)

	//Area de Vendidos

	grid := container.New(layout.NewGridLayout(4),
		widget.NewButton("Comprar", func() {
			compradosUi = ElementosComprados(&vendidos, &disponibles)
			compradosScroll := container.NewVScroll(compradosUi)
			compradosScroll.SetMinSize(fyne.NewSize(500, 300))
			zonaPrincipal.RemoveAll()
			zonaPrincipal.Add(ingresoDatos)
			zonaPrincipal.Add(compradosScroll)
		}),
		widget.NewButton("Vender", func() {
			nuevo := ElementosVendidos(&vendidos)
			nuevoScroll := container.NewVScroll(nuevo)
			nuevoScroll.SetMinSize(fyne.NewSize(500, 500))
			zonaPrincipal.RemoveAll()
			zonaPrincipal.Add(nuevoScroll)
		}),
<<<<<<< HEAD
		widget.NewButton("Vista", func() {
<<<<<<< HEAD
<<<<<<< HEAD
=======
		widget.NewButton("Ordenar", func() {
>>>>>>> 2b27c6c (cambios)
			nuevo := ElementosVista()
			nuevoScroll := container.NewVScroll(nuevo)
			nuevoScroll.SetMinSize(fyne.NewSize(500, 500))
			zonaPrincipal.RemoveAll()
			zonaPrincipal.AddObject(nuevoScroll)
		}),
		widget.NewButton("Disponibles", func() {
			nuevo := ElementosDisponibles(&disponibles)
			nuevoScroll := container.NewVScroll(nuevo)
			nuevoScroll.SetMinSize(fyne.NewSize(500, 500))
			zonaPrincipal.RemoveAll()
			zonaPrincipal.AddObject(nuevoScroll)
		}))
	content := container.NewVBox(hello, errorWidget, grid, zonaPrincipal)
=======
			nuevo:=ElementosVista()
			nuevoScroll:=container.NewVScroll(nuevo)
=======
			nuevo := ElementosVista()
			nuevoScroll := container.NewVScroll(nuevo)
>>>>>>> 3f2ca0e (format)
			nuevoScroll.SetMinSize(fyne.NewSize(500, 500))
			zonaPrincipal.RemoveAll()
			zonaPrincipal.Add(nuevoScroll)
		}),
		widget.NewButton("Disponibles", func() {
			nuevo := ElementosDisponibles(&disponibles)
			nuevoScroll := container.NewVScroll(nuevo)
			nuevoScroll.SetMinSize(fyne.NewSize(500, 500))
			zonaPrincipal.RemoveAll()
			zonaPrincipal.Add(nuevoScroll)
<<<<<<< HEAD
		}) )
	content:=container.NewVBox(hello,errorWidget,grid,zonaPrincipal)
>>>>>>> ba0016e (update)
=======
		}))
	content := container.NewVBox(hello, errorWidget, grid, zonaPrincipal)
>>>>>>> 3f2ca0e (format)
	w.SetContent(content)

	w.ShowAndRun()
}
