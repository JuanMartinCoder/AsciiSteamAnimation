package particles

import (
	"math"
	"slices"
	"strings"
	"time"
)

type Particle struct {
	lifetime int64
	speed    float64

	x float64
	y float64
}

type ParticleParams struct {
	MaxLife  int64
	MaxSpeed float64

	ParticleCount int

	X int
	Y int

	nextPosition func(particle *Particle, deltaMs int64)
	ascii        func(x, y int, count [][]int) string
	reset        func(particle *Particle, params *ParticleParams)
}

type ParticleSystem struct {
	ParticleParams
	particles []*Particle

	lastime int64
}

func CreateParticleSystem(params ParticleParams) ParticleSystem {
	particles := make([]*Particle, 0)
	for i := 0; i < params.ParticleCount; i++ {
		particles = append(particles, &Particle{})
	}

	return ParticleSystem{
		ParticleParams: params,
		lastime:        time.Now().UnixMilli(),
		particles:      particles,
	}
}

func (p *ParticleSystem) Start() {
	for _, r := range p.particles {
		p.reset(r, &p.ParticleParams)
	}
}

// Funcion para actualizar el sistema de particulas en base al delta de tiempo entre la ultima actualizacion y ahora
func (p *ParticleSystem) Update() {
	delta := time.Now().UnixMilli() - p.lastime
	p.lastime = time.Now().UnixMilli()

	for _, r := range p.particles {
		p.nextPosition(r, delta)

		// Si al actualizar la particula se va del rango, se resetea la particula
		if r.y >= float64(p.Y) || r.x >= float64(p.X) || r.lifetime <= 0 {
			p.reset(r, &p.ParticleParams)
		}
	}
}

func (p *ParticleSystem) Show() []string {
	counts := make([][]int, 0)

	for row := 0; row < p.Y; row++ {
		count := make([]int, 0)
		for col := 0; col < p.X; col++ {
			count = append(count, 0)
		}
		counts = append(counts, count)
	}

	for _, r := range p.particles {
		row := int(math.Floor(r.y))
		col := int(math.Round(r.x))

		counts[row][col]++
	}

	out := make([][]string, 0)
	for r, row := range counts {
		out2 := make([]string, 0)
		for c := range row {
			out2 = append(out2, p.ascii(r, c, counts))
		}
		out = append(out, out2)
	}

	slices.Reverse(out)
	outStr := make([]string, 0)
	for _, row := range out {
		outStr = append(outStr, strings.Join(row, ""))
	}

	return outStr
}
