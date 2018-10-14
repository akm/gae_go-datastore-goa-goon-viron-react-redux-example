const merge = require('webpack-merge');
const common = require('./webpack.common.js');

module.exports = merge(common({sourceMap: true}), {
  mode: 'development',
  devtool: 'source-map',
  devServer: {
    contentBase: './dist'
  }
});
