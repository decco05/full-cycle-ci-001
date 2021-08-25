package main

import "testing"

func TesteSoma(t *testing.T) {
   total := Soma(15, 15)

   if total != 30{
     t.Errorf("Resultado da soma é inválido. Resutaldo Esperado = %d, Soma da funcção = %ds", 30, total);
   }
   
}