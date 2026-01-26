
# Natas 20 (Over The Wire CTF)

Credentials:

- Username: `natas20`
- Password: `guVaZ3ET35LbgbFMoaN5tFcYT1jEP7UH`


# Problem

The `mywrite(sid, data)` method serializes data in an unsafe way, using
a `\n` as a delimiter for each `key value\n` pair. This leaves the `myread(sid)`
method vulnerable, because it is using a loop that relies on `foreach(explode("\n", data))`.

This allows an attacker to override a `key` using an injected `\n` character
in the posted name value.


# Solution

The `name` variable can be changed to `whatever\nadmin 1` in order to override
the `admin` key, which will lead to the `myread()` method building a wrong session
object.

```bash
cd ./solution;

go run main.go;
```

