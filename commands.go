package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/gographics/imagick/imagick"
)

var Commands = []cli.Command{
	commandCreate,
}

var commandCreate = cli.Command{
	Name:  "create",
	Usage: "",
	Description: `
`,
	Action: doCreate,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doCreate(c *cli.Context) {
	path := c.Args().Get(0)

	if path == "" {
		cli.ShowCommandHelp(c, "create")
		os.Exit(1)
	}

	generateIcon(path)
	convert("icon.tif")
	setIcon("icns.icns")
	sweep()
}

func generateIcon(filename string) {
	imagick.Initialize()
	defer imagick.Terminate()
	im := imagick.NewMagickWand()

	err := im.ReadImage(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "No such image: %s", filename)
		os.Exit(1)
	}
	mw := imagick.NewMagickWand()
	execPath := getExecPath()

	mw.ReadImage(execPath + "/folder.png")
	fm := mw.Clone()
	defer mw.Destroy()
	defer im.Destroy()
	defer fm.Destroy()
	mw.CompositeImage(im, imagick.COMPOSITE_OP_LINEAR_DODGE, 0, 0)
	mw.CompositeImage(fm, imagick.COMPOSITE_OP_COPY_OPACITY, 0, 0)
	mw.WriteImage("./icon.tif")
}

func convert(filename string) {
	exec.Command("tiff2icns", "-noLarge", filename, "icns.icns").Output()
}

func setIcon(filename string) {
	execPath := getExecPath()
	currentPath, err := filepath.Abs("./")
	err = exec.Command("sh", execPath+"/setIcon.sh", currentPath+"/"+filename, currentPath).Run()
	if err != nil {
		panic(err)
	}
}

func getExecPath() string {
	execPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return execPath
}

func sweep() {
	err := exec.Command("rm", "icns.icns", "icon.tif").Run()
	if err != nil {
		panic(err)
	}
}
