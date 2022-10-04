```
sdk install gradle 7.5.1
brew install gradle
gradle init
./gradlew run
./gradlew build
```

## Concepts
* tasks: inbuilt(init).  
* plugins

### tasks
* One task can depend on other task hence Gradle create task graph and run in that order.
* Task consist of input, output and actions.
* Some are inbuilt task like `init`

### Plugins
* `base` plugin has tasks
    * clean: Deletes the build directory and everything in it
    * check (lifecycle task): Plugins and build authors should attach their verification tasks, such as ones that run tests, to this lifecycle task using check.dependsOn(task).
    * assemble (lifecycle task): Plugins and build authors should attach tasks that produce distributions and other consumable artifacts to this lifecycle task. For example, jar produces the consumable artifact for Java libraries. Attach using dependsOn(task).
    * build (lifecycle task): Intended to build everything, including running all tests, producing the production artifacts and generating documentation. You will probably rarely attach concrete tasks directly to build as assemble and check are typically more appropriate.
    * buildConfiguration (task rule)
    * cleanTask (task rule)
* `application` plugin tasks. implicitly applies java plugin and distribution plugin.
    * `run` dependsOn `classes`
    * `startScripts` dependsOn `jar`
    * it does not add `Main-Class` to manifest, so jar are not executable. use distribution tar/zip for executable. or add manifest by jar config of java plugin
* `java` plugin
