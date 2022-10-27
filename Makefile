define runTests
	$(shell "for d in ./*/*; do (cd "$d" && go test); done")
endef

test:
	for d in ./*/ ; do (cd $d && echo $d && go test); done