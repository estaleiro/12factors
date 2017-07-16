# 12factors

This is 12Factors demonstration APP.

It's composed, now, by 3 Factors that will be demonstrated

## Factor 3

The environment variable ``MENSAGEM`` should be defined, otherwise calling '/factor3' will return an Error 500

## Factor 6

This Factor have 2 endpoints: ``/factor6/fs`` and ``/factor6/mc``

While the **FS** Endpoint stores the sessions in Server's Filesystem, the **MC** Endpoint will try to store this in a Memcached Server defined in environment variable MEMCACHE_HOST.

If the MEMCACHE_HOST variable is not defined, it'll try to connect to 'localhost:11211'.

## Factor 9

This factor have 3 endpoints: ``/factor9/status``, ``/factor9/habilita`` and ``/factor9/desabilita``.

Using the **desabilita** endpoint will cause the server to return an Error 500 when calling **status**

Using the **habilita** endpoint will cause the server to return an OK message when calling **status**

The default behaviour (when server starts) is to return an OK message
