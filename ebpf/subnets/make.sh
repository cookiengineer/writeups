#!/bin/bash

TARGET="";
BPFTOOL="$(which bpftool)";
CLANG="$(which clang)";
GO="$(which go)";
GO_LDFLAGS="";
LLC="$(which llc)";
ROOT="$(pwd)";
SUDO="$(which sudo)";

build_ebpf() {

	cd "${ROOT}";

	if [[ -f "${ROOT}/build/module.ll" ]]; then
		rm "${ROOT}/build/module.ll" 2> /dev/null;
	fi;

	if [[ -f "${ROOT}/build/module.bpfeb" ]]; then
		rm "${ROOT}/build/module.bpfeb" 2> /dev/null;
	fi;

	if [[ -f "${ROOT}/build/module.bpfel" ]]; then
		rm "${ROOT}/build/module.bpfel" 2> /dev/null;
	fi;

	if [[ -f "/tmp/bpftool.err" ]]; then
		rm "/tmp/bpftool.err" 2> /dev/null;
	fi;

	if [[ -f "/tmp/bpftool.out" ]]; then
		rm "/tmp/bpftool.out" 2> /dev/null;
	fi;

	# TODO: Add -DENABLE_DNSFILTER once it's ready
	${CLANG} -DENABLE_DEBUG -S -I"${ROOT}/headers" -target bpf -O2 -Wall -Wno-unused-function -Wno-unused-value -Wno-pointer-sign -Wno-compare-distinct-pointer-types -Werror -emit-llvm -g -c -o "${ROOT}/build/module.ll" "${ROOT}/source/module.c"

	if [[ "$?" == "0" ]]; then
		echo -e "- Generate eBPF LLVM code: ${os} [\e[32mok\e[0m]";
	else
		echo -e "- Generate eBPF LLVM code: ${os} [\e[31mfail\e[0m]";
		return 1;
	fi;

	if [[ -f "${ROOT}/build/module.ll" ]]; then

		${LLC} -march=bpfeb -mcpu=v1 -filetype=obj -o "${ROOT}/build/module.bpfeb" "${ROOT}/build/module.ll";

		if [[ "$?" == "0" ]]; then
			echo -e "- Generate eBPF module: big-endian [\e[32mok\e[0m]";
		else
			echo -e "- Generate eBPF module: big-endian [\e[31mfail\e[0m]";
			return 1;
		fi;

		${LLC} -march=bpfel -mcpu=v1 -filetype=obj -o "${ROOT}/build/module.bpfel" "${ROOT}/build/module.ll";

		if [[ "$?" == "0" ]]; then
			echo -e "- Generate eBPF module: little-endian [\e[32mok\e[0m]";
		else
			echo -e "- Generate eBPF module: little-endian [\e[31mfail\e[0m]";
			return 1;
		fi;

		if [[ -f "${ROOT}/build/module.bpfeb" ]] && [[ -f "${ROOT}/build/module.bpfel" ]]; then

			if [[ "${BPFTOOL}" != "" ]] && [[ "${SUDO}" != "" ]] && [[ -f "${ROOT}/build/module.bpfel" ]]; then

				arch="$(uname -m)";
				bpf_endian="";

				if [[ "${arch}" == "i386" ]] || [[ "${arch}" == "i686" ]]; then
					bpf_endian="little-endian";
				elif [[ "${arch}" == "amd64" ]] || [[ "${arch}" == "x86_64" ]]; then
					bpf_endian="little-endian";
				elif [[ "${arch}" == "armv6l" ]] || [[ "${arch}" == "armv7l" ]]; then
					bpf_endian="big-endian";
				elif [[ "${arch}" == "aarch64" ]] || [[ "${arch}" == "aarch64_be" ]] || [[ "${arch}" == "armv8b" ]] || [[ "${arch}" == "armv8l" ]]; then
					bpf_endian="big-endian";
				fi;

				if [[ "${bpf_endian}" == "little-endian" ]]; then

					${SUDO} rm -f "/sys/fs/bpf/whatever";
					${SUDO} ${BPFTOOL} --debug prog load "${ROOT}/build/module.bpfel" "/sys/fs/bpf/whatever" type xdp 2> /tmp/bpftool.err 1> /tmp/bpftool.out;

					if [[ "$?" == "0" ]]; then
						echo -e "- Verify eBPF module: ${bpf_endian} [\e[32mok\e[0m]";
						return 0;
					else
						echo -e "- Verify eBPF module: ${bpf_endian} [\e[31mfail\e[0m]";
						echo -e "! bpftool errors at \"/tmp/bpftool.err\".";
						echo -e "! bpftool output at \"/tmp/bpftool.out\".";
						return 1;
					fi;

				elif [[ "${bpf_endian}" == "big-endian" ]]; then

					${SUDO} rm -f "/sys/fs/bpf/whatever";
					${SUDO} ${BPFTOOL} --debug prog load "${ROOT}/build/module.bpfeb" "/sys/fs/bpf/whatever" type xdp 2> /tmp/bpftool.err 1> /tmp/bpftool.out;

					if [[ "$?" == "0" ]]; then
						echo -e "- Verify eBPF module: ${bpf_endian} [\e[32mok\e[0m]";
						return 0;
					else
						echo -e "- Verify eBPF module: ${bpf_endian} [\e[31mfail\e[0m]";
						echo -e "! bpftool errors at \"/tmp/bpftool.err\".";
						echo -e "! bpftool output at \"/tmp/bpftool.out\".";
						return 1;
					fi;

				else
					echo -e "- Verify eBPF module: big-endian [\e[31mfail\e[0m]";
					echo -e "- Verify eBPF module: little-endian [\e[31mfail\e[0m]";
					echo -e "Unsupported host CPU architecture \"${arch}\". [\e[31mfail\e[0m]";
					return 1;
				fi;

			else
				echo -e "- Verify eBPF module: big-endian [\e[31mfail\e[0m]";
				echo -e "- Verify eBPF module: little-endian [\e[31mfail\e[0m]";
				echo -e "Please install bpftools. [\e[31mfail\e[0m]";
				return 1;
			fi;

		else
			echo -e "- Verify eBPF module: big-endian [\e[31mfail\e[0m]";
			echo -e "- Verify eBPF module: little-endian [\e[31mfail\e[0m]";
			return 1;
		fi;

	else

		echo -e "- Generate eBPF module: big-endian [\e[31mfail\e[0m]";
		echo -e "- Generate eBPF module: little-endian [\e[31mfail\e[0m]";
		echo -e "- Verify eBPF module: big-endian [\e[31mfail\e[0m]";
		echo -e "- Verify eBPF module: little-endian [\e[31mfail\e[0m]";
		return 1;

	fi;

}

build_ebpf;

