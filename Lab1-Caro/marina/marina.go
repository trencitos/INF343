package main

import (
"context"
"log"
"time"

pb "marina/proto/grpc-server/proto"
"google.golang.org/grpc"
)

type Marina struct {
    pb.UnimplementedMarinaServer
    piratasCapturados  map[string]bool
    reputacion         map[string]int
    capturasRecientes  int
}


func BuscarPirata(pirata []*pb.PirataInfo) {


}

// procesa la venta del pirata
// verifica si el pirata aún es buscado
// verifica la reputación del cazarrecompensas
// rechaza si el pirata ya fue capturado o si la reputación es menor a 30

func (s *servidorMarina) ComprarPirata(ctx context.Context, pirata *pb.PirataCapturado) (*pb.ResultadoCompra, error) {
    pirata := req.Pirata
    idCazador := req.IdCazarrecompensas

	log.Printf("Recibiendo Pirata Capturado: %s con recompensa de %d berries", pirata.Nombre, pirata.Recompensa)

	//Verificar si ya fue buscado y entregado
    if m.piratasCapturados[pirata.Id] {
        return &pb.ResultadoEntrega{
            Exito: false,
            RecompensaObtenida: 0,
            Mensaje:"El pirata ya fue entregado anteriormente.",
        }, nil
    }

    // Verificar reputación cazarrecompensas
    if m.reputacion[idCazador] < 30 {
        return &pb.ResultadoEntrega{
            Exito: false,
            RecompensaObtenida: 0,
            Mensaje:"Entrega rechazada por mala reputación.",
        }, nil
    }

    recompensa := pirata.Recompensa

    // Reducir la recompensa si hay mucha actividad
    if m.capturasRecientes > 5 {
        recompensa = recompensa / 2
        log.Println("Recompensa reducida por alta actividad.")
    }

    // Registrar entrega
    m.piratasCapturados[pirata.Id] = true
    m.reputacion[idCazador]++   // bonificación por éxito
    m.capturasRecientes++

    return &pb.ResultadoEntrega{
        Exito:              true,
        RecompensaObtenida: recompensa,
        Mensaje: "✅ Entrega aceptada por la Marina.",
    }, nil
}

// Rutina para reiniciar capturas recientes cada minuto
func (m *MarinaServer) ResetCapturasPeriodicamente() {
    for {
        time.Sleep(1 * time.Minute)
        log.Println("Reiniciando contador de capturas recientes.")
        m.capturasRecientes = 0
    }
}
  

func main() {
log.Println("Iniciando marina...")
rand.Seed(time.Now().UnixNano())

// Crear un listener para el servidor gRPC
lis, err := net.Listen("tcp", ":50053")
if err != nil {
log.Fatalf("Servidor falla al escuchar: %v", err)
}

//crear el cliente gRPC para el gobierno
connGob, err := grpc.Dial("gobierno:50052", grpc.WithInsecure())
if err != nil {
log.Fatalf("No se pudo conectar al Gobierno: %v", err)
}
defer connGob.Close()
clientGob := pb.NewGobiernoServiceClient(connGob)

// Crear un servidor gRPC
s := grpc.NewServer()
pb.RegisterMarinaServiceServer(s, &servidorMarina{gobClient: clientGob})

log.Println("Servidor gRPC de la Marina escuchando en el puerto 50053...")
// Iniciar el servidor gRPC
if err := s.Serve(lis); err != nil {
log.Fatalf("Fallo al iniciar el servidor gRPC: %v", err)
}

defer s.Stop() // Detener el servidor al finalizar
}