package commands

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

var binPath string
var updateFlag bool

var Install = cli.Command{
	Name:  "install",
	Usage: "Install gendoc and its pre-req's into your PATH",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "binpath",
			Usage:       "Destination directory to install docs tools to",
			Value:       "/usr/local/bin/",
			Destination: &binPath,
		},
		cli.BoolFlag{
			Name:        "update",
			Usage:       "Check for updated releases",
			Destination: &updateFlag,
		},
	},
	Action: func(context *cli.Context) error {
		gendocFileToInstall := os.Args[0]
		gendocTo := binPath + "gendoc"
		if runtime.GOOS == "windows" {
			gendocTo = gendocTo + ".exe"
		}

		logrus.Debugf("os.Arg[0]: %s ~~ gendocTo %s\n", gendocFileToInstall, gendocTo)
		if updateFlag || os.Args[0] == gendocTo {
			// If the user is running setup from an already installed gendoc, assume update
			// TODO: if main.Version == today, maybe don't bother?
			fmt.Printf("Checking for newer version of gendoc.\n")
			resp, err := http.Get("https://github.com/SvenDowideit/gendoc/releases/latest")
			if err != nil {
				fmt.Printf("Error checking for latest version \n%s\n", err)
			} else {
				releaseUrl := resp.Request.URL.String()
				latestVersion := releaseUrl[strings.LastIndex(releaseUrl, "/")+1:]
				fmt.Printf("this version == %s, latest version == %s\n", context.App.Version, latestVersion)
				thisDate, _ := time.Parse("2006-01-02", context.App.Version)
				latestDate, _ := time.Parse("2006-01-02", latestVersion)

				if !latestDate.After(thisDate) {
					gendocFileToInstall = gendocTo
				} else {
					fmt.Printf("Downloading new version of gendoc.")
					gendocFile := "gendoc"
					if runtime.GOOS == "darwin" {
						gendocFile += "-osx"
					}
					if runtime.GOOS == "windows" {
						gendocFile += ".exec"
					}
					gendocFileToInstall := "gendoc-latest"
					if err := wget("https://github.com/SvenDowideit/gendoc/releases/download/"+latestVersion+"/"+gendocFile, gendocFileToInstall); err != nil {
						return err
					}
				}
			}
		}

		if gendocFileToInstall == gendocTo {
			fmt.Printf("Latest gendoc already installed at %s\n", gendocTo)
		} else {
			if err := install(gendocFileToInstall, gendocTo); err != nil {
				return err
			}
		}

		hugoFileToInstall := binPath + "hugo"
		if runtime.GOOS == "windows" {
			hugoFileToInstall = hugoFileToInstall + ".exe"
		}
		if _, err := os.Stat(hugoFileToInstall); !os.IsNotExist(err) {
			fmt.Println("Hugo is already installed")
			// TODO: add update code.
			return nil
		}

		// install hugo
		arch := runtime.GOARCH
		if arch == "amd64" {
			arch = "64bit"
		}
		if arch == "386" {
			arch = "32bit"
		}
		ext := "tgz"
		if runtime.GOOS == "windows" {
			ext = "zip"
		}
		goos := runtime.GOOS
		if goos == "darwin" {
			goos = "osx"
		}
		hugoarchive := "hugo_0.16_" + goos + "-" + arch + "." + ext
		if _, err := os.Stat(hugoarchive); os.IsNotExist(err) {
			if err := wget("https://github.com/spf13/hugo/releases/download/v0.16/"+hugoarchive, hugoarchive); err != nil {
				return err
			}
		}
		hugo := "./hugo"
		if ext == "tgz" {
			if err := processTGZ(hugoarchive, hugo); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("BOOM, zip files not coded yet")
		}
		if err := install(hugo, hugoFileToInstall); err != nil {
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

	// on OSX, the file got a quarantine xattr, (-c) clearing all
	// sorry - need to fix up the last 2 week's versions so we can upgrade :)
	if runtime.GOOS == "darwin" {
		if _, err := os.Stat(to); !os.IsNotExist(err) {
			cmd := exec.Command("sudo", "xattr", "-c", to)
			//PrintVerboseCommand(cmd)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			if err := cmd.Run(); err != nil {
				return err
			}
		}
	}
	//TODO ah, windows.
	// TODO check if its already there - or if that's where we're running from!

	// on OSX, the file gets a quarantine xattr, (-c) clearing all
	if runtime.GOOS == "darwin" {
		cmd := exec.Command("sudo", "xattr", "-c", from)
		//PrintVerboseCommand(cmd)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	cmd := exec.Command("sudo", "cp", from, to)
	//PrintVerboseCommand(cmd)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
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
			fmt.Printf("Found %s file\n", name)
			if filename == name {
				out, err := os.Create(name)
				if err != nil {
					return err
				}
				defer out.Close()
				io.Copy(out, tarReader)
				out.Chmod(0755)
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
