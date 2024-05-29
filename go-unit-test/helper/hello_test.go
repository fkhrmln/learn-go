package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Before Test")

	m.Run()

	fmt.Println("After Test")
}

func TestHelloFakhri(t *testing.T) {
	result := Hello("Fakhri")

	if result != "Hi Fakhri" {
		t.Fail()
	}

	fmt.Println("TestHelloFakhri Finish")
}

func TestHelloRifky(t *testing.T) {
	result := Hello("Rifky")

	if result != "Hi Rifky" {
		t.FailNow()
	}

	fmt.Println("TestHelloRifky Finish")
}

func TestHelloAudry(t *testing.T) {
	result := Hello("Audry")

	if result != "Hi Audry" {
		t.Error("TestHelloAudry should return Hello Audry")
	}

	fmt.Println("TestHelloAudry Finish")
}

func TestHelloBobby(t *testing.T) {
	result := Hello("Bobby")

	if result != "Hi Bobby" {
		t.Fatal("TestHelloBobby should return Hello Bobby")
	}

	fmt.Println("TestHelloBobby Finish")
}

func TestHelloWorldAssert(t *testing.T) {
	result := Hello("World")

	assert.Equal(t, "Hi World", result, "TestHelloWorldAssert should return Hello World")

	fmt.Println("TestHelloWorldAssert Finish")
}

func TestHelloWorldRequire(t *testing.T) {
	result := Hello("World")

	require.Equal(t, "Hi World", result, "TestHelloWorldRequire should return Hello World")

	fmt.Println("TestHelloWorldRequire Finish")
}

func TestHelloWorldSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("This test can't run in linux")
	}

	result := Hello("World")

	require.Equal(t, "Hi World", result, "TestHelloWorldSkip should return Hello World")

	fmt.Println("TestHelloWorldSkip Finish")
}

func TestSubHello(t *testing.T) {
	t.Run("TestSubHelloFakhri", func(t *testing.T) {
		result := Hello("Fakhri")

		require.Equal(t, "Hello Fakhri", result, "TestSubHelloFakhri should return Hello Fakhri")
	})

	t.Run("TestSubHelloRifky", func(t *testing.T) {
		result := Hello("Rifky")

		require.Equal(t, "Hello Rifky", result, "TestSubHelloRifky should return Hello Rifky")
	})

	t.Run("TestSubHelloAudry", func(t *testing.T) {
		result := Hello("Audry")

		require.Equal(t, "Hello Audry", result, "TestSubHelloAudry should return Hello Audry")
	})
}

func TestTableHello(t *testing.T) {
	tests := []struct {
		TestName string
		Name     string
		Expected string
	}{
		{
			"Fakhri",
			"TestSubHelloFakhri",
			"Hello Fakhri",
		},
		{
			"Rifky",
			"TestSubHelloRifky",
			"Hello Rifky",
		},
		{
			"Audry",
			"TestSubHelloAudry",
			"Hello Audry",
		},
	}

	for _, test := range tests {
		t.Run(test.TestName, func(t *testing.T) {
			result := Hello(test.Name)

			require.Equal(t, test.Expected, result, fmt.Sprintf("%s should return %s", test.TestName, test.Expected))
		})
	}
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("World")
	}
}

func BenchmarkSubHello(b *testing.B) {
	b.Run("BenchmarkSubHelloFakhri", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hello("Fakhri")
		}
	})

	b.Run("BenchmarkSubHelloRifky", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hello("Rifky")
		}
	})

	b.Run("BenchmarkSubHelloAudry", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Hello("Audry")
		}
	})
}

func BenchmarkTableHello(b *testing.B) {
	benchmarks := []struct {
		BenchmarkName string
		Name          string
	}{
		{
			BenchmarkName: "BenchmarkSubHelloFakhri",
			Name:          "Fakhri",
		},
		{
			BenchmarkName: "BenchmarkSubHelloRifky",
			Name:          "Rifky",
		},
		{
			BenchmarkName: "BenchmarkSubHelloAudry",
			Name:          "Audry",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.BenchmarkName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Hello(benchmark.Name)
			}
		})
	}
}
