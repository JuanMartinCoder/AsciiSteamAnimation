package particles

import (
	"math"
	"math/rand"
)

type Fire struct {
	ParticleSystem
}

func ascii(x, y int, counts [][]int) string {
	count := counts[x][y]
	if count < 3 {
		return " "
	}
	if count < 6 {
		return "."
	}
	if count < 9 {
		return ":"
	}
	if count < 12 {
		return "{"
	}
	return "}"
}

func reset(particle *Particle, params *ParticleParams) {
	particle.lifetime = int64(math.Floor(float64(params.MaxLife) * rand.Float64()))
	particle.speed = params.MaxSpeed * rand.Float64()

	// utilizar una distribucion normal para asignar una posicion a la particula
	// Mientras mas alejanda del centro de la animacion, menos "intensidad" va a tener
	particle.y = 0

	middle := math.Floor(float64(params.X) / 2)
	particle.x = middle + math.Max(-middle, math.Min(rand.NormFloat64(), middle))
}

func nextPosition(particle *Particle, deltaMs int64) {
	particle.lifetime -= deltaMs

	if particle.lifetime <= 0 {
		return
	}

	percent := (float64(deltaMs) / 1000.0)
	particle.y += particle.speed * percent
}

func NewFire(width, height int) Fire {
	params := ParticleParams{
		MaxLife:  7000,
		MaxSpeed: 1,

		ParticleCount: 1000,
		reset:         reset,
		nextPosition:  nextPosition,
		ascii:         ascii,

		X: width,
		Y: height,
	}

	return Fire{
		ParticleSystem: CreateParticleSystem(params),
	}
}
