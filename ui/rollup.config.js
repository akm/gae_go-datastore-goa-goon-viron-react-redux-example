// node-resolve will resolve all the node dependencies
import resolve from 'rollup-plugin-node-resolve';
import commonjs from 'rollup-plugin-commonjs';
import babel from 'rollup-plugin-babel';
import replace from 'rollup-plugin-replace';
import json from 'rollup-plugin-json';
import builtins from 'rollup-plugin-node-builtins';
import serve from 'rollup-plugin-serve'

export default {
  input: 'src/index.js',
  output: {
    file: 'dist/bundle.js',
    format: 'iife'
  },
  // external: [
  //   'react'
  // ],
  plugins: [
    // https://github.com/calvinmetcalf/rollup-plugin-node-builtins
    builtins(),
    // https://github.com/rollup/rollup-plugin-babel
    babel({
      exclude: 'node_modules/**'
    }),
    // https://github.com/rollup/rollup-plugin-commonjs
    commonjs({
      exclude: 'node_modules/process-es6/**',
      include: [
        //some react related modules i need
        'node_modules/fbjs/**',
        'node_modules/object-assign/**',
        'node_modules/react/**',
        'node_modules/react-dom/**',
        'node_modules/prop-types/**',
        'node_modules/create-react-class/**', // adding the module with that "default not exported by" message to this includes list, made that message go away
        'node_modules/axios/**',
      ]
    }),
    // https://github.com/rollup/rollup-plugin-node-resolve
    resolve(),
    // https://github.com/rollup/rollup-plugin-replace
    replace({
      'process.env.NODE_ENV': JSON.stringify( 'production' )
    }),
    // https://github.com/rollup/rollup-plugin-json
    json(),
    serve({
      contentBase: 'dist',
      port: 8080
    })
  ]
};
