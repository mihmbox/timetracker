/**
 * Created by mik on 19.9.15.
 */

var webpack = require('webpack');
var path = require('path');
var node_modules = path.resolve(__dirname, 'node_modules');

var publicPath = 'public/build';

// Libraries to include in vendor bundle and ignore any processing
var vendorScripts = {
    'react': path.resolve(node_modules, 'react/dist/react.min.js'),
    'jquery': path.resolve(node_modules, 'jquery/dist/jquery.min.js'),
    'materialize': path.resolve(node_modules, 'materialize-css/dist/js/materialize.min.js')
};
// Scripts without any dependencies(any dependencies will be ignored during processing)
var noParseModules = [
    vendorScripts['react'],
    vendorScripts['jquery']
];

var getEntryPath = function (name) {
    return path.resolve(__dirname, 'public/' + name);
};

module.exports = {
    entry: {
        'home': [getEntryPath('home.jsx')],
        'signin': [getEntryPath('signin.jsx')],
        'signup': [getEntryPath('signup.jsx')],
        'vendor': [
            'webpack/hot/dev-server',
            'webpack-dev-server/client?http://localhost:8080'
        ].concat(
            Object.keys(vendorScripts)
        )
    },
    resolve: {
        alias: vendorScripts
    },
    output: {
        path: path.resolve(__dirname, publicPath),
        filename: '[name].js',
        publicPath: publicPath
    },
    plugins: [
        new webpack.HotModuleReplacementPlugin(),
        new webpack.NoErrorsPlugin(), // don't reload if there is an erro
        new webpack.optimize.CommonsChunkPlugin(/* chunkName= */"vendor", /* filename= */"vendor.js"),
        new webpack.ProvidePlugin({
            $: 'jquery',
            jQuery: 'jquery'
        })
    ],
    module: {
        loaders: [{
            test: /\.jsx?$/, // A regexp to test the require path. accepts either js or jsx
            loader: 'babel', // The module to load. "babel" is short for "babel-loader"
            exclude: /node_modules/
        }],
        // Scripts without any dependencies(any dependencies will be ignored during processing)
        noParse: noParseModules
    },
    //externals: {
    //    //don't bundle the 'react' npm package with our bundle.js, but get it from a global 'React' variable
    //    'react': 'React'
    //},
    devServer: {
        contentBase: publicPath,
        port: 8080
    }
};