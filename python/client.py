import click
import grpc
from google.protobuf.wrappers_pb2 import StringValue

import library_pb2
import library_pb2_grpc

from server import SERVER_ADDRESS


@click.command()
@click.option('-s', '--server-address', default="localhost:8091", help='Server address')
def main(server_address):
    with grpc.insecure_channel(server_address) as channel:
        stub = library_pb2_grpc.LibraryStub(channel)

        print("*" * 30, "Listing all books:")
        response = stub.ListBooks(library_pb2.ListBooksRequest())
        for book in response:
            print(book)

        print("*" * 30, "Listing all 'Lord...' books:")
        request = library_pb2.FilterBooksRequest(
            title=StringValue(value="Lord"),
        )
        response = stub.FilterBooks(request)
        for book in response:
            print(book)

        print("*" * 30, "Listing all books by Williams:")
        request = library_pb2.FilterBooksRequest(
            authorName=StringValue(value="William"),
        )
        response = stub.FilterBooks(request)
        for book in response:
            print(book)


if __name__ == "__main__":
    main()
