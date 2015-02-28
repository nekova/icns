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
	commandGenerate,
	commandReset,
	commandSet,
}

var commandGenerate = cli.Command{
	Name:  "generate",
	Usage: "Generate your custom icon",
	Description: `
`,
	Action: doGenerate,
}

var commandSet = cli.Command{
	Name:  "set",
	Usage: "Set your custom icon",
	Description: `
`,
	Action: doSet,
}

var commandReset = cli.Command{
	Name:  "reset",
	Usage: "Reset your custom icon",
	Description: `
`,
	Action: doReset,
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

func doGenerate(c *cli.Context) {
	path := c.Args().Get(0)

	if path == "" {
		cli.ShowCommandHelp(c, "generate")
		os.Exit(1)
	}

	generateIcon(path)
}

func doSet(c *cli.Context) {
	filename := c.Args().Get(0)
	dir := absPath(c.Args().Get(1))

	if filename == "" {
		cli.ShowCommandHelp(c, "set")
		os.Exit(1)
	}

	ext := filepath.Ext(filename)
	if ext == ".icns" {
		setIcon(filename, dir)
	} else {
		generateIcon(filename)
		setIcon("icns.icns", dir)
	}
}

func doReset(c *cli.Context) {
	err := exec.Command("rm", "Icon\r").Run()
	if err != nil {
		panic(err)
	}
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
	convert("icon.tif")
	sweep("icon.tif")
}

func convert(filename string) {
	exec.Command("tiff2icns", "-noLarge", filename, "icns.icns").Output()
}

func setIcon(filename, dir string) {
	execPath := getExecPath()
	err := exec.Command("sh", execPath+"/setIcon.sh", dir+"/"+filename, dir).Run()
	if err != nil {
		panic(err)
	}
}

func sweep(filename string) {
	err := exec.Command("rm", filename).Run()
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

func absPath(path string) string {
	abspath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return abspath
}
