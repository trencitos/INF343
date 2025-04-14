package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	pb "cazarrecompensas/proto/grpc-server/proto"

	"google.golang.org/grpc"
)

type Pirata struct {
	ID           string
	Nombre       string
	Recompensa   int
	Peligrosidad string
	Estado       string
}

func LeerPiratasDesdeCSV(path string) ([]Pirata, error) {
	archivo, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error al abrir el archivo CSV: %w", err)
	}
	defer archivo.Close()

	lector := csv.NewReader(archivo)
	filas, err := lector.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error al leer el archivo CSV: %w", err)
	}

	var piratas []Pirata

	for i, fila := range filas {
		if i == 0 {
			continue // saltar encabezado
		}

		recompensa, err := strconv.Atoi(fila[2])
		if err != nil {
			return nil, fmt.Errorf("error en recompensa (fila %d): %w", i+1, err)
		}

		pirata := Pirata{
			ID:           fila[0],
			Nombre:       fila[1],
			Recompensa:   recompensa,
			Peligrosidad: fila[3],
			Estado:       fila[4],
		}

		piratas = append(piratas, pirata)
	}

	return piratas, nil
}

func main() {
	fmt.Println("Bienvenido al sistema de cazarrecompensas!")
	// log.Println("Iniciando cliente...")

	// Conectarse al servidor gRPC de submundo
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // //"10.35.168.119:50051" para cuando las corramos en las VM de la U
	if err != nil {
		log.Fatalf("No se pudo conectar al Submundo: %v", err)
	}
	defer conn.Close()

	// crear el cliente gRPC
	client := pb.NewSubmundoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	piratas, err := LeerPiratasDesdeCSV("piratas.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, p := range piratas {
		fmt.Printf("Pirata: %+v\n", p)
	}

	respuesta, err := client.ComprarPirata(ctx, pirata)
	if err != nil {
		log.Fatalf("Error al vender pirata al Submundo: %v", err)
	}

	log.Println("Resultado de la venta al Submundo:")
	log.Printf("Éxito: %v", respuesta.Exito)
	log.Printf("Mensaje: %s", respuesta.Mensaje)
	log.Printf("Monto pagado: %d", respuesta.MontoPagado)
	log.Printf("¿Fue fraude?: %v", respuesta.FueFraude)

	// reader := bufio.NewReader(os.Stdin) // Para capturar entrada del usuario
}
