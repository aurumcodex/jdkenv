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
