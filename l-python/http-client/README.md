```python
r = requests.get('https://api.github.com/user', auth=('user', 'pass'))
r.status_code: 200
r.headers['content-type']: 'application/json; charset=utf8'
r.encoding: 'utf-8'
r.text: '{"type":"User"...'
r.json(): {'private_gists': 419, 'total_private_repos': 77, ...}


import requests
r = requests.get('https://swapi.dev/api/starships/9/')
print(r.json())

data = {"name": "Obi-Wan Kenobi", ...}
r = requests.post('https://httpbin.org/post', json=data)
print(r.json())

r = requests.post('https://httpbin.org/post', data=data, cookies={'foo': 'bar', 'hello': 'world'}))

s = requests.Session()
s.get('https://httpbin.org/cookies/set/sessioncookie/123456789')
r = s.get('https://httpbin.org/cookies')
print(r.text)
```