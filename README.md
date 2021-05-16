# log
Logs for concurrent and modular systems

## How to use this

In the `version/` directory is a template that you want to copy to
the root of your repository.

Change the references in the `update/` directory to point to your new repo URL
version directory:

```go
import (
	"github.com/cybriq/log/version"
)
```

It is of course permissible to use this however you want, but the author's
recommendation is to create a file like this, that you drop in any package you 
want to access this logger from:

```go
package main

import (
	_l "github.com/cybriq/log"

	"github.com/cybriq/log/version"
)

var log = _l.Get(_l.Add(version.PathBase))
```

which will give you back a struct with memorable members pointing to log print
functions that you can use like this:

```go
var e error
if e = SomeFunction(); log.E.Chk(e) {
	return e
}
```

which will print the error at the location and pass it back.

There are other ways to use it, nothing stopping the direct creation of
the logger inside an arbitrary file of a package, just that, like everything
about this 'freedom' in the Go package specification, including especially init()
and var declarations that invoke functions, it can get lost, thus the suggestion
to put it in a file with a meaningful name so it's obvious where to look to 
change it.

And, whoever is using those init functions: Please, stop, and think about 
someone, you know, maintaining it after you. IMO Go should not allow it in any
file not named 'init.go'. Well, I like 'main.go' also. It is not for nothing that
a key goal of the cybriq packages author is to eliminate the manual structuring
of a package folder, and just have it natively structured in the editor.

Regarding the use of the letter `e` instead of `err` in the code above. Except 
for certain transcendental math, `e`, like `i`, logically connects to the concept 
of error, especially in a language where `err != nil` is so frequent it's a meme.
This is also why there is a 'check' function, as this ensures that the code 


**logs at the site of the error!!!!** 

so that you can jump straight to it and fix it,
instead of spending 20 minutes tracing back from the hole it emerges from, 
through the little warren of imports and invocations where it originates.