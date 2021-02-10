package observer

func ExampleObserver() {

	reader1 := NewReader("reader1")
	reader2 := NewReader("reader2")
	reader3 := NewReader("reader3")

	subject := NewSubject()
	subject.Add(reader1)
	subject.Add(reader2)
	subject.Add(reader3)

	subject.UpdateContext("observer mode")
	// Output:
	// receive observer mode
	// receive observer mode
	// receive observer mode
}
