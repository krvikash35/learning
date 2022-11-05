# class MyError(Exception):
#     pass

# def f1():
#     print("f1")
#     raise MyError("my error")

# def f2():
#     print("f2")
#     f1()


# def f3():
#     print("f3")
#     f2()


# try:
#     f3()
# except Exception as e:
#     print("caught")
#     raise Exception("chain Exception").with_traceback(e.__traceback__)

# ticket_tags="aaaduplicatefff"

# a =ticket_tags.find("duplicate") != -1
# print(a)
# b=20
# if b is not None:
#     print("not none")

x = lambda a : a + 10
print(x(5))