# lss-img
### Install
    go get github.com/GreenRaccoon23/lss-img
### Clone
    git clone https://github.com/GreenRaccoon23/lss-img.git
### Description
Command-line tool to find image files on the filesystem.  

    [chuck@norris ~]$ lss-img --help
    Usage: lss-img [options] <patterns-to-match>
    Options:
        -d "$PWD":
            Start search under a specific directory
        -x:
            Patterns to exclude from matches, separated by commas
        -f:
            Find matches based on the full path of files
              (by default, only the basenames of files are checked for matches)
        -v:
            Display slightly more output
        -r:
            Display the relative path to <path>
              (the full path is displayed by default)
        -b
            Display the basename only
              (the full path is displayed by default)
        -c
            Colorize output
        -w:
            Write output to file
        -h:
            Print this help 
 
Note: This program mostly checks file extensions, not file signatures.
