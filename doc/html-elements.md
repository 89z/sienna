# HTML elements

~~~html
<!doctype html>
<html lang="en">
<head>
   <meta charset="utf-8">
   <title>May June</title>
</head>
<body>
   <p>July August</p>
</body>
</html>
~~~

## doctype

Error: Start tag seen without seeing a doctype first.

<https://w3.org/tr/html5/syntax#the-doctype>

## html

Warning: Consider adding a `lang` attribute to the `html` start tag to declare
the language of this document.

## head

The character encoding of the HTML document was not declared. The document
will render with garbled text in some browser configurations if the document
contains characters from outside the US-ASCII range. The character encoding of
the page must be declared in the document or in the transfer protocol.

Error: Element `head` is missing a required instance of child element `title`.

## link

~~~html
<link rel="icon" href="/repo/favicon.png">
~~~

Note that SVG favicons are not supported with mobile browser. Alternative is
PNG.

- <https://flaticon.com>
- <https://gauger.io/fonticon>
- <https://ikonate.com>
- <https://materialdesignicons.com>

~~~html
<link rel="stylesheet" href="style.css">
~~~

## meta

~~~html
<meta name="viewport" content="width=device-width, initial-scale=1">
~~~

`initial-scale` is needed for Chrome Android.

## References

- <https://validator.w3.org>
- <https://w3.org/tr/html5/syntax#optional-tags>
