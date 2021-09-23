from concurrent.futures import  ThreadPoolExecutor

import grpc

import library_pb2
import library_pb2_grpc
import library


SERVER_ADDRESS = "127.0.0.1:8091"


class LibraryServicer(library_pb2_grpc.LibraryServicer):

    def ListBooks(self, request, context):
        request: library_pb2.ListBooksRequest

        for book in library.list_books():
            yield book

    def FilterBooks(self, request, context):
        request: library_pb2.FilterBooksRequest

        for book in library.filter_books(title=request.title.value, author=request.authorName.value):
            yield book


def main():
    server: grpc.Server = grpc.server(ThreadPoolExecutor())
    library_pb2_grpc.add_LibraryServicer_to_server(LibraryServicer(), server)
    server.add_insecure_port(SERVER_ADDRESS)
    print(f"Starting server listening at {SERVER_ADDRESS}...")
    server.start()
    server.wait_for_termination()


if __name__  == "__main__":
    main()
