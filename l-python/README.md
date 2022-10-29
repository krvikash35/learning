## dict['key'] vs dict.get('key')
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

## json nil
* nil are converted to None while loading `json.loads`
* None are converted to nil while dumping `json.dumps`

## greater than and less than

```
for a in range(11, 20):
    whatever
```

```
if 10 < a < 20:
    whatever
```
## web frameworks
* fastapi (lightweight, simple, express like, )
* flask
* django (heavy, complex, opinionated, annotation/decorator, magic)
* sanic
    * it allow to configure the no of sanic worker which is multi process arch rather multi thread.
    * unless we are doing cpu intensive rathare than io intensive then we should not really increae no of sanic worker.

## http client
* requests (sync call)
* aiohttp (async call)