import person_pb2
import sys
import traceback

file_name = "address_book_pb"

address_book = person_pb2.AddressBook()
address_book_read = person_pb2.AddressBook()

def add_dummy_address_book():
    print("adding dummy data to address book")
    person1 = address_book.people.add()
    person2 = address_book.people.add()

    person1.id = 1
    person1.name = "John"
    person1.email = "jdoe@example.com"
    phone1 = person1.phones.add()
    phone1.number = "555-4321"
    phone1.type = person_pb2.Person.HOME

    person2.id = 2
    person2.name = "Doe"
    person2.email = "jdoe@example.com"
    phone1 = person2.phones.add()
    phone1.number = "555-4322"
    phone1.type = person_pb2.Person.WORK
    print("adding dummy data to address book DONE", address_book)


def write_address_book_to_file():
    print("writting to file")
    try:
        with open(file_name, "wb") as f:
            f.write(address_book.SerializeToString())
            print("writting to file DONE")
    except Exception as e:
        print(f"error: {e} occoured while opening file: {file_name}")

def read_address_book_from_file():
    print("reading from file")
    try:
        with open(file_name, "rb") as f:
           address_book_read.ParseFromString(f.read())
           print("reading from file DONE", address_book_read)
    except Exception as e:
        print(f"error: {e} occured while reading file: {file_name}")

add_dummy_address_book()
write_address_book_to_file()
read_address_book_from_file()
