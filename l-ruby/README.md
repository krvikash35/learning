https://www.theodinproject.com/paths/full-stack-ruby-on-rails/courses/ruby


```
brew install rbenv ruby-build
rbenv install -L
rbenv install -l
rbenv install 3.2.2
rbenv versions
rbenv global 3.1.2
rbenv local 3.1.2
rbenv which ruby
```
RBS(ruby signature) a type anotation was introduced in ruby 3.0.

Editor support
* vscode is not good for ruby development, better use rubymine by jetbrain
* vscode need extra setup, to start with install extension `solargraph` but for ruby/rail, you need more setup
* install vs code extension
    * Solargraph ( need to `gem install solargraph; gem install yard`)
        * `solargraph download-core`, `yard gems`(optional), `yard config –gem-install-yri`
    * rubocop-rails

vscode solargraph vs ruby LSP
* LSP is good for ruby >= 3.0 else use solargraph



gem vs bundler
* with `gem` you have to install all the gem individually i.e
```
gem install rubocop
gem install xyz
```
* with `bundler`, you just have to declare all the gem in the file `Gemfile`, it will install all the gems with command `bundle install`
```
gem 'rubocop', '~> 1.56', require: false
gem 'rails', '~> 5.2.4'
```

require, load, autoload

* unless you require file will not be loaded & you might get uninitialized error. so we can configure autoload to `application.rb` file
```
config.autoload_paths += %W(#{config.root}/lib)
```

run a single test file or single test
```
rake test TEST=tests/functional/accounts_test.rb
rake test TEST=tests/functional/accounts_test.rb TESTOPTS="-n /paid accounts/"
```

Commands
```
gem list
gem install byebug
gem install byebug -v 10.0.2

rake test TEST=tests/functional/accounts_test.rb
rake test TEST=tests/functional/accounts_test.rb TESTOPTS="-n /paid accounts/"

rails server -b 0.0.0.0 -p 3000 -e development

ruby -S gem -v
jruby -S gem -v
jruby --debug
/opt/jruby -S gem -v
```

console debugging
* Use [this]( https://www.rubydoc.info/gems/pry-debugger-jruby/2.1.1) doc for navigation, breakpoint & expression
```
binding.pry -> put anywhere in code where u want 1st breakpoint
```

upcase! vs upcase - 1st will modify original variable i.e rewrite in memory , 2nd will return