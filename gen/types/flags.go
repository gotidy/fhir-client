// Copyright 2021 Evgeny Safonov
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.package main

package main

import "flag"

func ParseFlags() Config {
	config := Config{}
	flag.StringVar(&(config.Output), "o", "", "Output file, else output to STDOUT")
	flag.StringVar(&(config.Input), "i", "", "Input dir, else get from STDIN")
	flag.StringVar(&(config.Definitions), "d", "", "Definitions destination file")
	flag.Parse()
	return config
}
