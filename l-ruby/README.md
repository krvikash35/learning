
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

Editor support
* vscode is not good for ruby development, better use rubymine by jetbrain
* vscode need extra setup, to start with install extension `solargraph` but for ruby/rail, you need more setup
* install vs code extension
    * Solargraph ( need to `gem install solargraph;`)
    * rubocop-rails
* 

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