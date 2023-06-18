#!/bin/bash

FYNE="$(which fyne)";
ROOT="$(pwd)";


package_app() {

	local os="$1";

	cd "${ROOT}/source";

	${FYNE} package -os "${os}" -icon ../asset/icon.png;

}

package_app "linux";
# package_app "darwin";
# package_app "windows";

