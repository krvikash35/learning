
* pip is the preferred installer program. starting with python 3.4, it is included by default.
* venv is the standard tool for creating virtual environments, and has been part of Python since Python 3.3. Starting with Python 3.4, it defaults to installing pip into all created virtual environments

```
python -m pip install SomePackage
python -m pip install SomePackage==1.0.4    # specific version
python -m pip install "SomePackage>=1.0.4"  # minimum version
python -m pip install --upgrade SomePackage

pip install pip-tools==6.6.0 pip==22.0.4
pip-compile requirements/dev.pip
pip-sync requirements/dev.txt
```

virtual env
```
python3 -m venv tutorial_env
python3.8 -m venv tutorial_env
/usr/local/Cellar/python@3.8/3.8.13_1/bin/python3 -m venv tutorial_env

. tutorial_env/bin/activate
source tutorial_env/bin/activate
```