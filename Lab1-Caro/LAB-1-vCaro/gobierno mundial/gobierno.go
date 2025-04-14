package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	pb "cazarrecompensas/proto/grpc-server/proto"

	"google.golang.org/grpc"
)

type GobiernoMundial struct {
	pb.UnimplementedGobiernoMundialServer
	piratas           map[string]*pb.Pirata
	reputacion        map[string]int
	capturasTotales   int
	actividadSubmundo int
	mu                sync.Mutex
}

type Pirata struct {
	ID           string
	Nombre       string
	Recompensa   int
	Peligrosidad string
	Estado       string
}

// Lee piratas desde un archivo CSV al iniciar
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

// Listar piratas buscados
func (g *GobiernoMundial) ListarPiratas(ctx context.Context, req *pb.SolicitudVacia) (*pb.ListaPiratas, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	var lista []*pb.Pirata
	for _, p := range g.piratas {
		if strings.ToLower(p.Estado) == "buscado" {
			lista = append(lista, p)
		}
	}
	return &pb.ListaPiratas{Piratas: lista}, nil
}

// Validar que el pirata siga buscado
func (g *GobiernoMundial) ValidarCaptura(ctx context.Context, cap *pb.Captura) (*pb.Validacion, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	p, existe := g.piratas[cap.Pirata.Id]
	if !existe {
		return &pb.Validacion{Valido: false, Mensaje: "Pirata no registrado"}, nil
	}

	if strings.ToLower(p.Estado) != "buscado" {
		return &pb.Validacion{Valido: false, Mensaje: "Pirata ya capturado"}, nil
	}

	return &pb.Validacion{Valido: true, Mensaje: "Captura válida"}, nil
}

// Reporte de entrega: reputación y estado
func (g *GobiernoMundial) ReportarEntrega(ctx context.Context, r *pb.ReporteEntrega) (*pb.RespuestaEntrega, error) {
	g.mu.Lock()
	defer g.mu.Unlock()

	idC := r.IdCazarrecompensas
	resultado := r.Resultado

	// Ajustar reputación
	switch resultado {
	case "Exito":
		g.reputacion[idC]++
		g.capturasTotales++
	case "Fraude", "Escapado", "Confiscado":
		g.reputacion[idC]--
		if resultado == "Fraude" {
			g.actividadSubmundo++
		}
	}

	// Cambiar estado del pirata
	if p, ok := g.piratas[r.IdPirata]; ok {
		if resultado == "Exito" {
			p.Estado = "Capturado"
		} else {
			p.Estado = "Perdido"
		}
	}

	return &pb.RespuestaEntrega{Registrado: true}, nil
}

// Ajuste dinámico de recompensas (ejecútalo como goroutine)
func (g *GobiernoMundial) ajustarRecompensasPeriodicamente() {
	for {
		time.Sleep(30 * time.Second)

		g.mu.Lock()
		for _, p := range g.piratas {
			if strings.ToLower(p.Estado) == "buscado" {
				if g.capturasTotales > 10 {
					p.Recompensa = int32(math.Round(float64(p.Recompensa) * 0.8))
				}
				if g.actividadSubmundo > 5 {
					p.Recompensa = int32(math.Round(float64(p.Recompensa) * 1.2))
				}
			}
		}
		fmt.Println("Ajuste dinámico de recompensas aplicado.")
		g.capturasTotales = 0
		g.actividadSubmundo = 0
		g.mu.Unlock()
	}
}

// Alerta de tráfico ilegal (podrías integrarlo con una notificación a Marina)
func (g *GobiernoMundial) detectarTraficoIlegal() {
	for {
		time.Sleep(15 * time.Second)
		g.mu.Lock()
		if g.actividadSubmundo >= 5 {
			fmt.Println("⚠️ ALERTA: Alta actividad ilegal en el Submundo detectada.")
		}
		g.mu.Unlock()
	}
}

func main() {
	piratas, err := LeerPiratasDesdeCSV("piratas.csv")
	if err != nil {
		log.Fatalf("Error al cargar CSV: %v", err)
	}

	servidor := &GobiernoMundial{
		piratas:           piratas,
		reputacion:        make(map[string]int),
		capturasTotales:   0,
		actividadSubmundo: 0,
	}

	// Lanzar tareas de monitoreo
	go servidor.ajustarRecompensasPeriodicamente()
	go servidor.detectarTraficoIlegal()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error al escuchar: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGobiernoMundialServer(grpcServer, servidor)

	fmt.Println("👑 Gobierno Mundial escuchando en el puerto 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}
}
