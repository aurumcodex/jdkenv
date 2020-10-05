package util

import (
	"testing"

	toml "github.com/pelletier/go-toml"
)

func TestBuildCorrettoURL(t *testing.T) {
	const (
		url8  = "https://corretto.aws/downloads/latest/amazon-corretto-8-x64-linux-jdk.tar.gz"
		url11 = "https://corretto.aws/downloads/latest/amazon-corretto-11-x64-linux-jdk.tar.gz"
	)

	c, e := toml.LoadFile("../jdk_list.toml")
	if e != nil {
		t.Errorf("Unable to locate file jdk_list.toml; err = %v", e)
	}

	baseURLCorretto := c.Get("corretto.base_url").(string)
	fileURL8 := c.Get("corretto.8.file_url").(string)
	fileURL11 := c.Get("corretto.11.file_url").(string)

	result := BuildString(baseURLCorretto, fileURL8)
	if result != url8 {
		t.Errorf("BuildString(baseURLCorretto, fileURL8) = %v; wanted : %v", result, url8)
	}

	result = BuildString(baseURLCorretto, fileURL11)
	if result != url11 {
		t.Errorf("BuildString(baseURLCorretto, fileURL11) = %v; wanted : %v", result, url11)
	}
}

func TestBuildLibericaURL(t *testing.T) {
	const (
		url8  = "https://download.bell-sw.com/java/8u265+1/bellsoft-jdk8u265+1-linux-amd64-full.tar.gz"
		url11 = "https://download.bell-sw.com/java/11.0.8+10/bellsoft-jdk11.0.8+10-linux-amd64-full.tar.gz"
		url15 = "https://download.bell-sw.com/java/15+36/bellsoft-jdk15+36-linux-amd64-full.tar.gz"
	)

	c, e := toml.LoadFile("../jdk_list.toml")
	if e != nil {
		t.Errorf("Unable to locate file jdk_list.toml; err = %v", e)
	}

	baseURLLiberica := c.Get("liberica.base_url").(string)
	extURL8 := c.Get("liberica.8.ext_url").(string)
	extURL11 := c.Get("liberica.11.ext_url").(string)
	extURL15 := c.Get("liberica.15.ext_url").(string)
	fileURL8 := c.Get("liberica.8.filename").(string)
	fileURL11 := c.Get("liberica.11.filename").(string)
	fileURL15 := c.Get("liberica.15.filename").(string)

	result := BuildString(baseURLLiberica, extURL8, fileURL8)
	if result != url8 {
		t.Errorf("BuildString(baseURLLiberica, extURL8, fileURL8) = %v; wanted : %v", result, url8)
	}

	result = BuildString(baseURLLiberica, extURL11, fileURL11)
	if result != url11 {
		t.Errorf("BuildString(baseURLLiberica, extURL11, fileURL11) = %v; wanted : %v", result, url11)
	}

	result = BuildString(baseURLLiberica, extURL15, fileURL15)
	if result != url15 {
		t.Errorf("BuildString(baseURLLiberica, extURL15, fileURL15) = %v; wanted : %v", result, url15)
	}
}

func TestBuildOracleURL(t *testing.T) {
	const (
		url8  = "https://download.java.net/openjdk/jdk8u41/ri/openjdk-8u41-b04-linux-x64-14_jan_2020.tar.gz"
		url11 = "https://download.java.net/openjdk/jdk11/ri/openjdk-11+28_linux-x64_bin.tar.gz"
		url15 = "https://download.java.net/openjdk/jdk15/ri/openjdk-15+36_linux-x64_bin.tar.gz"
	)

	c, e := toml.LoadFile("../jdk_list.toml")
	if e != nil {
		t.Errorf("Unable to locate file jdk_list.toml; err = %v", e)
	}

	baseURLOracle := c.Get("oracle.base_url").(string)
	extURL8 := c.Get("oracle.8.ext_url").(string)
	extURL11 := c.Get("oracle.11.ext_url").(string)
	extURL15 := c.Get("oracle.15.ext_url").(string)
	fileURL8 := c.Get("oracle.8.filename").(string)
	fileURL11 := c.Get("oracle.11.filename").(string)
	fileURL15 := c.Get("oracle.15.filename").(string)

	result := BuildString(baseURLOracle, extURL8, fileURL8)
	if result != url8 {
		t.Errorf("BuildString(baseURLOracle, extURL8, fileURL8) = %v; wanted : %v", result, url8)
	}

	result = BuildString(baseURLOracle, extURL11, fileURL11)
	if result != url11 {
		t.Errorf("BuildString(baseURLOracle, extURL11, fileURL11) = %v; wanted : %v", result, url11)
	}

	result = BuildString(baseURLOracle, extURL15, fileURL15)
	if result != url15 {
		t.Errorf("BuildString(baseURLOracle, extURL15, fileURL15) = %v; wanted : %v", result, url15)
	}
}

func TestBuildOpenJDKURL(t *testing.T) {
	const (
		url8  = "https://github.com/AdoptOpenJDK/openjdk8-binaries/releases/download/jdk8u265-b01/OpenJDK8U-jdk_x64_linux_hotspot_8u265b01.tar.gz"
		url11 = "https://github.com/AdoptOpenJDK/openjdk11-binaries/releases/download/jdk-11.0.8%2B10/OpenJDK11U-jdk_x64_linux_hotspot_11.0.8_10.tar.gz"
		url14 = "https://github.com/AdoptOpenJDK/openjdk14-binaries/releases/download/jdk-14.0.2%2B12/OpenJDK14U-jdk_x64_linux_hotspot_14.0.2_12.tar.gz"
		url15 = "https://github.com/AdoptOpenJDK/openjdk15-binaries/releases/download/jdk-15%2B36/OpenJDK15U-jdk_x64_linux_hotspot_15_36.tar.gz"
	)

	c, e := toml.LoadFile("../jdk_list.toml")
	if e != nil {
		t.Errorf("Unable to locate file jdk_list.toml; err = %v", e)
	}

	baseURLOpenJDK := c.Get("openjdk.base_url").(string)
	releaseURL := c.Get("openjdk.release_path").(string)
	extURL8 := c.Get("openjdk.8.ext_url").(string)
	extURL11 := c.Get("openjdk.11.ext_url").(string)
	extURL14 := c.Get("openjdk.14.ext_url").(string)
	extURL15 := c.Get("openjdk.15.ext_url").(string)
	tag8 := c.Get("openjdk.8.tag").(string)
	tag11 := c.Get("openjdk.11.tag").(string)
	tag14 := c.Get("openjdk.14.tag").(string)
	tag15 := c.Get("openjdk.15.tag").(string)
	fileURL8 := c.Get("openjdk.8.filename").(string)
	fileURL11 := c.Get("openjdk.11.filename").(string)
	fileURL14 := c.Get("openjdk.14.filename").(string)
	fileURL15 := c.Get("openjdk.15.filename").(string)

	result := BuildString(baseURLOpenJDK, extURL8, releaseURL, tag8, fileURL8)
	if result != url8 {
		t.Errorf("BuildString(baseURLOpenJDK, extURL8, releaseURL, tag8, fileURL8) = %v; wanted : %v", result, url8)
	}

	result = BuildString(baseURLOpenJDK, extURL11, releaseURL, tag11, fileURL11)
	if result != url11 {
		t.Errorf("BuildString(baseURLOpenJDK, extURL11, releaseURL, tag11, fileURL11) = %v; wanted : %v", result, url11)
	}

	result = BuildString(baseURLOpenJDK, extURL14, releaseURL, tag14, fileURL14)
	if result != url14 {
		t.Errorf("BuildString(baseURLOpenJDK, extURL14, releaseURL, tag14, fileURL14) = %v; wanted : %v", result, url14)
	}

	result = BuildString(baseURLOpenJDK, extURL15, releaseURL, tag15, fileURL15)
	if result != url15 {
		t.Errorf("BuildString(baseURLOpenJDK, extURL15, releaseURL, tag15, fileURL15) = %v; wanted : %v", result, url15)
	}
}

func TestBuildOpenJDKOpenJ9URL(t *testing.T) {
	const (
		url8  = "https://github.com/AdoptOpenJDK/openjdk8-binaries/releases/download/jdk8u265-b01_openj9-0.21.0/OpenJDK8U-jre_x64_linux_openj9_8u265b01_openj9-0.21.0.tar.gz"
		url11 = "https://github.com/AdoptOpenJDK/openjdk11-binaries/releases/download/jdk-11.0.8%2B10_openj9-0.21.0/OpenJDK11U-jdk_x64_linux_openj9_11.0.8_10_openj9-0.21.0.tar.gz"
		url14 = "https://github.com/AdoptOpenJDK/openjdk14-binaries/releases/download/jdk-14.0.2%2B12_openj9-0.21.0/OpenJDK14U-jdk_x64_linux_openj9_14.0.2_12_openj9-0.21.0.tar.gz"
	)

	c, e := toml.LoadFile("../jdk_list.toml")
	if e != nil {
		t.Errorf("Unable to locate file jdk_list.toml; err = %v", e)
	}

	baseURLOpenJDK := c.Get("openjdk.base_url").(string)
	releaseURL := c.Get("openjdk.release_path").(string)
	extURL8 := c.Get("openjdk.8.ext_url").(string)
	extURL11 := c.Get("openjdk.11.ext_url").(string)
	extURL14 := c.Get("openjdk.14.ext_url").(string)
	tag8 := c.Get("openjdk.8.openj9.tag").(string)
	tag11 := c.Get("openjdk.11.openj9.tag").(string)
	tag14 := c.Get("openjdk.14.openj9.tag").(string)
	fileURL8 := c.Get("openjdk.8.openj9.filename").(string)
	fileURL11 := c.Get("openjdk.11.openj9.filename").(string)
	fileURL14 := c.Get("openjdk.14.openj9.filename").(string)

	result := BuildString(baseURLOpenJDK, extURL8, releaseURL, tag8, fileURL8)
	if result != url8 {
		t.Errorf("BuildString(baseURLOpenJDK, extURL8, releaseURL, tag8, fileURL8) = %v; wanted : %v", result, url8)
	}

	result = BuildString(baseURLOpenJDK, extURL11, releaseURL, tag11, fileURL11)
	if result != url11 {
		t.Errorf("BuildString(baseURLOpenJDK, extURL11, releaseURL, tag11, fileURL11) = %v; wanted : %v", result, url11)
	}

	result = BuildString(baseURLOpenJDK, extURL14, releaseURL, tag14, fileURL14)
	if result != url14 {
		t.Errorf("BuildString(baseURLOpenJDK, extURL14, releaseURL, tag14, fileURL14) = %v; wanted : %v", result, url14)
	}
}
