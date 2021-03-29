# GitHub

~~~
curl.exe --netrc-file C:\msys64\home\Steven\.netrc `
https://api.github.com/rate_limit
authorization: Basic BASE64_PASS
{
  "rate": {
    "limit": 5000,
    "used": 0,
    "remaining": 5000,
    "reset": 1617052862
  }
}
~~~
