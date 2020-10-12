/*
Package util contians various helping functions for jdkenv
(i.e. downloading and extracting compressed tarballs).

Copyright © 2020 Nathan Adams <aurumcodex@protonmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package util

const (
	// Corretto flag for setting/using Corretto as JDK of choice.
	Corretto int = iota

	// Liberica flag for setting/using Liberica as JDK of choice.
	Liberica int = iota

	// OpenJDK flag for setting/using OpenJDK as JDK of choice.
	OpenJDK int = iota

	// OpenJ9 flag for setting/using OpenJdk with OpenJ9 JVM as JDK of choice.
	OpenJ9 int = iota

	// Oracle flag for setting/using Oracle as JDK of choice.
	Oracle int = iota

	// Azul flag for setting/using Azul Zulu as JDK of choice.
	Azul int = iota

	// Sapmachine flag for setting/using Sapmachine as JDK of choice.
	Sapmachine int = iota
)

// Constants for determining what to put into the `.env` file in `~/.jdkenv`
// and also the filenames to download to.
const (
	JavaHomePath string = `export PATH="${JAVA_HOME}/bin:${PATH}`

	Corretto8Home  string = `export JAVA_HOME="${HOME}/.jdkenv/amazon-corretto-8"`
	Corretto11Home string = `export JAVA_HOME="${HOME}/.jdkenv/amazon-corretto-11"`

	Liberica8Home  string = `export JAVA_HOME="${HOME}/.jdkenv/bellsoft-liberica-8"`
	Liberica11Home string = `export JAVA_HOME="${HOME}/.jdkenv/bellsoft-liberica-11"`
	Liberica15Home string = `export JAVA_HOME="${HOME}/.jdkenv/bellsoft-liberica-15"`

	Oracle8Home  string = `export JAVA_HOME="${HOME}/.jdkenv/oracle-openjdk-8"`
	Oracle11Home string = `export JAVA_HOME="${HOME}/.jdkenv/oracle-openjdk-11"`
	Oracle15Home string = `export JAVA_HOME="${HOME}/.jdkenv/oracle-openjdk-15"`

	OpenJDK8Home  string = `export JAVA_HOME="${HOME}/.jdkenv/adoptopenjdk-8"`
	OpenJDK11Home string = `export JAVA_HOME="${HOME}/.jdkenv/adoptopenjdk-11"`
	OpenJDK14Home string = `export JAVA_HOME="${HOME}/.jdkenv/adoptopenjdk-14"`
	OpenJDK15Home string = `export JAVA_HOME="${HOME}/.jdkenv/adoptopenjdk-15"`

	OpenJ98Home  string = `export JAVA_HOME="${HOME}/.jdkenv/adoptopenjdk-openj9-8"`
	OpenJ911Home string = `export JAVA_HOME="${HOME}/.jdkenv/adoptopenjdk-openj9-11"`
	OpenJ914Home string = `export JAVA_HOME="${HOME}/.jdkenv/adoptopenjdk-openj9-14"`
)
