package main

import (
	"context"
	"encoding/csv"
	"log"
	"net"
	"os"
	"strconv"
	"sync"

	pb "gobierno/proto/grpc-server/proto"
	"google.golang.org/grpc"
)

type cazarrecompensas struct {
	Id            string
	VentaMarina   int32
	VentaSubmundo int32
}

type gobiernoServer struct {
	pb.UnimplementedGobiernoServiceServer
	mu      sync.Mutex
	piratas []*pb.PirataInfo
	ventas  map[string]bool
	reputacion map[int32]int32
}

// Cargar piratas desde CSV al iniciar
func (s *gobiernoServer) cargarDatosIniciales(archivo string) error {
	file, err := os.Open(archivo)
	if err != nil {
		log.Fatalf("Error al abrir CSV: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	for _, record := range records[1:] { // Saltar encabezado
		recompensa, _ := strconv.ParseUint(record[2], 10, 64)
		s.piratas = append(s.piratas, &pb.PirataInfo{
			Id:            record[0],
			Nombre:        record[1],
			Recompensa:    recompensa,
			Peligrosidad:  record[3],
			Estado:        record[4],
		})
	}
	return nil
}

// Servicio gRPC: Obtener lista de piratas buscados
func (s *gobiernoServer) ObtenerListaPiratas(ctx context.Context, req *pb.SolicitudLista) (*pb.ListaPiratas, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var buscados []*pb.PirataInfo
	for _, p := range s.piratas {
		if p.Estado == "Buscado" {
			buscados = append(buscados, p)
		}
	}
	return &pb.ListaPiratas{Piratas: buscados}, nil
}

// Servicio gRPC: Reportar captura de pirata y actualizar reputación
func (s *gobiernoServer) ReportarCaptura(ctx context.Context, req *pb.ReporteCaptura) (*pb.ConfirmacionCaptura, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	reputacion, ok := s.reputacion[req.CazarrecompensasId]
	if !ok {
		reputacion = 100 // reputacion inicial
	}

	switch {
		case req.Estado == "Perdido":
	}
}

func (s *gobiernoServer) ModificarEstadoPirata(ctx context.Context, req *pb.ReporteCaptura) {
	for _, p := range s.piratas {
		if p.Id == req.PirataId && p.Estado == "Buscado" {
			p.Estado = req.Estado
			return nil
		}
	}
	return fmt.Errorf("Pirata con ID %d no encontrado o no está buscado", req.PirataId)
}

func main() {
	server := &gobiernoServer{}
	err := server.cargarDatosIniciales("piratas.csv") // Ruta del CSV
	if err!= nil {
		log.Fatalf("Error al cargar CSV: %v", err)
	}

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGobiernoServiceServer(s, server)
	log.Println("Gobierno Mundial escuchando en el puerto 50052")
	s.Serve(lis) // Iniciar el servidor gRPC
}