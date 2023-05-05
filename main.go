package main

import (
	"image/color"
	"strconv"
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
)

func createButtonCallback(n int) func() {
    return func() {
        fmt.Println(n)
    }
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

func ElementosComprados(l* List) *fyne.Container {
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
		zona.AddObject(widget.NewLabel("Modelo: "+p.coche.modelo))
		zona.AddObject(widget.NewLabel("Marca: "+p.coche.marca))
		zona.AddObject(widget.NewLabel("Color: "+p.coche.color))
		zona.AddObject(widget.NewLabel(fmt.Sprintf("Año: %d",p.coche.año)))
		zona.AddObject(widget.NewButton("Vender",func (){}))
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


	for i:=0;i<2;i++{
		cochevacio:=new(Coche)
		cochevacio.año=i+1
		disponibles.InsertarCoche(*cochevacio)
	}

	a := app.New()
	w := a.NewWindow("Venta de Autos")

	hello := widget.NewLabel("Venta de Autos")
	error:=widget.NewLabel("")
	zonaPrincipal:=container.NewVBox()
	
	//Area de Disponibles
	disponiblesUi:=ElementosDisponibles(&disponibles)
	disponiblesScroll:=container.NewVScroll(disponiblesUi)
	disponiblesScroll.SetMinSize(fyne.NewSize(500, 500))
	//zonaPrincipal=container.NewVBox(error ,disponiblesScroll)
	zonaPrincipal=container.NewVBox(disponiblesScroll)

	//Area de ingreso
	ingresoDatos:=container.New(layout.NewVBoxLayout())
	marca := widget.NewEntry()
	marca.SetPlaceHolder("Ingrese la Marca")
	modelo := widget.NewEntry()
	modelo.SetPlaceHolder("Ingrese el modelo")
	año := widget.NewEntry()
	año.SetPlaceHolder("Ingrese el año")
	color := widget.NewEntry()
	color.SetPlaceHolder("Ingrese la color")
	botonAñadir:=widget.NewButton("Comprar Coche",func (){})
	botonRandom:=widget.NewButton("Comprar Coche al azar",func (){})
	botonAñadir=widget.NewButton("Comprar Coche",func (){
				añoInt,err:=strconv.Atoi(año.Text);
				if err!=nil{return}
				err1:=Comprar(añoInt,color.Text,marca.Text,modelo.Text,&disponibles,&comprados)
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
				ingresoDatos=container.New(layout.NewVBoxLayout(),
					marca,modelo,color,año,botonAñadir,botonRandom)
				compradosUi:=ElementosComprados(&comprados)
				compradosScroll:=container.NewVScroll(compradosUi)
				compradosScroll.SetMinSize(fyne.NewSize(500, 300))
				zonaPrincipal.RemoveAll()
				zonaPrincipal.AddObject(ingresoDatos)
				zonaPrincipal.AddObject(compradosScroll)
	})
	botonRandom=widget.NewButton("Comprar Coche al azar",func (){
		coche:=new(Coche)
		coche.Random()
		ComprarCoche(coche,&disponibles,&comprados)
	})
	ingresoDatos=container.New(layout.NewVBoxLayout(),
			marca,modelo,color,año,botonAñadir,botonRandom)


	//Area de Comprados
	//compradosUi:=container.New(layout.NewVBoxLayout())
	//compradosScroll:=container.NewVScroll(compradosUi)
	//compradosScroll.SetMinSize(fyne.NewSize(500, 500))

	//for i:=0;i<100;i++{
	//	zona:=container.New(layout.NewHBoxLayout())
	//	zona.AddObject(widget.NewLabel("Papu"))
	//	zona.AddObject(widget.NewButton(fmt.Sprintf("Hola %d", i), createButtonCallback(i)))
	//	compradosUi.AddObject(zona)
	//}
	//compradosUi:=ElementosDisponibles(&disponibles)
	//compradosScroll:=container.NewVScroll(compradosUi)
	//compradosScroll.SetMinSize(fyne.NewSize(500, 500))
	//zonaPrincipal:=container.NewVBox(compradosScroll)

	grid := container.New(layout.NewGridLayout(3), 
	widget.NewButton("Comprar", func() {
			zonaPrincipal.RemoveAll()
			compradosUi:=ElementosComprados(&comprados)
			compradosScroll:=container.NewVScroll(compradosUi)
			compradosScroll.SetMinSize(fyne.NewSize(500, 300))
			//zonaPrincipal.AddObject(error)
			zonaPrincipal.AddObject(ingresoDatos)
			zonaPrincipal.AddObject(compradosScroll)
		}),
		widget.NewButton("Vender", func() {zonaPrincipal.RemoveAll()}),
		widget.NewButton("Disponibles", func() {
			zonaPrincipal.RemoveAll()
			nuevo:=ElementosDisponibles(&disponibles)
			nuevoScroll:=container.NewVScroll(nuevo)
			nuevoScroll.SetMinSize(fyne.NewSize(500, 500))
			//zonaPrincipal.AddObject(error)
			zonaPrincipal.AddObject(nuevoScroll)
		}) )
	content:=container.NewVBox(hello,error,grid,zonaPrincipal)
	w.SetContent(content)

	w.ShowAndRun()
}

