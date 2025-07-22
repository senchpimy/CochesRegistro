# Venta de Autos

![Badge](https://img.shields.io/badge/License-MIT-blue) ![Badge](https://img.shields.io/badge/Version-1.0.0-green)

**Venta de Autos** es una aplicación de escritorio interactiva escrita en Go utilizando el framework [Fyne](https://fyne.io/). Permite gestionar la compra, venta y visualización de autos clasificados por tamaño (Grande, Mediano, Pequeño), y controlar el inventario disponible, vendido y comprado.

---

## 🧩 Características

* 🚗 Compra y venta de autos con formulario de datos.
* 🔄 Adquisición de coches al azar.
* 📊 Vista organizada de autos por categoría.
* 📦 Gestión de autos disponibles, comprados y vendidos.
* ❌ Detección y manejo de errores (por ejemplo, al ingresar año inválido).
* 🖱️ Interfaz visual totalmente interactiva gracias a Fyne.

---

## 🖼️ Interfaz

La aplicación está organizada en varias secciones:

* **Formulario de compra**: Permite ingresar la marca, modelo, año, color y tamaño del auto.
* **Listado de autos disponibles**: Espacios que pueden llenarse con nuevas compras.
* **Vista de autos comprados**: Autos organizados por tamaño, con botones para venderlos.
* **Listado de autos vendidos**: Autos eliminados del sistema con sus datos.
* **Panel de navegación**: Botones para cambiar entre vistas.

---

## 🚀 Instalación y ejecución

### Requisitos

* Go 1.18 o superior
* Fyne (`go get fyne.io/fyne/v2`)

### Clonar y ejecutar

```bash
git clone https://github.com/tu-usuario/venta-autos-go.git
cd venta-autos-go
go run .
```

---

## 🛠️ Estructura Interna

* `main()`: Punto de entrada, inicializa la interfaz principal.
* `ElementosComprados()`, `ElementosVendidos()`, `ElementosDisponibles()`, `ElementosVista()`: Renderizan visualmente el estado actual del inventario.
* `Comprar()`, `ComprarCoche()`: Funciones lógicas para mover autos entre listas.
* `List`, `Node`, `Coche`: Estructuras utilizadas para representar los datos del sistema.

---

## 📦 Dependencias

* [Fyne v2](https://pkg.go.dev/fyne.io/fyne/v2): Framework de GUI para Go.
* Librerías estándar de Go (`image/color`, `errors`, `fmt`, etc.)

---

## 🪪 Licencia

Este proyecto está licenciado bajo la licencia MIT. Consulta el archivo `LICENSE` para más detalles.
