## Commands to run

```bash
poetry install

poetry run python -m grpc_tools.protoc -I../protos --python_out=. --grpc_python_out=. ../protos/library.proto

poetry run python server.py

poetry run python client.py
````