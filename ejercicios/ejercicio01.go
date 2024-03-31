package ejercicios

import (
	"strconv"
)

func DevolverEntero(entero string)(enteronumero int, mensaje string){	
	
	enteronumero,errorconvert := strconv.Atoi(entero)
	if errorconvert == nil {
		if enteronumero > 100 {
			mensaje = "Es mayor a 100" 
		} else if enteronumero < 100 {
			mensaje = "Es menor a 100"
		}else {
			mensaje = "Es igual a 100"
		}
	} else{
		mensaje = errorconvert.Error()
	}

	return enteronumero,mensaje
}

