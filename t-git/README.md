```
git fetch
git merge
git rebase
git pull
git pull --rebase

git clone --depth 1 --branch <tag> <repo>
git rm --cached <file>
git rm -r --cached <folder>
git revert -m 1 <commit>
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

 config
 ```
git config --local -l                           // get all local config keys
git config --global -l                          // get all global config keys
git config --global user.name "vikash kumar"    // set config key globally
git config --local user.name "vikash kumar"     // set config key locally
git config --unset                              // remove key from config
git config --global --edit                      // edit global config in editor
git config --local --edit                       // edit local config in editor

git config --global url.ssh://git@source.golabs.io/.insteadOf https://source.golabs.io/
 ```

```
git revert sha1
git revert sh1 sha2
git revert --no-commit sh1 sha2
git revert --no-commit HEAD~2..
```

```
git reset --soft HEAD~1
git reset --hard HEAD~1

```

```
git branch -D branchname
```

```
git stash
git stash list
git stash pop
git stash pop stash@{n}

git stash push -m "stashname"
git stash save "stashname" //deprecated
git stash apply stash@{n}
git stash apply stash^{/my_stash_name}

```
 ## gitflow vs github flow

 ### github flow
 ![](./github-flow.png)
 * simple workflow suited for small agile **team** and  **Project/software** that has no concept of release and/or release happens through continious delivery. i.e microservice
 * Mainline branches: mainly two branch are used in this approach
    * master: 
        * parmanent branch, which is always considered stable, ready to deploy to prod anytime.
        * continious deployment happens through master branch.
    * feature: 
        * temporary branch, dev create for each feature, test(unit, QA), merge to master then delete/destroy this branch.
        * hotfix also go throguh feature branch.



 ### git flow
 ![](./git-flow.png)
 * Complex workflow suited for software that has concept of release i.e mobile app.
 * Branches
    * master
        * parmanent branch. Always contain baseline of latest released. release and hotfix are merged to master once done.
    * develop
        * parmanent branch. contain all the developed features. release and feature are created from this branch.
        * It work as integration test, all feature are combined and tested together.
    * release
        * once release is done then it is merged to master and develop branch.
    * hotfix
        * temporary branch. hotfix are bug fix to software that are already released. dev create seperate hotfix branch from release tag and merge to develop and master once done.
    * feature
        * temporary branch. dev create seperate feature branch for each fature. continiously commit, merge it to develop once done.
* since there are multiple branch, it become complex and require more work though it is suitable for particular usecase.

### Problem
* There are multiple feature `f1`, `f2`, `f3` etc being worked parallely and individually by devs. `qa1` want to test `f1` and `qa2` want to test `f2`. `f1` and `f2` both belong to same service with additional constraint that service has only one environment(`staging`) to test at the moment. So here, we have concurrency issue, both feature can't be deployed to single resource(staging env) at same time. only one feature can be deployed at a time and tested. So `qa1` is blocked on `qa2` to finish his testing. vice versa.
* Problem is not about enabling parallel feature testing but enabling QA productivity such that they are not blocked on each other because parallel feature testing can be problem for even one single QA if he demands to test multiple feature at same time.


#### approaches
* **planning: assignment**: what if f1 and f2 was assigned to qa1, it would not have been problem as he would be testing sequenctially.
    * Assign/distribute the feature testing task to QA in such a way that they don't block each other.
    * Since each service has one testing env, assign each QA a different service.
* **planning: cordination**: even if situation arise where they are blocked on each other for testing f1 and f2. Then they can cordinate  to each other on when they can use testing env, meanwhile they can test some other stuff. staging environment is common resource that we all need on shared basis, so better cordination can solve this problem.
* **use gitflow**: we can use gitflow to solve this requirement but it does not suite our team and software development dynamics. It will bring unnecessary complexity and overhead.
* **merge f1 and f2**: this is what have started doing at adhoc basis which is unstandard,  has lots of flaw, many unanswered question.
    * merge f1 to f2 and deploy f2? merge f1 and f2 to fk and deploy fk? create new branch from master and cherry pick both f1 and f2 then deploy common branch? so many question
    * since now we have more that two branch for same feature, we will have to make fix in multiple branch which again has multiple approach.
    * Sometime this might not be feasible where git(api/lib) are leverage to provide the higher user layer functionality. Might require a bit of work to provide this feature.
    * This will increase the effort and complexity, hence need to check the cost and benefit.
    * Sometime this might not solve all the usecase. i.e for cross services feature which might lead to merging other  features indirectly.