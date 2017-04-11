const path = require('path');

// Exports our modules 
module.exports = {
    target: 'electron-main',
    node : {
        __dirname: false
    },
    entry : {
       electron: './electron/app.js',
       component: './electron/entry/entry.js'
    },
    module: {
        rules: [{
            test: /\.(jsx|js)$/,
            exclude : '/(node_modules)/',
            use: {
                loader: 'babel-loader',
                options: {
                    presets: ['env','react']
                }
            }
        }]
    },
    resolve : {
        extensions : ['.js', '.jsx']
    },
    output : {
        filename : '[name].js',
        path: path.resolve(__dirname, './electron/dist')
    },
};