package skyline

import "testing"

func benchmarkSkyline(b *testing.B, size int) {
	b.Helper()

	for range b.N {
		p := Packer{}
		p.Initialize(size, size)

		for range size * size {
			_, _, err := p.AddRect(1, 1)
			if err != nil {
				panic(err)
			}
		}
	}

	b.ReportAllocs()
}

func BenchmarkSkyline_100(b *testing.B)   { benchmarkSkyline(b, 10) }
func BenchmarkSkyline_400(b *testing.B)   { benchmarkSkyline(b, 20) }
func BenchmarkSkyline_2500(b *testing.B)  { benchmarkSkyline(b, 50) }
func BenchmarkSkyline_10000(b *testing.B) { benchmarkSkyline(b, 100) }
