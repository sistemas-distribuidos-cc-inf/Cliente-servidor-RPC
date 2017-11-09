package main

import (
   "net/rpc"
   "fmt"
   "log"
   "os"
)
   
type Args struct {
   Cpf string
}

func main() {

   var cpfString string
   fmt.Scanf("%s", &cpfString)
   args := &Args{Cpf: cpfString}
   var resposta bool
    
   cliente, erro := rpc.DialHTTP("tcp", "localhost:8080");

   if erro != nil {
      log.Fatal("Erro chamando servidor: ", erro);
      os.Exit(1)
   }

   erro = cliente.Call("Validador.ValidaCpf", args, &resposta)
   if erro != nil {
      log.Fatal("Erro chamando função: ", erro);
      os.Exit(1)
   }

   if resposta {

      fmt.Println("O CPF enviado é válido!")
   }else {

      fmt.Println("CPF inválido!")
   }

}
