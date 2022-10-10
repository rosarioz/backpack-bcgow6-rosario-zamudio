package main

/*
Sus costos adicionales son:
Pequeño: El costo del producto (sin costo adicional)
Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.

Requerimientos:
- Crear una estructura “tienda” que guarde una lista de productos. 
- Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
- Crear una interface “Producto” que tenga el método “CalcularCosto”
- Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
- Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto.
- Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
Interface Producto:
El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
Interface Ecommerce:
 - El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
 - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda
*/

type ecommerce struct{
	products []string
}

type product struct{
	kind string
	name string
    price float64
}

type Product interface{
	CostsCalculator()
}

type Ecommerce interface{
	Total()
	Add()
}

func newProduct(kind string, name string, price float64) product{
	return product{kind: kind, name: name, price: price}
}

func(p product) CostsCalculator() float64{
	switch p.kind{
	case "pequeño":
		return p.price
	case "mediano":
        return p.price * 1.03
	case "grane":
		return p.price * 1.06 + 2500.0
	default:
        return 0
    }
}

