// node-resolve will resolve all the node dependencies
import resolve from 'rollup-plugin-node-resolve';
import commonjs from 'rollup-plugin-commonjs';
import babel from 'rollup-plugin-babel';
import replace from 'rollup-plugin-replace';
import json from 'rollup-plugin-json';
import builtins from 'rollup-plugin-node-builtins';
import serve from 'rollup-plugin-serve'

const rollupArgs = process.argv.slice(2)

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
      include: 'node_modules/**',
      exclude: 'node_modules/process-es6/**',
      namedExports: {
        'node_modules/react/index.js': ['Children', 'Component', 'PropTypes', 'createElement'],
        'node_modules/react-dom/index.js': ['render']
      },
    }),
    // https://github.com/rollup/rollup-plugin-node-resolve
    resolve(),
    // https://github.com/rollup/rollup-plugin-replace
    replace({
      'global': 'window',
      'process.env.NODE_ENV': JSON.stringify( 'production' )
    }),
    // https://github.com/rollup/rollup-plugin-json
    json(),
    // https://www.npmjs.com/package/rollup-plugin-serve
    rollupArgs.includes('-w') && serve({
      contentBase: 'dist',
      port: 8080
    })
  ]
};
