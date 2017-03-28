package main

import (
	"log"
	"testing"
)

func TestKPCos(t *testing.T) {
	for k := 0; k < 4; k++ {
		for p := 0; p < 6; p++ {
			log.Println("cos", p, k, kpCos(k, p))
			log.Println("sin", p, k, kpSin(k, p))
		}
	}
}
