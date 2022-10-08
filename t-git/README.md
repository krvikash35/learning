```
git fetch
git merge
git rebase
git pull
git pull --rebase
```

* pull=fetch+merge or fetch+rebase. by default pull uses merge strategy
* rebase create neat clean flat history compare to merge


remote
 ```
git remote -v
git remote add <remote_name> <remote_url>
git remote remove <remote_name>
git remote set-url <remote_url>
 ```

remote urls
 ```
https://github.com/krvikash35/learning.git
https://username:password@github.com/krvikash35/learning.git
https://a%40b.com:password@github.com/krvikash35/learning.git //url encode special char
git@github.com:krvikash35/learning.git
 ```