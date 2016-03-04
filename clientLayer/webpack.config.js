var path = require('path');
var webpack = require('webpack')

module.exports = {
  cache: true,
  entry: './client.js',
  output: {
    path: __dirname,
    publicPath: '/',
    filename: 'client.build.js',
  },
  module: {
    loaders: [{
      test: /\.jsx?$/,
      exclude: /(node_modules)/,
      loader: 'babel?stage=0'
    }]
  },
  resolve: {
    extensions: ['', '.js', '.jsx']
  },
  plugins: []
};