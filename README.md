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
