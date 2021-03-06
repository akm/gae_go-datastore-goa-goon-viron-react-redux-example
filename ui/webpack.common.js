const HtmlWebpackPlugin = require('html-webpack-plugin')
const CleanWebpackPlugin = require('clean-webpack-plugin')
const path = require('path');

module.exports = ({sourceMap=false} = {}) => { 
  return {
    entry: "./src/index.tsx",
    output: {
      filename: "bundle.js",
      path: path.resolve(__dirname, "dist")
    },
    resolve: {
      extensions: [".ts", ".tsx", ".js", ".json"]
    },
    module: {
      rules: [
        { test: /\.tsx?$/, loader: "awesome-typescript-loader" },
        { enforce: "pre", test: /\.js$/, loader: "source-map-loader" },
        { test: /\.hbs$/, loader: 'handlebars-loader' },
        {
          test: /\.css$/,
          use: [
            'style-loader',
            {
              loader: 'css-loader',
              options: {
                url: false,
                sourceMap: sourceMap
              }
            }
          ]
        },
        {
          test: /src\/styles\/.*\.scss$/,
          use: [
            "style-loader",
            {
              loader: 'css-loader',
              options: {
                url: true,
                sourceMap: sourceMap,
                importLoaders: 2
              }
            },
            {
              loader: 'sass-loader',
              options: {sourceMap: sourceMap } 
            }
          ]
        },
        {
          test: /src\/styles\/.*\.(gif|png|jpg|eot|wof|woff|woff2|ttf|svg)$/,
          loader: 'url-loader'
        },
      ]
    },
    plugins: [
      new HtmlWebpackPlugin({template: 'src/templates/index.hbs'}),
      new CleanWebpackPlugin(["dist"]),
    ]
  }
};
