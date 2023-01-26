Feature: List books present in database
    Scenario: user sends get request to gw to list books started from offset and with limit
        Given book database contains books and user provides how much he wants to get
        When user requests to list books starting from offset and number of books not greater than limit
        Then user gets response with books

Feature: Create book and store in in database
    Scenario: user sends post request to gw to create a book
        Given user provides name and author of a book
        When user requests to create a book with given name and author
        Then user gets response with newly created book