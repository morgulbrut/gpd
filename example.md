%title: gdp - Sample Presentation
%author: morgulbrut
%date: 2016-02-07

---------------

# mdp 


A command-line based markdown presentation tool. 

_Basic controls:_

next slide      *Enter*, *Space*, *Page Down*, *j*, *l*,
                *Down Arrow*, *Right Arrow*

previous slide  *Backspace*, *Page Up*, *h*, *k*,
                *Up Arrow*, *Left Arrow*

quit            *q*
reload          *r*
slide N         *1..9*
first slide     *Home*, *g*
last slide      *End*, *G*

---------------

# Supported markdown formatting 

The input file is split into multiple slides by
horizontal rules (hr). A hr consisting of at
least 3 *-*. It can also contain spaces but
no other characters.

Each of these represents the start of a new slide.

\---

---------------

# Supported markdown formatting 

First-level headers can be prefixed by single *#*.

\# first-level

becomes

# first-level

---------------

# Supported markdown formatting 

Second-level headers can be prefixed by *##*.

\# second-level

becomes

# second-level


---------------

# Supported markdown formatting's 

Inline codes are surrounded with backticks.

C program starts with \`main()\`.

becomes

C program starts with `main()`.

---------------

# Supported markdown formatting 

You can also use [github](https://guides.github.com/features/mastering-markdown/#GitHub-flavored-markdown) flavored markdown's
code block. Use at least three backticks to open
and at least as many or more backticks for closing.

\```
\int main(int argc, char \*argv[]) {
\    printf("%s\\n", "Hello world!");
\}
\```

becomes

```
int main(int argc, char *argv[]) {
    printf("%s\n", "Hello world!");
}
```

Language hint will be ignored

---------------

# Supported markdown formatting 

Quotes are auto-detected by preceding *>*.

Multiple *>* are interpreted as nested quotes.

\> quote
\>> nested quote 1
\> > nested quote 2

becomes

> quote
>> nested quote 1
> > nested quote 2

-------------------

# Supported markdown formatting 

Inline highlighting is supported as followed:

\- *\** colors text as red
\- *\_* underlines text

\_some\_ \*highlighted\* \_\*text\*\_

becomes

_some_ *highlighted* _*text*_

---------------

# Supported markdown formatting 

Backslashes force special markdown characters
like *\**, *\_*, *#* and *>* to be printed as
normal characters.

\\\*special\\\*

becomes

\*special\*

---------------

# Supported markdown formatting 

Leading *\** or *-* indicate lists.

list
\* major
\    - minor
\        - \*important\*
\          detail
\    - minor

becomes

list
* major
    - minor
        - *important*
          detail
    - minor

---------------

# Supported markdown formatting 

A single *\<br\>* or *^* in a line indicates mdp
to stop the output on that position.

This can be used to show bullet points
line by line.

*\<br\>* is also not displayed in HTML converted
output.

Agenda
<br>
* major
<br>
    * minor
<br>
* major
  ^
    * minor
      ^
        * detail

---------------

# Supported markdown formatting 

Leading ** indicates centering.

\ # test 
\ ## test 
\ test
\ \_\*test\*\_ 

becomes

# test 
## test 
test
_*test*_ 

---------------

 # Supported markdown formatting 

URL in pandoc style are supported:

\[Google](http://www.google.com/)

becomes

[Google](http://www.google.com/)

---------------

 ## More information about markdown 

can be found in the [markdown documentation](http://daringfireball.net/projects/markdown/).

---------------

 # Support for UTF-8 special characters 

Here are some examples.

ae = ä, oe = ö, ue = ü, ss = ß
upsilon = Ʊ, phi = ɸ

▛▀▀▀▀▀▀▀▀▀▜
▌rectangle▐
▙▄▄▄▄▄▄▄▄▄▟


---------------

# Suspend your presentation for hands-on examples 

Use *Ctrl + z* to suspend the presentation.

Use *fg* to resume it.