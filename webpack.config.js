/**
 * Created by mik on 19.9.15.
 */

var webpack = require('webpack');
var path = require('path');

var publicPath = 'public/build';
module.exports = {
    entry: [
        //'webpack/hot/dev-server',
        'webpack-dev-server/client?http://localhost:8080',
        path.resolve(__dirname, 'public/main.js')
    ],
    output: {
        path: path.resolve(__dirname, publicPath), //'public/build'),
        filename: 'bundle.js',
        publicPath: publicPath//'/public/build/'
    },
    plugins: [
        new webpack.HotModuleReplacementPlugin(),
        new webpack.NoErrorsPlugin()
    ],
    devServer: {
        contentBase: publicPath,
        //host: hostname,
        port: 8080,
        //historyApiFallback: true,
        //proxy: {
        //    //'/api/v1/*': config.get('api_proxy'),
        //    //'/font/*': config.get('api_proxy'),
        //}
    }
};