import json


data = {
    "cities1": [{
        "name": "n1",
        "state": "s1"
    }, {"name": "n1",
        "state": "s1"
        }]
}



# for f in data["cities"]:
#     print(f)
# b ="cities" in data and ['[f for f in data.get("cities")]
b = [f for f in data["cities"]] if "cities" in data else []
print(b)

a = {

}
print(a.get("name"))
# print(a["name"])

o1 = {
    "name": "vikash",
    "age": None
}

print(o1)
print(json.dumps(o1))

o2 = json.loads(json.dumps(o1))
print(o2)
print(o2["name"])
print(o2["age"])
