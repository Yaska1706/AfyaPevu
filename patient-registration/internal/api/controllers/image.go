package controllers

import (
	"fmt"
	"time"
)

type Image struct {
	Id       int       `json:"id"`
	Location string    `json:"location"`
	Path     string    `json:"path"`
	Date     time.Time `json:"date"`
}

type Images []Image

var currentId int
var images Images

func init() {
	RepoCreateImage(Image{})
	RepoCreateImage(Image{})
}

func RepoFindImage(id int) Image {
	for _, t := range images {
		if t.Id == id {
			return t
		}
	}
	// yoksa geriye bos dondur
	return Image{}
}

func RepoCreateImage(t Image) Image {
	currentId += 1
	t.Id = currentId
	images = append(images, t)
	return t
}

func RepoDestroyImage(id int) error {
	for i, t := range images {
		if t.Id == id {
			images = append(images[:i], images[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("could not find Image with id of %d to delete", id)
}
