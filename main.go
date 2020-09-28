package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	//"strings"

	spin "github.com/briandowns/spinner"
	aurora "github.com/logrusorgru/aurora/v3"
	arc "github.com/mholt/archiver/v3"
	ini "gopkg.in/ini.v1"
)

const (
	version string = "0.1.0"

	javaHome string = "JAVA_HOME"
	path     string = "PATH"
)

var (
	au     aurora.Aurora
	colors = flag.Bool("no-color", false, "false to enable color or true to disable color")
	v      = flag.Bool("version", false, "prints version")
	j      = flag.String("jdk-list", "", "path to an .ini file containing a list of known JDK distributions")
)

func init() {
	flag.Parse()
	au = aurora.NewAurora(!*colors)

	if *v {
		fmt.Printf("jdkenv :: version %v\n", version)
		os.Exit(0)
	}
}

func main() {
	// TODO: add in cli argument parsing for easier/faster setting of jdks
	// todo: add a second config file with system installed jdks and the current setting
	s := spin.New(charSet, 100*time.Millisecond)
	// s.Prefix = "  "
	// s.Suffix = " Extracting..."
	// s.Color("bold", "yellow")
	// s.Stop()

	// load up the jdk sources file
	jdkCfg, jdkErr := ini.LoadSources(ini.LoadOptions{
		SkipUnrecognizableLines: true,
	}, "jdks.ini")

	if jdkErr != nil {
		fmt.Fprintf(os.Stderr, "Failed to load JDK source list; Error: %v", jdkErr)
	}

	// url := jdkCfg.Section("openjdk").Key("base_url")

	item := BuildString(
		jdkCfg.Section("corretto.8").Key("base_url").String(),
		jdkCfg.Section("corretto.8").Key("filename").String(),
	)

	var urlList []string

	urlList = append(urlList, item)
	urlList = append(urlList, BuildString(
		jdkCfg.Section("corretto.11").Key("base_url").String(),
		jdkCfg.Section("corretto.11").Key("filename").String(),
	))
	urlList = append(urlList, jdkCfg.Section("liberica.11").Key("base_url").String())
	urlList[2] += jdkCfg.Section("liberica.11").Key("ext_url").String()
	urlList[2] += jdkCfg.Section("liberica.11").Key("filename").String()

	Menu() // this *needs* to be a void function that does nothing but render some options via switches

	// sections := jdkCfg.SectionStrings()

	for _, s := range urlList {
		fmt.Println(s)
		Download(s)
	}
	// str := BuildString(jdkCfg.Section())

	// Download(urlList[0])

	// fmt.Println("  Extracting", jdkCfg.Section("corretto.11").Key("filename").String())
	s.Prefix = "  "
	s.Suffix = fmt.Sprintf("%v %v", au.Bold(au.Cyan(" Extracting:")), jdkCfg.Section("corretto.11").Key("filename").String())
	s.FinalMSG = fmt.Sprintf("   %v %v\n", au.Bold(au.Green(" Extracted")), jdkCfg.Section("corretto.11").Key("filename").String())

	s.Color("bold", "yellow")
	// s.Stop()

	s.Start()
	e := arc.Unarchive(jdkCfg.Section("liberica.11").Key("filename").String(), "./test/")
	s.Stop()
	if e != nil {
		fmt.Fprintf(os.Stderr, "couldn't extract file: %v\n", e)
		os.Exit(1)
	}
}
