package mlib

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager failed")
	}
	m0 := &MusicEntry{
		"1",
		"My heart will go on",
		"Celion Dion",
		"http://127.0.0.1",
		"MP3",
	}

	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("MusicManager.Add() failed")
	}

	m := mm.Find(m0.Name)
	fmt.Println(m0, m0.Name, &m0.Name)

	if m == nil {
		t.Error("MusicManager.Find() failed")
	}
	if m.ID != m0.ID || m.Artist != m0.Artist || m.Name != m0.Name || m.Source != m0.Source || m.Type != m0.Type {
		t.Error("MusicManager.Find() failed.")
	}

	m, err := mm.Get(0)
	if err == nil {
		fmt.Println("get Success")
	}

	if m == nil {
		t.Error("MusicManager.Get() failed")
	}

	m = mm.Remove(0)
	if m == nil || mm.Len() != 0 {
		t.Error("MusicManager.Remove() failed")
	}

}
