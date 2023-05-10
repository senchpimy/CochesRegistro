package main

import (
	"errors"
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
			compradosUi.AddObject(content)
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
		}
	})
	compradosUi.AddObject(marca)
	compradosUi.AddObject(modelo)
	compradosUi.AddObject(año)
	compradosUi.AddObject(colorW)
	compradosUi.AddObject(tamaño)
	compradosUi.AddObject(botonAñadir)
	p := l.Head.Next
	i := 0
	zona := container.New(layout.NewHBoxLayout())
	lugaresVacios := widget.NewLabel("Coches Vendidos")
	zona.AddObject(lugaresVacios)
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 10
	compradosUi.AddObject(zona)
	compradosUi.AddObject(line)
	for p != nil {
		zona := container.New(layout.NewHBoxLayout())
		line := canvas.NewLine(color.White)
		line.StrokeWidth = 2
		zona.AddObject(widget.NewLabel("Modelo: " + p.coche.modelo))
		zona.AddObject(widget.NewLabel("Marca: " + p.coche.marca))
		zona.AddObject(widget.NewLabel("Color: " + p.coche.color))
		zona.AddObject(widget.NewLabel("Tamaño: " + p.coche.tamaño))
		zona.AddObject(widget.NewLabel(fmt.Sprintf("Año: %d", p.coche.año)))
		compradosUi.AddObject(zona)
		compradosUi.AddObject(line)
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
	zona.AddObject(lugaresVacios)
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 10
	compradosUi.AddObject(zona)
	compradosUi.AddObject(line)
	for p != nil {
		zona := container.New(layout.NewHBoxLayout())
		//zona.AddObject(widget.NewLabel(fmt.Sprintf("Lugar #%d",i)))
		zona.AddObject(widget.NewLabel("Vacio"))
		line := canvas.NewLine(color.White)
		line.StrokeWidth = 2
		compradosUi.AddObject(zona)
		compradosUi.AddObject(line)
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
		zona.AddObject(widget.NewButton("Vender", createButtonCallback(i, &compradosGrandes, vendidos, disponibles)))
		line := canvas.NewLine(color.White)
		line.StrokeWidth = 2
		compradosUi.AddObject(zona)
		compradosUi.AddObject(line)
		p = p.Next
		i += 1
	}
	compradosUi.AddObject(widget.NewLabel("Medianos"))
	for p2 != nil {
		zona := container.New(layout.NewHBoxLayout())
		zona.AddObject(widget.NewLabel("Modelo: " + p2.coche.modelo))
		zona.AddObject(widget.NewLabel("Marca: " + p2.coche.marca))
		zona.AddObject(widget.NewLabel("Color: " + p2.coche.color))
		zona.AddObject(widget.NewLabel("Tamaño: " + p2.coche.tamaño))
		zona.AddObject(widget.NewLabel(fmt.Sprintf("Año: %d", p2.coche.año)))
		zona.AddObject(widget.NewButton("Vender", createButtonCallback(i, &compradosMedianos, vendidos, disponibles)))
		line := canvas.NewLine(color.White)
		line.StrokeWidth = 2
		compradosUi.AddObject(zona)
		compradosUi.AddObject(line)
		p2 = p2.Next
		i += 1
	}
	compradosUi.AddObject(widget.NewLabel("Pequeños"))
	for p3 != nil {
		zona := container.New(layout.NewHBoxLayout())
		zona.AddObject(widget.NewLabel("Modelo: " + p3.coche.modelo))
		zona.AddObject(widget.NewLabel("Marca: " + p3.coche.marca))
		zona.AddObject(widget.NewLabel("Color: " + p3.coche.color))
		zona.AddObject(widget.NewLabel("Tamaño: " + p3.coche.tamaño))
		zona.AddObject(widget.NewLabel(fmt.Sprintf("Año: %d", p3.coche.año)))
		zona.AddObject(widget.NewButton("Vender", createButtonCallback(i, &compradosChicos, vendidos, disponibles)))
		line := canvas.NewLine(color.White)
		line.StrokeWidth = 2
		compradosUi.AddObject(zona)
		compradosUi.AddObject(line)
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
			errorWidget.SetText(err.Error())
			fmt.Println("Error")
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
		zonaPrincipal.AddObject(ingresoDatos)
		zonaPrincipal.AddObject(compradosScroll)
	})
	botonRandom = widget.NewButton("Comprar Coche al azar", func() {
		coche := new(Coche)
		coche.Random()
		tamañoSel := coche.tamaño
		err1 := errors.New("")
		if tamañoSel == "Grande" {
			err1 = ComprarCoche(coche, &disponibles, &compradosGrandes)
		} else if tamañoSel == "Mediano" {
			err1 = ComprarCoche(coche, &disponibles, &compradosMedianos)
		} else if tamañoSel == "Pequeño" {
			err1 = ComprarCoche(coche, &disponibles, &compradosChicos)
		} else {
			fmt.Printf("Por algun motivo el valor no mathceo $%s$\n", tamañoSel)
		}

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
		zonaPrincipal.AddObject(ingresoDatos)
		zonaPrincipal.AddObject(compradosScroll)
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
			zonaPrincipal.AddObject(ingresoDatos)
			zonaPrincipal.AddObject(compradosScroll)
		}),
		widget.NewButton("Vender", func() {
			nuevo := ElementosVendidos(&vendidos)
			nuevoScroll := container.NewVScroll(nuevo)
			nuevoScroll.SetMinSize(fyne.NewSize(500, 500))
			zonaPrincipal.RemoveAll()
			zonaPrincipal.AddObject(nuevoScroll)
		}),
		widget.NewButton("Vista", func() {
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
	w.SetContent(content)

	w.ShowAndRun()
}
