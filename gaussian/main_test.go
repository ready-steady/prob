package gaussian

import (
	"math"
	"testing"

	"github.com/ready-steady/probability"
	"github.com/ready-steady/probability/uniform"
	"github.com/ready-steady/support/assert"
)

func TestCDF(t *testing.T) {
	x := []float64{
		-4.0, -3.5, -3.0, -2.5, -2.0, -1.5, -1.0, -0.5,
		0.0, 0.5, 1.0, 1.5, 2.0, 2.5, 3.0, 3.5, 4.0,
	}

	F := []float64{
		6.209665325776139e-03,
		1.222447265504470e-02,
		2.275013194817922e-02,
		4.005915686381709e-02,
		6.680720126885809e-02,
		1.056497736668553e-01,
		1.586552539314571e-01,
		2.266273523768682e-01,
		3.085375387259869e-01,
		4.012936743170763e-01,
		5.000000000000000e-01,
		5.987063256829237e-01,
		6.914624612740131e-01,
		7.733726476231317e-01,
		8.413447460685429e-01,
		8.943502263331446e-01,
		9.331927987311419e-01,
	}

	assert.AlmostEqual(probability.CDF(New(1, 2), x), F, t)
}

func TestInvCDF(t *testing.T) {
	F := []float64{
		0.00, 0.05, 0.10, 0.15, 0.20, 0.25, 0.30, 0.35, 0.40, 0.45, 0.50,
		0.55, 0.60, 0.65, 0.70, 0.75, 0.80, 0.85, 0.90, 0.95, 1.00,
	}

	x := []float64{
		math.Inf(-1),
		-1.411213406737868e+00,
		-1.320387891386150e+00,
		-1.259108347373447e+00,
		-1.210405308393228e+00,
		-1.168622437549020e+00,
		-1.131100128177010e+00,
		-1.096330116601892e+00,
		-1.063336775783950e+00,
		-1.031415336713768e+00,
		-1.000000000000000e+00,
		-9.685846632862315e-01,
		-9.366632242160501e-01,
		-9.036698833981082e-01,
		-8.688998718229899e-01,
		-8.313775624509796e-01,
		-7.895946916067714e-01,
		-7.408916526265525e-01,
		-6.796121086138498e-01,
		-5.887865932621319e-01,
		math.Inf(1),
	}

	assert.AlmostEqual(probability.InvCDF(New(-1, 0.25), F), x, t)
}

func BenchmarkCDF(b *testing.B) {
	distribution := New(0, 1)
	x := probability.Sample(distribution, 1000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		probability.CDF(distribution, x)
	}
}

func BenchmarkInvCDF(b *testing.B) {
	distribution := New(0, 1)
	F := probability.Sample(uniform.New(0, 1), 1000)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		probability.InvCDF(distribution, F)
	}
}

func BenchmarkSample(b *testing.B) {
	distribution := New(0, 1)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		distribution.Sample()
	}
}
