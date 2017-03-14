# Define

Very simple library to translate word from german to english. Will print general information - speech part and meanings. Nothing fancy for now. 

PRs are welcome.

## Usage

If you have `go` installed: 

`go get -u github.com/ideahitme/define`

**Use:**
```
define <german-word>
```

**For example:**

```
define sprechen

//Output

Definition: 1
Sprechen: noun neutral
speech,speaking
===================
Definition: 2
sprechen: verb
to talk,to speak,to converse,to put up
===================
```

Planning to add binaries prebuilt for Mac and Linux in the release page

## Contribute

Contribution is especially needed for:
1.  adding more elaborated and smarter outputs
2.  adding more sources (for now only `bab.la`)

## License

The MIT License (MIT)

Copyright (c) 2016 Yerken Tussupbekov

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.