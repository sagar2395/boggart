/*
=======================
boggart
=======================

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
You should have received a copy of the GNU General Public License
along with this program.  If not, see http://www.gnu.org/licenses/.

	@Repository:  https://github.com/edoardottt/boggart
	@Author:      edoardottt, https://www.edoardoottavianelli.it
	@License: https://github.com/edoardottt/boggart/blob/main/LICENSE
*/

package main

import (
	"log"

	"github.com/edoardottt/boggart/pkg/template"
	"github.com/edoardottt/boggart/server/honeypot"
)

func main() {
	tmpl, err := ReadTemplate("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	//Check Template format
	if _, err := template.CheckTemplate(tmpl); err != nil {
		log.Fatal(err)
	}
	honeypot.Raw(tmpl)
}
