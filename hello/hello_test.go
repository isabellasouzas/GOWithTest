package main

import "testing"

func TestHello(t *testing.T) {
    
    assertCorrectMessage := func(t *testing.T, got, want string){
        t.Helper()
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    }

    t.Run("toa  person", func(t *testing.T) {
    got := Hello("Chris", "")
    want := "Hello, Chris"
    assertCorrectMessage(t, got, want)

    if got != want {
        t.Errorf("got %q want %q", got, want)
        }
    })

    t.Run("empty string", func(t *testing.T) {
        got := Hello("", " ")
        want := "Hello, world" 
        assertCorrectMessage(t, got, want)
        
        if got != want{
            t.Errorf("got %q want %q", got, want)
        }
    })


    t.Run("in Spanish", func(t *testing.T) {
        got := Hello("Elodie", spanish)
        want := "Hola, Elodie"
        assertCorrectMessage(t, got, want)
    })


    t.Run("in French", func(t *testing.T) {
        got := Hello("Isabella", french)
        want := "Bonjour, Isabella"
        assertCorrectMessage(t, got, want)

    })

}
