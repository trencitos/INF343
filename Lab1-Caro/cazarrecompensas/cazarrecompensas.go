package main

import (
	"context"
	"log"
	"time"
	"math/rand"

	pb "cazarrecompensas/proto/grpc-server/proto"
	"google.golang.org/grpc"
)

// Función para reportar captura al gobierno mundial y actualizar reputación
func reportarPirata(pirata *pbPirataCapturado) {
	conn, err := grpc.Dial("gobierno:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("No se pudo conectar al Gobierno: %v", err)
	}
	defer conn.Close()

	client := pb.NewGobiernoServiceClient(conn)
	resp, err := client.ReportarCaptura(context.Background(), &pb.ReporteCaptura{
		CazarrecompensaId: "CR123",
		PirataId:         pirata.Id,
		Estado:           pirata.Estado,
		Caso:        pirata.Caso,
		Entrega:     pirata.Entrega,
	})

	if err != nil {
		log.Printf("Error al enviar reporte de captura al gobierno: %v", err)
		return
	}

	reputacion := resp.Reputacion

	log.Printf("se reportó la captura de %s al gobierno. Reputación actual: %d", pirata.Nombre, reputacion)
}

func VenderAlSubmundo(pirata *pb.PirataCapturado) *pb.ResultadoCompra {
	// Conectarse al servidor gRPC de submundo
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) //"10.35.168.119:50051" para cuando las corramos en las VM de la U
	if err != nil {
		log.Fatalf("No se pudo conectar al Submundo: %v", err)
	}
	defer conn.Close()

	// crear el cliente gRPC para el submundo
	client := pb.NewSubmundoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	/*
	req := &pb.PirataCapturado{
		Id:         pirata.Id,
		Nombre:     pirata.Nombre,
		Recompensa: pirata.Recompensa,
		CazarrecompensasId: "CR123",
		Peligrosidad: pirata.Peligrosidad,
	}
	*/
	res, err := client.ComprarPirata(ctx, pirata)
	if err != nil {
		log.Fatalf("Error al vender pirata al Submundo: %v", err)
		return &pb.ResultadoCompra{
			Estado: "Error",
			Mensaje: "Error al vender pirata al Submundo",
		}
	}

	log.Printf("Resultado venta: %s | Mensaje: %s", res.Estado, res.Mensaje)

	switch res.Estado {
		case "Exito":
			log.Printf("💰 Pirata vendido al Submundo por %d berries", res.MontoPagado)
			// a lo mejor aquí se puede reportar el estado y luego actualizar rep
		case "Fraude":
			// a lo mejor aquí se puede reportar el estado y luego actualizar rep
			log.Printf("🚨 Fraude: El cazarrecompensas no recibió pago.")
	}
	return res
}

func EntregarAMarina(pirata *pb.PirataCapturado) *pb.ResultadoCompra {
	// Conectarse al servidor gRPC de la Marina
	log.Printf("Entregando a %s a la Marina...", pirata.Nombre)
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure()) //"10.35.168.119:50051" para cuando las corramos en las VM de la U
	if err != nil {
		log.Fatalf("No se pudo conectar a la Marina: %v", err)
	}
	defer conn.Close()

	// crear el cliente gRPC para la marina
	client := pb.NewMarinaServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	
	res, err := client.ComprarPirata(ctx, pirata)
	if err != nil {
		log.Fatalf("Error al vender pirata a la Marina: %v", err)
		return &pb.ResultadoCompra{
			Estado: "Error",
			Mensaje: "Error al vender pirata a la Marina",
		}
	}

	return &pb.ResultadoCompra{
		Estado:  "Entregado",
		Mensaje: "Pirata entregado exitosamente a la Marina.",
		Nombre:  pirata.Nombre,
	}	
}

func ProbabilidadEscape(peligrosidad string) float32 {
	var probEscape float32

	if peligrosidad == "Alta" {
		probEscape = 0.3
	} else if peligrosidad == "Media" {
		probEscape = 0.2
	} else {
		probEscape = 0.1
	}
	return probEscape
}

func simularTransporte(pirata *pb.PirataInfo) *pb.ResultadoCompra {
	log.Printf("🛡️ Iniciando transporte del pirata %s...", pirata.Nombre)
	res := &pb.ResultadoCompra{
		Nombre:     pirata.Nombre,
	}
	probEscape := ProbabilidadEscape(pirata.Peligrosidad)
	time.Sleep(2 * time.Second) // Simular tiempo de transporte

	if rand.Float32() < probEscape {
		log.Printf("🚨 El pirata %s logró escapar durante el transporte!", pirata.Nombre)
		res.Estado = "Perdido"
		res.Caso = "Escapado"
		res.Mensaje = "El pirata logró escapar durante el transporte."
		// a lo mejor aquí se puede reportar el estado y luego actualizar rep
		return res
	}

	if pirata.Recompensa > 200000000 && rand.Float32() < 0.35 {
		log.Println("¡Ataque de mercenarios! El pirata fue rescatado.")
		res.Estado = "Perdido"
		res.Caso = "Ataque"
		res.Mensaje = "¡Ataque de mercenarios! El pirata fue rescatado."
		// a lo mejor aquí se puede reportar el estado y luego actualizar rep
		return res
	}

	piratacapturado := &pb.PirataCapturado{
		Id:         pirata.Id,
		Nombre:     pirata.Nombre,
		Recompensa: int32(pirata.Recompensa),
		CazarrecompensasId: "CR123",
		Peligrosidad: pirata.Peligrosidad,
	}

	numero := rand.Intn(2)
	switch numero {
	case 0:
		log.Println("💰 Decisión: Vender al Submundo")
		return VenderAlSubmundo(piratacapturado)
	default:
		log.Println("⚖️ Decisión: Entregar a la Marina")
		return EntregarAMarina(piratacapturado)
	}
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
		
		// intentar capturar al pirata
		// verificar redada ? si hay redada reportar estado al gobierno con los datos
		// del pirata y si no hay redada ver lógica para venderlo al submundo o marina
		// reportar pirata y entonces venderlo al submundo o entregarlo a la marina
		// simular el transporte del pirata
		// resultado := simularTransporte(pirata)
		// enviarInformeGobierno(pirata, res)
		
		time.Sleep(15 * time.Second) // Esperar 15 segundos antes de la siguiente consulta
	}
	log.Println("Fin del programa.")
}



