# injectP
Tool for bug hunters The tool is similar to the qsreplace tool, but with some improvements for those who do not know the qsreplace "https://github.com/tomnomnom/qsreplace" tool, it takes a list of links with a parameter and injects it with a payload .
The following example shows the difference between the two tools

```
> echo "http://www.testme.com/search?q=cats&lang=en&userid=12345" | qsreplace <xss>

http://www.testme.com/search?lang=%3Cxss%3E&q=%3Cxss%3E&userid=%3Cxss%3E
```

```
> echo "http://www.testme.com/search?q=cats&lang=en&userid=12345" | injectP '%3Cxss%3E'

http://www.testme.com/search?q=%3Cxss%3E&lang=en&userid=12345
http://www.testme.com/search?q=cats&lang=%3Cxss%3E&userid=12345
http://www.testme.com/search?q=cats&lang=en&userid=%3Cxss%3E

```

# install 

`go get -u github.com/raoufmaklouf/injectP`

# usage
`cat endpointFile.txt | injectP 'xPayload'`

# example

```
> cat EndPointFile.txt

http://www.testme.com/prod/abc?size=S&lang=en
http://www.testme.com/user/bob?q=show&userid=12345
http://www.testme.com/user/jack?q=remove&id=123
```
```
> cat EndPointFile.txt | injectP "%22%3E%3Cscript%3Ealert(1)%3C/script%3E"

http://www.testme.com/prod/abc?size=%22%3E%3Cscript%3Ealert(1)%3C/script%3E&lang=en
http://www.testme.com/prod/abc?size=S&lang=%22%3E%3Cscript%3Ealert(1)%3C/script%3E
http://www.testme.com/user/bob?q=%22%3E%3Cscript%3Ealert(1)%3C/script%3E&userid=12345
http://www.testme.com/user/bob?q=show&userid=%22%3E%3Cscript%3Ealert(1)%3C/script%3E
http://www.testme.com/user/jack?q=%22%3E%3Cscript%3Ealert(1)%3C/script%3E&id=123
http://www.testme.com/user/jack?q=remove&id=%22%3E%3Cscript%3Ealert(1)%3C/script%3E

```
