package commands

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"net/http"

	"github.com/codegangsta/cli"
)

var binPath string

var Setup = cli.Command{
	Name:  "setup",
	Usage: "Install gendoc and its pre-req's into your PATH",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "binpath",
			Usage:       "Destination directory to install docs tools to",
			Value:       "/usr/local/bin/",
			Destination: &binPath,
		},
	},
	Action: func(context *cli.Context) error {
		//install gendoc
		if err := install(os.Args[0], binPath); err != nil {
			return err
		}

		// install hugo
		hugotgz := "hugo_0.16_linux-64bit.tgz"
		if _, err := os.Stat(hugotgz); os.IsNotExist(err) {
			if err := wget("https://github.com/spf13/hugo/releases/download/v0.16/hugo_0.16_linux-64bit.tgz", hugotgz); err != nil {
				return err
			}
		}
		hugo := "./hugo"
		if err := processTGZ(hugotgz, hugo); err != nil {
			return err
		}
		if err := install(hugo, binPath); err != nil {
			return err
		}

		return nil
	},
}

func wget(from, to string) error {
	fmt.Printf("Downloading %s into %s\n", from, to)
	resp, err := http.Get(from)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(to)
	if err != nil {
		return err 
	}
	defer out.Close()
	io.Copy(out, resp.Body)
	return nil
}

func install(from, to string) error {
	fmt.Printf("Installing %s into %s\n", from, to)
	
	//TODO ah, windows.
	// TODO check if its already there - or if that's where we're running from!
	

        cmd := exec.Command("sudo", "cp", from, to)
	//PrintVerboseCommand(cmd)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func processTGZ(srcFile, filename string) error {
	f, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer f.Close()

	gzf, err := gzip.NewReader(f)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(gzf)

	i := 0
	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		name := header.Name

		switch header.Typeflag {
		case tar.TypeDir:
			continue
		case tar.TypeReg:
			fmt.Println("(", i, ")", "Name: ", name)
			if filename == name {
				out, err := os.Create(name)
				if err != nil {
					return err 
				}
				defer out.Close()
				io.Copy(out, tarReader)
				return nil
			}
		default:
			fmt.Printf("%s : %c %s %s\n",
				"Yikes! Unable to figure out type",
				header.Typeflag,
				"in file",
				name,
			)
		}

		i++
	}
	return nil
}
