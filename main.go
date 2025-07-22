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

func createButtonCallback(n int, comprados *List, vendidos *List, disponibles *List) func() {
	return func() {
		nodo := comprados.Obtener(n)
		if nodo != nil {
			vendidos.InsertarNodo(nodo)
			disponibles.InsertarNodo(new(Node))
			compradosUi.RemoveAll()
			content := ElementosComprados(vendidos, disponibles)
			compradosUi.Add(content)
		}
	}
}

func ElementosVendidos(l *List) *fyne.Container {
	compradosUi := container.New(layout.NewVBoxLayout())
	marca := widget.NewEntry()
	marca.SetPlaceHolder("Ingrese la Marca")
	modelo := widget.NewEntry()
	modelo.SetPlaceHolder("Ingrese el modelo")
	año := widget.NewEntry()
	año.SetPlaceHolder("Ingrese el año")
	colorW := widget.NewEntry()
	colorW.SetPlaceHolder("Ingrese el color")
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
		zonaPrincipal.Add(nuevoScroll)
	})
	compradosUi.Add(marca)
	compradosUi.Add(modelo)
	compradosUi.Add(año)
	compradosUi.Add(colorW)
	compradosUi.Add(tamaño)
	compradosUi.Add(botonAñadir)

	p := l.Head.Next
	i := 0
	zona := container.New(layout.NewHBoxLayout())
	lugaresVacios := widget.NewLabel("Coches Vendidos")
	zona.Add(lugaresVacios)
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 10
	compradosUi.Add(zona)
	compradosUi.Add(line)

	for p != nil {
		zona := container.New(layout.NewHBoxLayout())
		line := canvas.NewLine(color.White)
		line.StrokeWidth = 2
		zona.Add(widget.NewLabel("Modelo: " + p.coche.modelo))
		zona.Add(widget.NewLabel("Marca: " + p.coche.marca))
		zona.Add(widget.NewLabel("Color: " + p.coche.color))
		zona.Add(widget.NewLabel("Tamaño: " + p.coche.tamaño))
		zona.Add(widget.NewLabel(fmt.Sprintf("Año: %d", p.coche.año)))
		compradosUi.Add(zona)
		compradosUi.Add(line)
		p = p.Next
		i += 1
	}
	lugaresVacios.SetText(fmt.Sprintf("%d Coches han sido vendidos", i))
	return compradosUi
}

func ElementosDisponibles(l *List) *fyne.Container {
	p := l.Head.Next
	compradosUi := container.New(layout.NewVBoxLayout())
	i := 0
	zona := container.New(layout.NewHBoxLayout())
	lugaresVacios := widget.NewLabel("")
	zona.Add(lugaresVacios)
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 10
	compradosUi.Add(zona)
	compradosUi.Add(line)
	for p != nil {
		zona := container.New(layout.NewHBoxLayout())
		zona.Add(widget.NewLabel("Vacio"))
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

func ElementosVista() *fyne.Container {
	p := compradosGrandes.Head.Next
	p2 := compradosMedianos.Head.Next
	p3 := compradosChicos.Head.Next

	compradosUi := container.New(layout.NewVBoxLayout())
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 10
	compradosUi.Add(widget.NewLabel("Grandes"))
	compradosUi.Add(line)
	for p != nil {
		zona := container.New(layout.NewHBoxLayout())
		zona.Add(widget.NewLabel("Modelo: " + p.coche.modelo))
		zona.Add(widget.NewLabel("Marca: " + p.coche.marca))
		zona.Add(widget.NewLabel("Color: " + p.coche.color))
		zona.Add(widget.NewLabel("Tamaño: " + p.coche.tamaño))
		zona.Add(widget.NewLabel(fmt.Sprintf("Año: %d", p.coche.año)))
		lineDetalle := canvas.NewLine(color.White)
		lineDetalle.StrokeWidth = 2
		compradosUi.Add(zona)
		compradosUi.Add(lineDetalle)
		p = p.Next
	}

	compradosUiMed := container.New(layout.NewVBoxLayout())
	compradosUiMed.Add(widget.NewLabel("Medianos"))
	compradosUiMed.Add(line)
	for p2 != nil {
		zonaMed := container.New(layout.NewHBoxLayout())
		zonaMed.Add(widget.NewLabel("Modelo: " + p2.coche.modelo))
		zonaMed.Add(widget.NewLabel("Marca: " + p2.coche.marca))
		zonaMed.Add(widget.NewLabel("Color: " + p2.coche.color))
		zonaMed.Add(widget.NewLabel("Tamaño: " + p2.coche.tamaño))
		zonaMed.Add(widget.NewLabel(fmt.Sprintf("Año: %d", p2.coche.año)))
		lineDetalle := canvas.NewLine(color.White)
		lineDetalle.StrokeWidth = 2
		compradosUiMed.Add(zonaMed)
		compradosUiMed.Add(lineDetalle)
		p2 = p2.Next
	}

	compradosUiPeq := container.New(layout.NewVBoxLayout())
	compradosUiPeq.Add(widget.NewLabel("Pequeños"))
	compradosUiPeq.Add(line)
	for p3 != nil {
		zonaPeq := container.New(layout.NewHBoxLayout())
		zonaPeq.Add(widget.NewLabel("Modelo: " + p3.coche.modelo))
		zonaPeq.Add(widget.NewLabel("Marca: " + p3.coche.marca))
		zonaPeq.Add(widget.NewLabel("Color: " + p3.coche.color))
		zonaPeq.Add(widget.NewLabel("Tamaño: " + p3.coche.tamaño))
		zonaPeq.Add(widget.NewLabel(fmt.Sprintf("Año: %d", p3.coche.año)))
		lineDetalle := canvas.NewLine(color.White)
		lineDetalle.StrokeWidth = 2
		compradosUiPeq.Add(zonaPeq)
		compradosUiPeq.Add(lineDetalle)
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
	zona.Add(lugaresVacios)
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 10
	compradosUi.Add(zona)
	compradosUi.Add(line)

	compradosUi.Add(widget.NewLabel("Grandes"))
	for p != nil {
		zona := container.New(layout.NewHBoxLayout())
		zona.Add(widget.NewLabel("Modelo: " + p.coche.modelo))
		zona.Add(widget.NewLabel("Marca: " + p.coche.marca))
		zona.Add(widget.NewLabel("Color: " + p.coche.color))
		zona.Add(widget.NewLabel("Tamaño: " + p.coche.tamaño))
		zona.Add(widget.NewLabel(fmt.Sprintf("Año: %d", p.coche.año)))
		zona.Add(widget.NewButton("Vender", createButtonCallback(i, &compradosGrandes, vendidos, disponibles)))
		line := canvas.NewLine(color.White)
		line.StrokeWidth = 2
		compradosUi.Add(zona)
		compradosUi.Add(line)
		p = p.Next
		i += 1
	}

	compradosUi.Add(widget.NewLabel("Medianos"))
	for p2 != nil {
		zona := container.New(layout.NewHBoxLayout())
		zona.Add(widget.NewLabel("Modelo: " + p2.coche.modelo))
		zona.Add(widget.NewLabel("Marca: " + p2.coche.marca))
		zona.Add(widget.NewLabel("Color: " + p2.coche.color))
		zona.Add(widget.NewLabel("Tamaño: " + p2.coche.tamaño))
		zona.Add(widget.NewLabel(fmt.Sprintf("Año: %d", p2.coche.año)))
		zona.Add(widget.NewButton("Vender", createButtonCallback(i, &compradosMedianos, vendidos, disponibles)))
		line := canvas.NewLine(color.White)
		line.StrokeWidth = 2
		compradosUi.Add(zona)
		compradosUi.Add(line)
		p2 = p2.Next
		i += 1
	}

	compradosUi.Add(widget.NewLabel("Pequeños"))
	for p3 != nil {
		zona := container.New(layout.NewHBoxLayout())
		zona.Add(widget.NewLabel("Modelo: " + p3.coche.modelo))
		zona.Add(widget.NewLabel("Marca: " + p3.coche.marca))
		zona.Add(widget.NewLabel("Color: " + p3.coche.color))
		zona.Add(widget.NewLabel("Tamaño: " + p3.coche.tamaño))
		zona.Add(widget.NewLabel(fmt.Sprintf("Año: %d", p3.coche.año)))
		zona.Add(widget.NewButton("Vender", createButtonCallback(i, &compradosChicos, vendidos, disponibles)))
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

var vendidosVacio = new(Coche)
var vendidosNodoVacio = Node{*vendidosVacio, nil}
var vendidos = List{&vendidosNodoVacio}

func main() {
	disponiblesVacio := new(Coche)
	disponiblesNodoVacio := Node{*disponiblesVacio, nil}
	disponibles := List{&disponiblesNodoVacio}

	for i := 0; i < 100; i++ {
		cochevacio := new(Coche)
		cochevacio.año = i + 1
		disponibles.InsertarCoche(*cochevacio)
	}

	a := app.New()
	w := a.NewWindow("Venta de Autos")
	w.Resize(fyne.NewSize(800, 600))

	hello := widget.NewLabel("Venta de Autos")

	disponiblesUi := ElementosDisponibles(&disponibles)
	disponiblesScroll := container.NewVScroll(disponiblesUi)
	disponiblesScroll.SetMinSize(fyne.NewSize(500, 500))
	zonaPrincipal = container.NewVBox(disponiblesScroll)

	marca := widget.NewEntry()
	marca.SetPlaceHolder("Ingrese la Marca")
	modelo := widget.NewEntry()
	modelo.SetPlaceHolder("Ingrese el modelo")
	año := widget.NewEntry()
	año.SetPlaceHolder("Ingrese el año")
	color := widget.NewEntry()
	color.SetPlaceHolder("Ingrese el color")
	tamaño := widget.NewSelect([]string{"Grande", "Mediano", "Pequeño"}, func(val string) { tamañoSeleccionado = val })

	var botonAñadir *widget.Button
	var botonRandom *widget.Button

	refreshCompradosView := func() {
		compradosUi = ElementosComprados(&vendidos, &disponibles)
		compradosScroll := container.NewVScroll(compradosUi)
		compradosScroll.SetMinSize(fyne.NewSize(500, 300))
		zonaPrincipal.RemoveAll()
		zonaPrincipal.Add(ingresoDatos)
		zonaPrincipal.Add(compradosScroll)
	}

	botonAñadir = widget.NewButton("Comprar Coche", func() {
		añoInt, err := strconv.Atoi(año.Text)
		if err != nil {
			errorWidget.SetText("Error en el año: " + err.Error())
			return
		}
		var err1 error
		if tamañoSeleccionado == "Grande" {
			err1 = Comprar(añoInt, color.Text, marca.Text, modelo.Text, tamañoSeleccionado, &disponibles, &compradosGrandes)
		} else if tamañoSeleccionado == "Mediano" {
			err1 = Comprar(añoInt, color.Text, marca.Text, modelo.Text, tamañoSeleccionado, &disponibles, &compradosMedianos)
		} else if tamañoSeleccionado == "Pequeño" {
			err1 = Comprar(añoInt, color.Text, marca.Text, modelo.Text, tamañoSeleccionado, &disponibles, &compradosChicos)
		} else {
			errorWidget.SetText("Debe seleccionar un tamaño")
			return
		}
		if err1 != nil {
			errorWidget.SetText(err1.Error())
		} else {
			errorWidget.SetText("")
			marca.SetText("")
			modelo.SetText("")
			año.SetText("")
			color.SetText("")
		}
		refreshCompradosView()
	})

	botonRandom = widget.NewButton("Comprar Coche al azar", func() {
		coche := new(Coche)
		coche.Random()
		tamañoSel := coche.tamaño
		var err1 error
		if tamañoSel == "Grande" {
			err1 = ComprarCoche(coche, &disponibles, &compradosGrandes)
		} else if tamañoSel == "Mediano" {
			err1 = ComprarCoche(coche, &disponibles, &compradosMedianos)
		} else if tamañoSel == "Pequeño" {
			err1 = ComprarCoche(coche, &disponibles, &compradosChicos)
		}
		if err1 != nil {
			errorWidget.SetText(err1.Error())
		} else {
			errorWidget.SetText("")
		}
		refreshCompradosView()
	})

	ingresoDatos = container.New(layout.NewVBoxLayout(),
		marca, modelo, color, año, tamaño, botonAñadir, botonRandom)

	grid := container.New(layout.NewGridLayout(4),
		widget.NewButton("Comprar", func() {
			refreshCompradosView()
		}),
		widget.NewButton("Vender", func() {
			nuevo := ElementosVendidos(&vendidos)
			nuevoScroll := container.NewVScroll(nuevo)
			nuevoScroll.SetMinSize(fyne.NewSize(500, 500))
			zonaPrincipal.RemoveAll()
			zonaPrincipal.Add(nuevoScroll)
		}),
		widget.NewButton("Vista", func() {
			nuevo := ElementosVista()
			nuevoScroll := container.NewVScroll(nuevo)
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
		}))

	content := container.NewVBox(hello, errorWidget, grid, zonaPrincipal)
	w.SetContent(content)
	w.ShowAndRun()
}
