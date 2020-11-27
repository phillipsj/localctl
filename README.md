# localctl

localctl, which I call *local cuddle*, is a tool for installing applications in your .local directory on Linux. 

Installing tools in your .local directory keeps from cluttering your system and doesn't require administrative privileges which reduces risk to your system. 

## Usage


You can test it out with the following script that just echos **"It worked!"**.

First install the script using the *get* command.

```Bash
$ localctl get https://gist.githubusercontent.com/phillipsj/2880b7cf10ec1bb3afe0758a4e9e0236/raw/b20f8094a2322dd5fba0fcccb0a46495adc90364/myscript.sh
Installing to /home/phillipsj/.local/bin ...
Download finished!
```

Now let's execute the script we downloaded.

```Bash
$ myscript
It worked!
```
