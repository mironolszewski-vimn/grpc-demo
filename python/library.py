from typing import List, Optional

import library_pb2


AUTHORS = [
    library_pb2.Author(id=1, name="John Ronald Reuel Tolkien"),
    library_pb2.Author(id=2, name="William Golding"),
    library_pb2.Author(id=3, name="William Shakespeare"),
]

BOOKS = [
    library_pb2.Book(id=1, title="The Hobbit", authors=[AUTHORS[0]]),
    library_pb2.Book(id=2, title="The Lord Of The Rings", authors=[AUTHORS[0]]),
    library_pb2.Book(id=3, title="Lord Of The Flies", authors=[AUTHORS[1]]),
    library_pb2.Book(id=4, title="Hamlet", authors=[AUTHORS[2]]),
]


def list_books() -> List[library_pb2.Book]:
    return [b for b in BOOKS]


def filter_books(
    title: Optional[str] = None,
    author: Optional[str] = None,
) -> List[library_pb2.Book]:
    books = [b for b in BOOKS]

    if title:
        books = [b for b in books if title.lower() in b.title.lower()]

    if author:
        books = [b for b in books if any(author.lower() in a.name.lower() for a in b.authors)]

    return books
