# user api - example

### setup

1. copy env file
    ```bash
        cp .env.example .env
    ```
2. generate rand key and set `JWT_SECRET`
    ```bash
        openssl rand -base64 20
        # iQtx3LJnzpOkNkYEOZhH60j/FIs=
    ```