package util


import "testing"

func Test1(t *testing.T) {
	data := " POST /hello/robo HTTP/1.1\n" +
		"Host: localhost:8080\n" +
		"api_key: 132138120980\n" +
		"Cache-Control: no-cache\n" +
		"\n" +
		"hello world\n"
	s := Server{}
	re := s.Parse([]byte(data))
	println("methods =>")
	println(re.method+ "\n")
	println("bodys =>")
	println(re.body)

	println("headers =>")
	for k, v := range re.header {
		print(k + " : " + v)
	}
	//t.Log(re)
}