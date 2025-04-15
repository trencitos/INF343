package main

import (
	"context"
	"log"
	"time"
	"math/rand"

	pb "cazarrecompensas/proto/grpc-server/proto"
	"google.golang.org/grpc"
)

// Función para reportar captura
func reportarPirata(pirataID string) {
	conn, err := grpc.Dial("gobierno:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar al Gobierno: %v", err)
	}
	defer conn.Close()

	client := pb.NewGobiernoServiceClient(conn)
	resp, err := client.ReportarCaptura(context.Background(), &pb.ReporteCaptura{
		CazarrecompensaId: "CR123",
		PirataId:          pirataID,
	})

	if err != nil {
		log.Printf("Error en reporte: %v", err)
		return
	}

	if resp.Exito {
		log.Println("✅ Captura confirmada por el Gobierno")
	} else {
		log.Printf("❌ Fallo en captura: %s", resp.Mensaje)
	}
}

func VenderAlSubmundo(pirata *pb.PirataCapturado) {
	// Conectarse al servidor gRPC de submundo
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) //"10.35.168.119:50051" para cuando las corramos en las VM de la U
	if err != nil {
		log.Fatalf("No se pudo conectar al Submundo: %v", err)
	}
	defer conn.Close()
	/*
	// crear el cliente gRPC para el submundo
	client := pb.NewSubmundoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	*/
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.Println("Bienvenido al sistema de cazarrecompensas!")

	// Consultar lista de piratas buscados
	connGob, err := grpc.Dial("localhost:50052", grpc.WithInsecure()) // 50052 es el puerto del gobierno mundial
	if err != nil {
		log.Fatalf("No se pudo conectar al Gobierno: %v", err)
	}
	defer connGob.Close()
	client := pb.NewGobiernoServiceClient(connGob)

	for {
		log.Println("Consultando lista de piratas buscados...")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		res, err := client.ObtenerListaPiratas(ctx, &pb.SolicitudLista{
			CazarrecompensaId: "CR123", // ID del cazarrecompensas
		})
		if err != nil {
			log.Fatalf("Error al obtener lista de piratas: %v", err)
		}

		cancel()

		if len(res.Piratas) == 0 {
			log.Println("No hay piratas buscados en este momento.")
			return // break
		}

		log.Println("Lista de piratas buscados:")
		for _, pirata := range res.Piratas {
			log.Printf("ID: %s, Nombre: %s, Recompensa: %d, Peligrosidad: %s", pirata.Id, pirata.Nombre, pirata.Recompensa, pirata.Peligrosidad)
		}

		// elegir un pirata al azar de la lista
		pirata := res.Piratas[rand.Intn(len(res.Piratas))]
		log.Printf("Pirata elegido: %s, con recompensa de %d", pirata.Nombre, pirata.Recompensa)
		
		numero := rand.Intn(2) // 0 o 1
		switch numero {
		case 0: // submundo
			log.Println("💰 Decisión: Vender al Submundo")
			// VenderAlSubmundo(pirata.Id)
		default: // marina
			log.Println("⚖️ Decisión: Entregar a la Marina")
			// venderMarina(pirata.Id)
		}
		
		time.Sleep(15 * time.Second) // Esperar 15 segundos antes de la siguiente consulta
	}

	/*
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
	*/
}



