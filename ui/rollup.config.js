// node-resolve will resolve all the node dependencies

import babel from 'rollup-plugin-babel'
import serve from 'rollup-plugin-serve'

export default {
  input: 'src/index.js',
  output: {
    file: 'dist/bundle.js',
    format: 'iife'
  },
  plugins: [
    babel({
      exclude: 'node_modules/**'
    }),
    // https://github.com/thgh/rollup-plugin-serve#options
    serve({
      contentBase: 'dist',
      port: 8080
    })
  ]
};
