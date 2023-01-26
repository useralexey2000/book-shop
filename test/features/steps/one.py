from behave import *
import requests
import json

@given('book database contains books')
def impl_bk(context):
    # TODO populate databese with data
    context.headers = {'content-type': 'application/json'}
    context.payload = {
        "limit": 3,
        "offset": 0,
    }

@When('user requests to list books starting from offset and number of books not greater than limit')
def impl_bk(context):
    context.res = requests.get('http://bookserv-svc:8080/books/list', params=context.payload, headers=context.headers)

@Then('user gets response with books')
def impl_bk(context):
    assert context.res.status == 200


@given('user provides name and author of a book')
def impl_bk(context):
    context.headers = {'content-type': 'application/json'}
    context.payload = {
        "name": "book1",
        "author_name": "author1",
    }

@When('user requests to create a book with given name and author')
def impl_bk(context):
    context.res = requests.post('http://bookserv-svc:8080/books/create', data=context.payload, headers=context.headers)


@Then('user gets response with newly created book')
def impl_bk(context):
    assert context.res.status == 200