
# symbol key
person = { name: "Vikash" }
puts person[:name]  # => "Vikash"

# string key
person = { "name" => "Vikash" }
puts person["name"]  # => "Vikash" 

# symbol key
person = { name: => "vikash"}
puts person["name"]  # => "Vikash" 