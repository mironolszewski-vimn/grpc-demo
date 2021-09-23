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
        response_stream = stub.ListBooks(library_pb2.ListBooksRequest())
        for book in response_stream:
            print(book)

        print("*" * 30, "Listing all 'Lord...' books:")
        request = library_pb2.FilterBooksRequest(
            title=StringValue(value="Lord"),
        )
        response_stream = stub.FilterBooks(request)
        for book in response_stream:
            print(book)

        print("*" * 30, "Listing all books by Williams:")
        request = library_pb2.FilterBooksRequest(
            authorName=StringValue(value="William"),
        )
        response_stream = stub.FilterBooks(request)
        for book in response_stream:
            print(book)

        print("*" * 30, "Subscribing for book updates...")
        import pdb; pdb.set_trace()
        request = library_pb2.SubscribeForBookUpdatesRequest()
        response_stream = stub.SubscribeForBookUpdates(request)
        for book in response_stream:
            print("A new book just appeared!", book)


if __name__ == "__main__":
    main()
