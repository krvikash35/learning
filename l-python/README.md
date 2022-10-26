dict['key'] vs dict.get('key')
* `.get` does not raise `keyError` as it has default value `None`
```
dict.get("key", "default")
```

check if dict has key
```
"key" in dict
```

"ternary" expressions:
```
i = 5 if a > 7 else 0
```

* nil are converted to None while loading `json.loads`
* None are converted to nil while dumping `json.dumps`

## web frameworks
* fastapi
* flask
* django