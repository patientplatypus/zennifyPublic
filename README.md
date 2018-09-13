# zennify.me

## To Run: 

### frontend

#### Project setup
```
npm install
```

#### Compiles and hot-reloads for development
```
npm run serve
```

### backend: 

- `docker-compose down && docker-compose build --no-cache && docker-compose up` after cding into `zennify/backend`.

## You will also need ...

To get IPFS to work you will need to install and run the rendezvous server which can be found here - `https://github.com/libp2p/js-libp2p-websocket-star-rendezvous`. This is a websocket server that allows the talk page to use pub-sub capabilities. To start after installation, open a new terminal window and enter `rendezvous`. Please note that there is currently a memory leak in the server that gives the following warning: 

`(node:64740) MaxListenersExceededWarning: Possible EventEmitter memory leak detected. 11 disconnect listeners added. Use emitter.setMaxListeners() to increase limit`

The maintainers of the library are aware of the bug and believe that is happening with their socket.io implementation. Will possibly need to migrate to another websocket server or implement an original one in the future.