# goduper a rather fast duplicate finder

## what?
find duplicates in a given directory as fast as possible

it has 2 options:

⏩**1-quick**: partial md5 checksum , first few kb(s) , 4096(4kb) to be exact

💻**2-full**: as name implies, full md5 checksum, better accuracy compared to quick



## how?

usage: 

`gpduper /path/to/your/directory`

build:

just run 

`go build goduper.go`

and it's compiled 😃

add to path if you need to

for example,do a cp thing-y on a gnu/linux system:

`sudo cp goduper /usr/local/bin`

or any other path for that matter

## future

there are a buch of stuff missing or to be improved , I would like to correct these stuff overtime:

**1-bad cross platform support**

**2 -no option to remove found duplicates properly**

**3 - verbose output but no error or exception handling**`

**4 - improve performance even more, bacuase why not?**


