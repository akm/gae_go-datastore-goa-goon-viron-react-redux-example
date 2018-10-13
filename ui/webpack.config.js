const HtmlWebpackPlugin = require('html-webpack-plugin')
const CleanWebpackPlugin = require('clean-webpack-plugin')
const path = require('path');

const MODE = 'development';
const enabledSourceMap = (MODE === 'development');

module.exports = {
  mode: MODE,
  entry: "./src/index.tsx",
  output: {
    filename: "bundle.js",
    path: path.resolve(__dirname, "dist")
  },
  devtool: "source-map", // for debug
  resolve: {
    extensions: [".ts", ".tsx", ".js", ".json"]
  },
  module: {
    rules: [
      { test: /\.tsx?$/, loader: "awesome-typescript-loader" },
      { enforce: "pre", test: /\.js$/, loader: "source-map-loader" },
      { test: /\.hbs$/, loader: 'handlebars-loader' },
      {
        test: /\.css/,
        use: [
          'style-loader',
          {
            loader: 'css-loader',
            options: {
              url: false,
              minimize: true,
              sourceMap: enabledSourceMap,
            },
          },
        ],
      },

    ]
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: 'src/templates/index.hbs'
    }),
    new CleanWebpackPlugin(["dist"]),
  ]
};
