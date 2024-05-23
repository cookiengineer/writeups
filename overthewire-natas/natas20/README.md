
# Natas 20 (Over The Wire CTF)

Credentials:

- Username: `natas20`
- Password: `guVaZ3ET35LbgbFMoaN5tFcYT1jEP7UH`


# Problem

The `print_credentials()` function can be used to get admin privileges by exploiting the `myread($sid)` function. This works because we can inject malicious data into the `mywrite($sid,$data)` function, as the data is never checked.


# Solution

We can set the input of the `name`-field to something like `stuff\nadmin 1` This will trigger the bug in 

``` php

    foreach(explode("\n", $data) as $line) {
        debug("Read [$line]");
    $parts = explode(" ", $line, 2);
    if($parts[0] != "") $_SESSION[$parts[0]] = $parts[1];
    }
```
and we will get the admin rights as it goes through the file line by line and sets the SESSION such that the `print_credentials()`-function is satisfied
 


```bash
cd ./solution;

go run main.go;
```

