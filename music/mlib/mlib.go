package mlib

import "errors"

//MusicEntry music item
type MusicEntry struct {
	ID     string
	Name   string
	Artist string
	Source string
	Type   string
}

//MusicManager c
type MusicManager struct {
	musics []MusicEntry
}

//给MusicManager类定义方法

//Len get the number of musics
func (m *MusicManager) Len() int {
	return len(m.musics)
}

//Get get music by index,return a pointer of the music and a err
func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range.")
	}
	return &m.musics[index], nil
}

//Find get music by name,return a pointer of the music
func (m *MusicManager) Find(name string) *MusicEntry {
	if len(m.musics) == 0 {
		return nil
	}

	for _, music := range m.musics {
		if music.Name == name {
			return &music
		}
	}
	return nil
}

//Add add a music
func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

//Remove remove a music by index, return a pointer of the removed music
func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index > len(m.musics) {
		return nil
	}

	removedMusic := &m.musics[index]
	m.musics = append(m.musics[:index], m.musics[index+1:]...)
	return removedMusic
}

//RemoveByName remove a music by name, return a pointer of the removed music
func (m *MusicManager) RemoveByName(name string) *MusicEntry {
	var index int
	for i, v := range m.musics {
		if name == v.Name {
			index = i
		}
	}
	removedMusic := &m.musics[index]
	m.musics = append(m.musics[:index], m.musics[index+1:]...)
	return removedMusic
}

//NewMusicManager create a music manager,Return a *MusicManager
func NewMusicManager() *MusicManager { //全局的New函数，新建一个对象，返回该对象指针
	return &MusicManager{
		make([]MusicEntry, 0),
	}
}
