brew install python@3.8
/usr/local/opt/python@3.8/bin/python3
/usr/local/opt/python@3.9/bin/python3
ln -s /usr/local/Cellar/python@3.8/3.8.13_1/bin/python3 /usr/local/bin/python3.8

install multiple version of python using pyenv, conda, poetry
brew install pyenv
pyenv install|uninstall 3.8.13
pyenv shell|local|global 3.8.3

brew update
brew install python

python3 -m ensurepip --default-pip
python3 --version
pip3 --version