# Docker Basics

## Set Up
### MySQL
1. Install the MySQL docker image
```bash
    docker pull mysql
```
2. Create a MySQL container
```bash
    docker run --name mysql-tasks -e MYSQL_ROOT_PASSWORD=12345 -d mysql:latest
```
3. Get the IP Address of this new container
```bash
    docker inspect mysql-tasks
```
You will get a huge JSON, at the end you will find the IP ADDRESS:
```json
    [
    {
        "Id": "3d715d5a776a544da17a119585545882d859dc562c7d592acf147dbd822c823c",
        "Created": "2024-01-13T16:38:03.867737457Z",
        ...
            "Networks": {
                "bridge": {
                    "IPAMConfig": null,
                    "Links": null,
                    "Aliases": null,
                    "MacAddress": "02:42:ac:11:00:02",
                    "NetworkID": "f1e11e4150566901a7c5a6658f17cbd7a63980ec309f124a7c18d83df6e46a9b",
                    "EndpointID": "0d6b23641f001d72da2c53fa938b804b053fde71edf2a0084ea330f6f6cbd742",
                    "Gateway": "172.17.0.1",
                    "IPAddress": "172.17.0.2", // HERE!!!
                    "IPPrefixLen": 16,
                    "IPv6Gateway": "",
                    "GlobalIPv6Address": "",
                    "GlobalIPv6PrefixLen": 0,
                    "DriverOpts": null
                }
            }
        }
    }
]
```
4. 