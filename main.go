package main

import (
	"bufio"
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var FolderPadre = "path/to/directory"
var ListaCanciones []string
type PackageDetails struct {
	Nombre string
	Artista string
}
var PackageList []PackageDetails

// Abre una carpeta de canciones
// Abre una de las subcarpetas
// Abre el archivo .sm o .scc disponible
// Consigue el atributo #TITLE:
// Consigue el atributo #ARTIST:
// Cierra ese archivo
// Continua con la siguiente carpeta
// Termina al recorrer todas las carpetas

func main() {
    files, err := ioutil.ReadDir(FolderPadre)
    if err != nil {
        log.Fatal(err)
    }
    for _, file := range files {
        //fmt.Println(file.Name(), file.IsDir())
		if file.IsDir() {
			FindStepFile(file)
		}
    }
	for _, item := range ListaCanciones {
		//fmt.Println(item)
		ReadStepFile(item)
	}
	jItem, _ := json.Marshal(PackageList)
	//fmt.Println(string(jItem))
	_ = ioutil.WriteFile("test2.json", jItem, 0644)
}

func FindStepFile(folder fs.FileInfo){
	files, err := ioutil.ReadDir(FolderPadre + "/" + folder.Name())
    if err != nil {
        log.Fatal(err)
    }
	for _, file := range files {
		if strings.Contains(file.Name(), ".sm") {
			ListaCanciones = append(ListaCanciones, FolderPadre + "/" + folder.Name() + "/" + file.Name())
			break
		} else if strings.Contains(file.Name(), ".ssc") {
			ListaCanciones = append(ListaCanciones, FolderPadre + "/" + folder.Name() + "/" + file.Name())
			break
		}
    }
}

func ReadStepFile(filepath string)  {
	dat, err := os.Open(filepath)
	check(err)
	//fmt.Print(string(dat))
	defer dat.Close()
	scanner := bufio.NewScanner(dat)
	item := PackageDetails{"", ""}
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "#TITLE:") && item.Nombre == "" {
			ss := strings.Split(scanner.Text(), "#TITLE:")
			sss := strings.Split(ss[1], ";")
			//fmt.Println(sss[0])
			item.Nombre = sss[0]
		}
		if strings.Contains(scanner.Text(), "#ARTIST:") && item.Artista == "" {
			dd := strings.Split(scanner.Text(), "#ARTIST:")
			ddd := strings.Split(dd[1], ";")
			//fmt.Println(ddd[0])
			item.Artista = ddd[0]
		}
		if item.Nombre != "" && item.Artista != "" {
			PackageList = append(PackageList, item)
			//jItem, _ := json.Marshal(item)
			//fmt.Println(string(jItem))
			break			
		}
    }
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}