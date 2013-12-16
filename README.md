fontxtract
==========

fontxtract is a simple commandline utility which parses a directory of zip archives and extracts TTF & OTF font files to a specified directory.

USAGE

Linux:

	fontxtract path-to-zips/ path-to-fonts/

Windows:

	fontxtract.exe path-to-zips\ path-to-fonts\

BUILDING

Building & running fontxtract from source is easy. It requires Go (http://golang.org) & Git (http://git-scm.com). Simply clone the source repository, run ‘go build’ in the source directory to build the executable, then extract the fonts!

Linux bash shell example:

	$ git clone https://github.com/lee8oi/fontxtract.git
	$ cd fontxtract
	$ go build
	$ ./fontxtract zippath/ fontpath/

Windows command prompt example:

	> git clone https://github.com/lee8oi/fontxtract.git
	> cd fontxtract
	> go build
	> fontxtract zippath\ fontpath\