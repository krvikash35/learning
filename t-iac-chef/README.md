chef
* [how to write simple chef cookbook](https://www.digitalocean.com/community/tutorials/how-to-create-simple-chef-cookbooks-to-manage-infrastructure-on-ubuntu)
* knife, receipe, cookbook, runlist
* comp
    * Chef Server: central location that stores configuration recipes, cookbooks, and node and workstation definitions.
    * Chef Nodes: Each node represents a separate, contained machine environment that can be on physical hardware or virtualized. each contain chef client to communicate with server.
    * Chef Workstations: where Chef configuration details are created or edited. The configuration files are then pushed to the Chef server, where they will be available to deploy to any nodes


```
cd ~/chef-repo
knife cookbook create nginx
knife cookbook create apt
knife cookbook upload apt
knife cookbook upload nginx
knife cookbook upload -a

knife node list
knife node edit client1

sudo chef-client
sudo chef-client -o 'recipe[recipeName]' # to avoid restart of other process, run specific recipe

cd cookbooks/nginx
ls
attributes  CHANGELOG.md  definitions  files  libraries  metadata.rb  providers  README.md  recipes  resources	templates
```

ndoe edit
```
#node edit
{
  "name": "client1",
  "chef_environment": "_default",
  "normal": {
    "tags": [

    ]
  },
  "run_list": [
    "recipe[nginx]"
  ]
}
```

ngnix cookbook
```
#nano ~/chef-repo/cookbooks/ngnix/recipes/default.rb
include_recipe “apt”
package 'nginx' do
  action :install
end

service 'nginx' do
  action [ :enable, :start ]
end

cookbook_file "/usr/share/nginx/www/index.html" do
  source "index.html"
  mode "0644"
endo
```

apt cookbook
```
#nano ~/chef-repo/cookbooks/apt/recipes/default.rb
execute "apt-get update" do
  command "apt-get update"
end
```