Use this script (Go binary) to generate a service authentication token from a service account private key.

Can output to stdout or to a file.

Usage:
```
$ ./dcos-auth login -k priv.key -u sa -m 192.168.10.31
eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ1aWQiOiJzYSIsImV4cCI6MTUyNzA5MDA2OH0.YUG2zp9pnJE1TJHfdYqv0ExqxFKYcurL8B1VHDvNCFRQDlBH9N9v-0ohNB1PoVWZynWZt9QXAA2E_12OWUCpW6ghp4jAyVmABhcyxqpOf8jwP0CHYxyxcM2QZiKYlO7eIru8dsNVvXH9unsKv0iOTFTzJbUTqFAoSH82PMix30ORPRawwzOrkP0TJ3F0CTQFf_wWX5R1naZ2PJCaxpRmk11DnU5oTTJQSteCopBbiryN3lZJaAS9nmGoA_mULU5ysTcL95C4DwjQIpmqwM17xkvg0Iht-52SJz6Uq9rLS2JoA29whmRQ7YnGqZPnvQu4TjK0YH4CmtuZ8UDocoAWNw

$ ./dcos-auth login -k priv.key -u sa -m 192.168.10.31 -o token
$ cat token
eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJ1aWQiOiJzYSIsImV4cCI6MTUyNzA5MDA3MH0.ZP1SAhV75DLYi-CcVfco-uv55HbfPSAis1SjYh-I1uN7lqwfN4w6gEbUaqc2452-DrucKkxv7W8ScsHip1UKwe2HlxYrQSAigT-ztSbGw-toxhQeuDOa086pifraCdNXeevjHdtno6kRfrXHewPDeukVqbmr2_uizwwy3Hq1-0diCq3RojZ-q5ljdTONCBpXPl9ElEdQipxrlrPdAljGf8e-COAmob0hGw4pCWCOWYLjoWq86jD0nKNkFtr80O47RIqYFbM6_mdHd_swTBlkdVn9nxWJS8Z0vLreDvjO-0kEFwRtR29YefMZMoOXD7hltJQOf7I8iR3A_3HIg-WMAw
```

Also has a handful of other helper methods, such as one to generate a service login token.