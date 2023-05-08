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

func createButtonCallback(n int, comprados* List,vendidos* List,disponibles* List) func() {
    return func() {
	nodo:=comprados.Obtener(n)	
	if nodo!=nil{
		vendidos.InsertarNodo(nodo)
		disponibles.InsertarNodo(new(Node))
		compradosUi.RemoveAll()
		content:=ElementosComprados(comprados,vendidos,disponibles)
		compradosUi.AddObject(content)
	}
    }
}

func ElementosVendidos(l* List) *fyne.Container {
	p := l.Head.Next
	compradosUi:=container.New(layout.NewVBoxLayout())
	i:=0
	zona:=container.New(layout.NewHBoxLayout())
	lugaresVacios:=widget.NewLabel("Coches Vendidos")
	zona.AddObject(lugaresVacios)
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 10
	compradosUi.AddObject(zona)
	compradosUi.AddObject(line)
	for p != nil {
		zona:=container.New(layout.NewHBoxLayout())
		line := canvas.NewLine(color.White)
		line.StrokeWidth = 2
		zona.AddObject(widget.NewLabel("Modelo: "+p.coche.modelo))
		zona.AddObject(widget.NewLabel("Marca: "+p.coche.marca))
		zona.AddObject(widget.NewLabel("Color: "+p.coche.color))
		zona.AddObject(widget.NewLabel("Tamaño: "+p.coche.tamaño))
		zona.AddObject(widget.NewLabel(fmt.Sprintf("Año: %d",p.coche.año)))
		compradosUi.AddObject(zona)
		compradosUi.AddObject(line)
		p = p.Next
		i+=1
	}
	lugaresVacios.SetText(fmt.Sprintf("%d Coches han sido vendidos",i))
	return compradosUi
}

func ElementosDisponibles(l* List) *fyne.Container {
	p := l.Head.Next
	compradosUi:=container.New(layout.NewVBoxLayout())
	i:=0
	zona:=container.New(layout.NewHBoxLayout())
	lugaresVacios:=widget.NewLabel("")
	zona.AddObject(lugaresVacios)
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 10
	compradosUi.AddObject(zona)
	compradosUi.AddObject(line)
	for p != nil {
		zona:=container.New(layout.NewHBoxLayout())
		//zona.AddObject(widget.NewLabel(fmt.Sprintf("Lugar #%d",i)))
		zona.AddObject(widget.NewLabel("Vacio"))
		line := canvas.NewLine(color.White)
		line.StrokeWidth = 2
		compradosUi.AddObject(zona)
		compradosUi.AddObject(line)
		p = p.Next
		i+=1
	}
	lugaresVacios.SetText(fmt.Sprintf("Hay %d lugares vacios",i))
	return compradosUi
}

func ElementosComprados(l* List,vendidos* List,disponibles* List) *fyne.Container {
	p := l.Head.Next
	compradosUi:=container.New(layout.NewVBoxLayout())
	i:=0
	zona:=container.New(layout.NewHBoxLayout())
	lugaresVacios:=widget.NewLabel("")
	zona.AddObject(lugaresVacios)
	line := canvas.NewLine(color.White)
	line.StrokeWidth = 10
	compradosUi.AddObject(zona)
	compradosUi.AddObject(line)
	for p != nil {
		zona:=container.New(layout.NewHBoxLayout())
		zona.AddObject(widget.NewLabel("Modelo: "+p.coche.modelo))
		zona.AddObject(widget.NewLabel("Marca: "+p.coche.marca))
		zona.AddObject(widget.NewLabel("Color: "+p.coche.color))
		zona.AddObject(widget.NewLabel("Tamaño: "+p.coche.tamaño))
		zona.AddObject(widget.NewLabel(fmt.Sprintf("Año: %d",p.coche.año)))
		zona.AddObject(widget.NewButton("Vender",createButtonCallback(i,l,vendidos,disponibles)))
		line := canvas.NewLine(color.White)
		line.StrokeWidth = 2
		compradosUi.AddObject(zona)
		compradosUi.AddObject(line)
		p = p.Next
		i+=1
	}
	lugaresVacios.SetText(fmt.Sprintf("Hay %d coches comprados",i))

	return compradosUi
}

var zonaPrincipal = container.NewVBox()
var ingresoDatos = container.NewVBox()
var compradosUi = container.NewVBox()
var tamañoSeleccionado=""

func main() {
	disponiblesVacio:=new(Coche)
	compradosVacio:=new(Coche)
	vendidosVacio:=new(Coche)

	disponiblesNodoVacio:=Node{*disponiblesVacio,nil}
	compradosNodoVacio:=Node{*compradosVacio,nil}
	vendidosNodoVacio:=Node{*vendidosVacio,nil}

	disponibles:=List{&disponiblesNodoVacio}
	comprados:=List{&compradosNodoVacio}
	vendidos:=List{&vendidosNodoVacio}
	Show(&vendidos)


	for i:=0;i<100;i++{
		cochevacio:=new(Coche)
		cochevacio.año=i+1
		disponibles.InsertarCoche(*cochevacio)
	}

	a := app.New()
	w := a.NewWindow("Venta de Autos")

	hello := widget.NewLabel("Venta de Autos")
	error:=widget.NewLabel("")
	
	//Area de Disponibles
	disponiblesUi:=ElementosDisponibles(&disponibles)
	disponiblesScroll:=container.NewVScroll(disponiblesUi)
	disponiblesScroll.SetMinSize(fyne.NewSize(500, 500))
	//zonaPrincipal=container.NewVBox(error ,disponiblesScroll)
	zonaPrincipal=container.NewVBox(disponiblesScroll)

	//Area de ingreso
	marca := widget.NewEntry()
	marca.SetPlaceHolder("Ingrese la Marca")
	modelo := widget.NewEntry()
	modelo.SetPlaceHolder("Ingrese el modelo")
	año := widget.NewEntry()
	año.SetPlaceHolder("Ingrese el año")
	color := widget.NewEntry()
	color.SetPlaceHolder("Ingrese la color")
	tamaño:=widget.NewSelect([]string{"Grande","Mediano","Pequeño"},func(val string){tamañoSeleccionado=val})
	botonAñadir:=widget.NewButton("Comprar Coche",func (){})
	botonRandom:=widget.NewButton("Comprar Coche al azar",func (){})
	botonAñadir=widget.NewButton("Comprar Coche",func (){
		añoInt,err:=strconv.Atoi(año.Text);
		if err!=nil{return}
		err1:=Comprar(añoInt,color.Text,marca.Text,modelo.Text,tamañoSeleccionado,&disponibles,&comprados)
		if err1!=nil{
			error.SetText(err1.Error())
		}
		marca = widget.NewEntry()
		marca.SetPlaceHolder("Ingrese la Marca")
		modelo = widget.NewEntry()
		modelo.SetPlaceHolder("Ingrese el modelo")
		año = widget.NewEntry()
		año.SetPlaceHolder("Ingrese el año")
		color = widget.NewEntry()
		color.SetPlaceHolder("Ingrese la color")
		tamaño=widget.NewSelect([]string{"Grande","Mediano","Pequeño"},func(val string){tamañoSeleccionado=val})
		ingresoDatos:=container.New(layout.NewVBoxLayout(),
			marca,modelo,color,año,tamaño,botonAñadir,botonRandom)
		compradosUi=ElementosComprados(&comprados,&vendidos,&disponibles)
		compradosScroll:=container.NewVScroll(compradosUi)
		compradosScroll.SetMinSize(fyne.NewSize(500, 300))
		zonaPrincipal.RemoveAll()
		zonaPrincipal.AddObject(ingresoDatos)
		zonaPrincipal.AddObject(compradosScroll)
	})
	botonRandom=widget.NewButton("Comprar Coche al azar",func (){
		coche:=new(Coche)
		coche.Random()
		err1:=ComprarCoche(coche,&disponibles,&comprados)
		if err1!=nil{
			error.SetText(err1.Error())
		}
		marca = widget.NewEntry()
		marca.SetPlaceHolder("Ingrese la Marca")
		modelo = widget.NewEntry()
		modelo.SetPlaceHolder("Ingrese el modelo")
		año = widget.NewEntry()
		año.SetPlaceHolder("Ingrese el año")
		color = widget.NewEntry()
		color.SetPlaceHolder("Ingrese la color")
		tamaño=widget.NewSelect([]string{"Grande","Mediano","Pequeño"},func(val string){tamañoSeleccionado=val})
		ingresoDatos:=container.New(layout.NewVBoxLayout(),
			marca,modelo,color,año,tamaño,botonAñadir,botonRandom)
		compradosUi=ElementosComprados(&comprados,&vendidos,&disponibles)
		compradosScroll:=container.NewVScroll(compradosUi)
		compradosScroll.SetMinSize(fyne.NewSize(500, 300))
		zonaPrincipal.RemoveAll()
		zonaPrincipal.AddObject(ingresoDatos)
		zonaPrincipal.AddObject(compradosScroll)
	})
	ingresoDatos=container.New(layout.NewVBoxLayout(),
			marca,modelo,color,año,tamaño,botonAñadir,botonRandom)


	//Area de Vendidos

	grid := container.New(layout.NewGridLayout(3), 
	widget.NewButton("Comprar", func() {
			compradosUi=ElementosComprados(&comprados,&vendidos,&disponibles)
			compradosScroll:=container.NewVScroll(compradosUi)
			compradosScroll.SetMinSize(fyne.NewSize(500, 300))
			zonaPrincipal.RemoveAll()
			zonaPrincipal.AddObject(ingresoDatos)
			zonaPrincipal.AddObject(compradosScroll)
		}),
		widget.NewButton("Vender", func() {
			nuevo:=ElementosVendidos(&vendidos)
			nuevoScroll:=container.NewVScroll(nuevo)
			nuevoScroll.SetMinSize(fyne.NewSize(500, 500))
			zonaPrincipal.RemoveAll()
			zonaPrincipal.AddObject(nuevoScroll)
		}),
		widget.NewButton("Disponibles", func() {
			nuevo:=ElementosDisponibles(&disponibles)
			nuevoScroll:=container.NewVScroll(nuevo)
			nuevoScroll.SetMinSize(fyne.NewSize(500, 500))
			zonaPrincipal.RemoveAll()
			zonaPrincipal.AddObject(nuevoScroll)
		}) )
	content:=container.NewVBox(hello,error,grid,zonaPrincipal)
	w.SetContent(content)

	w.ShowAndRun()
}

