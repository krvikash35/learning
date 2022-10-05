```
pytest --cov=lib --cov=actions --cov=worker ./tests/
pytest -q test_sysexit.py //quiet mode
```

```
def func(x):
    return x + 1  
def test_answer():
    assert func(3) == 5
```

```
import pytest
def f():
    raise SystemExit(1)
def test_mytest():
    with pytest.raises(SystemExit):
        f()
```

```
@pytest.fixture
def my_fruit():
    return Fruit("apple")
    
@pytest.fixture
def fruit_basket(my_fruit):
    return [Fruit("banana"), my_fruit]
    
def test_my_fruit_in_basket(my_fruit, fruit_basket):
    assert my_fruit in fruit_basket
```