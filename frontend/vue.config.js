module.exports = {
    configureWebpack: { 
        module: {
            rules: [{
                test: /\.rs$/,
                use: [{
                loader: 'wasm-loader'
                }, {
                loader: 'rust-native-wasm-loader'
                }]
            }]
        },
        output: { globalObject: 'this', }, 
    },
    devServer: {
        port: 8081
    }
}