package main

import (
	"context"
	"log"
	"time"
	"fmt"

	pb "cazarrecompensas/proto/grpc-server/proto"

	"google.golang.org/grpc"
)

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

	// Ejemplo de pirata capturado
	pirata := &pb.PirataCapturado{
		Id:         "123",
		Nombre:     "Roronoa Zoro",
		Recompensa: 230_000_000,
		Peligrosidad: "ALTA", 
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



