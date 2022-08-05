package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "HelloWorld(Vika)",
			request: "Vika",
		},
		{
			name:    "HelloWorld(Putri)",
			request: "Putri",
		},
		{
			name:    "HelloWorld(Ariyanti)",
			request: "Ariyanti",
		},
		{
			name:    "HelloWorld(Vika Putri)",
			request: "Ariyanti",
		},
		{
			name:    "HelloWorld(Vika Putri Ariyanti)",
			request: "Ariyanti",
		},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}

}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Vika", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Vika")
		}
	})
	b.Run("Putri", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Putri")
		}
	})

}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Vika")
	}

}

func BenchmarkHelloWorldPutri(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Putri")
	}

}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "TestHelloOWorld(Vika)",
			request:  "Vika",
			expected: "Hello Vika",
		},
		{
			name:     "TestHelloOWorld(Putri)",
			request:  "Putri",
			expected: "Hello Putri",
		},
		{
			name:     "TestHelloOWorld(Ariyanti)",
			request:  "Ariyanti",
			expected: "Hello Ariyanti",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}

}

func TestSubTes(t *testing.T) {
	t.Run("Vika", func(t *testing.T) {
		result := HelloWorld("Vika")
		require.Equal(t, "Hello Vika", result)
	})
	t.Run("Putri", func(t *testing.T) {
		result := HelloWorld("Putri")
		require.Equal(t, "Hi Putri", result)
	})
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("sebelum unit test")
	m.Run()
	//after
	fmt.Println("setelah unit test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can not run on Windows")
	}
	result := HelloWorld("Vika")
	require.Equal(t, "Hello Vika", result)
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Vika")
	require.Equal(t, "Hello Vika", result)
	fmt.Println(("TestHelloWorld with Assert Done"))
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Vika")
	assert.Equal(t, "Hello Vika", result)
	fmt.Println(("TestHelloWorld with Require Done"))
}

func TestHelloOWorld(t *testing.T) {
	result := HelloWorld("Vika")
	if result != "Hello Vika" {
		//error
		//panic("Result is not 'Hello Vika")
		//t.Fail()
		t.Error("Result must be 'Hello Vika'")
	}
	fmt.Println("Test Helloworld Done")
}
func TestHelloOWorldPutri(t *testing.T) {
	result := HelloWorld("Putri")
	if result != "Hello Putri" {
		//error
		//panic("Result is not 'Hello Putri")
		//t.FailNow()
		t.Fatal("Result must be 'Hello Putri'")
	}
	fmt.Println("Test HelloworldPutri Done")
}

//go test
//go test -v
//go test -v ./...
