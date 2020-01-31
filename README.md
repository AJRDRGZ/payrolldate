# Payroll dates

Go package for calculate payroll dates.

Install

```bash
    go get -u github.com/AJRDRGZ/payrolldate
```

## Español

El paquete permite contar los días entre fechas teniendo en cuenta que todos los meses tienen 30 días. Esto es debido a la ley colombiana para el proceso de nómina.

### Ejemplo de uso
```go
package main

import (
    "fmt"

    pd "github.com/AJRDRGZ/payrolldate"
) 

func main() {
	start := pd.Date("2020-02-05")
	end := pd.Date("2020-03-03")
	days := pd.Days360(start, end)
	fmt.Println(days)

	// Output:
	// 29
}
```