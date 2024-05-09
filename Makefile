
MovieNight.AppDir:
	mkdir -p MovieNight.AppDir/subprogram
	go build -o MovieNight.AppDir/AppRun ./AppRun
	go build -o MovieNight.AppDir/subprogram/subprogram ./subprogram