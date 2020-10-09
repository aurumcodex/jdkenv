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
	// Corretto flag for setting/using Corretto as JDK of choice
	Corretto int = iota

	// Liberica flag for setting/using Liberica as JDK of choice
	Liberica int = iota

	// OpenJDK flag for setting/using OpenJDK as JDK of choice
	OpenJDK int = iota

	// OpenJ9 flag for setting/using OpenJdk with OpenJ9 JVM as JDK of choice
	OpenJ9 int = iota

	// Oracle flag for setting/using Oracle as JDK of choice
	Oracle int = iota
)

// Constants for determining what to put into the `.env` file in `~/.jdkenv`
// and also the filenames to download to.
const (
	Corretto8Home string = `export JAVA_HOME="${HOME}/.jdkenv/amazon-corretto-8"`
	Corretto8Path string = `export PATH=""`

	Corretto11Home string = `export JAVA_HOME="${HOME}/.jdkenv/amazon-corretto-11"`
	Corretto11Path string = `export PATH=""`

	Liberica8Home string = `export JAVA_HOME="${HOME}/.jdkenv/bellsoft-liberica-8"`
	Liberica8Path string = `export PATH=""`

	Liberica11Home string = `export JAVA_HOME="${HOME}/.jdkenv/bellsoft-liberica-11"`
	Liberica11Path string = `export PATH=""`

	Liberica15Home string = `export JAVA_HOME="${HOME}/.jdkenv/bellsoft-liberica-15"`
	Liberica15Path string = `export PATH=""`

	Oracle8Home string = `export JAVA_HOME="${HOME}/.jdkenv/oracle-openjdk-8"`
	Oracle8Path string = `export PATH=""`

	Oracle11Home string = `export JAVA_HOME="${HOME}/.jdkenv/oracle-openjdk-11"`
	Oracle11Path string = `export PATH=""`

	Oracle15Home string = `export JAVA_HOME="${HOME}/.jdkenv/oracle-openjdk-15"`
	Oracle15Path string = `export PATH=""`

	OpenJDK8Home string = `export JAVA_HOME="${HOME}/.jdkenv/adoptopenjdk-8"`
	OpenJDK8Path string = `export PATH=""`

	OpenJDK11Home string = `export JAVA_HOME="${HOME}/.jdkenv/adoptopenjdk-11"`
	OpenJDK11Path string = `export PATH=""`

	OpenJDK14Home string = `export JAVA_HOME="${HOME}/.jdkenv/adoptopenjdk-14"`
	OpenJDK14Path string = `export PATH=""`

	OpenJDK15Home string = `export JAVA_HOME="${HOME}/.jdkenv/adoptopenjdk-15"`
	OpenJDK15Path string = `export PATH=""`

	OpenJ98Home string = `export JAVA_HOME="${HOME}/.jdkenv/adoptopenjdk-openj9-8"`
	OpenJ98Path string = `export PATH=""`

	OpenJ911Home string = `export JAVA_HOME="${HOME}/.jdkenv/adoptopenjdk-openj9-11"`
	OpenJ911Path string = `export PATH=""`

	OpenJ914Home string = `export JAVA_HOME="${HOME}/.jdkenv/adoptopenjdk-openj9-14"`
	OpenJ914Path string = `export PATH=""`
)
