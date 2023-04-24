package views

import (
	"log"

	rice "github.com/GeertJohan/go.rice"
	"gopkg.in/yaml.v2"
)

type MenuItem struct {
	Name    string     `yaml:"name"`
	URL     string     `yaml:"url"`
	SubMenu []MenuItem `yaml:"sub_menu"`
}

func GetMenuItems() []MenuItem {
	box := rice.MustFindBox("sideMenu").HTTPBox()
	data, err := box.Bytes("menu_super_admin.yaml")
	if err != nil {
		log.Fatalf("error reading menu data: %v", err)
	}

	var items []MenuItem
	err = yaml.Unmarshal(data, &items)
	if err != nil {
		log.Fatalf("error unmarshaling menu data: %v", err)
	}

	return items
}
