package main

import (
   "net/rpc"
   "net/http"
   "net"
   "fmt"
   "log"
   "strings"
   "strconv"
)
   

type Args struct {
   Cpf string
}

type Validador bool

func (t *Validador) ValidaCpf(args *Args, res *bool) error {
    var cpfInteger = []int{}
	calcDig1 := 0
	calcDig2 := 0
	
	cpfString := strings.SplitAfter((*args).Cpf, "")
	
	for _, i := range cpfString {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
        cpfInteger = append(cpfInteger, j)
    }
	count := 0
	for i := 10; i > 1; i-- {

		calcDig1 += cpfInteger[count]*i
		count++
	}
	count = 0
	for i := 11; i > 1; i-- {
		
		calcDig2 += cpfInteger[count]*i
		count++
	}
    checkDig1 := false
    checkDig2 := false
    if ((calcDig1*10)%11) >= 10 {
        checkDig1 = 0 == cpfInteger[9] 
    } else {
        checkDig1 = ((calcDig1*10)%11) == cpfInteger[9]
    }

    if ((calcDig2*10)%11) >= 10 {
        checkDig2 = 0 == cpfInteger[10] 
    } else {
        checkDig2 = ((calcDig2*10)%11) == cpfInteger[10]
    }
	
	
	if checkDig1 && checkDig2 {
		*res = true
	} else {
		*res = false
	}
    return nil
}
 

func main() {
   servicoValidarCpf := new(Validador)
   rpc.Register(servicoValidarCpf)
   rpc.HandleHTTP()
   porta := ":8080"

   l, err := net.Listen("tcp", porta)
   fmt.Println("Aguardando chamada na porta ", porta);

   if err != nil {
      log.Fatal("ERRO: ", err)
   }
   http.Serve(l, nil)
}
