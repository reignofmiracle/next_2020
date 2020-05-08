package interface_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

func TestWriter(t *testing.T) {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	w.Write([]byte("hellow"))
	//w = time.Second // failed
}

func TestWriter2(t *testing.T) {
	os.Stdout.Write([]byte("hello"))
	//os.Stdout.Close()

	var w io.Writer
	w = os.Stdout
	w.Write([]byte("hello"))
	//w.Close() // failed
}

func TestInterface(t *testing.T) {
	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one": 1}
	any = new(bytes.Buffer)
	fmt.Println(any)
}

type Artifact interface {
	Title() string
	Creators() []string
	Created() time.Time
}

type ArtifactImpl struct {
}

func (r *ArtifactImpl) Title() string {
	return "ArtifactImpl"
}
func (r *ArtifactImpl) Creators() []string {
	return []string{"ArtifactImpl", "a", "b"}
}
func (r *ArtifactImpl) Created() time.Time {
	return time.Now()
}

func TestArtifact(t *testing.T) {
	var a Artifact = new(ArtifactImpl)
	fmt.Println(a.Title())
	fmt.Println(a.Creators())
	fmt.Println(a.Created())
}
